package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/hashicorp/hcl/hcl/printer"
	"github.com/hashicorp/hcl/json/parser"
)

func main() {
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
