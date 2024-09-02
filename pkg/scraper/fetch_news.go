package main

import (
	"fmt"
	"log"

	"github.com/gocolly/colly"
)

// Article represents the structure of the news article data.
type Article struct {
	Title   string // Title of the news article
	Content string // Content of the news article
}

// fetchNews fetches news articles from any news website.
//
// Parameters:
//   - url (string): The URL of the news website to scrape.
//   - titleSelector (string): The CSS selector to find article titles.
//   - contentSelector (string): The CSS selector to find article content.
//   - linkPrefix (string): The prefix to add to relative links.
//
// Returns:
//   - ([]Article, error): A slice of Article structs containing the title and content of news articles,
//     and an error if any occurs during the process.
func fetchNews(url, titleSelector, contentSelector, linkPrefix string) ([]Article, error) {
	var articles []Article

	// Create a new collector
	c := colly.NewCollector()

	// Set the User-Agent to mimic a web browser
	c.UserAgent = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/91.0.4472.124 Safari/537.36"

	// Set up a callback for when an HTML element with the specified titleSelector is found
	c.OnHTML(titleSelector, func(e *colly.HTMLElement) {
		title := e.Text
		link := e.Attr("href")
		if link != "" && linkPrefix != "" {
			fullLink := linkPrefix + link
			c.Visit(fullLink)
			articles = append(articles, Article{Title: title})
		} else {
			articles = append(articles, Article{Title: title})
		}
	})

	// Set up a callback for when an HTML element with the specified contentSelector is found
	c.OnHTML(contentSelector, func(e *colly.HTMLElement) {
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

// main is the entry point of the application. It fetches news articles from a specified news website
// and prints the title and content of each article.
//
// It calls the fetchNews function with the required parameters, handles any errors that occur,
// and outputs the results to the console.
func main() {
	// Example usage for New York Times
	url := "https://www.nytimes.com"
	titleSelector := ".css-66vf3i"          // CSS selector for the article title
	contentSelector := ".css-1fanzo5"       // CSS selector for the article content
	linkPrefix := "https://www.nytimes.com" // Prefix to complete the relative URLs

	articles, err := fetchNews(url, titleSelector, contentSelector, linkPrefix)
	if err != nil {
		log.Fatal(err)
	}

	for _, article := range articles {
		fmt.Printf("Title: %s\n", article.Title)
		fmt.Printf("Content: %s\n", article.Content)
		fmt.Println()
	}

	// Example usage for BBC
	bbcURL := "https://www.bbc.com/news"
	bbcTitleSelector := ".gs-c-promo-heading__title"
	bbcContentSelector := ".ssrcss-uf6wea-RichTextComponentWrapper"
	bbcLinkPrefix := "https://www.bbc.com"

	bbcArticles, err := fetchNews(bbcURL, bbcTitleSelector, bbcContentSelector, bbcLinkPrefix)
	if err != nil {
		log.Fatal(err)
	}

	for _, article := range bbcArticles {
		fmt.Printf("Title: %s\n", article.Title)
		fmt.Printf("Content: %s\n", article.Content)
		fmt.Println()
	}
}
