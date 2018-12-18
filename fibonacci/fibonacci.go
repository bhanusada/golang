package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	number, _ := strconv.Atoi(os.Args[1])
	generate(number)
}

func generate(i int) {
	curr, prev := 1, 0
	fmt.Println(prev)
	for j := 0; j < i; j++ {
		fib := curr + prev
		prev = curr
		curr = fib
		fmt.Println(prev)
	}
}
