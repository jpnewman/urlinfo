// Go Program to get page information from URLs.
package main

import (
	"io"
	"os"
	"time"

	"github.com/jpnewman/urlinfo/profiling"
	"github.com/jpnewman/urlinfo/report"
)

// Report Report
var Report = report.New()

func main() {
	LogInit("urlinfo.log")

	defer profiling.Elapsed("Program Done")([]io.Writer{os.Stdout, RootLogger.Out})
	Logger.Infof("Program Started: %s", os.Args)

	args := parseArgs()

	err := profiling.StartCPUProfiling(args.cpuProfile)
	if err != nil {
		Logger.Fatal(err)
	}

	Report.SetFormatter(*args.reportFormat)
	Report.SetOutputFile(args.reportFile)

	Report.PrintHeader("URLInfo")

	urls, errs := readURLFile(args.urlFile, 5)
	printFileDetails(urls, errs)

	processURLs(urls, &processURLsArgs{
		httpTimeoutMilliseconds: time.Duration(time.Duration(*args.httpTimeout) * time.Millisecond),
		numberOfWorkers:         *args.numberOfWorkers,
		getHeadOny:              *args.getHeadOny,
		dontFollowRedirects:     *args.dontFollowRedirects,
		dryRun:                  *args.dryRun,
	})

	defer profiling.StopCPUProfiling(args.cpuProfile)

	err = profiling.ProfileMem(args.memProfile, "Done")
	if err != nil {
		Logger.Fatal(err)
	}

	LogFileClose()
}
