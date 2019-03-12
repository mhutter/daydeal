package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func main() {
	run(os.Stdout)
}

func run(out io.Writer) {
	log.SetOutput(out)

	doc, err := goquery.NewDocument("https://www.daydeal.ch/")
	if err != nil {
		log.Fatal(err)
	}

	title := doc.Find(".product-description__title1").First().Text()
	subtitle := doc.Find(".product-description__title2").First().Text()

	price := doc.Find(".product-pricing__prices-new-price").First().Text()
	originalPrice := doc.Find(".product-pricing__prices-old-price").First().Text()
	originalPrice = strings.TrimSpace(originalPrice)
	//extract the external price comparison source (1,2,3)
	priceSource := originalPrice[len(originalPrice)-1:]
	//remove the the external price comparison form the price string
	originalPrice = originalPrice[:len(originalPrice)-1]

	percentage := doc.Find(".product-progress__availability").First().Text()

	fmt.Fprintf(out, "\n    %s\n    %s\n\n", title, subtitle)
	fmt.Fprintf(out, "Für %s %s (%s)\n", price, originalPrice, priceSource)
	fmt.Fprintf(out, "Noch %s verfügbar\n", percentage)
}
