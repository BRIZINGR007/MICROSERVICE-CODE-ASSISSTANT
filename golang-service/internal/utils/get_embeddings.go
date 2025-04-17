package utils

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
)

// fetchEmbedding makes the GET request and returns a slice of float32
func FetchEmbedding(query string) ([]float32, error) {
	baseURL := os.Getenv("BASE_URL")
	if baseURL == "" {
		log.Fatal("BASE_URL is not set in the environment")
	}

	embeddingURL := baseURL + "/api/app002-ai-service/encoder/get-embedding"

	// Add query parameter
	params := url.Values{}
	params.Add("query", query)

	// Construct full URL
	fullURL := fmt.Sprintf("%s?%s", embeddingURL, params.Encode())

	// Create request
	req, err := http.NewRequest("GET", fullURL, nil)
	log.Printf("ERRORR  ", err)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}
	req.Header.Set("Accept", "application/json")

	// Perform request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error performing request: %w", err)
	}
	defer resp.Body.Close()

	// Read response
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response: %w", err)
	}

	// Parse JSON response
	var embedding []float32
	if err := json.Unmarshal(body, &embedding); err != nil {
		return nil, fmt.Errorf("error parsing JSON: %w", err)
	}

	return embedding, nil
}
