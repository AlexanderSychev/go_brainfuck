package main

import (
	"log"
	"os"
)

// VERSION is constant which contains application semantic version
const VERSION = "0.0.1"

func main() {
	app := newCLIApp()

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
