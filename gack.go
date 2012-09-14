// gack is a version of Ack written in Go.
package main

import (
	"path/filepath"
	"regexp"
)

func shouldIgnoreDir(path string) bool {
	basename := filepath.Base(path)
	for _, ignore := range ignore_dirs {
		if matched, _ := regexp.MatchString(ignore, basename); matched {
			return true
		}

	}
	return false
}

func shouldIgnoreFile(path string) bool {
	basename := filepath.Base(path)
	for _, ignore := range ignore_files {
		if matched, _ := regexp.MatchString(ignore, basename); matched {
			return true
		}
	}

	keep := false

	for _, langRegex := range language_files {
		if matched := langRegex.MatchString(basename); matched {
			keep = true
			break
		}
	}
	if keep {
		return false
	}
	return true
}

func extendLine(line []byte, lineBytes []byte) []byte {
	for _, b := range lineBytes {
		line = append(line, b)
	}
	return line
}
