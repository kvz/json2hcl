package main

import (
  "os"
  "io/ioutil"
  "github.com/Acconut/hcl/json/parser"
  "github.com/Acconut/hcl/hcl/printer"
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
