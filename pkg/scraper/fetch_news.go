package scraper

import (
	"fmt"
	"log"
	"strings"

	"github.com/gocolly/colly"
)

// Article represents the structure of the news article data.
type Article struct {
	Title   string // Title of the news article
	Content string // Content of the news article
}

// fetchGenericNews attempts to scrape news articles from any news website by guessing common selectors.
//
// Parameters:
//   - url (string): The URL of the website to scrape.
//
// Returns:
//   - ([]Article, error): A slice of Article structs containing the title and content of news articles,
//     and an error if any occurs during the process.
func fetchGenericNews(url string) ([]Article, error) {
	var articles []Article

	// Create a new collector
	c := colly.NewCollector(
		colly.AllowedDomains("www.nytimes.com", "www.bbc.com", "www.cnn.com"),                                                                  // Add other domains as needed
		colly.UserAgent("Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/91.0.4472.124 Safari/537.36"), // Mimic browser user agent
	)

	// Find article titles by guessing common title tags
	c.OnHTML("h1, h2, h3, .headline, .title", func(e *colly.HTMLElement) {
		// Guess this is an article title
		title := e.Text

		// Only continue if the title is not empty
		if strings.TrimSpace(title) == "" {
			return
		}

		// Attempt to visit the link if it exists
		link := e.Attr("href")
		if link != "" {
			if link[0] == '/' {
				link = e.Request.AbsoluteURL(link)
			}
			c.Visit(link)
		}

		// Create a new article with the found title
		articles = append(articles, Article{Title: title})
	})

	// Find article content by guessing common content tags
	c.OnHTML("p, .content, .article-body, .story-body, .ssrcss-uf6wea-RichTextComponentWrapper", func(e *colly.HTMLElement) {
		content := e.Text
		if len(articles) > 0 && strings.TrimSpace(content) != "" {
			// Append content to the last article in the list
			articles[len(articles)-1].Content += content + "\n"
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

// main is the entry point of the application.
func Fetchmain() []Article {
	NYurl := "https://www.nytimes.com"
	NYarticles, err := fetchGenericNews(NYurl)
	if err != nil {
		log.Fatal(err)
	}

	for _, article := range NYarticles {
		fmt.Printf("Title: %s\n", article.Title)
		fmt.Printf("Content: %s\n", article.Content)
		fmt.Println()
	}
	fmt.Println("\n \n \n \n  BBC:")

	BBCurl := "https://www.bbc.com"
	BBCarticles, err := fetchGenericNews(BBCurl)
	if err != nil {
		log.Fatal(err)
	}

	for _, article- := range BBCarticles {
		fmt.Printf("Title: %s\n", article.Title)
		fmt.Printf("Content: %s\n", article.Content)
		fmt.Println()
	}
	// Combine both slices of articles
	combinedArticles := append(NYarticles, BBCarticles...)
	return combinedArticles

}
