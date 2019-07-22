package main

import (
	"net/http"
)

func handler(res http.ResponseWriter, req *http.Request) {
	//fmt.Fprint(res, "Hello world!")
	res.Write([]byte("Hello World"))
}

func main() {
	http.HandleFunc("/test", handler)
	http.Handle("/home", http.StripPrefix("/home", http.FileServer(http.Dir("./build"))))
	http.ListenAndServe(":8080", nil)
}
