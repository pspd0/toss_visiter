package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/PuerkitoBio/goquery"
)

func checkErr(err error) {
	if err != nil {
		fmt.Println(err)
	}
}

func main() {
	var id string
	fmt.Print("Enter Your Toss id: ")

	fmt.Scan(&id)

	fmt.Println("")

	for {
		req, err := http.Get("https://toss.me/" + id)
		checkErr(err)

		defer req.Body.Close()

		html, err := goquery.NewDocumentFromReader(req.Body)
		checkErr(err)

		value := html.Find("span.css-a774xx").Text()

		fmt.Println("\x1b[1m\x1b[32m" + value + "\x1b[0m")

		time.Sleep(time.Second)
	}
}
