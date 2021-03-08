package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
)

func main() {
	csvFilename := flag.String("csv", "problems.csv", "a csv file in the format of 'question, answer")
	flag.Parse()

	file, err := os.Open(*csvFilename)
	if err != nil {
		fmt.Println("failed to Open csv file: %s \n", *csvFilename)
		os.Exit(1)
	}
	fmt.Println("success Open csv file: %s \n", *csvFilename)
	_ = file

	r := csv.NewReader(file)
	lines, err := r.ReadAll()
	if err != nil {
		fmt.Println("Failed to parse the provided CSV file")
	}
	fmt.Println(lines)
}
