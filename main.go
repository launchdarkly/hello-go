package main

import (
	"fmt"
	ld "gopkg.in/launchdarkly/go-client.v2"
	"time"
)

func main() {
	// TODO : Enter your LaunchDarkly API key here
	client, err := ld.MakeClient("YOUR_API_KEY", 5*time.Second)
	if err != nil {
		fmt.Println(err.Error())
	}

	key := "bob@example.com"
	first := "Bob"
	last := "Loblaw"
	custom := map[string]interface{}{"groups": "beta_testers"}

	user := ld.User{Key: &key,
		FirstName: &first,
		LastName:  &last,
		Custom:    &custom,
	}

	// TODO : Enter the key for your feature flag here
	show_feature, err := client.BoolVariation("YOUR_FEATURE_FLAG_KEY", user, false)
	if err != nil {
		fmt.Println(err.Error())
	}

	if show_feature {
		// application code to show the feature
		fmt.Println("Showing your feature to " + *user.Key)
	} else {
		// the code to run if the feature is off
		fmt.Println("Not showing your feature to " + *user.Key)
	}
	client.Close()
}
