package main

import (
	"errors"
	"flag"
	"fmt"
)

type inputArgs struct {
	HeaderParam bool
	InputParam  string
}

func ParseArgs() (inputArgs, error) {
	args := parseCliArgs()
	argsValidationResult := validateArgs(args)

	return args, argsValidationResult
}

func parseCliArgs() inputArgs {
	headerParam := flag.Bool("header", false, "Indicates if the csv file to display has headers")
	inputParam := flag.String("input", "", "The path of the csv file to display")

	flag.Usage = func() {
		fmt.Println("Usage: csv-formater [options]")
		flag.PrintDefaults()
	}

	flag.Parse()

	return inputArgs{*headerParam, *inputParam}
}

func validateArgs(args inputArgs) error {
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
