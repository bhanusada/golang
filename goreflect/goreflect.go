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

	//fmt.Println(xsd.TransactionCode)

	/*

		fmt.Println(t.Field(1))

		f, found := t.FieldByName("RelatedOrderNumber")
		if !found {
			log.Print("Not found")
		}
		fmt.Println()

		fmt.Println(f.Tag.Get("validate"))
	*/
	//xsd1 := map[string]interface{}(xsd)
	filtertags(xsd.OrderDetails.OrderDetails)

}

func filtertags(xsd interface{}) {
	fmt.Println(xsd)
	t := reflect.TypeOf(xsd)

	for i := 0; i < t.NumField(); i++ {
		fmt.Println()
		f1 := t.Field(i)
		//fmt.Println(f1.Type)
		fmt.Println(f1.Name)
		switch reflect.TypeOf(f1).String() {
		case "string":
			fmt.Println("string")
		case "int", "int16", "int32", "int64":
			fmt.Println("int")
		case "bool":
			fmt.Println("bool")
		default:
			fmt.Println("Others")
			filtertags(f1)
		}
	}

}
