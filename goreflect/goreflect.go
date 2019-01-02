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
	//xsdjson, _ := json.Marshal(&xsd)
	//fmt.Println(string(xsdjson))
	filtertags(xsd)

}

func filtertags(xsd interface{}) {
	//fmt.Println(xsd)
	t := reflect.ValueOf(xsd)
	//typeofT := t.Type()
	if t.Kind() == reflect.Ptr && !t.IsNil() {
		t = t.Elem()
	}
	//fmt.Println(t)
	//fmt.Println(t.NumField())
	for i := 0; i < t.NumField(); i++ {
		//fmt.Println(i)
		f1 := t.Field(i)
		//fmt.Println(f1)
		//fmt.Println(f1.String())
		//fmt.Println(t.Field(i).Kind().String())
		switch t.Field(i).Kind().String() {
		case "string":
			//fmt.Println("string")
			if reflect.TypeOf(xsd).Field(i).Tag.Get("validate") == "required" {
				fmt.Println(reflect.TypeOf(xsd).Field(i).Name)
			}
		case "int", "int16", "int32", "int64":
			//fmt.Println("int")
			if reflect.TypeOf(xsd).Field(i).Tag.Get("validate") == "required" {
				fmt.Println(reflect.TypeOf(xsd).Field(i).Name)
			}
		case "float32", "float64":
			//fmt.Println("float")
			if reflect.TypeOf(xsd).Field(i).Tag.Get("validate") == "required" {
				fmt.Println(reflect.TypeOf(xsd).Field(i).Name)
			}
		case "bool":
			//fmt.Println("bool")
			if reflect.TypeOf(xsd).Field(i).Tag.Get("validate") == "required" {
				fmt.Println(reflect.TypeOf(xsd).Field(i).Name)
			}
		case "slice":
			fmt.Println("slice")
			slicestruct(f1.Interface())
		case "struct":
			filtertags(f1.Interface())
		default:
			fmt.Println("Others")
			//fmt.Println(f1.Kind().String())
			//filtertags(reflect.ValueOf(f1))
		}
	}

}

func slicestruct(v interface{}) {
	s := reflect.ValueOf(v)

	for i := 0; i < s.Len(); i++ {
		if reflect.TypeOf(s.Index(i)).Kind() == reflect.Struct {
			filtertags(s.Index(i))
		}
		//fmt.Println(s.Index(i))
	}
}
