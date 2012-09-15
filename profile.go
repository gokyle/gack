package main

import (
	"log"
	"os"
	"runtime/pprof"
)

var f *os.File
var memerr error

func initProfile() {
	if *cpuprofile != "" {
		log.Println("[-] will profile")
		f, err := os.Create(*cpuprofile)
		if err != nil {
			log.Fatal(err)
		}
		profiled = true
		pprof.StartCPUProfile(f)
	}

	if *memprofile != "" {
		f, memerr = os.Create(*memprofile)
		if memerr != nil {
			log.Fatal(memerr)
		}
	}
}

func killProfile() {
	log.Println("[-] profile stopped")
	if *cpuprofile != "" {
		log.Println("[-] cpu profile written to ", *cpuprofile)
		pprof.StopCPUProfile()
	}

	if *memprofile != "" {
		log.Println("[-] heap profile written to ", *memprofile)
		pprof.WriteHeapProfile(f)
		f.Close()
	}
}
