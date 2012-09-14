package main

import (
        "regexp"
)

var query *regexp.Regexp

const fileMax = 1024
const numWorkers = 4
var fileChannel chan string       // stores the files that need to be scanned
var resultChannel chan *Result      // stores results

type ResultLine struct {
        Line string
        Lineno int
}

type Result struct {
	Path   string
        Results []ResultLine
}

// constants and globals for the gack utility
var ignore_dirs = []string{
	"^bzr$",            // Bazaar
	"^cdv$",            // Codeville
	"^~.dep$",          // Interface builder
	"^~.dot$",          // Interface builder
	"^~.nib$",          // Interface builder
	"^~.plst$",         // Interface builder
	"^.git$",           // Git
	"^.hg$",            // Mercurial
	"^.pc$",            // Quilt
	"^.svn$",           // Subversion
	"^_MTN$",           // Monotone
	"^blib$",           // Perl module building
	"^CVS$",            // CVS
	"^RCS$",            // RCS
	"^SCCS$",           // SCCS
	"^_darcs$",         // darcs
	"^_sgback$",        // Vault/Fortress
	"^autom4te.cache$", // autoconf
	"^cover_db$",       // Devel::Cover
	"^_build$",         // Module::Build
}

var ignore_files = []string {
        "^[.].*\\w+[.]swp$",
}

var language_files map[string][]*regexp.Regexp

// config vars
var configFilesOnly bool
var configFilesOnlyRegex string

