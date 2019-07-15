package main

import (
	"fmt"
	"math"
	"strings"

	"rsc.io/pdf"
)

type segmentStruct struct {
	name string
	list []fieldStruct
}
type fieldStruct struct {
	name     string
	id       string
	desc     string
	req      string
	datatype string
	minmax   string
}

func (f *fieldStruct) setid(id string) {
	f.id = id
}
func (f *fieldStruct) setdesc(desc string) {
	f.desc = desc
}
func (f *fieldStruct) setreq(req string) {
	f.req = req
}
func (f *fieldStruct) setdatatype(datatype string) {
	f.datatype = datatype
}
func (f *fieldStruct) setminmax(minmax string) {
	f.minmax = minmax
}

func main() {
	r, _ := pdf.Open("./856_HDPRO_ 4060 [04022019].pdf")
	segment := ""
	field := ""
	num := 0
	var edi []segmentStruct
	//var fields []map[string]string
	//var temp fieldStruct
	var seg segmentStruct
	var fl *fieldStruct
	//var segments []map[string]interface{}
	//keys := []string{"id", "desc", "required", "datatype", "minmax"}
	for i := 4; i < r.NumPage(); i++ {
		p := r.Page(i)
		words := findWords(p.Content().Text)
		for _, v := range words {
			if v.FontSize > 18 && len(v.S) < 4 {
				segment = v.S
				seg = segmentStruct{}
				seg.name = v.S

				//seg.list = []fieldStruct{}
				//segments = append(segments, map[string]interface{}{"name": v.S, "list": []map[string]string{}})
				continue
			}
			if len(segment) > 0 {
				if len(v.S) >= len(segment) {
					if v.S[:len(segment)] == segment {
						field = v.S
						num = 0
						seg.list = append(seg.list, fieldStruct{})
						//fields = append(fields, map[string]string{})
						seg.list[len(seg.list)-1].name = v.S
						fl = &seg.list[len(seg.list)-1]
						edi = append(edi, seg)
						continue
					}
				}
			}
			if len(field) > 0 && num < 10 {
				num++
				//temp = seg.list[len(seg.list)-1]
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
				/*for _, val := range keys {
					_, ok := temp[val]
					if !ok {
						seg.list[len(seg.list)-1][val] = v.S
						break
					}
				}*/
				/*x := reflect.ValueOf(seg.list[len(seg.list)-1])
				for i := 1; i < x.NumField(); i++ {
					if x.Field(i).Interface() == "" {
						x.Elem().FieldByIndex(i).SetString()
					}
				}*/
				if strings.Contains(v.S, "Description:") {
					num = 10
					continue
				}
			}
		}
	}
	for _, val1 := range edi {
		fmt.Println(val1)
	}
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
