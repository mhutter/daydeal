package daydeal_test

import (
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/PuerkitoBio/goquery"
	"github.com/mhutter/daydeal"
)

const html = `
<span class="product-description__title1">Gopher</span>
<span class="product-description__title2">Must-Have!</span>
<div class="product-pricing__prices col col-6 col-sm-auto">
	<h2 class="product-pricing__prices-new-price js-deal-price">CHF 199.–</h2>
	<strong class="product-pricing__prices-old-price">
		statt
		<span class="js-old-price">CHF 299.–</span>
		<span class="product-pricing__prices-old-price-annotation">2</span>
	</strong>
</div>
<strong class="product-progress__availability">88%</strong>
<span class="js-clock" data-next-deal="2019-03-04 08:00:00"></span>
`

func TestDaydeal(t *testing.T) {
	t.Parallel()
	doc, _ := goquery.NewDocumentFromReader(strings.NewReader(html))
	d := daydeal.NewDaydeal(doc)
	assert.Equal(t, "Gopher", d.Title())
	assert.Equal(t, "Must-Have!", d.Subtitle())
	assert.Equal(t, "CHF 199.–", d.NewPrice())
	assert.Equal(t, "CHF 299.–", d.OldPrice())
	assert.Equal(t, "2", d.PriceSource())
	assert.Equal(t, "88%", d.Percentage())
	assert.Equal(t, time.Date(2019, 3, 4, 8, 0, 0, 0, time.UTC), d.NextDeal().UTC())
	assert.IsType(t, time.Duration(0), d.NextDealIn())
}
