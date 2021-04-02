package main

import (
	"bufio"
	"fmt"
	"github.com/mvdan/xurls"
	"github.com/spf13/cobra"
	"io/ioutil"
	"log"
	"os"
)

var (
	pattern, _ = xurls.StrictMatchingScheme("https?://")
	verbose    = false
	prog       = &cobra.Command{
		Use:   "parseurls [*]",
		Short: "Command line tool for parse urls from text files or stdin",
		Run: func(c *cobra.Command, args []string) {
			err := run(c, args)
			if err != nil {
				log.Fatal(err)
			}
		},
	}
)

func init() {
	prog.Flags().SortFlags = false
	prog.PersistentFlags().SortFlags = false
	prog.PersistentFlags().BoolVarP(
		&verbose,
		"verbose",
		"v",
		false,
		"verbose",
	)
}

func run(c *cobra.Command, files []string) error {
	if len(files) == 0 {
		scan := bufio.NewScanner(os.Stdin)

		for scan.Scan() {
			text := scan.Text()
			printURLs(text)
		}

		return scan.Err()
	}

	for _, fp := range files {
		if verbose {
			log.Printf("processing %s", fp)
		}
		content, err := ioutil.ReadFile(fp)
		if err != nil {
			return err
		}
		printURLs(string(content))
	}

	return nil
}

func printURLs(text string) {
	if text == "" {
		return
	}
	for _, url := range pattern.FindAllString(text, -1) {
		_, _ = fmt.Fprintln(os.Stdout, url)
	}
}

func main() {
	_ = prog.Execute()
}