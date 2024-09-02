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

// fetchBBCNews fetches news articles from the BBC News website.
//
// It creates a new Colly collector, sets the User-Agent to mimic a web browser,
// and defines callbacks for HTML elements to extract article titles and content.
// It visits the specified URL and returns a slice of Article structs and any error encountered.
//
// Parameters:
//   - url (string): The URL of the BBC News website to scrape.
//
// Returns:
//   - ([]Article, error): A slice of Article structs containing the title and content of news articles,
//     and an error if any occurs during the process.
func fetchBBCNews(url string) ([]Article, error) {
	var articles []Article

	// Create a new collector
	c := colly.NewCollector()

	// Set the User-Agent to mimic a web browser
	//c.UserAgent = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/91.0.4472.124 Safari/537.36"

	// Set up a callback for when an HTML element with class 'gs-c-promo-heading' is found
	c.OnHTML(".gs-c-promo-heading", func(e *colly.HTMLElement) {
		title := e.Text
		link := e.Attr("href")
		if link != "" {
			fullLink := "https://www.bbc.com" + link
			c.Visit(fullLink)
			articles = append(articles, Article{Title: title})
		}
	})

	// Set up a callback for when an HTML element with class 'ssrcss-uf6wea-RichTextComponentWrapper' is found
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

// main is the entry point of the application. It fetches news articles from the BBC News website
// and prints the title and content of each article.
//
// It calls the fetchBBCNews function with the BBC News URL, handles any errors that occur,
// and outputs the results to the console.
func main() {
	url := "https://www.nytimes.com"
	articles, err := fetchBBCNews(url)
	if err != nil {
		log.Fatal(err)
	}

	//fmt.Print(articles)
	//fmt.Println("here")
	for _, article := range articles {
		fmt.Printf("Title: %s\n", article.Title)
		fmt.Printf("Content: %s\n", article.Content)
		fmt.Println()
	}
}
