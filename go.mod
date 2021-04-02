module github.com/launchdarkly/hello-go

go 1.14

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
	gopkg.in/launchdarkly/go-sdk-common.v2 latest
	gopkg.in/launchdarkly/go-server-sdk.v5 latest
)
