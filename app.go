package daydeal

import (
	"flag"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/PuerkitoBio/goquery"
)

// Kind of deal (Daydeal of the Day, Daydeal of the Week, ...) to be
// displayed
type Kind uint8

// Different kinds of deals to be fetched and displayed
const (
	DaydealDay Kind = iota
	DaydealWeek
)

var (
	version = "dev"
	commit  = "none"
	date    = "unknown"
)

var urls = map[Kind]string{
	DaydealDay:  "https://www.daydeal.ch/",
	DaydealWeek: "https://www.daydeal.ch/deal-of-the-week",
}

// App parses arguments and calls the other components accordingly.
type App struct {
	Out io.Writer
}

// NewApp returns a new App instance with sensible defaults.
func NewApp(out io.Writer) App {
	return App{Out: out}
}

// FetchDeal fetches the daydeal with the kind given.
func FetchDeal(k Kind) (deal Daydeal, err error) {
	url := urls[k]
	if url == "" {
		return deal, fmt.Errorf("Unknow deal %v", k)
	}

	doc, err := goquery.NewDocument(url)
	return NewDaydeal(doc), err
}

// Run the Daydeal app
func (a App) Run(args []string) {
	var (
		fs             = flag.NewFlagSet("daydeal", flag.ExitOnError)
		week           = false
		displayVersion = false
		kind           = DaydealDay
	)

	fs.BoolVar(&week, "w", week, "Fetch deal of the week instead")
	fs.BoolVar(&displayVersion, "v", displayVersion,
		"Display program version and exit")

	if err := fs.Parse(args); err != nil {
		if err == flag.ErrHelp {
			fmt.Fprintf(os.Stderr, "Usage of %s:\n", fs.Name())
			fs.PrintDefaults()
			return
		}

		log.Fatalf("Could not parse flags: %v", err)
	}

	if displayVersion {
		fmt.Fprintf(a.Out, "%v, commit %v, built on %v\n",
			version, commit, date)
		return
	}

	if week {
		kind = DaydealWeek
	}

	deal, err := FetchDeal(kind)
	if err != nil {
		log.Fatalf("Could not fetch deal: %v", err)
	}

	p := NewPrinter(deal)
	p.PrintTo(a.Out)
}
