package main

import (
	"fmt"
	"log"
	"github.com/lestrrat/go-libxml2/xsd"
)

func main() {
	schema, err := xsd.Parse("./Order_Canonical_850_MincronPO.xsd")

	if err != nil {
		log.Fatal(err)
	}

	defer schema.Free()
	if err := schema.Validate("./mincronPO.xml"); err != nil {
		for _, e := range err.(SchemaValidationErr).Error() {
			fmt.Println(e.Error())
		}
	}
}

