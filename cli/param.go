package main

import (
	"errors"
	"flag"
	"fmt"
)

type InputArgs struct {
	HeaderParam bool
	InputParam  string
}

func ParseArgs() (InputArgs, error) {
	args := parseCliArgs()
	argsValidationResult := validateArgs(args)

	return args, argsValidationResult
}

func parseCliArgs() InputArgs {
	headerParam := flag.Bool("header", false, "Indicates if the csv file to display has headers")
	inputParam := flag.String("input", "", "The path of the csv file to display")

	flag.Usage = func() {
		fmt.Println("Usage: csv-formater [options]")
		flag.PrintDefaults()
	}

	flag.Parse()

	return InputArgs{*headerParam, *inputParam}
}

func validateArgs(args InputArgs) error {
	inputParamErrors := isInputParamValid(args.InputParam)

	if inputParamErrors != nil {
		return inputParamErrors
	}

	return nil
}

func isInputParamValid(input string) error {
	if input == "" {
		return errors.New("--input option is mandatory")
	}
	return nil
}
