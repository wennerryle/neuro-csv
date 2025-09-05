package main

import (
	"fmt"
	"neuro-csv/parser"
)

func main() {
	flags := parser.ParseFlags()

	fmt.Printf("%+v\n", flags)
}
