package main

import (
	"fmt"
	"os"
	"sync"
	"time"

	"github.com/launchdarkly/go-sdk-common/v3/ldcontext"
	"github.com/launchdarkly/go-sdk-common/v3/ldvalue"
	ld "github.com/launchdarkly/go-server-sdk/v7"
)

func showBanner() {
	fmt.Print("\n        ██       \n" +
		"          ██     \n" +
		"      ████████   \n" +
		"         ███████ \n" +
		"██ LAUNCHDARKLY █\n" +
		"         ███████ \n" +
		"      ████████   \n" +
		"          ██     \n" +
		"        ██       \n")
}

func main() {
	// Set sdkKey to your LaunchDarkly SDK key.
	var sdkKey = os.Getenv("LAUNCHDARKLY_SERVER_KEY")

	if sdkKey == "" {
		showMessage("Please edit main.go to set sdkKey to your LaunchDarkly SDK key first")
		os.Exit(1)
	}

	ldClient, _ := ld.MakeClient(sdkKey, 5*time.Second)
	if ldClient.Initialized() {
		showMessage("SDK successfully initialized!")
	} else {
		showMessage("SDK failed to initialize")
		os.Exit(1)
	}

	// Set up the evaluation context. This context should appear on your LaunchDarkly contexts dashboard
	// soon after you run the demo.
	context := ldcontext.NewBuilder("example-user-key").
		Name("Sandy").
		Build()

	// Set featureFlagKey to the feature flag key you want to evaluate.
	var featureFlagKey = "sample-feature"

	if os.Getenv("LAUNCHDARKLY_FLAG_KEY") != "" {
		featureFlagKey = os.Getenv("LAUNCHDARKLY_FLAG_KEY")
	}

	flagValue, err := ldClient.BoolVariation(featureFlagKey, context, false)
	if err != nil {
		showMessage("error: " + err.Error())
	}

	showMessage(fmt.Sprintf("The '%s' feature flag evaluates to %t.", featureFlagKey, flagValue))

	if flagValue {
		showBanner()
	}

	updateCh := ldClient.GetFlagTracker().AddFlagValueChangeListener(featureFlagKey, context, ldvalue.Null())
	go func() {
		for event := range updateCh {
			fmt.Sprintf("The '%s' feature flag evaluates to %t.", featureFlagKey, event.NewValue)
		}
	}()

	wg := sync.WaitGroup{}
	wg.Add(1)
	wg.Wait()
}
