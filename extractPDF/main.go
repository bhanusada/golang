package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"regexp"
	"sort"
	"strconv"
	"strings"

	"rsc.io/pdf"
)

var stdout *bufio.Writer

type Word struct {
	Position int
	Text     pdf.Text
}

type TextLine struct {
	Words []Word
}

type TimetableEntry struct {
	Day  int
	From string
	To   string
}

func main() {
	//log.SetFlags(0)
	//log.SetPrefix("zeitkonto: ")

	//if len(os.Args) != 2 {
	//fmt.Fprintf(os.Stderr, "Usage: zeitkonto <pdf file>\n")
	//os.Exit(2)
	//}
	//fileName := os.Args[1]

	pdfFile, err := pdf.Open("../template/Interline856spec.pdf")
	if err != nil {
		log.Fatal("Error openining file", err)
	}

	var textLines []*TextLine

	page := pdfFile.Page(1)
	content := page.Content()
	words := findWords(content.Text)

	//add first line
	currentLine := new(TextLine)
	textLines = append(textLines, currentLine)
	//declare position markers
	lastY := 0.0
	wordPos := 0

	//read all words into the textLines slice
	for _, w := range words {
		//add next line if Y coord changes
		if w.Y < lastY && lastY != 0 {
			currentLine = new(TextLine)
			textLines = append(textLines, currentLine)
			wordPos = 0
		}

		//append word to the line and save position  marker
		currentLine.Words = append(currentLine.Words, Word{wordPos, w})
		wordPos++
		lastY = w.Y
	}

	//for _, v := range textLines {
	//	fmt.Println(*v)
	//}

	rDay, _ := regexp.Compile("^([0-9]{2}) (Mo|Di|Mi|Do|Fr|Sa|So)")
	rTime, _ := regexp.Compile("^[0-9]{2}:[0-9]{2}$")

	var timetable []TimetableEntry

	for _, line := range textLines {
		timetableEntry := TimetableEntry{0, "", ""}

		for _, word := range line.Words {

			if word.Position == 0 {
				match := rDay.FindStringSubmatch(word.Text.S)
				if len(match) > 0 {
					day, err := strconv.Atoi(match[1])
					if err == nil {
						timetableEntry.Day = day
					}
				}
			} else if word.Position == 2 {
				timetableEntry.From = word.Text.S
			} else if word.Position == 3 {
				timetableEntry.To = word.Text.S
			}
		}

		if timetableEntry.Day != 0 && rTime.MatchString(timetableEntry.From) && rTime.MatchString(timetableEntry.To) {
			timetable = append(timetable, timetableEntry)
		}
	}

	for _, timetableEntry := range timetable {
		fmt.Fprintf(os.Stdout, "%d: %s-%s\n", timetableEntry.Day, timetableEntry.From, timetableEntry.To)
	}
}

/* from https://github.com/rsc/arm/blob/master/armspec/spec.go */
func findWords(chars []pdf.Text) (words []pdf.Text) {
	// Sort by Y coordinate and normalize.
	const nudge = 1
	sort.Sort(pdf.TextVertical(chars))
	old := -100000.0
	for i, c := range chars {
		if c.Y != old && math.Abs(old-c.Y) < nudge {
			chars[i].Y = old
		} else {
			old = c.Y
		}
	}

	// Sort by Y coordinate, breaking ties with X.
	// This will bring letters in a single word together.
	sort.Sort(pdf.TextVertical(chars))

	// Loop over chars.
	for i := 0; i < len(chars); {
		// Find all chars on line.
		j := i + 1
		for j < len(chars) && chars[j].Y == chars[i].Y {
			j++
		}
		var end float64
		// Split line into words (really, phrases).
		for k := i; k < j; {
			ck := &chars[k]
			s := ck.S
			end = ck.X + ck.W
			charSpace := ck.FontSize / 6
			wordSpace := ck.FontSize * 2 / 3
			l := k + 1
			for l < j {
				// Grow word.
				cl := &chars[l]
				if sameFont(cl.Font, ck.Font) && math.Abs(cl.FontSize-ck.FontSize) < 0.1 && cl.X <= end+charSpace {
					s += cl.S
					end = cl.X + cl.W
					l++
					continue
				}
				// Add space to phrase before next word.
				if sameFont(cl.Font, ck.Font) && math.Abs(cl.FontSize-ck.FontSize) < 0.1 && cl.X <= end+wordSpace {
					s += " " + cl.S
					end = cl.X + cl.W
					l++
					continue
				}
				break
			}
			f := ck.Font
			f = strings.TrimSuffix(f, ",Italic")
			f = strings.TrimSuffix(f, "-Italic")
			words = append(words, pdf.Text{f, ck.FontSize, ck.X, ck.Y, end - ck.X, s})
			k = l
		}
		i = j
	}
	fmt.Println(words)
	return words
}

func sameFont(f1, f2 string) bool {
	f1 = strings.TrimSuffix(f1, ",Italic")
	f1 = strings.TrimSuffix(f1, "-Italic")
	f2 = strings.TrimSuffix(f1, ",Italic")
	f2 = strings.TrimSuffix(f1, "-Italic")
	return strings.TrimSuffix(f1, ",Italic") == strings.TrimSuffix(f2, ",Italic") || f1 == "Symbol" || f2 == "Symbol" || f1 == "TimesNewRoman" || f2 == "TimesNewRoman"
}
