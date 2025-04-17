package controllers

import (
	"net/http"

	"github.com/BRIZINGR007/app-002-code-assistant/internal/services"
	"github.com/BRIZINGR007/go-service-utils/helpers"
	"github.com/gin-gonic/gin"
)

func GetAllChats(c *gin.Context) {
	codeBaseId := c.Query("codeBaseId")
	if codeBaseId == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Missing required codeBaseId parameter"})
		return
	}

	context := helpers.GetGinContextHeadersStruct(c)
	allChats, err := services.GetAllChatsForUserByCodeBase(context.UserId, codeBaseId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve chats", "details": err.Error()})
		return
	}

	c.JSON(http.StatusOK, allChats)
}
