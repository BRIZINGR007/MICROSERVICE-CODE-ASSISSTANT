package repositories

import (
	"context"
	"fmt"
	"time"

	"github.com/BRIZINGR007/app-002-code-assistant/internal/db/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type UserRepository struct {
	Collection *mongo.Collection
}

func (r *UserRepository) UserIndexes(ctx context.Context) error {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	indexModel := mongo.IndexModel{
		Keys:    bson.M{"email": 1},
		Options: options.Index().SetUnique(true),
	}

	_, err := r.Collection.Indexes().CreateOne(ctx, indexModel)
	return err
}

func (r *UserRepository) GetCodeBaseData(ctx context.Context, email string) ([]models.CodeBaseData, error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	var user models.User
	err := r.Collection.FindOne(ctx, bson.M{"email": email}).Decode(&user)
	if err != nil {
		return nil, err
	}

	return user.CodebaseData, nil
}

func (r *UserRepository) AddUser(ctx context.Context, user *models.User) error {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	_, err := r.Collection.InsertOne(ctx, user)
	return err
}
func (r *UserRepository) GetUserByEmail(ctx context.Context, email string) (*models.User, error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()
	var user models.User
	err := r.Collection.FindOne(ctx, bson.M{"email": email}).Decode(&user)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}

func (r *UserRepository) AddCodeBaseData(ctx context.Context, email string, codeBaseData *models.CodeBaseData) error {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	filter := bson.M{"email": email}
	update := bson.M{
		"$push": bson.M{
			"codebasedata": codeBaseData,
		},
	}

	result, err := r.Collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return err
	}

	if result.MatchedCount == 0 {
		return fmt.Errorf("no user found with email: %s", email)
	}

	return nil
}
