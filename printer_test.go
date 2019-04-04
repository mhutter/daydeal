package daydeal_test

import (
	"bytes"
	"testing"
	"time"

	"github.com/mhutter/daydeal"
	"github.com/stretchr/testify/assert"
)

type tp struct {
	VTitle       string
	VSubtitle    string
	VNewPrice    string
	VOldPrice    string
	VPriceSource string
	VPercentage  string
	VNextDeal    time.Time
	VNextDealIn  time.Duration
}

var deal = tp{
	VTitle:       "TestTitle",
	VSubtitle:    "TestSubtitle",
	VNewPrice:    "TestNewPrice",
	VOldPrice:    "TestOldPrice",
	VPriceSource: "TestPriceSource",
	VPercentage:  "TestPercentage",
	VNextDeal:    time.Date(2019, 3, 5, 8, 0, 0, 0, time.UTC),
	VNextDealIn:  13 * time.Hour,
}

var expectedOutput = `
    TestTitle
    TestSubtitle

Für TestNewPrice statt TestOldPrice (TestPriceSource)
Noch TestPercentage verfügbar
Nächster Deal am: Tue Mar  5 08:00:00 (in 13h0m)
`

func TestPrintTo(t *testing.T) {
	t.Parallel()
	buf := new(bytes.Buffer)

	daydeal.NewPrinter(buf).Print(deal)

	assert.Equal(t, expectedOutput, buf.String())
}

func (t tp) Title() string             { return t.VTitle }
func (t tp) Subtitle() string          { return t.VSubtitle }
func (t tp) NewPrice() string          { return t.VNewPrice }
func (t tp) OldPrice() string          { return t.VOldPrice }
func (t tp) PriceSource() string       { return t.VPriceSource }
func (t tp) Percentage() string        { return t.VPercentage }
func (t tp) NextDeal() time.Time       { return t.VNextDeal }
func (t tp) NextDealIn() time.Duration { return t.VNextDealIn }
