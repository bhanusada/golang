package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/bhanu/gomongo"
	_ "github.com/bhanu/gomongo"
	"github.com/bhanu/xsdparser"
)

type xmlroot struct {
	XMLName xml.Name
}

func main() {
	//mux := http.NewServeMux()
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("got the request.")
		w.Write([]byte("helloworld"))
	})
	uh := uploadHandler()
	http.Handle("/upload", uh)
	http.ListenAndServe(":8080", nil)
}

func uploadHandler() http.Handler {
	fmt.Println("request")
	fn := func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "GET":
			fmt.Printf("Request method : %s", r.Method)
		case "POST":
			fmt.Printf("Request method : %s", r.Method)
			fmt.Println()
			body, _ := ioutil.ReadAll(r.Body)
			//fmt.Println(string(body))
			extractxmlroot(&body)
			parsedxml, err := xsdparser.Parse(&body)
			if err != nil {
				log.Fatal(err)
			}
			fmt.Println(parsedxml)
		}
		fmt.Fprint(w, "processed")
	}
	return http.HandlerFunc(fn)
}

func extractxmlroot(xmlfile *[]byte) {
	var v = xmlroot{}

	err := xml.Unmarshal(*xmlfile, &v)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(v.XMLName.Local)

	result := gomongo.DetermineDocFlow(v.XMLName.Local)

	fmt.Println(result)
}
