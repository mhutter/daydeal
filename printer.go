package daydeal

import (
	"fmt"
	"io"
	"strings"
	"time"
)

// Printer is responsible to output a deal to various sources
type Printer struct {
	Deal Printable
}

// Printable interface defines all methods required to print a deal
type Printable interface {
	Title() string
	Subtitle() string
	NewPrice() string
	OldPrice() string
	PriceSource() string
	Percentage() string
	NextDeal() time.Time
	NextDealIn() time.Duration
}

// NewPrinter returns a new Printer instance
func NewPrinter(d Printable) Printer {
	return Printer{Deal: d}
}

// PrintTo writes the deal to the Writer given
func (p Printer) PrintTo(out io.Writer) {
	fmt.Fprintln(out, "")
	fmt.Fprintf(out, "    %s\n", p.Deal.Title())
	fmt.Fprintf(out, "    %s\n", p.Deal.Subtitle())
	fmt.Fprintln(out, "")
	fmt.Fprintf(out, "Für %s statt %s (%s)\n",
		p.Deal.NewPrice(), p.Deal.OldPrice(), p.Deal.PriceSource())
	fmt.Fprintf(out, "Noch %s verfügbar\n", p.Deal.Percentage())
	fmt.Fprintf(out, "Nächster Deal am: %s (in %s)\n",
		p.Deal.NextDeal().Format("Mon Jan _2 15:04:05"),
		strings.TrimSuffix(p.Deal.NextDealIn().String(), "0s"))
}
