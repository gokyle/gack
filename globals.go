package main

// constants and globals for the gack utility
import (
	"flag"
	"regexp"
)

var query *regexp.Regexp
var root = "."
var exitStatus = 1
var walkDone = false

// profile vars
var cpuprofile = flag.String("cpuprofile", "", "write cpu profile to file")
var memprofile = flag.String("memprofile", "", "write memory profile to this file")
var profiled = false

// control the number of works and files that we can open at once
const fileMax = 4096
const numWorkers = fileMax

// channels for workers
var fileChannel chan string    // stores the files that need to be scanned
var resultChannel chan *Result // stores results

var ignore_dir_strings = []string{
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
var ignore_dirs = []*regexp.Regexp{}
var ignore_files = []string{}
var language_files map[string]*regexp.Regexp

// config vars
var configFilesOnly bool
var configFilesOnlyRegex string

// types

// ResultLine stores a single line of matching data.
type ResultLine struct {
	Line   string
	Lineno int
}

// A result represents a file and any matches it has.
type Result struct {
	Path    string
	Results []ResultLine
}
