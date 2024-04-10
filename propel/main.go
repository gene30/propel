package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/gocolly/colly"
)

type post struct {
	Title     string `json:"title"`
	Author    string `json:"author,omitempty"`
	Subreddit string `json:"subreddit"`
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
			p.Author = "old.reddit.com/" + author[1:]
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

	jwriter, err := os.Create("posts.json")
	if err != nil {
		debug(err.Error())
		return
	}
	defer jwriter.Close()

	encoder := json.NewEncoder(jwriter)
	encoder.SetIndent("", "  ")
	if err := encoder.Encode(posts); err != nil {
		debug(err.Error())
		return
	}

	fmt.Println("propel wrote to json")
}

func debug(message string) {
	f, err := os.OpenFile("text.log",
		os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		debug(err.Error())
		return
	}
	defer f.Close()

	logger := log.New(f, "Version 1.2 ", log.LstdFlags)
	logger.Println(message)
}
