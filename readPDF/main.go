package main

import (
	"fmt"
	"math"
	"strings"

	"rsc.io/pdf"
)

func main() {
	r, _ := pdf.Open("../template/Interline856spec.pdf")
	p := r.Page(8)
	//fmt.Println(r.NumPage())
	//fmt.Printf("%v", p.Content())
	//fmt.Println(r.Trailer().Key("Root").String())
	//fmt.Println(p.V.TextFromUTF16())

	for _, v := range p.Content().Text {
		fmt.Printf("%#v", v)
	}
	words := findWords(p.Content().Text)

	for _, v := range words {
		fmt.Println(v.S)
	}
}
{Font:"Arial-BoldItalicMT", FontSize:8.9858405, X:453.10301349950015, Y:354.614715, W:7.9884122045, S:"m"}
{Font:"Arial-BoldItalicMT", FontSize:8.9858405, X:463.81413537550014, Y:354.614715, W:4.996127318, S:"a"}
{Font:"Arial-BoldItalicMT", FontSize:8.9858405, X:468.81026269350014, Y:354.614715, W:2.9922848865000002, S:"t"}
{Font:"Arial-BoldItalicMT", FontSize:8.9858405, X:474.5252572515002, Y:354.614715, W:4.996127318, S:"8"}


func findWords(chars []pdf.Text) (words []pdf.Text) {
	// Sort by Y coordinate and normalize.
	/*const nudge = 1
	sort.Sort(pdf.TextVertical(chars))
	old := -100000.0
	for i, c := range chars {
		if c.Y != old && math.Abs(old-c.Y) < nudge {
			chars[i].Y = old
		} else {
			old = c.Y
		}
	}*/

	// Sort by Y coordinate, breaking ties with X.
	// This will bring letters in a single word together.
	//sort.Sort(pdf.TextVertical(chars))
	//fmt.Println(chars)
	// Loop over chars.
	for i := 0; i < len(chars); {
		fmt.Printf("%#v\n", chars[i])
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
			end = ck.X + ck.W 500
			charSpace := ck.FontSize / 6 1.5
			wordSpace := ck.FontSize * 2 / 3 6
			l := k + 1
			for l < j {
				// Grow word.
				cl := &chars[l]
				if cl.Font == ck.Font && math.Abs(cl.FontSize-ck.FontSize) < 0.1 && cl.X <= end+charSpace {
					s += cl.S
					end = cl.X + cl.W
					l++
					continue
				}
				// Add space to phrase before next word.
				if cl.Font == ck.Font && math.Abs(cl.FontSize-ck.FontSize) < 0.1 && cl.X <= end+wordSpace {
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
	//fmt.Println(words)
	return words
}

func sameFont(f1, f2 string) bool {
	//f1 = strings.TrimSuffix(f1, ",Italic")
	//f1 = strings.TrimSuffix(f1, "-Italic")
	//f2 = strings.TrimSuffix(f1, ",Italic")
	//f2 = strings.TrimSuffix(f1, "-Italic")
	//return strings.TrimSuffix(f1, ",Italic") == strings.TrimSuffix(f2, ",Italic") || f1 == "Symbol" || f2 == "Symbol" || f1 == "TimesNewRoman" || f2 == "TimesNewRoman"
	return f1 == f2
}
