package utils

import (
	"fmt"
	"strings"

	"github.com/BRIZINGR007/app-002-code-assistant/internal/db/models"
)

type LlamaPayload struct {
	Prompt      string  `json:"prompt"`
	MaxGenLen   int     `json:"max_gen_len"`
	Temperature float64 `json:"temperature"`
	TopP        float64 `json:"top_p"`
}

func GetLlamaPayloadForBedrock(userPrompt string, inputTokens int, systemPrompt string) LlamaPayload {
	specialTokens := map[string]string{
		"bos_token":           "<|begin_of_text|>",
		"content_start_token": "<|start_header_id|>%s<|end_header_id|>",
		"content_end_token":   "<|eot_id|>",
		"eos_token":           "<|end_of_text|>",
	}

	structuredPrompt := specialTokens["bos_token"]

	if systemPrompt == "" {
		structuredPrompt += fmt.Sprintf(specialTokens["content_start_token"], "system")
		structuredPrompt += systemPrompt
		structuredPrompt += specialTokens["content_end_token"]
	}

	structuredPrompt += fmt.Sprintf(specialTokens["content_start_token"], "user")
	structuredPrompt += userPrompt
	structuredPrompt += specialTokens["content_end_token"]
	structuredPrompt += "<|start_header_id|>assistant<|end_header_id|>"

	return LlamaPayload{
		Prompt:      structuredPrompt,
		MaxGenLen:   256,
		Temperature: 0.6,
		TopP:        0.6,
	}
}

func GetCodeAssistPrompt(userQuery string, codeContexts []models.CodeContext) string {
	codeAssistBasePrompt := `You are an AI programming assistant tasked with maintaining a codebase.
You will be provided with multiple code snippets along with their file paths. Carefully analyze the code and answer the user query using only the information from the provided snippets.
Only return your answer.

User Query: %s

Code Snippets: %s

Answer:`

	var codeChunksBuilder strings.Builder
	codeChunksBuilder.WriteString("[")
	for i, ctx := range codeContexts {
		codeChunk := fmt.Sprintf(`{"file_path": "%s", "code_chunk": "%s"}`, ctx.FilePath, ctx.Code)
		codeChunksBuilder.WriteString(codeChunk)
		if i != len(codeContexts)-1 {
			codeChunksBuilder.WriteString(", ")
		}
	}
	codeChunksBuilder.WriteString("]")

	finalPrompt := fmt.Sprintf(codeAssistBasePrompt, userQuery, codeChunksBuilder.String())
	return finalPrompt
}
