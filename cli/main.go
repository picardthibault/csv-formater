package main

import (
	"errors"
	"flag"
	"fmt"
	"os"
)

func main() {
	hasHeader := flag.Bool("header", false, "Indicates if the csv file to display has headers")
	input := flag.String("input", "", "The path of the csv file to display")

	flag.Usage = func() {
		fmt.Println("Usage: csv-formater [options]")
		flag.PrintDefaults()
	}

	flag.Parse()

	inputValidationError := isInputArgValid(*input)
	if inputValidationError != nil {
		fmt.Println(inputValidationError)
		os.Exit(1)
	}

	fmt.Println("input:", *input)
	fmt.Println("hasHeader:", *hasHeader)
}

func isInputArgValid(input string) error {
	if input == "" {
		return errors.New("--input option is mandatory")
	}
	return nil
}
