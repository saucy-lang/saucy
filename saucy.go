package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/antlr/antlr4/runtime/Go/antlr"

	"github.com/saucy-lang/saucy/parser"
)

const (
	version = "0.0.1"

	usage = `Usage: saucy (options) [script 1]...[script n]`

	help = `=====================================================
 saucy - The Saucy Programming Language Interpreter
 Copyright 2018 Christopher Simpkins
 Apache License, Version 2.0

 Source: https://github.com/saucy-lang/saucy
=====================================================
 Usage:
  - With command line arguments:
      $ saucy (options) [script 1]...[script n]
 Options:
  -h, --help           Interpreter help
      --usage          Interpreter usage
  -v, --version        Interpreter version`
)

var versionShort, versionLong, helpShort, helpLong, usageLong *bool

func init() {
	// define available command line flags
	versionShort = flag.Bool("v", false, "Interpreter version")
	versionLong = flag.Bool("version", false, "Interpreter version")
	helpShort = flag.Bool("h", false, "Interpreter Help")
	helpLong = flag.Bool("help", false, "Interpreter Help")
	usageLong = flag.Bool("usage", false, "Interpreter Usage")
}

type TreeShapeListener struct {
	*parser.BaseSaucyListener
}

func NewTreeShapeListener() *TreeShapeListener {
	return new(TreeShapeListener)
}

func (this *TreeShapeListener) EnterEveryRule(ctx antlr.ParserRuleContext) {
	fmt.Println(ctx.GetText())
}

func main() {
	flag.Parse()

	// parse command line flags and handle them
	switch {
	case *versionShort, *versionLong:
		fmt.Printf("saucy v%s\n", version)
		os.Exit(0)
	case *helpShort, *helpLong:
		fmt.Println(help)
		os.Exit(0)
	case *usageLong:
		fmt.Println(usage)
		os.Exit(0)
	}

	sourceFile := flag.Args()[0]

	input, _ := antlr.NewFileStream(sourceFile)
	lexer := parser.NewSaucyLexer(input)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	p := parser.NewSaucyParser(stream)
	p.AddErrorListener(antlr.NewDiagnosticErrorListener(true))
	p.BuildParseTrees = true
	tree := p.File_input()
	antlr.ParseTreeWalkerDefault.Walk(NewTreeShapeListener(), tree)
}
