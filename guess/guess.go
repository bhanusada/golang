package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

var randomNumber int64

func main() {
	generateNumber()
	input()
}

func generateNumber() {
	seconds := time.Now().Unix()
	rand.Seed(seconds)
	randomNumber = rand.Int63n(101)
}

func input() {
	fmt.Println("Guess a number between 1 and 100.")
	userInput := bufio.NewReader(os.Stdin)
	for t := 3; t > 0; t-- {
		fmt.Print("You have ", t, " guesses left!!! Enter a number : ")
		readString, error := userInput.ReadString('\n')
		if error != nil {
			log.Fatal(error)
		}
		convStringToNumber, error := strconv.ParseInt(strings.TrimSpace(readString), 10, 64)
		if error != nil {
			log.Fatal(error)
		}
		if compare(&convStringToNumber) {
			break
		}
	}
	fmt.Println("Random number is ", randomNumber)
}

func compare(givenNumber *int64) bool {
	var result bool
	if *givenNumber > randomNumber {
		fmt.Println("Oops! Your guess was HIGH")
		result = false
	} else if *givenNumber < randomNumber {
		fmt.Println("Oops! Your guess was LOW")
		result = false
	} else {
		fmt.Println("Good job!")
		result = true
	}
	return result
}
