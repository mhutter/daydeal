package daydeal

import (
	"fmt"
	"io"

	"github.com/PuerkitoBio/goquery"
)

// URLs of the various daydeal sources
const (
	URLDaydealDay  = "https://www.daydeal.ch/"
	URLDaydealWeek = "https://www.daydeal.ch/deal-of-the-week"
)

// App parses arguments and calls the other components accordingly.
type App struct {
	Out io.Writer
	URL string
}

// NewApp returns a new App instance with sensible defaults.
func NewApp(out io.Writer, url string) App {
	return App{Out: out, URL: url}
}

// FetchDeal fetches the daydeal with the url given.
func FetchDeal(url string) (deal Daydeal, err error) {
	doc, err := goquery.NewDocument(url)
	return NewDaydeal(doc), err
}

// Run the Daydeal app
func (a App) Run() error {
	deal, err := FetchDeal(a.URL)
	if err != nil {
		return fmt.Errorf("Could not fetch deal: %v", err)
	}

	NewPrinter(a.Out).Print(deal)
	return nil
}
