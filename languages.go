package main

import (
	"fmt"
	"regexp"
        "strings"
)

// contains functions for including or excluding files by language

func addFile(regex, filename string) string {
        if len(regex) > 0 {
                regex += "|"
        }
        re_string := fmt.Sprintf("(^%s$)", filename)
        return regex + re_string
}

func addFiles(files string) *regexp.Regexp {
        regex := ""
        for _, file := range strings.Fields(files) {
               regex = addFile(regex, file)
        }

        return regexp.MustCompile(regex)
}

func addExtension(regex, ext string) string {
	if len(regex) > 0 {
		regex += "|"
	}
	re_string := fmt.Sprintf("(\\w+[.]%s$)", ext)
	return regex + re_string
}

func buildLanguageExtensions(exts []string) *regexp.Regexp {
	var regex string
	for _, ext := range exts {
		regex = addExtension(regex, ext)
	}

	extRegex := regexp.MustCompile(regex)
	return extRegex
}

func addLanguage(language, extensions string) {
	regex := buildLanguageExtensions(strings.Fields(extensions))
	language_files[language] = regex
}

func initLanguages() {
	language_files = make(map[string]*regexp.Regexp, 0)
	addLanguage("actionscript", "as mxml")
	addLanguage("ada", "ada adb ads")
	addLanguage("asm", "asm s S")
	addLanguage("batch", "bat cmd")
	addLanguage("cc", "c h xs")
	addLanguage("cfmx", "cfc cfm cfml")
	addLanguage("clojure", "clj")
	addLanguage("cpp", "cpp cc cxx m hpp hh h hxx")
	addLanguage("csharp", "cs")
	addLanguage("css", "css")
	addLanguage("delphi", "pas int dfm nfm dof dpk dproj groupproj bdsgroup bdsproj")
	addLanguage("elisp", "el")
	addLanguage("erlang", "erl hrl")
	addLanguage("fortran", "f f77 f90 f95 f03 for ftn fpp")
	addLanguage("go", "go")
	addLanguage("groovy", "groovy gtmpl gpp grunit")
	addLanguage("haskell", "hs lhs")
	addLanguage("hh", "h")
	addLanguage("html", "htm html shtml xhtml")
	addLanguage("java", "java properties")
	addLanguage("js", "js")
	addLanguage("jsp", "jsp jspx jhtm jhtml")
	addLanguage("lisp", "lisp lsp")
	addLanguage("lua", "lua")
	addLanguage("mason", "mas mhtml mpl mtxt")
	addLanguage("objc", "m h")
	addLanguage("objcpp", "mm h")
	addLanguage("ocaml", "ml mli")
	addLanguage("parrot", "pir pasm pmc ops pod pg tg")
	addLanguage("perl", "pl pm pm6 pod t psgi")
	addLanguage("php", "php phpt php3 php4 php5 phtml")
	addLanguage("plone", "pt cpt metadata cpy py")
	addLanguage("python", "py")
	addLanguage("ruby", "rb rhtml rjs rxml erb rake spec")
	addLanguage("scala", "scala")
	addLanguage("scheme", "scm ss")
	addLanguage("shell", "sh bash csh tcsh ksh zsh")
	addLanguage("smalltalk", "st")
	addLanguage("sql", "sql ctl")
	addLanguage("tcl", "tcl itcl itk")
	addLanguage("tex", "tex cls sty")
	addLanguage("tt", "tt tt2 ttml")
	addLanguage("vb", "bas cls frm ctl vb resx")
	addLanguage("verilog", "v vh sv")
	addLanguage("vhdl", "vhd vhdl")
	addLanguage("vim", "vim")
	addLanguage("xml", "xml dtd xsl xslt ent")
	addLanguage("yaml", "yaml yml")
        language_files["make"] = addFiles("Makefile")
        language_files["projects"] = addFiles("README([.]\\w+)? LICENSE " +
                "COPYRIGHT NOTICE NOTES TODO")
        language_files["rake"] = addFiles("Rakefile")
}
