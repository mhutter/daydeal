package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func main() {
	doc, err := goquery.NewDocument("https://www.daydeal.ch/")
	if err != nil {
		log.Fatal(err)
	}

	title := doc.Find("h1.meta-first-line").First().Text()
	subtitle := doc.Find("span.meta-second-line").First().Text()

	price := doc.Find("span.price").First().Text()
	originalPrice := doc.Find("div.originalPrice span").First().Text()
	originalPrice = strings.TrimSpace(originalPrice)
	originalPrice = strings.TrimSuffix(originalPrice, "*")
	originalPrice = strings.TrimSpace(originalPrice)

	percentage := doc.Find("span.percentage").First().Text()

	fmt.Printf("\n    %s\n    %s\n\n", title, subtitle)
	fmt.Printf("Für CHF %s anstatt %s\n", price, originalPrice)
	fmt.Printf("Noch %s%% verfügbar\n", percentage)
}
