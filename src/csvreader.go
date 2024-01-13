package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
)

// ReadCSV reads a CSV file from the given file path and processes its contents.
func ReadCSV(filePath string) {
	file, err := os.Open(filePath)

	if err != nil {
		fmt.Printf("Failed to open %s", filePath)
		return
	}
	defer file.Close()

	reader := csv.NewReader(file)

	for {
		_, err := reader.Read()

		if err == io.EOF {
			break
		}

		if err != nil {
			fmt.Printf("There was an error reading a line: %s", err.Error())
			continue
		}
	}
}
