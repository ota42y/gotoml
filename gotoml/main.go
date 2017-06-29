package main

import (
	"flag"
	"fmt"
	"io"
	"os"

	"github.com/ota42y/gotoml"
)

func main() {
	var name = flag.String("name", "Foo", "the name of the struct")
	var pkg = flag.String("pkg", "main", "the name of the package for the generated code")

	flag.Parse()

	var input io.Reader = os.Stdin
	output, err := gotoml.Generate(input, *name, *pkg)
	if err != nil {
		fmt.Fprintln(os.Stderr, "error parsing", err)
		os.Exit(1)
	} else {
		fmt.Print(string(output))
	}
}
