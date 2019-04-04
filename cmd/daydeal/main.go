package main

import (
	"fmt"
	"os"

	"github.com/mhutter/daydeal"
	"gopkg.in/alecthomas/kingpin.v2"
)

var (
	version = "dev"
	commit  = "none"
	date    = "unknown"
)

func main() {
	app := kingpin.New("daydeal", "Display the current deal from daydeal.ch")
	week := app.Flag("week", "Display deal of the week instead").Short('w').Bool()

	app.Version(fmt.Sprintf("%v, commit %v, built at %v", version, commit, date))
	app.VersionFlag.Short('v')
	app.HelpFlag.Short('h')

	if _, err := app.Parse(os.Args[1:]); err != nil {
		fmt.Printf("Could not parse flags: %v\n", err)
		os.Exit(1)
	}

	url := daydeal.URLDaydealDay
	if *week {
		url = daydeal.URLDaydealWeek
	}
	daydeal.NewApp(os.Stdout, url).Run()
}
