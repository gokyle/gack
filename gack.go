// gack is a version of Ack written in Go.
package main

import (
	"path/filepath"
	"regexp"
)

func initIgnoreDirs() {
	for _, dir := range ignore_dir_strings {
		ignore_dirs = append(ignore_dirs, regexp.MustCompile(dir))
	}
}

func shouldIgnoreDir(path string) bool {
	basename := filepath.Base(path)
	for _, ignore := range ignore_dirs {
		if ignore.MatchString(basename) {
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

	for _, langRegex := range language_files {
		if langRegex.MatchString(basename) {
			return false
		}
	}
	return true
}

func extendLine(line []byte, lineBytes []byte) []byte {
	for _, b := range lineBytes {
		line = append(line, b)
	}
	return line
}
