package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
)

func main() {
	csvFileName := flag.String("csv", "problems.csv", "")
	flag.Parse()

	file, err := os.Open(*csvFileName)

	if err != nil {
		fmt.Printf("file can't be opened")
	}

	r := csv.NewReader(file)
	lines, err := r.ReadAll()
	if err != nil {
		println("failed to read file")
	}

	problems := parseLine(lines)

	counter := 0
	for i, problem := range problems {
		var answer string
		fmt.Printf("problem #%d: %s \n", i, problem.q)
		fmt.Scanf("%s \n", &answer)
		if answer == problem.a {
			counter++
		}
	}
	fmt.Printf("total score: %d", counter)
}

type problem struct {
	q string
	a string
}

func parseLine(lines [][]string) []problem {
	problems := make([]problem, len(lines))
	for i, line := range lines {
		problems[i] = problem{
			q: line[0],
			a: line[1],
		}
	}
	return problems
}
