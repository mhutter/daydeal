package daydeal

import (
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
)

const (
	titleSelector       = ".product-description__title1"
	subtitleSelector    = ".product-description__title2"
	newPriceSelector    = ".product-pricing__prices-new-price"
	oldPriceSelector    = ".product-pricing__prices-old-price .js-old-price"
	priceSourceSelector = ".product-pricing__prices-old-price-annotation"
	percentageSelector  = ".product-progress__availability"
	nextDealSelector    = ".js-clock"
	nextDealAttr        = "data-next-deal"
	nextDealFmt         = "2006-01-02 15:04:05"
)

// Daydeal holds all information about the current deal and knows how to
// extract information from it.
type Daydeal struct {
	*goquery.Document
}

// NewDaydeal returns a new Daydeal instance with the given document.
func NewDaydeal(doc *goquery.Document) Daydeal {
	return Daydeal{Document: doc}
}

// Title displays the title of the deal (ie name of the item sold)
func (d Daydeal) Title() string {
	return d.getText(titleSelector)
}

// Subtitle displays the subtitle of the deal (ie a description of the
// item sold)
func (d Daydeal) Subtitle() string {
	return d.getText(subtitleSelector)
}

// NewPrice displays the new (reduced) price
func (d Daydeal) NewPrice() string {
	return d.getText(newPriceSelector)
}

// OldPrice displays the "original" price of the item
func (d Daydeal) OldPrice() string {
	val := d.getText(oldPriceSelector)
	return strings.TrimRight(val, "123")
}

// PriceSource displays the price source according to daydeal.ch:
// 1 Konkurrenzvergleich
// 2 Selbstvergleich BRACK.CH
// 3 Einführungspreis
func (d Daydeal) PriceSource() string {
	return d.getText(priceSourceSelector)
}

// NextDeal returns the time & date of the next deal.
func (d Daydeal) NextDeal() time.Time {
	nextDealUTC := d.Document.Find(nextDealSelector).AttrOr(nextDealAttr, "")
	nextDeal, _ := time.ParseInLocation(nextDealFmt, nextDealUTC, time.UTC)
	return nextDeal.Local()
}

// NextDealIn returns the duration to the next deal.
func (d Daydeal) NextDealIn() time.Duration {
	return d.NextDeal().Sub(time.Now()).Truncate(time.Minute)
}

// Percentage displays how many items (in percent obviously) are left
func (d Daydeal) Percentage() string {
	return d.getText(percentageSelector)
}

func (d Daydeal) getText(selector string) string {
	val := d.Document.Find(selector).First().Text()
	return strings.TrimSpace(val)
}
