module github.com/launchdarkly/hello-go

go 1.18

// READ THIS!
//
// This file has been deliberately checked in with "latest" instead of a specific
// version for the SDK packages, so that the Go compiler will pull in the latest
// published releases when the application is built. In a real application, go.mod
// would specify exact versions, but we want users who are running this demo to
// always get the latest version.
//
// Do not check in changes to this file with these packages pinned to a specific
// version. If you do that, the CI build will detect it and raise an error:
// "Undesirable specific version dependency!" You should then edit the file to
// set the versions of gopkg.in/launchdarkly/go-sdk-common.v2 and
// gopkg.in/launchdarkly/go-server-sdk.v5 back to "latest" and commit the change.

require (
	github.com/launchdarkly/go-sdk-common/v3 latest
	github.com/launchdarkly/go-server-sdk/v7 latest
)

require (
	github.com/google/uuid v1.1.1 // indirect
	github.com/gregjones/httpcache v0.0.0-20171119193500-2bcd89a1743f // indirect
	github.com/josharian/intern v1.0.0 // indirect
	github.com/launchdarkly/ccache v1.1.0 // indirect
	github.com/launchdarkly/eventsource v1.6.2 // indirect
	github.com/launchdarkly/go-jsonstream/v3 v3.0.0 // indirect
	github.com/launchdarkly/go-sdk-events/v3 v3.0.0 // indirect
	github.com/launchdarkly/go-semver v1.0.2 // indirect
	github.com/launchdarkly/go-server-sdk-evaluation/v3 v3.0.0 // indirect
	github.com/mailru/easyjson v0.7.6 // indirect
	github.com/patrickmn/go-cache v2.1.0+incompatible // indirect
	golang.org/x/exp v0.0.0-20220823124025-807a23277127 // indirect
	golang.org/x/sync v0.0.0-20190911185100-cd5d95a43a6e // indirect
)
