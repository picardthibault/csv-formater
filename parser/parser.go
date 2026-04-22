package parser

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
)

func ParseCsv(filePath string, header bool) {
	content := readCsvFile(filePath)

	fmt.Print(content)
}

func readCsvFile(filePath string) [][]string {
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
