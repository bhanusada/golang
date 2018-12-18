package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	samplexsd "github.com/bhanu/xsdparser/go_Sample"
)

func main() {
	/*
		var (
			sd  *xsd.Schema
			err error
		)

		xsd.PkgGen.BaseCodePath = "./"
		if sd, err = xsd.LoadSchema("sample.xsd", true); err != nil {
			log.Fatal(err)
		}
	*/

	xmlfile, err := os.Open("sample.xml")
	if err != nil {
		log.Fatal(err)
	}
	defer xmlfile.Close()
	byteValue, _ := ioutil.ReadAll(xmlfile)

	v := samplexsd.TxsdOrderCanonical{}

	err = xml.Unmarshal([]byte(byteValue), &v)

	v.Walk()

	fmt.Println(v.XsdGoPkgHasElem_OrderDetails.OrderDetails.OrderHeader.XsdGoPkgHasElem_OrderNumber.OrderNumber)

	if err != nil {
		log.Fatalln(err)
	}

	fmt.Printf("%#v", v)

}
