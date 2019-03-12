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
		"Noch \\d+% verfügbar",
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
