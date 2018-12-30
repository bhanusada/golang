package main

import (
	"encoding/json"
	"encoding/xml"
	"errors"
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
			msg := make(map[string]interface{})
			fmt.Printf("Request method : %s", r.Method)
			msg["message"] = "Get request."
			res, _ := json.Marshal(msg)
			fmt.Fprint(w, string(res))
		case "POST":
			msg := make(map[string]interface{})
			fmt.Printf("Request method : %s", r.Method)
			fmt.Println()
			body, _ := ioutil.ReadAll(r.Body)
			doctype, err := extractAndValidate(&body)
			if err != nil {
				msg["message"] = "XML invalid."
			} else {
				fmt.Println(doctype)
				msg["message"] = "XML valid."
			}
			res, _ := json.Marshal(msg)
			fmt.Fprint(w, string(res))
		}
	}
	return http.HandlerFunc(fn)
}

func extractAndValidate(xmlfile *[]byte) (string, error) {
	rootElement := extractxmlroot(xmlfile)
	Notfound := findDocFlow(rootElement)
	fmt.Printf("Notfound ? %t\n", Notfound)
	if Notfound {
		return "", errors.New("Not found")
	}
	return validate(xmlfile)
}

func extractxmlroot(xmlfile *[]byte) string {
	var v = xmlroot{}
	err := xml.Unmarshal(*xmlfile, &v)
	if err != nil {
		log.Fatal(err)
	}
	return string(v.XMLName.Local)
}

func findDocFlow(xmlrootElement string) bool {
	return gomongo.DetermineDocFlow(xmlrootElement)
}

func validate(xmlfile *[]byte) (string, error) {
	return xsdparser.ValidateXML(xmlfile)
}
