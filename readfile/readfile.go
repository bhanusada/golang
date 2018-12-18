package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	myslice := []string{}
	file, err := os.Open("data.txt")
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Println("len", len(myslice))
		myslice = append(myslice, scanner.Text())
	}
	err = file.Close()
	if err != nil {
		log.Fatal(err)
	}
	if scanner.Err() != nil {
		log.Fatal(scanner.Err())
	}
	fmt.Println(myslice)
}
