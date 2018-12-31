package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"log"
	"reflect"

	"github.com/bhanu/xsdparser/xsdtype"
)

func main() {

	var xsd xsdtype.OrderCanonical

	file, _ := ioutil.ReadFile("../xsdparser/sample.xml")

	err := xml.Unmarshal(file, &xsd)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(xsd.TransactionCode)

	t := reflect.TypeOf(xsd)

	fmt.Println(t.Field(1))

	f, found := t.FieldByName("RelatedOrderNumber")
	if !found {
		log.Print("Not found")
	}
	fmt.Println()

	fmt.Println(f.Tag.Get("validate"))

	//xsd1 := map[string]interface{}(xsd)

	for i := 0; i < t.NumField(); i++ {
		f1 := t.Field(i)
		fmt.Println(f1.Type)
	}

}
