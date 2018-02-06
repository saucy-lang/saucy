package main

import (
	"flag"
	"fmt"
	"os"
)

const (
	version = "0.0.1"

	usage = `Usage: saucy (options) [script 1]...[script n]`

	help = `=====================================================
 saucy - The Saucy Programming Language Interpreter
 Copyright 2018 The Saucy Authors
 MIT License

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
}
