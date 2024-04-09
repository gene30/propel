package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gocolly/colly"
)

func main() {
	fmt.Println("Propel Launched")
	debug("Propel Launched.")
	cheese, err := os.ReadFile("website.txt")
	final := "https://old.reddit.com/r/" + string(cheese)
	debug(final)
	if err != nil {
		debug(err.Error())
	}
	c := colly.NewCollector()
	colly.AllowedDomains("old.reddit.com")
	c.OnHTML("div.top-matter", func(h *colly.HTMLElement) {

		fmt.Println("title:", h.ChildText("a[tabindex]"))
	})
	c.OnHTML("p.tagline", func(h *colly.HTMLElement) {

		fmt.Println("author:", h.ChildText("a[href]"))
	})

	c.Visit(final)
}

func debug(message string) {
	f, err := os.OpenFile("text.log",
		os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		debug(err.Error())
	}
	defer f.Close()

	logger := log.New(f, "Version 1.0b ", log.LstdFlags)
	logger.Println(message)
}
