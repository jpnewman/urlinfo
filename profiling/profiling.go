// Package profiling is a small package for profiling.
package profiling

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime/pprof"
)

// StartCPUProfiling Start CPU Profiling
func StartCPUProfiling(cpuProfile *string) error {
	var err error
	if *cpuProfile != "" {
		f, errCreate := os.Create(*cpuProfile)
		err = errCreate
		pprof.StartCPUProfile(f)
	}

	return err
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

		f, err := os.Create(fmt.Sprintf("%s_%s%s", fileTitle, tag, extension))
		pprof.WriteHeapProfile(f)
		f.Close()
		return err
	}

	return err
}
