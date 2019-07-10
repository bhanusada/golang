package main

import (
	"fmt"

	"rsc.io/pdf"
)

func main() {
	file, _ := pdf.Open("../readfile/Interline-HD856.pdf")
	pages := file.NumPage()
	fmt.Println(pages)
	page := file.Page(6)
	pageContent := ""
	for _, cont := range page.Content().Text {
		pageContent += cont.S
	}
	fmt.Println(pageContent)
	//content := page.Content()
	//fmt.Println(content.Rect)

}
