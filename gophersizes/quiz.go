package quiz

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
)

// Question - Quiz question
type Question struct {
	question string
	answer   string
}

// PrintSomething - Stub function to print something
func PrintSomething() string {
	csvFile, _ := os.Open("quiz.csv")

	reader := csv.NewReader(bufio.NewReader(csvFile))

	for {
		line, err := reader.Read()

		if err == io.EOF {
			break
		} else if err != nil {
			log.Fatal(err)
		}

		fmt.Print(line)
	}

	// Generate Questions/Answers
	// Iterate over questions, and prompt user to answer
	// Conditionally print correct answer or success message based on correction
	// Store successes and failures
	// Display Score w/ percentage
	return "Works"
}
