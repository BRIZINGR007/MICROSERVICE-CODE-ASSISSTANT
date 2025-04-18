package repositories

import (
	"context"
	"time"

	"github.com/BRIZINGR007/app-002-code-assistant/internal/db/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type ChatRepository struct {
	Collection *mongo.Collection
}

func (r *ChatRepository) ChatIndexes(ctx context.Context) error {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	codeBaseIdIndex := mongo.IndexModel{
		Keys: bson.M{"code_base_id": 1},
	}

	compoundIndex1 := mongo.IndexModel{
		Keys: bson.D{
			{Key: "user_id", Value: 1},
			{Key: "code_base_id", Value: 1},
		},
	}
	_, err := r.Collection.Indexes().CreateMany(ctx, []mongo.IndexModel{
		codeBaseIdIndex,
		compoundIndex1,
	})
	return err

}

func (r *ChatRepository) AddChat(ctx context.Context, chat *models.Chat) error {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()
	_, err := r.Collection.InsertOne(ctx, chat)
	return err
}

func (r *ChatRepository) GetAllChats(ctx context.Context, userId string, codeBaseId string) ([]models.Chat, error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	filter := bson.M{
		"user_id":     userId,
		"codebase_id": codeBaseId,
	}

	cursor, err := r.Collection.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var chats []models.Chat
	if err := cursor.All(ctx, &chats); err != nil {
		return nil, err
	}

	return chats, nil
}

func (r *ChatRepository) DeleteChatsByCodeBaseId(ctx context.Context, codeBaseId string) error {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	filter := bson.M{"codebase_id": codeBaseId}

	_, err := r.Collection.DeleteMany(ctx, filter)
	if err != nil {
		return err
	}
	return nil
}
