package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

const (
	urlDay  = "https://www.daydeal.ch/"
	urlWeek = "https://www.daydeal.ch/deal-of-the-week"
)

var (
	fetchDOTW = false
)

func init() {
	flag.BoolVar(&fetchDOTW, "w", fetchDOTW, "Fetch the deal of the Week instead")
	flag.Parse()
}

func main() {
	url := urlDay
	if fetchDOTW {
		url = urlWeek
	}
	run(os.Stdout, url)
}

func run(out io.Writer, url string) {
	log.SetOutput(out)

	doc, err := goquery.NewDocument(url)
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
