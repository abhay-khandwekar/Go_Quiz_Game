package main

import (
	"bytes"
	"strings"
	"testing"
)

func TestShuffleProblems(t *testing.T) {
	problems := []problem{
		problem{
			q: "q1",
			a: "a1",
		},
		problem{
			q: "q2",
			a: "a2",
		},
		problem{
			q: "q3",
			a: "a3",
		},
		problem{
			q: "q4",
			a: "a4",
		},
	}

	shuffleProblems(problems)

	if problems[0].q == "q1" {
		t.Errorf("Problem shuffle might have failed. Expecting problem[0].q not to be 'q1', but got 'q1' instead.")
	}
}

func TestScanAnswer(t *testing.T) {
	var stdin bytes.Buffer
	stdin.Write([]byte("HI\n"))
	answerCh := make(chan string)

	go scanAnswer(&stdin, answerCh)
	answer := <-answerCh

	if answer != strings.ToLower("HI\n") {
		t.Errorf("Expecting value 10, but got %v", answer)
	}
}

func TestParseLines(t *testing.T) {
	data := [][]string{
		{"1+2", "3"},
		{"3+5", "8"},
		{"4+1", "5"},
		{"4+3", "7"},
	}

	problems := parseLines(data)

	if problems == nil {
		t.Error("Unable to parse problem data of the form [][]string.")
	}

	if len(problems) != 4 {
		t.Errorf("Expecting problem count of 4, but got %d", len(problems))
	}
}

func TestGetProblemsFromCSVFile(t *testing.T) {
	fileName := "test.csv"
	problems, err := getProblemsFromCSVFile(&fileName)

	if err != nil {
		t.Errorf("Unable to read 'test.csv' file. Error: %v", err)
	}

	if len(problems) != 4 {
		t.Errorf("Expecting problem count of 4, but got %d", len(problems))
	}
}
