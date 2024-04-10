package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"

	"github.com/gocolly/colly"
)

type post struct {
	Title     string `csv:"title"`
	Author    string `csv:"author"`
	Subreddit string `csv:"subreddit"`
}

func main() {
	debug("Propel Launched.")

	cheese, err := os.ReadFile("website.txt")
	if err != nil {
		debug(err.Error())
		return
	}
	subreddit := string(cheese)

	final := "https://old.reddit.com/r/" + subreddit + "/new"

	debug(final)

	c := colly.NewCollector()
	colly.AllowedDomains("old.reddit.com")

	var posts []post

	c.OnHTML("div.top-matter", func(h *colly.HTMLElement) {
		title := h.ChildText("a[tabindex]")
		p := post{
			Title:     title,
			Subreddit: subreddit,
		}

		if author := h.ChildAttr("a[tabindex]", "href"); author != "" {
			p.Author = author
		}

		posts = append(posts, p)
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting:", r.URL)
	})

	err = c.Visit(final)
	if err != nil {
		debug(err.Error())
		return
	}

	csvit(posts)
}

func debug(message string) {
	f, err := os.OpenFile("text.log",
		os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		debug(err.Error())
		return
	}
	defer f.Close()

	logger := log.New(f, "Version 1.1b ", log.LstdFlags)
	logger.Println(message)
}

func csvit(posts []post) {
	j := csv.NewWriter(os.Stdout)
	defer j.Flush()

	if err := j.Write([]string{}); err != nil {
		debug(err.Error())
		return
	}

	for _, p := range posts {
		record := []string{p.Title, p.Author}
		supurb(record)
	}
}

func supurb(message []string) {
	f, err := os.OpenFile("result.log",
		os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		debug(err.Error())
		return
	}
	defer f.Close()

	logger := log.New(f, "", log.Flags())
	logger.Println(message)
}
