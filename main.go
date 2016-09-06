package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/Acconut/hcl/hcl/printer"
	"github.com/Acconut/hcl/json/parser"
)

// VERSION is what is returned by the `-v` flag
var VERSION = "v0.0.4"

func main() {
	version := flag.Bool("version", false, "Prints current app version")
	flag.Parse()
	if *version {
		fmt.Println(VERSION)
		os.Exit(0)
	}

	if err := convert(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func convert() error {
	input, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		return fmt.Errorf("unable to read from stdin: %s", err)
	}

	ast, err := parser.Parse([]byte(input))
	if err != nil {
		return fmt.Errorf("unable to parse JSON: %s", err)
	}

	err = printer.Fprint(os.Stdout, ast)
	if err != nil {
		return fmt.Errorf("unable to print HCL: %s", err)
	}

	return nil
}
