package daydeal_test

import (
	"bytes"
	"regexp"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/mhutter/daydeal"
)

var testRe = regexp.MustCompile(`\n` +
	`    [^\n]+\n` +
	`    [^\n]+\n` +
	`\n` +
	`Für CHF [0-9.–]+ statt CHF [0-9.–]+ \(\d\)\n` +
	`Noch \d+% verfügbar\n` +
	`Nächster Deal am: \w{3} \w{3} \d{2} \d{2}:\d{2}:\d{2} \(in (\d{1,3}h)?\d\d?m\)`,
)

func TestRunDay(t *testing.T) {
	t.Parallel()
	testRun(t, []string{})
}

func TestRunWeek(t *testing.T) {
	t.Parallel()
	testRun(t, []string{"-w"})
}

func testRun(t *testing.T, args []string) {
	buf := new(bytes.Buffer)
	daydeal.NewApp(buf).Run(args)
	assert.Regexp(t, testRe, buf)
}

func TestFetchDealUnknownKind(t *testing.T) {
	t.Parallel()
	_, err := daydeal.FetchDeal(daydeal.Kind(13))
	assert.EqualError(t, err, "Unknow deal 13")
}
