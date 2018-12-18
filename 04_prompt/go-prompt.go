package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	getTime()
	fmt.Print("Enter a grade: ")
	reader := bufio.NewReader(os.Stdin)
	// readString, _ := reader.ReadString('\n') // blank identifier to ignore error
	userInput, err := reader.ReadString('\n')

	handleError(err)

	grade, error := strconv.ParseFloat(strings.TrimSpace(userInput), 64)

	handleError(error)

	if grade > 60 {
		fmt.Println("Pass")
	} else {
		fmt.Println("Fail")
	}
}

func getTime() {
	// now := time.Now()
	fmt.Println(time.Now())
}

func handleError(error) {
	if error != nil {
		log.Fatal(error)
	}
}
