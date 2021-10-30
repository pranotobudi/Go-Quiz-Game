package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
)

func main() {
	fileName := flag.String("csv", "problems.csv", "")
	flag.Parse()

	file, err := os.Open(*fileName)
	if err != nil {
		fmt.Println("file not found")
	}
	reader := csv.NewReader(file)

	lines, err := reader.ReadAll()
	if err != nil {
		fmt.Println("can't read the file")
	}

	problems := parseLine(lines)
	counter := 0
	for i, problem := range problems {
		var answer string
		fmt.Printf("problem #%d: %s \n", i+1, problem.q)
		fmt.Scanf("%s \n", &answer)
		if answer == problem.a {
			counter++
		}
	}

	fmt.Println("Your score: %d ", counter)

}

type problem struct {
	q string
	a string
}

func parseLine(lines [][]string) []problem {
	var problems = make([]problem, len(lines))

	for i, line := range lines {
		problems[i] = problem{
			q: line[0],
			a: line[1],
		}
	}
	return problems
}
