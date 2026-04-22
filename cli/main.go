package main

import (
	"fmt"
	"os"

	"github.com/picardthibault/csv-formater/parser"
)

func main() {
	inputArgs, inputValidationError := ParseArgs()

	if inputValidationError != nil {
		fmt.Println(inputValidationError)
		os.Exit(1)
	}

	fmt.Println("inputParam:", inputArgs.InputParam)
	fmt.Println("headerParam:", inputArgs.HeaderParam)

	parser.ParseCsv(inputArgs.InputParam, inputArgs.HeaderParam)
}
