package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"time"
)

// this file contains the frontend code for the gack utility

func init() {
	fileChannel = make(chan string, fileMax)
	resultChannel = make(chan *Result, fileMax)
	initLanguages()
}

func main() {
	if len(os.Args) == 1 {
		usage()
	}

	if err := configure(); err != nil {
		fmt.Println("[!] ", err)
		os.Exit(1)
	}

	fmt.Println("[-] query: ", query)
	if profiled {
		fmt.Println("[+] profiling enabled")
	}

	workers := make(chan int, numWorkers)
	go func() {
		for i := 0; i < numWorkers; i++ {
			go fileScanner(workers)
			workers <- i
		}
	}()

	go func() {
		err := filepath.Walk(root, walker)
		if err != nil {
			fmt.Println("[!] ", err)
			os.Exit(1)
		}
		walkDone = true
	}()

	go parseResults()

	for !walkDone || len(fileChannel) > 0 {
		time.Sleep(1 * time.Millisecond)
	}
	close(fileChannel)

	for len(workers) > 0 || len(resultChannel) > 0 {
		time.Sleep(1 * time.Millisecond)
	}
	close(resultChannel)
	if profiled {
		killProfile()
	}

	os.Exit(exitStatus)
}

func usage() {
	fmt.Printf("usage: %s <query>\n", filepath.Base(os.Args[0]))
	os.Exit(1)
}

func configure() (err error) {
	flag.BoolVar(&configFilesOnly, "f", false,
		"Only list source files.")
	flag.StringVar(&configFilesOnlyRegex, "g", "",
		"Only list source files that match the specified regex.")
	flag.Parse()

	if configFilesOnlyRegex != "" {
		query = regexp.MustCompile(configFilesOnlyRegex)
		configFilesOnly = true
	} else if configFilesOnly {
		query = regexp.MustCompile(".*")
	} else if !configFilesOnly {
		query = regexp.MustCompile(flag.Args()[0])
	}

	initProfile()
	return err
}

func parseResults() {
	for !walkDone || len(resultChannel) > 0 || len(fileChannel) > 0 {
		res, ok := <-resultChannel
		if !ok {
			break
		}

		if exitStatus == 1 {
			exitStatus = 0
		}

		fmt.Println(res.Path)
		for _, line := range res.Results {
			fmt.Printf("%d:%s\n", line.Lineno, line.Line)
		}
		if len(res.Results) > 0 {
			fmt.Println("")
		}
	}
}
