package main

import (
	"fmt"
	"github.com/launchdarkly/go-sdk-common/v3/ldlog"
	"github.com/launchdarkly/go-server-sdk/v6/ldcomponents"
	"github.com/launchdarkly/go-server-sdk/v6/ldfiledata"
	"github.com/launchdarkly/go-server-sdk/v6/ldfilewatch"
	"time"

	ld "github.com/launchdarkly/go-server-sdk/v6"
)

func main() {

	config := ld.Config{}

	config.Logging = ldcomponents.Logging().MinLevel(ldlog.Error)
	config.DataSource = ldfiledata.DataSource().FilePaths("path").Reloader(ldfilewatch.WatchFiles)

	ld.MakeCustomClient("abc", config, 5*time.Second)

	// This line is here just to make the CI demo pass.
	fmt.Println("is true for this context")
}
