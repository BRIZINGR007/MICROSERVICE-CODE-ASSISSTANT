package utils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/BRIZINGR007/app-002-code-assistant/internal/db/models"
)

type Context struct {
	Code     string `json:"Code"`
	FilePath string `json:"FilePath"`
}
type LLMRequest struct {
	Query    string    `json:"query"`
	Contexts []Context `json:"contexts"`
}

// FetchLLMResponse makes a POST request to the LLM API and returns the response
func FetchLLMResponse(userQuery string, codeContexts []models.CodeContext) (string, error) {
	baseURL := os.Getenv("BASE_URL")
	llmURL := baseURL + "/api/app002-ai-service/decoder/get-llmresponse"

	// Convert from models.CodeContext to the simpler Context structure required by the API
	contexts := make([]Context, 0, len(codeContexts))
	for _, ctx := range codeContexts {
		contexts = append(contexts, Context{
			Code:     ctx.Code,
			FilePath: ctx.FilePath,
		})
	}

	// Create the request body
	requestBody := LLMRequest{
		Query:    userQuery,
		Contexts: contexts,
	}

	// Marshal the request body to JSON
	jsonData, err := json.Marshal(requestBody)
	if err != nil {
		return "", fmt.Errorf("error marshaling request data: %w", err)
	}

	// Create the HTTP request
	req, err := http.NewRequest("POST", llmURL, bytes.NewBuffer(jsonData))
	if err != nil {
		return "", fmt.Errorf("error creating request: %w", err)
	}

	// Set headers
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")

	// Send the request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf("error sending request: %w", err)
	}
	defer resp.Body.Close()

	// Check response status
	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	// Read and parse the response
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("error reading response body: %w", err)
	}

	// You may want to parse the JSON response into a struct
	// For now, returning the raw response as a string
	return string(respBody), nil
}
