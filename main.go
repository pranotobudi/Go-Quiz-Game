package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strings"
)

func main() {
	csvFilename := flag.String("csv", "problems.csv", "a csv file in the format of 'question, answer")
	flag.Parse()

	file, err := os.Open(*csvFilename)
	if err != nil {
		exit(fmt.Sprintf("failed to Open csv file: %s \n", *csvFilename))
	}
	fmt.Println(fmt.Sprintf("success Open csv file: %s \n", *csvFilename))
	_ = file

	r := csv.NewReader(file)
	lines, err := r.ReadAll()
	if err != nil {
		exit("Failed to parse the provided CSV file")
	}
	fmt.Println(lines)
	problems := parseLine(lines)
	totalScore := 0
	for i, problem := range problems {
		fmt.Printf("problem #%d: %s \n", i+1, problem.q)
		var answer string
		fmt.Scanf("%s \n", &answer)
		if answer == problem.a {
			fmt.Println("true")
			totalScore++
		} else {
			fmt.Println("false")
		}
	}

	fmt.Printf("your total score: %d", totalScore)
}

func parseLine(lines [][]string) []problem {
	var problemArray = make([]problem, len(lines))
	for i, line := range lines {
		problemArray[i] = problem{
			q: line[0],
			a: strings.TrimSpace(line[1]),
		}
	}
	return problemArray
}

type problem struct {
	q string
	a string
}

func exit(msg string) {
	fmt.Println(msg)
	os.Exit(1)
}
