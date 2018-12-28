package xsdparser

import (
	"encoding/xml"
	"fmt"
	"log"

	samplexsd "github.com/bhanu/xsdparser/go_Sample"
	validator "gopkg.in/go-playground/validator.v9"

	//samplexsd "github.com/bhanu/xsdparser/go_order"
	xsd "github.com/metaleap/go-xsd"
)

var validate *validator.Validate
var parsedxml samplexsd.TxsdOrderCanonical

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

//ValidateXML method to validate a given XML against specific XSD
func ValidateXML(xmlfile *[]byte) (string, error) {

	err := xml.Unmarshal(*xmlfile, &parsedxml)

	if err != nil {
		log.Fatal(err)
	}

	validate = validator.New()

	err = validate.Struct(parsedxml)

	fmt.Printf("Transaction code : %s", parsedxml.TransactionCode.String())

	if err != nil {
		if _, ok := err.(*validator.InvalidValidationError); ok {
			fmt.Println(err)
			return "", err
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
		return "", err
	}

	doctype := parsedxml.OrderDetails.OrderHeader.DocumentTypeCode.String()

	return doctype, nil
}
