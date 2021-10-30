package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
)

func main() {
	fileNamePtr := flag.String("csv", "problems.csv", "input your file name")
	flag.Parse()
	file, err := os.Open(*fileNamePtr)
	if err != nil {
		fmt.Println("failed to open file CSV file: %s", *fileNamePtr)
		os.Exit(1)
	}

	r := csv.NewReader(file)
	lines, err := r.ReadAll()
	if err != nil {
		fmt.Println("failed to read file")
	}

	problems := parsingLine1d(lines)
	counter := 0
	for i, problem := range problems {
		fmt.Printf("question #%d: %s \n", i+1, problem.q)
		var answer string
		fmt.Scanf("%s \n", &answer)
		if answer == problem.a {
			fmt.Println("correct!")
			counter++
		}
	}
	fmt.Println("your score is: ", counter, "out of: ", len(problems))
	fmt.Println(problems)
}

type qa struct {
	q string
	a string
}

func parsingLine1d(lines [][]string) []qa {
	problems := []qa{}

	for _, line := range lines {
		// fmt.Println(line)
		problem := qa{}
		problem.q = line[0]
		problem.a = line[1]
		problems = append(problems, problem)
	}

	return problems
}
