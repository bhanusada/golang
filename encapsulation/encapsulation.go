package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/bhanu/encapsulation/calendar"
)

func main() {
	event := calendar.Event{}
	input := bufio.NewReader(os.Stdin)

	fmt.Print("Enter event title : ")
	titleInput, err := input.ReadString('\n')
	handle(err)
	title := strings.TrimSpace(titleInput)
	event.SetTitle(title)

	fmt.Print("Enter year : ")
	yearInput, err := input.ReadString('\n')
	handle(err)
	year, err := strconv.Atoi(strings.TrimSpace(yearInput))
	handle(err)
	err = event.SetYear(year)
	handle(err)

	fmt.Print("Enter month : ")
	monthInput, err := input.ReadString('\n')
	handle(err)
	month, err := strconv.Atoi(strings.TrimSpace(monthInput))
	handle(err)
	err = event.SetMonth(month)
	handle(err)

	fmt.Print("Enter day : ")
	dayInput, err := input.ReadString('\n')
	handle(err)
	day, err := strconv.Atoi(strings.TrimSpace(dayInput))
	handle(err)
	err = event.SetDay(day)
	handle(err)

	printDate(event)
}

func handle(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func printDate(event calendar.Event) {
	fmt.Printf("Title : %s\n", event.GetTitle())
	fmt.Printf("Entered year %d\n", event.GetYear())
	fmt.Printf("Entered month %d\n", event.GetMonth())
	fmt.Printf("Entered day %d\n", event.GetDay())
}
