package main

import (
	"encoding/json"
	"fmt"
	"math"
	"strings"

	"rsc.io/pdf"
)

type segmentStruct struct {
	Name string        `json:"name"`
	List []fieldStruct `json:"list"`
}
type fieldStruct struct {
	Name     string `json:"name"`
	Id       string `json:"id"`
	Desc     string `json:"desc"`
	Req      string `json:"req"`
	Datatype string `json:"datatype"`
	Minmax   string `json:"minmax"`
}

func (f *fieldStruct) setid(id string) {
	f.Id = id
}
func (f *fieldStruct) setdesc(desc string) {
	f.Desc = desc
}
func (f *fieldStruct) setreq(req string) {
	f.Req = req
}
func (f *fieldStruct) setdatatype(datatype string) {
	f.Datatype = datatype
}
func (f *fieldStruct) setminmax(minmax string) {
	f.Minmax = minmax
}

func main() {
	r, _ := pdf.Open("./856_HDPRO_ 4060 [04022019].pdf")
	segment := ""
	field := ""
	num := 0
	var edi []segmentStruct
	var seg *segmentStruct
	var fl *fieldStruct
	for i := 4; i < r.NumPage(); i++ {
		p := r.Page(i)
		words := findWords(p.Content().Text)
		for _, v := range words {
			if v.FontSize > 18 && len(v.S) < 4 {
				segment = v.S
				edi = append(edi, segmentStruct{Name: v.S})
				seg = &edi[len(edi)-1]
				continue
			}
			if len(segment) > 0 {
				if len(v.S) >= len(segment) {
					if v.S[:len(segment)] == segment {
						field = v.S
						num = 0
						seg.List = append(seg.List, fieldStruct{})
						seg.List[len(seg.List)-1].Name = v.S
						fl = &seg.List[len(seg.List)-1]
						continue
					}
				}
			}
			if len(field) > 0 && num < 10 {
				num++
				if num == 1 {
					fl.setid(v.S)
				}
				if num == 2 {
					fl.setdesc(v.S)
				}
				if num == 3 {
					fl.setreq(v.S)
				}
				if num == 4 {
					fl.setdatatype(v.S)
				}
				if num == 5 {
					fl.setminmax(v.S)
				}
				if strings.Contains(v.S, "Description:") {
					num = 10
					continue
				}
			}
		}
	}
	edijson, err := json.Marshal(edi)
	if err != nil {
		fmt.Println("error")
		fmt.Println(err)
	} else {
		fmt.Println(string(edijson))
	}
	//for _, val1 := range edi {
	//	fmt.Println(val1)
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
