package main

import (
	"flag"
	"os"

	logging "github.com/jpnewman/urlinfo/logging"
)

type args struct {
	urlFile             *string
	httpTimeout         *int
	numberOfWorkers     *int
	getHeadOny          *bool
	dontFollowRedirects *bool
	reportFormat        *string
	reportFile          *string
	cpuProfile          *string
	memProfile          *string
	dryRun              *bool
}

func parseArgs() args {
	r := args{}
	r.urlFile = flag.String("urlFile", "", "[Required] Path to file containing a list of URLs.")
	r.httpTimeout = flag.Int("httpTimeout", 3000, "Http Timeout in Milliseconds")
	r.numberOfWorkers = flag.Int("numberOfWorkers", 5, "Number of workers")
	r.getHeadOny = flag.Bool("getHeadOny", false, "Get HTTP Headers only")
	r.dontFollowRedirects = flag.Bool("dontFollowRedirects", false, "Don't Follow HTTP Redirects")
	r.reportFormat = flag.String("reportFormat", "Console", "Report format")
	r.reportFile = flag.String("reportFile", "", "Report file")
	r.cpuProfile = flag.String("cpuprofile", "", "write cpu profile to file")
	r.memProfile = flag.String("memprofile", "", "write memory profile to this file")
	r.dryRun = flag.Bool("dryrun", false, "Dry-Run")
	flag.Parse()

	if *r.urlFile == "" {
		logging.Logger.Error("Program argument '-urlFile' not specified!")
		flag.PrintDefaults()
		os.Exit(1)
	}

	return r
}
