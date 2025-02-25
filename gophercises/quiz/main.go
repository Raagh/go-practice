package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

func main() {
	csvFilename := flag.String("csv", "problems.csv", "a csv file in the format of 'question,answer'")
	timeLimit := flag.Int("timeLimit", 30, "time limit to answer all questions")
	randomOrder := flag.Bool("randomOrder", false, "show questions in random order")
	flag.Parse()

	file, err := os.Open(*csvFilename)
	if err != nil {
		fmt.Printf("It was no possible to open the file: %s\n", *csvFilename)
		os.Exit(1)
	}

	reader := csv.NewReader(file)
	lines, err := reader.ReadAll()
	if err != nil {
		fmt.Println("There was a problem reading the csv")
		os.Exit(1)
	}

	if *randomOrder {
		rand.Shuffle(len(lines), func(i, j int) {
			lines[i], lines[j] = lines[j], lines[i]
		})
	}

	problems := parseProblems(lines)

	fmt.Printf("Press enter when you are ready to start")
	var ready string
	fmt.Scanf("%s\n", &ready)

	correctProblems := 0

	// MY SOLUTION

	// stopper := time.AfterFunc(time.Second*time.Duration(*timeLimit), func() {
	// 	fmt.Printf("Time is UP! You scored %d out of %d\n", correctProblems, len(problems))
	// 	os.Exit(0)
	// })
	//
	// for i, p := range problems {
	// 	fmt.Printf("Problem #%d: %s\n", i+1, p.question)
	// 	var answer string
	// 	fmt.Scanf("%s\n", &answer)
	//
	// 	if answer == p.answer {
	// 		correctProblems++
	// 	}
	// }
	//
	// stopper.Stop()
	// fmt.Printf("You scored %d out of %d\n", correctProblems, len(problems))

	// COURSE SOLUTION

	timer := time.NewTimer(time.Duration(*timeLimit) * time.Second)
	for i, p := range problems {
		fmt.Printf("Problem #%d: %s\n", i+1, p.question)
		answerCh := make(chan string)
		go func() {
			var answer string
			fmt.Scanf("%s\n", &answer)
			answerCh <- answer
		}()

		select {
		case <-timer.C:
			fmt.Printf("Time is UP! You scored %d out of %d\n", correctProblems, len(problems))
			return
		case answer := <-answerCh:
			if answer == p.answer {
				correctProblems++
			}
		}
	}

	fmt.Printf("You scored %d out of %d\n", correctProblems, len(problems))
}

func parseProblems(lines [][]string) []problem {
	problems := make([]problem, len(lines))
	for i, v := range lines {
		problems[i] = problem{
			question: v[0],
			answer:   strings.TrimSpace(v[1]),
		}
	}

	return problems
}

type problem struct {
	question string
	answer   string
}
