package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	input := []int{23, 43, 6}
	reverse(input)
	reverseString(os.Args[1])
}

func reverse(s []int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
	fmt.Println(s)
}

func reverseString(s string) {
	splitInput := strings.Split(s, "")

	var tempString string
	for i := len(splitInput); i > 0; i-- {
		tempString = tempString + strings.Join(splitInput[i-1:i], "")
	}

	for i, j := 0, len(splitInput)-1; i < j; i, j = i+1, j-1 {
		splitInput[i], splitInput[j] = splitInput[j], splitInput[i]
	}
	fmt.Println(tempString)
	fmt.Println(strings.Join(splitInput, ""))
}
