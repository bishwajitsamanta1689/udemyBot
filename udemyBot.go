package main

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"log"
	"net/http"
)

// Select Content for searching WebPage

func main() {

	res, err := http.Get("https://www.udemy.com/")
	if err != nil {
		log.Fatal(err)
	}
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	doc.Find("p[class]").Each(func(i int, selection *goquery.Selection) {
		fmt.Println(goquery.OuterHtml(selection))
	})
}
