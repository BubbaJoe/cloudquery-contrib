package main

import (
	"github.com/cloudquery/cloudquery/plugins/source/snyk/resources/plugin"
	"github.com/cloudquery/plugin-sdk/v3/serve"
)

const sentryDSN = "https://c2101fe3c5fd4f91a095b2b37dc6364a@o1396617.ingest.sentry.io/4504333793165313"

func main() {
	serve.Source(plugin.Snyk(), serve.WithSourceSentryDSN(sentryDSN))
}
