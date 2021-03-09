package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"time"
)

func main() {
	//read command line
	csvFilename := flag.String("csv", "problems.csv", "")
	timeLimit := flag.Int("time_limit", 3, "")
	flag.Parse()

	//parse file
	file, err := os.Open(*csvFilename)
	if err != nil {
		println("file can't be opened")
	}

	r := csv.NewReader(file)
	lines, err := r.ReadAll()
	if err != nil {
		println("failed to read file")
	}

	//show question
	problems := parseLines(lines)
	//receive answer
	//show correct Number
	fmt.Println(problems)
	var counter = 0
	timer := time.NewTimer(time.Second * time.Duration(*timeLimit))

	for i, problem := range problems {
		answerCh := make(chan string)

		go func() {
			var answer string
			fmt.Printf("problem #%d: %s \n", i, problem.q)
			fmt.Scanf("%s \n", &answer)
			answerCh <- answer
		}()

		select {
		case <-timer.C:
			fmt.Printf("total score: %d \n", counter)
			return
		case answer := <-answerCh:
			if answer == problem.a {
				counter++
			}
		}
	}
}

type problem struct {
	q string
	a string
}

func parseLines(lines [][]string) []problem {
	problems := make([]problem, len(lines))

	for i, line := range lines {
		problems[i] = problem{
			q: line[0],
			a: line[1],
		}
	}

	return problems

}
