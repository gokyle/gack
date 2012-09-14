package main

import (
        "fmt"
        "os"
        "regexp"
)

// contains functions for including or excluding files by language

func addExtension(regex, ext string) string {
        if len(regex) > 0 {
                regex += "|"
        }
        re_string := fmt.Sprintf("(^[.].*\\w+[.]%s$)", ext)
        return regex + re_string
}

func buildLanguageExtensions(exts []string) *regexp.Regexp {
        var regex string
        for _, ext := range exts {
                regex = addExtension(regex, ext)
        }

        extRegex, err := regexp.Compile(regex)
        if err != nil {
                fmt.Println("[!] cannot compile regex: ", err)
                os.Exit(1)
        }

        return extRegex
}
