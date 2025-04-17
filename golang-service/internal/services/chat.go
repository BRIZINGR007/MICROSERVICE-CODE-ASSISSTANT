package services

import (
	"context"

	"github.com/BRIZINGR007/app-002-code-assistant/internal/db/models"
	"github.com/BRIZINGR007/app-002-code-assistant/internal/db/repositories"
)

func AddChat(chat *models.Chat) error {
	repo := repositories.GetChatRepository()

	err := repo.AddChat(context.Background(), chat)
	return err
}

func GetAllChatsForUserByCodeBase(userId string, codeBaseId string) ([]models.Chat, error) {
	repo := repositories.GetChatRepository()
	all_chats, err := repo.GetAllChats(context.Background(), userId, codeBaseId)
	if err != nil {
		return nil, err
	}
	return all_chats, nil

}
