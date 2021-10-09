package cmd

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"

	"github.com/alecthomas/kong"
	"github.com/mvdan/xurls"
)

type options struct {
	Verbose bool `help:"Verbose printing."`

	Txt []string `arg:"" optional:""`
}

type ParseURLs struct {
	options
}

func (c *ParseURLs) Run() (err error) {
	kong.Parse(&c.options,
		kong.Name("parseurls"),
		kong.Description("Command line tool to parse urls from text files or stdin"),
		kong.UsageOnError(),
	)

	if len(c.Txt) == 0 {
		scan := bufio.NewScanner(os.Stdin)
		for scan.Scan() {
			text := scan.Text()
			printURLs(text)
		}
		return scan.Err()
	}

	for _, txt := range c.Txt {
		if c.Verbose {
			log.Printf("processing %s", txt)
		}
		content, err := ioutil.ReadFile(txt)
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

	for _, ref := range xurls.Strict.FindAllString(text, -1) {
		if !strings.HasPrefix(ref, "http") {
			continue
		}
		_, _ = fmt.Fprintln(os.Stdout, ref)
	}
}
