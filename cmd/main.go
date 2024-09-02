package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

// main is the entry point of the application. It sends a text for sentiment analysis
// to both TextBlob and VADER endpoints, then prints and interprets the results.
func main() {
	text := map[string]string{"text": "The service was truly a great and amazing!"}
	jsonData, _ := json.Marshal(text)

	// Analyze using TextBlob
	textblobResult, err := analyzeSentiment("http://localhost:5000/textblob", jsonData)
	if err != nil {
		fmt.Println("TextBlob Error:", err)
		return
	}

	// Analyze using VADER
	vaderResult, err := analyzeSentiment("http://localhost:5000/vader", jsonData)
	if err != nil {
		fmt.Println("VADER Error:", err)
		return
	}

	// Print results from both analyzers
	fmt.Println("TextBlob Sentiment Analysis:")
	fmt.Println(textblobResult)
	fmt.Println(interpretSentiment(textblobResult))
	fmt.Println()

	fmt.Println("VADER Sentiment Analysis:")
	fmt.Println(vaderResult)
	fmt.Println(interpretSentiment(vaderResult))
}

// analyzeSentiment sends a JSON-encoded text to the specified sentiment analysis service
// and returns the analysis result as a map. It returns an error if the request fails.
func analyzeSentiment(url string, jsonData []byte) (map[string]interface{}, error) {
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}

	return result, nil
}

// interpretSentiment interprets the sentiment scores from the sentiment analysis service
// and returns a human-readable description of the sentiment based on the compound score.
// It returns "Positive", "Negative", or "Neutral" based on the value of the compound score.
func interpretSentiment(sentimentScores map[string]interface{}) string {
	// Extracting and converting the values from the map
	compound, ok := sentimentScores["compound"].(float64)
	if !ok {
		return "Error: Unable to interpret sentiment scores."
	}

	// Determining the sentiment based on the compound score
	var sentiment string
	if compound >= 0.05 {
		sentiment = "Positive"
	} else if compound <= -0.05 {
		sentiment = "Negative"
	} else {
		sentiment = "Neutral"
	}

	return fmt.Sprintf("Sentiment Analysis Result:\n- Compound Score: %.4f\n- Sentiment: %s", compound, sentiment)
}
