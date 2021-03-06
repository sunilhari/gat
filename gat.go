package main

import (
	"os"
	"path/filepath"

	"github.com/alecthomas/chroma/lexers"
	"github.com/alecthomas/chroma/quick"
	"github.com/fatih/color"
)

func main() {
	if len(os.Args) != 2 {
		color.Red("Program Exactly requires 1 input file !!")
		os.Exit(0)
	}
	fpath := os.Args[1]
	fileInfo, err := os.Stat(fpath)
	if err != nil {
		color.Red("NonExistent Path !!")
		os.Exit(0)
	}

	if fileInfo.IsDir() {
		color.Yellow("Requires an input file")
		os.Exit(0)
	}
	// simplest file read for now
	dat, err := os.ReadFile(fpath)
	check(err)
	var content = string(dat)
	base := filepath.Base(fpath)
	lexer := lexers.Match(base)
	quick.Highlight(os.Stdout, content, lexer.Config().Name, "terminal256", "monokai")
	color.Green("--------END-------")

}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
