// Package profiling is a small package for profiling.
package profiling

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"runtime/pprof"
)

// StartCPUProfiling Start CPU Profiling
func StartCPUProfiling(cpuProfile *string) []error {
	var errs []error
	if *cpuProfile != "" {
		f, errCreate := os.Create(*cpuProfile)
		errs = append(errs, errCreate)

		errStartProf := pprof.StartCPUProfile(f)
		errs = append(errs, errStartProf)
	}

	return errs
}

// StopCPUProfiling Stop CPU Profiling
func StopCPUProfiling(cpuProfile *string) {
	if *cpuProfile != "" {
		pprof.StopCPUProfile()
	}
}

// ProfileMem Profile Memory
func ProfileMem(memProfile *string, tag string) error {
	var err error
	if *memProfile != "" {
		var filename = *memProfile
		var extension = filepath.Ext(filename)
		var fileTitle = filename[0 : len(filename)-len(extension)]

		runtime.GC()

		f, err := os.Create(fmt.Sprintf("%s_%s%s", fileTitle, tag, extension))
		pprof.WriteHeapProfile(f)
		f.Close()
		return err
	}

	return err
}
