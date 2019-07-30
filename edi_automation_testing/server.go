package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"

	pdf "github.com/bhanu/edi_automation_testing/pdfParser"
	"github.com/gorilla/mux"
)

type response struct {
	PDF struct {
		Processed bool        `json:"processed"`
		Data      interface{} `json:"data"`
	} `json:"pdf"`
	TXO struct {
		Processed bool `json:"processed"`
	} `json:"txo"`
	Data struct {
		Processed bool `json:"processed"`
	} `json:"data"`
}

func handler(res http.ResponseWriter, req *http.Request) {
	//fmt.Fprint(res, "Hello world!")
	res.Write([]byte("Hello World"))
}

func fileHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	log.Printf("Got the requuest")

	var res response

	r.ParseMultipartForm(10 << 20)

	for key, value := range r.MultipartForm.Value {
		log.Printf("%s:%s", key, value)
	}

	for _, fileHeaders := range r.MultipartForm.File {
		for _, fileHeader := range fileHeaders {
			file, _ := fileHeader.Open()
			path := fmt.Sprintf("files/%s", fileHeader.Filename)
			buf, _ := ioutil.ReadAll(file)
			ioutil.WriteFile(path, buf, os.ModePerm)
			if strings.HasSuffix(fileHeader.Filename, ".pdf") {
				pdf.ParsePDF(fileHeader.Filename)
				var pdfData []map[string]interface{}
				jsonData, _ := ioutil.ReadFile("./edi.json")
				json.Unmarshal(jsonData, &pdfData)
				fmt.Printf("%#v", pdfData)
				res.PDF.Data = pdfData
				res.PDF.Processed = true
			}
			res.Data.Processed = true
			res.TXO.Processed = true
		}
	}
	//responseData, _ := json.Marshal(res)
	json.NewEncoder(w).Encode(res)
	//w.Write(responseData)
	//fmt.Fprintf(w, res)
}

func main() {
	/*c := cors.New(cors.Options{
		AllowedMethods:     []string{"GET", "POST", "OPTIONS"},
		AllowedOrigins:     []string{"*"},
		AllowCredentials:   true,
		AllowedHeaders:     []string{"Content-Type", "Bearer", "Bearer ", "content-type", "Origin", "Accept"},
		OptionsPassthrough: true,
	})*/

	r := mux.NewRouter()

	r.HandleFunc("/upload", fileHandler).Methods(http.MethodGet, http.MethodPut, http.MethodPost, http.MethodPatch, http.MethodOptions)
	r.Use(mux.CORSMethodMiddleware(r))

	r.HandleFunc("/test", handler).Methods("GET")

	r.PathPrefix("/").Handler(http.StripPrefix("/", http.FileServer(http.Dir("build"))))

	//mux := http.NewServeMux()
	//mux.HandleFunc("/test", handler)
	//mux.HandleFunc("/upload", fileHandler)

	/*rawHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "%s", r.URL.Query().Get("param"))
	})*/

	//http.Handle("/home", http.StripPrefix("/home", http.FileServer(http.Dir("./build"))))

	//handler := c.Handler(mux)
	http.ListenAndServe(":8080", r)
}
