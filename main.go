package main

import (
	"fmt"
	"log"
	"os/exec"
)

func main() {
	text := "The movie was fantastic!"
	cmd := exec.Command("python3", "vader_sentiment.py", text)
	output, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatalf("Error: %v\nOutput: %s\n", err, string(output))
	}
	fmt.Printf("Sentiment Analysis: %s\n", output)
}
