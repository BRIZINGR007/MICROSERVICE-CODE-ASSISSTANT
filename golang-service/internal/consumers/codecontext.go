package consumers

import (
	"encoding/json"
	"log"
	"strings"

	"github.com/BRIZINGR007/app-002-code-assistant/internal/db/models"
	"github.com/BRIZINGR007/app-002-code-assistant/internal/domain"
	"github.com/BRIZINGR007/app-002-code-assistant/internal/services"
)

func AddCodeContext(payload []byte) error {
	var codebData models.CodeBaseData

	decoder := json.NewDecoder(strings.NewReader(string(payload)))
	decoder.DisallowUnknownFields()

	if err := decoder.Decode(&codebData); err != nil {
		log.Printf("Invalid payload: %v", err)
		return err
	}

	if err := domain.ValidateCodeBaseData(&codebData); err != nil {
		log.Printf("Validation failed: %v", err)
		return err
	}

	fileContents := services.ExtractCodeContext(&codebData)

	codeContextWithEmbeddings, err := services.PrepareDataForCodeContext(&fileContents, codebData.CodeBaseId)
	if err != nil {
		log.Printf("Failed to generate embeddings: %v", err)
		return err
	}

	if err := services.PopulateCodeContext(&codeContextWithEmbeddings); err != nil {
		log.Printf("Error during insert: %v", err)
		return err
	}

	log.Println("Successfully inserted all the documents")
	return nil
}
