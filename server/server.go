package main

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/bhanu/xsdparser"
)

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
			fmt.Println(string(body))

			xsdparser.Parse(body)
		}

		fmt.Fprint(w, "processed")
	}

	return http.HandlerFunc(fn)

}
