package main

import (
	"io"
	"log"
	"net/http"
	"os"
)

// Select Content for searching WebPage

func main() {

	res, err := http.Get("https://www.udemy.com/")
	if err != nil {
		log.Fatal(err)
	}
	file, err := os.Create("udemy.html")
	if err != nil {
		log.Fatal(err)
	}
	n, err := io.Copy(file, res.Body)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Number of Bytes Copied:", n)
}
