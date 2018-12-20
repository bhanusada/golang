package main

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/bhanu/xsdparser"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("got the request.")
		w.Write([]byte("helloworld"))
	})

	http.HandleFunc("/upload", handler)

	http.ListenAndServe(":8080", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case "GET":
		fmt.Printf("Request method : %s", r.Method)
	case "POST":
		fmt.Printf("Request method : %s", r.Method)
		fmt.Println()

		body, _ := ioutil.ReadAll(r.Body)
		fmt.Println(string(body))

		xsdparser.Parse(body)
	}

	fmt.Fprint(w, "processed")
}
