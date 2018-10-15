// Go Program to get page information from URLs.
package main

import (
	"os"
	"time"

	"github.com/jpnewman/urlinfo/profiling"
	"github.com/jpnewman/urlinfo/report"
)

// Report Report
var Report = report.New()

func main() {
	LogInit("urlinfo.log")

	args := parseArgs()

	defer func() {
		errDoneProf := profiling.ProfileMem(args.memProfile, "Done")
		if errDoneProf != nil {
			Logger.Fatal(errDoneProf)
		}
	}()

	defer profiling.TimeElapsed("Program Done")(Report.PrintMessage, LogPrintInfo)
	Logger.Infof("Program Started: %s", os.Args)

	errsProfiling := profiling.StartCPUProfiling(args.cpuProfile)
	for _, errProfiling := range errsProfiling {
		if errProfiling != nil {
			Logger.Fatal(errProfiling)
		}
	}
	defer profiling.StopCPUProfiling(args.cpuProfile)

	Report.SetFormatter(*args.reportFormat)
	Report.SetOutputFile(args.reportFile)

	Report.PrintHeader("URLInfo")

	urls, errs := readURLFile(*args.urlFile)
	printFileDetails(urls, errs)

	pArgs := processURLsArgs{
		httpTimeoutMilliseconds: time.Duration(time.Duration(*args.httpTimeout) * time.Millisecond),
		numberOfWorkers:         *args.numberOfWorkers,
		getHeadOny:              *args.getHeadOny,
		dontFollowRedirects:     *args.dontFollowRedirects,
		dryRun:                  *args.dryRun,
	}

	client := createHTTPClient(pArgs.httpTimeoutMilliseconds, pArgs.dontFollowRedirects)
	processURLs(urls, &pArgs, client)

	LogFileClose()
}
