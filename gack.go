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
        return false
}
