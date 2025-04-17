package utils

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"sync"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/bedrockruntime"
)

const (
	defaultRegion  = "us-east-1"
	defaultModelID = "meta.llama3-8b-instruct-v1:0"
)

var (
	client     *bedrockruntime.Client
	clientOnce sync.Once
	clientErr  error
)

// initBedrockRuntimeClient initializes the Bedrock runtime client once
func initBedrockRuntimeClient() (*bedrockruntime.Client, error) {
	clientOnce.Do(func() {
		region := os.Getenv("AWS_REGION")
		if region == "" {
			region = defaultRegion
		}

		cfg, err := config.LoadDefaultConfig(context.Background(),
			config.WithRegion(region),
			config.WithCredentialsProvider(credentials.NewStaticCredentialsProvider(
				os.Getenv("AWS_ACCESS_KEY_ID"),
				os.Getenv("AWS_SECRET_ACCESS_KEY"),
				"",
			)),
		)
		if err != nil {
			clientErr = err
			return
		}
		client = bedrockruntime.NewFromConfig(cfg)
	})
	return client, clientErr
}

// LlamaResponse represents possible response formats from the Llama3 model
type LlamaResponse struct {
	Generation string `json:"generation"`
	Completion string `json:"completion"`
	OutputText string `json:"outputText"`
	Results    []struct {
		OutputText string `json:"outputText"`
	} `json:"results"`
}

// CallLlama3 invokes the Llama3 model via AWS Bedrock and returns the generated text
func CallLlama3(ctx context.Context, payload LlamaPayload, modelID string) (string, error) {
	// Initialize client if not already done
	client, err := initBedrockRuntimeClient()
	if err != nil {
		return "", fmt.Errorf("bedrock client init failed: %w", err)
	}

	if modelID == "" {
		modelID = defaultModelID
	}

	// Marshal payload directly to bytes
	inputBytes, err := json.Marshal(payload)
	if err != nil {
		return "", fmt.Errorf("payload marshaling failed: %w", err)
	}

	// Invoke the model
	resp, err := client.InvokeModel(ctx, &bedrockruntime.InvokeModelInput{
		ModelId:     aws.String(modelID),
		ContentType: aws.String("application/json"),
		Body:        inputBytes,
	})
	if err != nil {
		return "", fmt.Errorf("model invocation failed: %w", err)
	}

	// Parse response
	var response LlamaResponse
	if err := json.Unmarshal(resp.Body, &response); err != nil {
		return "", fmt.Errorf("response decoding failed: %w", err)
	}

	// Extract text from whichever field exists
	if text := extractResponseText(response); text != "" {
		return text, nil
	}

	// If we got here, the response format was unexpected
	raw, _ := json.MarshalIndent(response, "", "  ")
	return "", fmt.Errorf("unexpected response format: %s", raw)
}

// extractResponseText checks all possible response fields for the generated text
func extractResponseText(r LlamaResponse) string {
	switch {
	case r.Generation != "":
		return r.Generation
	case r.Completion != "":
		return r.Completion
	case r.OutputText != "":
		return r.OutputText
	case len(r.Results) > 0 && r.Results[0].OutputText != "":
		return r.Results[0].OutputText
	default:
		return ""
	}
}
