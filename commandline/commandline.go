package main

import (
	"fmt"
	"os"
)

func main() {
	arguments := os.Args
	fmt.Print(arguments[1:])
}
