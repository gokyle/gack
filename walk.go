package main

import (
	"bufio"
	"io"
	"os"
	"path/filepath"
	"regexp"
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

func fileScanner(workers chan int) {
	for {
		path, ok := <-fileChannel
		if !ok {
			<-workers
			return
		} else if shouldIgnoreFile(path) {
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
		if err.Error() == "too many open files" {
			fileChannel <- path
		}
		return
	}
	defer file.Close()

	buf := bufio.NewReader(file)
	lineno := 0
	result := new(Result)
	result.Path = path
	longLine := false
	var line []byte

	for {
		err = nil

		lineBytes, isPrefix, err := buf.ReadLine()
		lineno++
		if io.EOF == err {
			break
		} else if err != nil {
			break
		} else if matched, _ := regexp.Match("\x00", lineBytes); matched {
			break
		} else if isPrefix {
			line = extendLine(line, lineBytes)
			longLine = true
			continue
		} else if longLine {
			line = extendLine(line, lineBytes)
			longLine = false
		} else {
			line = lineBytes
		}

		if matched := query.Match(line); matched {
			result.Results = append(result.Results,
				ResultLine{string(line), lineno})
		}

	}

	if len(result.Results) > 0 {
		resultChannel <- result
	}
}
