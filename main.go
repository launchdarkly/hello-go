package main

import (
	"fmt"
	ld "github.com/launchdarkly/go-client"
)

func main() {
	// TODO : Enter your LaunchDarkly API key here
	client := ld.MakeClient("YOUR_API_KEY")

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
	show_feature, _ := client.Toggle("YOUR_FEATURE_FLAG_KEY", user, false)

	if show_feature {
		// application code to show the feature
		fmt.Println("Showing your feature to " + *user.Key)
	} else {
		// the code to run if the feature is off
		fmt.Println("Not showing your feature to " + *user.Key)
	}
	client.Flush()
}
