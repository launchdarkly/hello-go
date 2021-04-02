package main

import (
	"fmt"
	"os"
	"time"

	"gopkg.in/launchdarkly/go-sdk-common.v2/lduser"
	ld "gopkg.in/launchdarkly/go-server-sdk.v5"
)

// Set sdkKey to your LaunchDarkly SDK key.
const sdkKey = ""

// Set featureFlagKey to the feature flag key you want to evaluate.
const featureFlagKey = "my-boolean-flag"

func showMessage(s string) { fmt.Printf("*** %s\n\n", s) }

func main() {
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

	// Set up the user properties. This user should appear on your LaunchDarkly users dashboard
	// soon after you run the demo.
	user := lduser.NewUserBuilder("example-user-key").
		Name("Sandy").
		Build()

	flagValue, err := ldClient.BoolVariation(featureFlagKey, user, false)
	if err != nil {
		showMessage("error: " + err.Error())
	}

	showMessage(fmt.Sprintf("Feature flag '%s' is %t for this user", featureFlagKey, flagValue))

	// Here we ensure that the SDK shuts down cleanly and has a chance to deliver analytics
	// events to LaunchDarkly before the program exits. If analytics events are not delivered,
	// the user properties and flag usage statistics will not appear on your dashboard. In a
	// normal long-running application, the SDK would continue running and events would be
	// delivered automatically in the background.
	ldClient.Close()
}
