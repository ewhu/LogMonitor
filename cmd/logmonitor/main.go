// cmd/logmonitor/main.go
package main

import (
	"flag"
	"log"
	"os"

	"logmonitor/internal/logmonitor"
)

func main() {
	verbose := flag.Bool("verbose", false, "Enable verbose logging")
	flag.Parse()

	app := logmonitor.NewApp(*verbose)
	if err := app.Run(); err != nil {
		log.Fatal(err)
	}
}
