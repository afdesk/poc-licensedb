package main

import (
	"github.com/go-enry/go-license-detector/v4/licensedb"

	"os"
	"runtime"
	"runtime/pprof"
)

func main() {
	licensedb.Preload()

	fMem, err := os.Create("mem.profile")
	if err != nil {
		panic("could not create memory profile: " + err.Error())
	}
	defer fMem.Close() // error handling omitted for example
	runtime.GC()       // get up-to-date statistics
	if err := pprof.WriteHeapProfile(fMem); err != nil {
		panic("could not write memory profile: " + err.Error())
	}
}
