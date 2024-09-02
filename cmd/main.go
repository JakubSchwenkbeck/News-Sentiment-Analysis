package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

func main() {
	text := map[string]string{"text": "The service was truly a great and amazing!"}
	jsonData, _ := json.Marshal(text)

	resp, err := http.Post("http://localhost:5000/vader", "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer resp.Body.Close()

	var result map[string]interface{}
	json.NewDecoder(resp.Body).Decode(&result)

	fmt.Println("Sentiment Analysis:", result)
	fmt.Println(interpretSentiment(result))

}

// interpretSentiment interprets sentiment scores and returns a human-readable sentiment description.
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
