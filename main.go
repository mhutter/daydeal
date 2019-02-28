package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
	"time"

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

func fmtDuration(d time.Duration) string {
	d = d.Round(time.Second)
	h := d / time.Hour
	d -= h * time.Hour
	m := d / time.Minute
	d -= m * time.Minute
	s := d / time.Second
	return fmt.Sprintf("%02d:%02d:%02d", h, m, s)
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

	// the website returns the moment of the next deal like "2006-01-02 15:04:05"
	// in UTC without announcingt that it is UTC.
	nextDealUTC := doc.Find("span.js-clock").AttrOr("data-next-deal", "")
	nextDeal, _ := time.ParseInLocation("2006-01-02 15:04:05", nextDealUTC, time.Local)

	nextDealIn := time.Until(nextDeal)
	nextDealInFmt := fmtDuration(nextDealIn)

	fmt.Fprintf(out, "\n    %s\n    %s\n\n", title, subtitle)
	fmt.Fprintf(out, "Für %s %s (%s)\n", price, originalPrice, priceSource)
	fmt.Fprintf(out, "Noch %s verfügbar\n", percentage)
	// Golang time formatting: https://flaviocopes.com/go-date-time-format/
	fmt.Fprintf(out, "Nächster Deal am: %s (in %s)\n", nextDeal.Format("Mon Jan _2 15:04:05"), nextDealInFmt)
}
