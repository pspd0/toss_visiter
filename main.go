package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/PuerkitoBio/goquery"
)

func main() {
	for {
		id := "rorack"

		req, err := http.Get("https://toss.me/" + id)

		if err != nil {
			fmt.Println(err)
		}

		defer req.Body.Close()

		html, _ := goquery.NewDocumentFromReader(req.Body)

		value := html.Find("span.css-a774xx").Text()

		fmt.Println(value)
		time.Sleep(time.Second)
	}
}
