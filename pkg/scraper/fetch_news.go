package main

import (
	"fmt"
	"log"

	"github.com/gocolly/colly"
)

// Article represents the structure of the news article data.
type Article struct {
	Title   string
	Content string
}

// fetchBBCNews fetches news articles from the BBC News website.
func fetchBBCNews(url string) ([]Article, error) {
	var articles []Article

	// Create a new collector
	c := colly.NewCollector()

	// Set up a callback for when a HTML element is found
	c.OnHTML(".gs-c-promo-heading", func(e *colly.HTMLElement) {
		title := e.Text
		link := e.Attr("href")
		if link != "" {
			fullLink := "https://www.bbc.com" + link
			c.Visit(fullLink)
			articles = append(articles, Article{Title: title})
		}
	})

	// Set up a callback for when a HTML element is found
	c.OnHTML(".ssrcss-uf6wea-RichTextComponentWrapper", func(e *colly.HTMLElement) {
		content := e.Text
		if len(articles) > 0 {
			articles[len(articles)-1].Content = content
		}
	})

	// Set up error handling
	c.OnError(func(r *colly.Response, err error) {
		log.Println("Request URL:", r.Request.URL, "failed with response:", r, "\nError:", err)
	})

	// Start the scraping
	err := c.Visit(url)
	if err != nil {
		return nil, err
	}

	return articles, nil
}

func FetchingWrapper() {
	url := "https://www.bbc.com/news"
	articles, err := fetchBBCNews(url)
	if err != nil {
		log.Fatal(err)
	}

	for _, article := range articles {
		fmt.Printf("Title: %s\n", article.Title)
		fmt.Printf("Content: %s\n", article.Content)
		fmt.Println()
	}
}
