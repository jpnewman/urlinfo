package main

import (
	"fmt"
	"os"
	"time"

	logging "github.com/jpnewman/urlinfo/logging"
	"github.com/jpnewman/urlinfo/profiling"
	"github.com/jpnewman/urlinfo/report"
)

// Report Report
var Report = report.New()

func main() {
	defer profiling.Elapsed("Program Done")()

	logging.LogInit("urlinfo.log")
	logging.Logger.Infof("Program Started: %s", os.Args)

	args := parseArgs()
	profiling.StartCPUProfiling(args.cpuProfile)

	Report.SetFormatter(args.reportFormat)
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
	defer profiling.ProfileMem(args.memProfile, "Done")

	fmt.Println("Done!!!")

	logging.LogFileClose()
}
