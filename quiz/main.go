package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

var allQuestions []question
var correctAnswersCounter int
var numberOfQuestions int
var timePerQuestion = 5

type question struct {
	text   string
	answer string
}

func readFile(filename string) {
	f, err := os.Open(filename)
	if err != nil {
		log.Fatalf("error during opening file: %d", err)
	}

	records, err := csv.NewReader(f).ReadAll()
	if err != nil {
		if err != io.EOF {
			log.Fatalf("error during reading file: %d", err)
		}
	}

	for _, r := range records {
		allQuestions = append(allQuestions, *createQuestion(&r[0]))
	}
}

func createQuestion(s *string) *question {
	spitted := strings.Split(*s, "=")
	qs, as := spitted[0], strings.ToLower(spitted[1])
	q := question{text: qs, answer: as}

	return &q
}

func getNumberOfQuestions() *int {
	fmt.Println("Please choose number of questions")
	var userInput string
	fmt.Scanln(&userInput)

	converted, err := strconv.Atoi(userInput)
	if err != nil {
		log.Fatalf("error during converting string to int: %d", err)
	}

	if !hasEnoughQuestions(&converted) {
		fmt.Printf("Number is too big. Max value is %d \n", len(allQuestions))
		return getNumberOfQuestions()
	}

	return &converted
}

func hasEnoughQuestions(userInput *int) bool {
	if *userInput > len(allQuestions) {
		return false
	}
	return true
}

func startQuiz(done chan bool) {
	for i := 0; i < numberOfQuestions; i++ {
		printQuestion(&i)
		var answer string
		fmt.Scanln(&answer)
		if isCorrectAnswer(allQuestions[i].answer, &answer) {
			correctAnswersCounter++
		}
	}
	done <- true
}

func countTime(done chan bool, numberOfQuestions *int, timePerQuestion *int) {
	ticker := time.NewTicker(time.Second * time.Duration(*numberOfQuestions) * time.Duration(*timePerQuestion))
	go func(c <-chan time.Time) {
		<-c
		gameOver()
		done <- false
	}(ticker.C)
}

func printQuestion(id *int) {
	actualQuestion := allQuestions[*id]
	fmt.Printf("Question %d \n%s = ? \n", *id+1, actualQuestion.text)
}

func isCorrectAnswer(correctAnswer string, userInput *string) bool {
	if correctAnswer == strings.ToLower(*userInput) {
		return true
	}
	return false
}

func showScore() {
	fmt.Println("Your score is", correctAnswersCounter)
}

func gameOver() {
	fmt.Println("Time's end")
}

func main() {
	readFile("quiz/source.csv")
	numberOfQuestions = *getNumberOfQuestions()

	done := make(chan bool)
	go startQuiz(done)
	go countTime(done, &numberOfQuestions, &timePerQuestion)

	<-done
	showScore()
}
