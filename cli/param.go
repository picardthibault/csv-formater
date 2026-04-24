package main

import (
	"errors"
	"flag"
	"fmt"
)

type Flags struct {
	Headers bool
	File    string
}

func ParseParams() (Flags, error) {
	args := parseFlags()
	err := isFlagValid(args)

	return args, err
}

func parseFlags() Flags {
	headerFlag := flag.Bool("header", false, "Indicates if the csv file to display has headers")
	fileFlag := flag.String("file", "", "The path of the csv file to display")

	flag.Usage = func() {
		fmt.Println("Usage: csv-formater [options]")
		flag.PrintDefaults()
	}

	flag.Parse()

	return Flags{*headerFlag, *fileFlag}
}

func isFlagValid(args Flags) error {
	err := isFileFlagValid(args.File)
	if err != nil {
		return err
	}

	return nil
}

func isFileFlagValid(file string) error {
	if file == "" {
		return errors.New("--file option is mandatory !")
	}
	return nil
}
