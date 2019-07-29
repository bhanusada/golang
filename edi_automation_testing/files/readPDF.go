package readPDF

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"math"
	"strings"

	"rsc.io/pdf"
)

type segmentStruct struct {
	Name string        `json:"name"`
	Desc string        `json:"desc"`
	Max  string        `json:"max"`
	Loop string        `json:"loop"`
	Req  string        `json:"req"`
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

//PDFParser to parse PDF
func PDFParser(fname string) {
	//r, _ := pdf.Open("./856_HDPRO_ 4060 [04022019].pdf")
	r, _ := pdf.Open("../edi_automation_testing/files/" + fname)
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
			if v.FontSize > 20 && len(v.S) < 4 {
				segment = v.S
				edi = append(edi, segmentStruct{Name: v.S})
				seg = &edi[len(edi)-1]
				continue
			}
			if v.FontSize > 18 {
				seg.Desc = seg.Desc + " " + v.S
				seg.Desc = strings.TrimSpace(seg.Desc)
				continue
			}
			if len(segment) > 0 {
				if strings.HasPrefix(v.S, "Max: ") {
					seg.Max = v.S[5:]
					continue
				}
				if strings.HasPrefix(v.S, "Loop: ") {
					seg.Loop = v.S[6:]
					continue
				}
				if strings.Contains(v.S, "Optional") {
					seg.Req = "N/A"
					continue
				}
				if strings.Contains(v.S, "Mandatory") {
					seg.Req = "M"
					continue
				}
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
				switch num {
				case 1:
					fl.setid(v.S)
				case 2:
					fl.setdesc(v.S)
				case 3:
					fl.setreq(v.S)
				case 4:
					fl.setdatatype(v.S)
				case 5:
					fl.setminmax(v.S)
				}
				if strings.Contains(v.S, "Description:") {
					num = 10
					continue
				}
			}
		}
	}
	buf := new(bytes.Buffer)
	enc := json.NewEncoder(buf)
	enc.SetEscapeHTML(false)
	enc.SetIndent("", " ")
	_ = enc.Encode(edi)
	_ = ioutil.WriteFile("edi.json", buf.Bytes(), 0644)
	log.Printf("written json")
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
