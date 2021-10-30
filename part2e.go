package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"time"
)

func main() {
	fileNamePtr := flag.String("csv", "problems.csv", "input your file name")
	timeLimit := flag.Int("limit", 10, "input your time limit")
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

	problems := parsingLine2e(lines)
	counter := 0

	timer := time.NewTimer(time.Duration(*timeLimit) * time.Second)

	for i, problem := range problems {
		fmt.Printf("question #%d: %s \n", i+1, problem.q)
		var answer string
		answerCh := make(chan string)
		go func() {
			fmt.Scanf("%s \n", &answer)
			answerCh <- answer
		}()

		select {
		case <-timer.C:
			fmt.Println("your score is: ", counter, "out of: ", len(problems))
			return
		case <-answerCh:
			if answer == problem.a {
				fmt.Println("correct!")
				counter++
			}
		}
	}
	fmt.Println("your score is: ", counter, "out of: ", len(problems))
	fmt.Println(problems)
}

type qa struct {
	q string
	a string
}

func parsingLine2e(lines [][]string) []qa {
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
