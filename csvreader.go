package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
)

func ReadCSV(filePath string) {
	file, err := os.Open(filePath)

	if err != nil {
		fmt.Printf("Failed to open %s", filePath)
		return
	}
	defer file.Close()

	reader := csv.NewReader(file)

	for {
		record, err := reader.Read()

		if err == io.EOF {
			break
		}

		if err != nil {
			fmt.Printf("There was an error reading a line: %s", err.Error())
			continue
		}

		fmt.Println("row: ", record)
	}

	GeneratePuzzleSQL()
}
