package main

import (
        "bufio"
        "io"
        "log"
	"os"
        "regexp"
	"path/filepath"
)

func walker(path string, info os.FileInfo, err error) error {
	if info.Mode().IsDir() == false {
                fileChannel <- path
	} else {
                if shouldIgnoreDir(path) {
                        return filepath.SkipDir
                }
        }
	return nil
}

func fileScanner() {
        for {
                path := <-fileChannel
                if shouldIgnoreFile(path) {
                        continue
                } else if configFilesOnly && query.MatchString(path) {
                        res := new(Result)
                        res.Path = path
                        resultChannel <- res
                } else if !configFilesOnly {
                        go scanFile(path)
                }
        }
}

func scanFile(path string) {
        file, err := os.Open(path)
	if err != nil {
		log.Println("[!] ReadConfig open: ", err)
		return
	}
	defer file.Close()

	buf := bufio.NewReader(file)
        lineno := 0
        result := new(Result)
        result.Path = path

	for {
		err = nil

		lineBytes, isPrefix, err := buf.ReadLine()
                lineno++
		if io.EOF == err {
			break
		} else if err != nil {
			log.Println("[!] ReadConfig read line: ", err)
			return
		} else if matched, _ := regexp.Match("\x00", lineBytes); matched {
                        return
                } else if isPrefix {
			log.Println("[!] ReadConfig line unexpectedly long (",
				path, ")")
			return
		}

                if query.Match(lineBytes) {
                        result.Results = append(result.Results, 
                                ResultLine{string(lineBytes), lineno})
                }
                         
        }

        if len(result.Results) > 0 {
                resultChannel <- result
        }
}
