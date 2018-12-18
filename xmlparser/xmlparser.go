package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

type xmlresult struct {
	orderCanonical orderCanonical
}

type orderCanonical struct {
	//XMLName          xml.Name       `xml:"orderCanonical"`
	TransactionCode  string         `xml:"transactionCode"`
	CurrentTimestamp string         `xml:"currentTimestamp"`
	OrderDetails     []orderDetails `xml:"orderDetails"`
}

type orderDetails struct {
	//XMLName     xml.Name      `xml:"orderDetails"`
	OrderHeader []orderHeader `xml:"orderHeader"`
}

type orderHeader struct {
	receivingLocationNumber       string
	receivingLocationTypeCode     string
	merchandisingVendorNumber     string
	merchandisingDepartmentNumber string
	distributionServiceTypeCode   string
	replenishmentMethodCode       string
	orderReasonCode               string
	orderNumber                   string
	orderClass                    string
	orderCreateDate               string
	orderCreateTimestamp          string
	lastUpdatedTimestamp          string
	expectedArrivalDate           string
	estimatedShipDate             string
}

func main() {
	v := orderCanonical{TransactionCode: "test"}

	xmlfile, err := os.Open("sample.xml")

	defer xmlfile.Close()
	/*
		scanner := bufio.NewScanner(xmlfile)
		for scanner.Scan() {
			myslice = myslice + scanner.Text()
		} */

	byteValue, _ := ioutil.ReadAll(xmlfile)

	//fmt.Println(myslice)
	err = xml.Unmarshal([]byte(byteValue), &v)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%#v", v)
}
