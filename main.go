package main

import (
  "os"
  "io/ioutil"
  "github.com/hashicorp/hcl/json/parser"
  "github.com/hashicorp/hcl/hcl/printer"
)

func main() {
  input, _ := ioutil.ReadAll(os.Stdin)

  ast, err := parser.Parse([]byte(input))
  if err != nil {
    panic(err)
  }

  err = printer.Fprint(os.Stdout, ast)
  if err != nil {
    panic(err)
  }
}
