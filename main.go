package main

import (
	"log"
	"os"

	"github.com/gonejack/parseurls/cmd"
)

func main() {
	log.SetOutput(os.Stdout)

	var c cmd.ParseURLs
	if e := c.Run(); e != nil {
		log.Fatal(e)
	}
}
