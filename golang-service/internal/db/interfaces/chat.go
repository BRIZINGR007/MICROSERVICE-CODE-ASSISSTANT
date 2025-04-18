package interfcaes

import (
	"context"

	"github.com/BRIZINGR007/app-002-code-assistant/internal/db/models"
	"github.com/BRIZINGR007/app-002-code-assistant/internal/db/repositories"
)

type ChatRepositoryInterface interface {
	ChatIndexes(ctx context.Context) error
	AddChat(ctx context.Context, chat *models.Chat) error
	GetAllChats(ctx context.Context, userId string, codeBaseId string) ([]models.Chat, error)
	DeleteChatsByCodeBaseId(ctx context.Context, codeBaseId string) error
}

var _ ChatRepositoryInterface = (*repositories.ChatRepository)(nil)
