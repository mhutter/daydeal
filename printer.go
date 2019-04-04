package daydeal

import (
	"fmt"
	"io"
	"strings"
	"time"
)

// Printer is responsible to output a deal to various sources
type Printer struct {
	Out io.Writer
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
func NewPrinter(out io.Writer) *Printer {
	return &Printer{Out: out}
}

// Print writes the given deal to p.Out
func (p Printer) Print(deal Printable) {
	fmt.Fprintln(p.Out, "")
	fmt.Fprintf(p.Out, "    %s\n", deal.Title())
	fmt.Fprintf(p.Out, "    %s\n", deal.Subtitle())
	fmt.Fprintln(p.Out, "")
	fmt.Fprintf(p.Out, "Für %s statt %s (%s)\n",
		deal.NewPrice(), deal.OldPrice(), deal.PriceSource())
	fmt.Fprintf(p.Out, "Noch %s verfügbar\n", deal.Percentage())
	fmt.Fprintf(p.Out, "Nächster Deal am: %s (in %s)\n",
		deal.NextDeal().Format("Mon Jan _2 15:04:05"),
		strings.TrimSuffix(deal.NextDealIn().String(), "0s"))
}
