package main

import (
	"log"
	"os"

	"github.com/gonejack/parseurls/cmd"
)

func init() {
	log.SetOutput(os.Stdout)
}

func main() {
	err := new(cmd.ParseURLs).Run()
	if err != nil {
		log.Fatal(err)
	}
}
