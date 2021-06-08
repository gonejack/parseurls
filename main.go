package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"net/url"
	"os"
	"strings"

	"github.com/mvdan/xurls"
	"github.com/spf13/cobra"
)

var (
	verbose = false
	weibo   = false
	prog    = &cobra.Command{
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

	flags := prog.PersistentFlags()
	{
		flags.SortFlags = false
		flags.BoolVarP(&weibo, "weibo", "", false, "weibo special treat")
		flags.BoolVarP(&verbose, "verbose", "v", false, "verbose")
	}
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

	for _, ref := range xurls.Strict.FindAllString(text, -1) {
		if !strings.HasPrefix(ref, "http") {
			continue
		}
		if weibo {
			u, err := url.Parse(ref)
			if err == nil && strings.Contains(u.Host, "weibo") {
				u.Host = "m.weibo.cn"
				ref = u.String()
			}
		}
		_, _ = fmt.Fprintln(os.Stdout, ref)
	}
}
func main() {
	_ = prog.Execute()
}
