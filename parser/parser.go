package parser

import (
	"encoding/csv"
	"log"
	"os"
)

type CSV struct {
	Headers []string
	Data    [][]string
}

func ParseCsv(filePath string, header bool) CSV {
	content := readCsv(filePath)

	if header {
		return CSV{content[0], content[1:]}
	} else {
		return CSV{[]string{}, content}
	}
}

func readCsv(filePath string) [][]string {
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	return records
}
