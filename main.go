package main

import (
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
        config = make(map[string]interface{}, 0)
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
        for i, flag := range os.Args {
                if flag[0] != '-' {
                        continue
                }

                switch flag {
                case "-f":
                        configFilesOnly = true
                        query, err = regexp.Compile(".*")
                case "-g":
                        configFilesOnly = true
                        query, err = regexp.Compile(os.Args[i+1])
                default:
                }

                if err != nil {
                        break
                }
        }

        if !configFilesOnly {
                query, err = regexp.Compile(os.Args[1])
        }

	return err
}

func parseResults() {
        for ; len(resultChannel) > 0 ; {
                res := <-resultChannel
                fmt.Println(res.Path)
                for _, line := range res.Results {
                        fmt.Printf("%d:%s\n", line.Lineno, line.Line)
                }
                if len(res.Results) > 0 {
                        fmt.Println("")
                }
        }
}
