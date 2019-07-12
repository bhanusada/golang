package main

import (
	"fmt"
	"math"
	"strings"

	"rsc.io/pdf"
)

type parameters struct {
	name     string
	id       int
	desc     string
	required string
	datatype string
	minmax   string
}

func main() {
	r, _ := pdf.Open("./856_HDPRO_ 4060 [04022019].pdf")
	segment := ""
	field := ""
	num := 0
	var fields []map[string]string
	//var fields []interface{}
	for i := 4; i < r.NumPage(); i++ {
		p := r.Page(i)
		words := findWords(p.Content().Text)

		for _, v := range words {
			//fmt.Println(v)
			if v.FontSize > 18 && len(v.S) < 4 {
				segment = v.S
				fmt.Println(segment)
				continue
			}
			if len(segment) > 0 {
				if len(v.S) >= len(segment) {
					if v.S[:len(segment)] == segment {
						fmt.Println(v.S)
						field = v.S
						num = 0
						//params = map[string]string{}
						fmt.Println(len(fields))
						fields[len(fields)] = map[string]string{}
						continue
						//segment = ""
					}
				}
			}
			if len(field) > 0 && num < 10 {
				num++
				temp := fields[len(fields)]
				for k, _ := range temp {
					_, ok := temp[k]
					if !ok {
						fields[len(fields)][k] = v.S
						break
					}
				}
				fields = append(fields, temp)
				if strings.Contains(v.S, "Description:") {
					num = 10
					continue
				}
				//fmt.Println(v.S)
			}
			fmt.Printf("%V", fields)
		}
	}

	//fmt.Println(r.NumPage())
	//fmt.Printf("%v", p.Content())
	//fmt.Println(r.Trailer().Key("Root").String())
	//fmt.Println(p.V.TextFromUTF16())

	//for _, v := range p.Content().Text {
	//	fmt.Printf("%#v", v)
	//}

}

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
		// fmt.Printf("%#v\n", chars[i])
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
			words = append(words, pdf.Text{Font: ck.Font, FontSize: ck.FontSize, X: ck.X, Y: ck.Y, W: end - ck.X, S: s})
			k = l
		}
		i = j
	}
	//fmt.Println(words)
	return words
}
