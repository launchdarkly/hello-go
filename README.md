### LaunchDarkly Sample Go Application  ###
We've built a simple console application that demonstrates how LaunchDarkly's SDK works.  Below, you'll find the basic build procedure, but for more comprehensive instructions, you can visit your [Quickstart page](https://app.launchdarkly.com/quickstart#/).
##### Build instructions  #####
1. Make sure you have [Godep](https://github.com/tools/godep) installed: Run `go get github.com/tools/godep`
2. Copy your API key and feature flag key from your LaunchDarkly dashboard into `main.go` 
3. Run `godep go build`