package main

import (
	"fmt"

	pdf "github.com/hhrutter/pdfcpu/pkg/pdfcpu"
	"github.com/hhrutter/pdfcpu/pkg/pdfcpu/validate"
)

func main() {
	ctx, _ := pdf.ReadFile("../readfile/Interline-HD856.pdf", pdf.NewDefaultConfiguration())

	validate.XRefTable(ctx.XRefTable)
	fmt.Printf("%v", ctx)

	o, _ := ctx.FindObject(0)
	fmt.Printf("%s", o.String())

	//_ := (*ctx.Table[53]).Object.String()

	//fmt.Println(ctx.XRefTable)
	/*	xrefobj, _ := ctx.XRefTable FindObject(1)

		fmt.Println(xrefobj.PDFString())

		data, _ := pdf.ExtractStreamData(ctx, 20)
		fmt.Println(data)  */

	/*config := pdf.NewDefaultConfiguration()

	rs, _ := os.Open("../readfile/Interline-HD856.pdf")

	defer rs.Close()

	fileInfo, _ := rs.Stat()

	ctx, err := pdf.Read(rs, fileIn, fileInfo.Size(), config)

	fmt.Println(ctx)

	pdf.ReadP */

}
