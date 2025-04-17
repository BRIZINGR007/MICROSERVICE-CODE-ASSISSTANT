package controllers

import (
	"log"
	"net/http"

	"github.com/BRIZINGR007/app-002-code-assistant/internal/db/models"
	"github.com/BRIZINGR007/app-002-code-assistant/internal/domain"
	"github.com/BRIZINGR007/app-002-code-assistant/internal/services"
	"github.com/BRIZINGR007/app-002-code-assistant/internal/utils"
	"github.com/BRIZINGR007/go-service-utils/helpers"
	"github.com/gin-gonic/gin"
)

func ExtractCode(c *gin.Context) {
	var req domain.RepoRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Request Format"})
		return
	}
	context := helpers.GetGinContextStringMap(c)
	codeBaseDataPayload := &models.CodeBaseData{
		CodeBaseId:   utils.GenerateUUID(),
		GitHubURL:    req.GitHubURL,
		Username:     req.Username,
		CodeBaseName: req.CodeBaseName,
		Token:        req.Token,
		Branch:       req.Branch,
		FolderPath:   req.FolderPath,
	}
	err := services.AddCodeBaseDataForUser(context["email"], codeBaseDataPayload)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "user does  not exist ."})
		return
	}
	fileContents := services.ExtractCodeContext(codeBaseDataPayload)
	log.Printf("File Contents :", fileContents)
	codeContextWithEmbeddings, err := services.PrepareDataForCodeContext(&fileContents, codeBaseDataPayload.CodeBaseId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "failed to  load  embeddings ."})
		return
	}

	if err := services.PopulateCodeContext(&codeContextWithEmbeddings); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "failed to  populate context . %v"})
		return
	}

	// err = queue.CodeAssistQueuePostMessageNF("add-codecontext", codeBaseDataPayload, context)
	// if err != nil {
	// 	c.JSON(http.StatusBadRequest, gin.H{"error": "Message Not Succesfully sent to Queue ."})
	// 	return
	// }
	c.JSON(http.StatusAccepted, gin.H{"success": "Proccessed ."})
}

func CodeBaseChat(c *gin.Context) {
	codeBaseId := c.Query("codeBaseId")
	query := c.Query("query")
	codeContext, err := services.CodeContextRecursiveRetriever(codeBaseId, query, nil, 0)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	answer, err := utils.FetchLLMResponse(query, codeContext)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Cannot Retrieve LLM Answer ."})
		return
	}
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Cannot Retrieve Bedrock Answer ."})
		return
	}
	context := helpers.GetGinContextHeadersStruct(c)
	chatPayload := models.Chat{
		ChatId:       utils.GenerateUUID(),
		UserID:       context.UserId,
		CodeBaseId:   codeBaseId,
		AIAnswer:     answer,
		UserQuestion: query,
	}
	err = services.AddChat(&chatPayload)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save chat", "details": err.Error()})
		return
	}
	c.JSON(http.StatusOK, chatPayload)

}
