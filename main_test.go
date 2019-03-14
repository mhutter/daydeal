package main

import (
	"bytes"
	"regexp"
	"testing"
)

var (
	testRe = regexp.MustCompile("\n" +
		"    [^\n]+\n" +
		"    [^\n]+\n" +
		"\n" +
		"Für CHF [0-9.–]+ statt CHF [0-9.–]+ \\(\\d\\)\n" +
		"Noch \\d+% verfügbar\n" +
		"Nächster Deal am: \\w{3} \\w{3} \\d{2} \\d{2}:\\d{2}:\\d{2} \\(in \\d{2,3}:\\d{2}:\\d{2}\\)",
	)
)

func TestDay(t *testing.T) {
	testRun(t, urlDay)
}

func TestWeek(t *testing.T) {
	testRun(t, urlWeek)
}

func testRun(t *testing.T, url string) {
	buf := new(bytes.Buffer)

	run(buf, url)

	if !testRe.Match(buf.Bytes()) {
		t.Errorf("Output does not match RE:\n %s", buf)
	}
}
