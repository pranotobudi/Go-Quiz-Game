package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"time"
)

type qa2d struct {
	q string
	a string
}

func main() {
	fileNamePtr := flag.String("csv", "problems.csv", "input your file name")
	timeLimit := flag.Int("time", 5, "input timer")
	flag.Parse()

	file, err := os.Open(*fileNamePtr)
	if err != nil {
		fmt.Println("failed to open file")
	}

	r := csv.NewReader(file)
	lines, err := r.ReadAll()
	if err != nil {
		fmt.Println("failed to open file")
	}
	problems := parseLine2d(lines)
	counter := 0

	timer := time.NewTimer(time.Second * time.Duration(*timeLimit))

	for i, problem := range problems {
		fmt.Printf("question #%d %s : ", i+1, problem.q)
		answerCh := make(chan string)
		go func() {
			var answer string
			n, err := fmt.Scanf("%s\n", &answer)
			if err != nil || n != 1 {
				// handle invalid input
				fmt.Println(n, err)
				return
			}
			answerCh <- answer
		}()
		select {
		case <-timer.C:
			fmt.Println("your score is: ", counter, "out of: ", len(problems))
			return
		case ans := <-answerCh:
			if ans == problem.a {
				counter++
			}
		}
	}
	fmt.Println(problems)
}

func parseLine2d(lines [][]string) []qa2d {
	problems := []qa2d{}

	for _, line := range lines {
		// fmt.Println(line)
		problem := qa2d{}
		problem.q = line[0]
		problem.a = line[1]
		problems = append(problems, problem)
	}

	return problems
}
