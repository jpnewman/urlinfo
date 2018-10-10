package profiling

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime/pprof"

	logging "github.com/jpnewman/urlinfo/logging"
)

// StartCPUProfiling Start CPU Profiling
func StartCPUProfiling(cpuProfile *string) {
	if *cpuProfile != "" {
		f, err := os.Create(*cpuProfile)
		if err != nil {
			logging.Logger.Fatal(err)
		}
		pprof.StartCPUProfile(f)
	}
}

// StopCPUProfiling Stop CPU Profiling
func StopCPUProfiling(cpuProfile *string) {
	if *cpuProfile != "" {
		pprof.StopCPUProfile()
	}
}

// ProfileMem Profile Memory
func ProfileMem(memProfile *string, tag string) {
	if *memProfile != "" {
		var filename = *memProfile
		var extension = filepath.Ext(filename)
		var fileTitle = filename[0 : len(filename)-len(extension)]

		f, err := os.Create(fmt.Sprintf("%s_%s%s", fileTitle, tag, extension))
		if err != nil {
			logging.Logger.Fatal(err)
		}
		pprof.WriteHeapProfile(f)
		f.Close()
		return
	}
}
