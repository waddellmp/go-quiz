package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	// Define flag for csv filename
	csvFilename := flag.String("csv", "problems.csv", "Input filename in the format of question,answer")

	// Define flag for time limit in minutes
	timeLimit := flag.Int("limit", 30, "Time limit for the quiz in seconds")

	// Parse the flags
	flag.Parse()

	// Read the fi stdin in readonly mode
	fi, err := os.Open(*csvFilename)

	// Check for open errors
	if err != nil {
		exit(fmt.Sprintf("Failed to open the CSV file %s", *csvFilename))
	}
	// Close the file
	defer fi.Close()

	// Make csv reader
	reader := csv.NewReader(fi)

	// Read all the file contents
	lines, err := reader.ReadAll()

	// Check for read errors
	if err != nil {
		exit(fmt.Sprintf("Failed to read the provided CSV file %s", *csvFilename))
	}

	// Map csv lines to a struct
	problems := makeProblems(lines)

	// Start the timer after the lines are parsed
	timer := time.NewTimer(time.Duration(*timeLimit) * time.Second)

	var correctCount int
	// Iterate over the problems and prompt the user for input
problemLoop:
	for i, p := range problems {
		// Print question
		fmt.Printf("Problem #%d: %s\n", i+1, p.question)

		// Read channel for our answer
		answerCh := make(chan string)

		go func() {
			// Create string pointer and pass into Scanf
			var answer string
			fmt.Scanf("%s\n", &answer)

			// Sending answer var through channel
			answerCh <- answer
		}()
		select {
		case <-timer.C:
			fmt.Println()
			break problemLoop
		case answer := <-answerCh:
			if answer == p.answer {
				correctCount++

			}
		}
	}
	fmt.Printf("Quiz is complete. Your got %d out of %d correct.\n", correctCount, len(problems))
}

type problem struct {
	question string
	answer   string
}

func makeProblems(lines [][]string) []problem {
	parsed := make([]problem, len(lines))
	for i, line := range lines {
		parsed[i] = problem{
			question: line[0],
			answer:   strings.TrimSpace(line[1]),
		}
	}
	return parsed
}

func exit(msg string) {
	fmt.Println(msg)
	os.Exit(1)
}
