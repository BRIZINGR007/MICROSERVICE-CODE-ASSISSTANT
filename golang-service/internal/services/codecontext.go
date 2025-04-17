package services

import (
	"context"
	"errors"
	"fmt"
	"log"

	"github.com/BRIZINGR007/app-002-code-assistant/internal/db/models"
	"github.com/BRIZINGR007/app-002-code-assistant/internal/db/repositories"
	"github.com/BRIZINGR007/app-002-code-assistant/internal/domain"
	"github.com/BRIZINGR007/app-002-code-assistant/internal/utils"
)

func ExtractCodeContext(payload *models.CodeBaseData) []domain.FileContent {
	fileContents, err := utils.ExtractCodeService(payload)
	if err != nil {
		fmt.Printf("Error in reading  FileContents ...")
		return nil
	}
	processedContents := make([]domain.FileContent, len(fileContents))
	for i, content := range fileContents {
		processedContent := content
		chunks := utils.SplitStringIntoChunks(content.Code)
		processedContent.Chunks = chunks
		processedContents[i] = processedContent
	}
	return processedContents
}

func PopulateCodeContext(codeContexts *[]models.CodeContext) error {
	repo := repositories.GetCodeAssistRepository()
	err := repo.BulkInsertCodeContext(context.Background(), codeContexts)
	if err != nil {
		return err
	}
	return nil
}

func PrepareDataForCodeContext(filecontents *[]domain.FileContent, codeBaseId string) ([]models.CodeContext, error) {
	var contexts []models.CodeContext
	for _, file := range *filecontents {
		for _, chunk := range file.Chunks {
			embedding, err := utils.FetchEmbedding(chunk)
			if err != nil {
				return nil, err
			}
			context := models.CodeContext{
				VectorId:   utils.GenerateUUID(),
				CodeBaseId: codeBaseId,
				HashId:     utils.GenerateHashFromString(chunk),
				FilePath:   file.FilePath,
				Code:       chunk,
				Embedding:  embedding,
			}
			contexts = append(contexts, context)
		}
	}
	return contexts, nil

}

func CodeContextRecursiveRetriever(codeBaseId string, query string, accumulated []models.CodeContext, recurseCount int) ([]models.CodeContext, error) {
	if recurseCount >= 5 {
		return accumulated, nil
	}
	repo := repositories.GetCodeAssistRepository()
	embedding, err := utils.FetchEmbedding(query)
	if err != nil {
		log.Printf("failed to get embeddings for app-002-ai-service")
		return nil, errors.New("failed to get embeddings for app-002-ai-service")
	}

	var vectors_ids []string
	if len(accumulated) > 0 {
		for _, context := range accumulated {
			vectors_ids = append(vectors_ids, context.VectorId)
		}
	}
	newContexts, err := repo.CodeContextRetriever(context.Background(), codeBaseId, embedding, vectors_ids)
	if err != nil {
		return nil, err
	}
	accumulated = append(accumulated, newContexts[0])
	if len(accumulated) < 5 {
		recurseCount += 1
		return CodeContextRecursiveRetriever(codeBaseId, query, accumulated, recurseCount)
	}

	return accumulated, nil
}
