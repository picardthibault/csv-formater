package main

import (
	"fmt"
	"os"

	"github.com/picardthibault/csv-formater/parser"
)

func main() {
	inputs, err := ParseParams()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	csv := parser.ParseCsv(inputs.File, inputs.Headers)
	fmt.Println("header:", csv.Headers)
	fmt.Println("data:", csv.Data)
}
