package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

func main() {
	text := map[string]string{"text": "The service was amazing!"}
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
}
