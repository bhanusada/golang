package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
)

func main() {
	mymap := map[string]int{}
	lines := readFile()
	for _, line := range lines {
		mymap[line]++
	}
	keysort := []string{}
	for key := range mymap {
		keysort = append(keysort, key)
	}
	sort.Strings(keysort)
	for _, val := range keysort {
		fmt.Println(val, mymap[val])
	}
}

func readFile() []string {
	myslice := []string{}
	file, err := os.Open("data.txt")
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		myslice = append(myslice, scanner.Text())
	}
	return myslice
}
