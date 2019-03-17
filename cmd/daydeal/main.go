package main

import (
	"os"

	"github.com/mhutter/daydeal"
)

func main() {
	daydeal.NewApp(os.Stdout).Run(os.Args[1:])
}
