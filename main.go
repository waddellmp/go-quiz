package main

import (
	"bufio"
	"encoding/csv"
	"flag"
	"fmt"
	"os"
)

var defaultFilename = "./problems.csv"

func main() {
	var correct int
	// Define flag for csv filename
	iFlag := flag.String("i", defaultFilename, "Input filename")
	flag.Parse()

	// Read the file stdin in readonly mode
	file, err := os.Open(*iFlag)
	if err != nil {
		fmt.Printf("No file found at %v\n", *iFlag)
		os.Exit(1)
	}
	// Close the file
	defer file.Close()

	if err != nil {
		fmt.Println("Error opening file")
		os.Exit(1)
	}

	// Read the file contents
	reader := csv.NewReader(file)

	records, err := reader.ReadAll()
	if err != nil {
		fmt.Println("Error reading file")
		os.Exit(1)
	}

	for index, record := range records {
		// Prompt the question
		fmt.Printf("Question #%v: %v\n", index, record[0])

		// ReadString() will block until the delimiter is entered
		reader := bufio.NewReader(os.Stdin)
		_, err := reader.ReadString('\n')
		// Trim out the newlines
		if err != nil {
			os.Exit(1)
		}

		correct++
	}
	fmt.Printf("Quiz complete. You answered %v out of %v correctly.", correct, len(records))
}
