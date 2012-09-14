package main

import (
        "flag"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
        "time"
)

var root = "."

// this file contains the frontend code for the gack utility

func init() {
        fileChannel = make(chan string, fileMax)
        resultChannel = make(chan *Result, fileMax)
}

func main() {
	if len(os.Args) == 1 {
		usage()
	}

	if err := configure(); err != nil {
		fmt.Println("[!] ", err)
		os.Exit(1)
	}

        go func() {
                for i := 0; i < numWorkers; i++ {
                        go fileScanner()
                }
        }()

	        err := filepath.Walk(root, walker)
	        if err != nil {
		        fmt.Println("[!] ", err)
		        os.Exit(1)
	        }

        for ; len(fileChannel) > 0 ; {
                time.Sleep(10 * time.Millisecond)
        }

        parseResults()
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
                query, err = regexp.Compile(os.Args[1])
        }
    

	return err
}

func parseResults() {
        results := 0
        for ; len(resultChannel) > 0 ; {
                res := <-resultChannel
                results++
                fmt.Println(res.Path)
                for _, line := range res.Results {
                        fmt.Printf("%d:%s\n", line.Lineno, line.Line)
                }
                if len(res.Results) > 0 {
                        fmt.Println("")
                }
        }

        if results == 0 {
                os.Exit(1)
        } else {
                os.Exit(0)
        }
}
