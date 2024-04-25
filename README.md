# LaunchDarkly sample Go application

We've built a simple console application that demonstrates how LaunchDarkly's SDK works.

Below, you'll find the basic build procedure. For more comprehensive instructions, you can visit your [Quickstart page](https://app.launchdarkly.com/quickstart#/) or the [Go SDK reference guide](https://docs.launchdarkly.com/sdk/server-side/go).

This demo requires Go 1.21 or higher.

## Build instructions

1. Set the environment variable `LAUNCHDARKLY_SERVER_KEY` to your LaunchDarkly SDK key. If there is an existing boolean feature flag in your LaunchDarkly project that you want to evaluate, set `LAUNCHDARKLY_FLAG_KEY` to the flag key; otherwise, a boolean flag of `sample-feature` will be assumed.
    ```bash
    export LAUNCHDARKLY_SERVER_KEY="1234567890abcdef"
    export LAUNCHDARKLY_FLAG_KEY="my-boolean-flag"
    ```

2. On the command line, run `go build`. Follow the instruction to run `go mod tidy` if `go build` said updates to go.mod is needed.

3. Run `./hello-go`

You should receive the message "The <flagKey> feature flag evaluates to <flagValue>.". The application will run continuously and react to the flag changes in LaunchDarkly.
