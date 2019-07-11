package main

import (
	"net/http"
)

func handler(res http.ResponseWriter, req *http.Request) {
	//fmt.Fprint(res, "Hello world!")
	res.Write([]byte("Hello World"))
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
