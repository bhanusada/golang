package main

import (
	"encoding/json"
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

	fmt.Println(xsd.OrderDetails.OrderDetails.OrderHeader.NewStoreFlag)

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
	xsdjson, _ := json.Marshal(&xsd)
	fmt.Println(string(xsdjson))
	filtertags(xsd)

}

func filtertags(xsd interface{}) {

	t := reflect.ValueOf(xsd)

	if t.Kind().String() == "ptr" && !t.IsNil() {
		t = t.Elem()
	}

	for i := 0; i < t.NumField(); i++ {
		f1 := t.Field(i)
		switch t.Field(i).Kind().String() {
		case "string":
			if reflect.TypeOf(xsd).Field(i).Tag.Get("validate") == "required" {
				if f1.Interface() == reflect.Zero(f1.Type()).Interface() {
					fmt.Printf("Field %s is missing.\n", reflect.TypeOf(xsd).Field(i).Name)
				}
			}
		case "int", "int16", "int32", "int64":
			if reflect.TypeOf(xsd).Field(i).Tag.Get("validate") == "required" {
				if f1.Interface() == reflect.Zero(f1.Type()).Interface() {
					fmt.Printf("Field %s is missing.\n", reflect.TypeOf(xsd).Field(i).Name)
				}
			}
		case "float32", "float64":
			if reflect.TypeOf(xsd).Field(i).Tag.Get("validate") == "required" {
				if f1.Interface() == reflect.Zero(f1.Type()).Interface() {
					fmt.Printf("Field %s is missing.\n", reflect.TypeOf(xsd).Field(i).Name)
				}
			}
		case "bool":
			if reflect.TypeOf(xsd).Field(i).Tag.Get("validate") == "required" {
				if f1.Bool() {
					fmt.Printf("bool %s is empty \n.", reflect.TypeOf(xsd).Field(i).Name)
				}
			}
		case "ptr":
			fmt.Println("pointer")
			fmt.Println(reflect.ValueOf(f1.Interface()).Elem())
			fmt.Println(reflect.TypeOf(xsd).Field(i).Name)
			if f1.IsNil() {
				fmt.Printf("pointer %s is empty \n.", reflect.TypeOf(xsd).Field(i).Name)
			}
		case "slice":
			if reflect.TypeOf(xsd).Field(i).Tag.Get("validate") == "required" {
				if f1.IsNil() {
					fmt.Printf("Slice %s is empty \n.", reflect.TypeOf(xsd).Field(i).Name)
				}
			}
			slicestruct(f1.Interface())
		case "struct":
			filtertags(f1.Interface())
		default:
			fmt.Println("Others")
		}
	}

}

func slicestruct(v interface{}) {

	s := reflect.ValueOf(v)

	for i := 0; i < s.Len(); i++ {
		switch s.Index(i).Kind().String() {
		case "struct":
			filtertags(s.Index(i).Interface())
		case "slice":
			slicestruct(s.Index(i).Interface())
		}
	}
}
