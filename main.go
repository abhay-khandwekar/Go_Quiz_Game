package main

import (
	"bufio"
	"encoding/csv"
	"flag"
	"fmt"
	"io"
	"math/rand"
	"os"
	"strings"
	"time"
)

type problem struct {
	q string
	a string
}

func main() {
	csvFileName := flag.String("csv", "problems.csv", "A CSV file in the format of 'question,answer'")
	duration := flag.Int("duration", 30, "Quiz duration in seconds")
	shuffle := flag.Bool("shuffle", true, "Boolen flag to shuffle problems")
	flag.Parse()

	problems, err := getProblemsFromCSVFile(csvFileName)
	if err != nil {
		exit(fmt.Sprintf("Unable to read CSV file. Error: %s\n", err))
	}

	if *shuffle {
		shuffleProblems(problems)
	}

	fmt.Println("Press 'Enter key' to start the quiz...")
	var startQuiz string
	fmt.Scanln(&startQuiz)

	qzt := time.NewTimer(time.Second * time.Duration(*duration))
	answerCh := make(chan string)

	correct := 0
quizLoopout:
	for i, p := range problems {
		fmt.Printf("Problem #%d: %s = ", i+1, p.q)
		go scanAnswer(os.Stdin, answerCh)

		select {
		case <-qzt.C:
			fmt.Printf("\nQuiz time is over.\n")
			break quizLoopout
		case answer := <-answerCh:
			if answer == p.a {
				correct++
			}
		}
	}

	fmt.Printf("You scored %d out of %d.\n", correct, len(problems))
}

func shuffleProblems(p []problem) {
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(p), func(i, j int) { p[i], p[j] = p[j], p[i] })
}

func scanAnswer(reader io.Reader, answerCh chan<- string) {
	rdr := bufio.NewReader(reader)
	answer, _ := rdr.ReadString('\n')
	answerCh <- strings.ToLower(answer)
}

func getProblemsFromCSVFile(csvFileName *string) ([]problem, error) {
	file, err := os.Open(*csvFileName)
	if err != nil {
		return nil, err
	}

	r := csv.NewReader(file)
	lines, err := r.ReadAll()
	if err != nil {
		return nil, err
	}

	return parseLines(lines), nil
}

func parseLines(lines [][]string) []problem {
	p := make([]problem, len(lines))

	for i, l := range lines {
		p[i] = problem{
			q: l[0],
			a: strings.ToLower(strings.TrimSpace(l[1])),
		}
	}

	return p
}

func exit(msg string) {
	fmt.Println(msg)
	os.Exit(1)
}
