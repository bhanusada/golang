package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	samplexsd "github.com/bhanu/xsdparser/go_Sample"
	"gopkg.in/go-playground/validator.v9"

	//samplexsd "github.com/bhanu/xsdparser/go_order"
	xsd "github.com/metaleap/go-xsd"
)

var validate *validator.Validate

func main() {

	validate = validator.New()

	command := os.Args[1]

	switch command {
	case "g":
		generateXSDTypes()
		fmt.Println("Generating XSD types.")
	case "v":
		validateXML()
		fmt.Println("Validating XSD.")
	default:
		fmt.Println("Inappropriate args.")
		validateXML()
	}

}

func generateXSDTypes() {
	var (
		sd  *xsd.Schema
		err error
	)

	xsd.PkgGen.BaseCodePath = "./"
	if _, err = xsd.LoadSchema("order.xsd", true); err != nil {
		log.Fatal(err)
	}

	sd.MakeGoPkgSrcFile()
}

func validateXML() {
	xmlfile, err := os.Open("sample.xml")
	if err != nil {
		log.Fatal(err)
	}
	defer xmlfile.Close()
	byteValue, _ := ioutil.ReadAll(xmlfile)

	v := samplexsd.TxsdOrderCanonical{}

	err = xml.Unmarshal([]byte(byteValue), &v)

	v.Walk()

	err = validate.Struct(v)
	if err != nil {
		if _, ok := err.(*validator.InvalidValidationError); ok {
			fmt.Println(err)
			return
		}

		for _, err := range err.(validator.ValidationErrors) {
			fmt.Println(err.Namespace())
			fmt.Println(err.Field())
			fmt.Println(err.StructNamespace())
			fmt.Println(err.StructField())
			fmt.Println(err.Tag())
			fmt.Println(err.ActualTag())
			fmt.Println(err.Kind())
			fmt.Println(err.Type())
			fmt.Println(err.Value())
			fmt.Println(err.Param())
			fmt.Println()
		}

		return
	}

	//fmt.Println(v.XsdGoPkgHasElem_OrderDetails.OrderDetails.OrderHeader.XsdGoPkgHasElem_OrderNumber.OrderNumber)
	//v.TransactionCode.Set("MODIFY")
	//fmt.Println(v.TransactionCode)

	if err != nil {
		log.Fatalln(err)
	}

	fmt.Printf("%#v", v)
}
