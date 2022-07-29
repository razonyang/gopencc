package main

import (
	"log"
	"os"

	"github.com/razonyang/gopencc/cmd"
)

func main() {
	app := cmd.New()
	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
