package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/hashicorp/hcl"
	"github.com/hashicorp/hcl/hcl/printer"
	jsonParser "github.com/hashicorp/hcl/json/parser"
)

// VERSION is what is returned by the `-v` flag
var Version = "development"

func main() {
	version := flag.Bool("version", false, "Prints current app version")
	reverse := flag.Bool("reverse", false, "Input HCL, output JSON")
	flag.Parse()
	if *version {
		fmt.Println(Version)
		return
	}

	var err error
	if *reverse {
		err = toJSON()
	} else {
		err = toHCL()
	}

	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func toJSON() error {
	input, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		return fmt.Errorf("unable to read from stdin: %s", err)
	}

	var v interface{}
	err = hcl.Unmarshal(input, &v)
	if err != nil {
		return fmt.Errorf("unable to parse HCL: %s", err)
	}

	json, err := json.MarshalIndent(v, "", "  ")
	if err != nil {
		return fmt.Errorf("unable to marshal json: %s", err)
	}

	fmt.Println(string(json))

	return nil
}

func toHCL() error {
	input, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		return fmt.Errorf("unable to read from stdin: %s", err)
	}

	ast, err := jsonParser.Parse([]byte(input))
	if err != nil {
		return fmt.Errorf("unable to parse JSON: %s", err)
	}

	err = printer.Fprint(os.Stdout, ast)
	if err != nil {
		return fmt.Errorf("unable to print HCL: %s", err)
	}

	return nil
}
