package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"time"
)

func main() {
	csvFileName := flag.String("csv", "problems.csv", "")
	timeOut := flag.Int("time", 10, "")
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
	timer := time.NewTimer(time.Second * time.Duration(*timeOut))
	for i, problem := range problems {
		fmt.Printf("problem #%d: %s \n", i, problem.q)
		fmt.Printf("iteration: #%d time: %d \n", i, time.Now().Second())
		var ansChannel = make(chan string)
		go func() {
			var answer string
			fmt.Scanf("%s \n", &answer)
			ansChannel <- answer
		}()
		select {
		case ans := <-ansChannel:
			fmt.Printf("in select>case ans, i: %d \n", i)
			if ans == problem.a {
				counter++
			}
		case <-timer.C:
			fmt.Printf("in select>case timer.C, i: %d \n", i)
			fmt.Printf("total score: %d", counter)
			return
		}
	}
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
