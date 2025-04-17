package utils

import (
	"regexp"
	"strings"
)

func SplitStringIntoChunks(input string) []string {
	re := regexp.MustCompile(`[^\S\r\n]+`)
	input = re.ReplaceAllString(input, " ")
	// split the string  by newlines  ...
	lines := strings.Split(input, "\n")
	var result []string
	var currentChunk strings.Builder
	wordCount := 0

	for _, line := range lines {
		words := strings.Fields(line)
		lineWordCount := len(words)
		if wordCount+lineWordCount > 300 && wordCount > 0 {
			result = append(result, currentChunk.String())
			currentChunk.Reset()
			wordCount = 0
		}

		// Add line to current chunk
		if currentChunk.Len() > 0 {
			currentChunk.WriteString("\n")
		}
		currentChunk.WriteString(line)
		wordCount += lineWordCount
	}
	if currentChunk.Len() > 0 {
		result = append(result, currentChunk.String())
	}

	return result

}
