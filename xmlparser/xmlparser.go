package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"gopkg.in/go-playground/validator.v9"
)

type xmlresult struct {
	orderCanonical orderCanonical
}

type orderCanonical struct {
	//XMLName          xml.Name       `xml:"orderCanonical"`
	TransactionCode  string       `xml:"transactionCode" validate:"required"`
	CurrentTimestamp string       `xml:"currentTimestamp" validate:"required"`
	OrderDetails     orderDetails `xml:"orderDetails" validate:"required"`
}

type orderDetails struct {
	//XMLName     xml.Name      `xml:"orderDetails"`
	OrderHeader orderHeader `xml:"orderHeader" validate:"required"`
}

type orderHeader struct {
	shippingLocationNumber        string
	shippingLocationTypeCode      string
	receivingLocationNumber       string `validate:"required"`
	receivingLocationTypeCode     string `validate:"required"`
	merchandisingVendorNumber     string
	merchandisingDepartmentNumber string
	distributionServiceTypeCode   string
	replenishmentMethodCode       string
	orderReasonCode               string
	orderNumber                   string `validate:"required"`
	orderClass                    string
	orderCreateDate               string
	orderCreateTimestamp          string
	lastUpdatedTimestamp          string
	expectedArrivalDate           string
	estimatedShipDate             string
	cancelEligibilityDate         string
	orderCancelTimestamp          string
	orderClosedTimestamp          string
	omtOrderTypeCode              string
	accountOrderTypeCode          string
	productOriginTypeCode         string
	productDestinationTypeCode    string
	orderRequestingSystemCode     string
	peggedOrderFlag               string
	specialOrderFlag              string
	inboundOrderTypeCode          string
	orderHeaderVersionNumber      string
	orderStatusCode               string
	orderOpenTimestamp            string
	orderOpenSystemUserId         string
	newStoreFlag                  string
	consignmentFlag               string
	commentText                   string
	wmsShippingNumber             string
	allocationId                  string
	srcOrderRefId                 string
	omtOrderRefTypeCode           string
	vendorCertifiedFlag           string
	orderTransmissionTypeCode     string
	relatedOrders                 string
}

var validate *validator.Validate

var v = orderCanonical{}

func main() {

	validate = validator.New()

	validateXML()
}

func validateXML() {

	xmlfile, err := os.Open("sample.xml")

	defer xmlfile.Close()

	byteValue, _ := ioutil.ReadAll(xmlfile)

	err = xml.Unmarshal([]byte(byteValue), &v)
	if err != nil {
		log.Fatal(err)
	}

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

	//fmt.Printf("%#v", v)
	copy()

}

func copy() {

	copyxml, copyerr := xml.Marshal(v)
	if copyerr != nil {
		fmt.Println(copyerr)
	}

	fmt.Println(string(copyxml[:len(copyxml)]))

}
