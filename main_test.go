package main

import (
	"bytes"
	"regexp"
	"testing"
)

func TestRun(t *testing.T) {
	testRe := regexp.MustCompile("\n" +
		"    [^\n]+\n" +
		"    [^\n]+\n" +
		"\n" +
		"Für CHF [0-9.–]+ statt CHF [0-9.–]+ \\(\\d\\)\n" +
		"Noch \\d+% verfügbar",
	)
	buf := new(bytes.Buffer)

	run(buf)

	if !testRe.Match(buf.Bytes()) {
		t.Errorf("Output does not match RE:\n %s", buf)
	}
}
