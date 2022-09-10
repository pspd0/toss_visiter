package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/PuerkitoBio/goquery"
)

func main() {
	var id string
	fmt.Print("Enter Your Toss id: ")

	fmt.Scan(&id)

	fmt.Println("")
	
	for i := 0; i < 5; i++ {
		req, err := http.Get("https://toss.me/" + id)

		if err != nil {
			fmt.Println(err)
		}

		defer req.Body.Close()

		html, _ := goquery.NewDocumentFromReader(req.Body)

		value := html.Find("span.css-a774xx").Text()

		fmt.Println("\x1b[1m\x1b[32m" + value + "\x1b[0m")

		time.Sleep(time.Second)
	}
}
