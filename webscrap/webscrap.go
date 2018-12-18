package main

import (
	"fmt"
	"log"

	"github.com/gocolly/colly"
)

func main() {

	log.Println("Processed started.")

	c := colly.NewCollector()

	err := c.Post("https://cn.sterlingcommerce.com/login.jsp", map[string]string{"username": "bhanu_prakash_sadavala@homedepot.com", "password": "TCSdec@18"})
	if err != nil {
		log.Fatal(err)
	}

	c.OnHTML("#j_username", func(e *colly.HTMLElement) {
		fmt.Fprint(e.Text, "bhanu")
		log.Println("element :", e.Text)
	})

	c.OnResponse(func(r *colly.Response) {
		log.Println("response received :", r.StatusCode)
	})

	c.OnRequest(func(r *colly.Request) {
		log.Println("request :", r.URL.String())
	})

	c.Visit("https://cn.sterlingcommerce.com/")
}
