package daydeal_test

import (
	"bytes"
	"io"
	"regexp"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/mhutter/daydeal"
)

var testRe = []*regexp.Regexp{
	regexp.MustCompile(`^\n$`),
	regexp.MustCompile(`^    [^\n]+\n$`),
	regexp.MustCompile(`^    [^\n]+\n$`),
	regexp.MustCompile(`^\n$`),
	regexp.MustCompile(`^Für CHF [0-9.–]+ statt CHF [0-9.–]+ \(\d\)\n$`),
	regexp.MustCompile(`^Noch \d+% verfügbar\n$`),
	regexp.MustCompile(`^Nächster Deal am: \w{3} \w{3} [ \d]\d \d{2}:\d{2}:\d{2} \(in (\d{1,3}h)?\d\d?m\)\n$`),
}

func TestRunDay(t *testing.T) {
	t.Parallel()
	testRun(t, daydeal.URLDaydealDay)
}

func TestRunWeek(t *testing.T) {
	t.Parallel()
	testRun(t, daydeal.URLDaydealWeek)
}

func testRun(t *testing.T, url string) {
	buf := new(bytes.Buffer)

	assert.Nil(t, daydeal.NewApp(buf, url).Run())

	for _, re := range testRe {
		line, err := buf.ReadString('\n')
		assert.Nil(t, err)
		assert.Regexp(t, re, line)
	}

	// Ensure we're at the end of the buffer
	_, err := buf.ReadString('\n')
	assert.Equal(t, io.EOF, err)
}

func TestFetchDealUnknownKind(t *testing.T) {
	t.Parallel()
	buf := new(bytes.Buffer)
	err := daydeal.NewApp(buf, "does absolutely not exist").Run()
	assert.Error(t, err)
}
