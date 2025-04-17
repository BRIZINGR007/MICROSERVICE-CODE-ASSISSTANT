package repositories

import (
	"context"
	"time"

	"github.com/BRIZINGR007/app-002-code-assistant/internal/db/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type CodeContextRepository struct {
	Collection *mongo.Collection
}

func (r *CodeContextRepository) CodeContextIndexes(ctx context.Context) error {
	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	indexModel := mongo.IndexModel{
		Keys:    bson.M{"codebase_id": 1},
		Options: options.Index().SetUnique(false),
	}

	_, err := r.Collection.Indexes().CreateOne(ctx, indexModel)
	return err
}

func (r *CodeContextRepository) BulkInsertCodeContext(ctx context.Context, codeContexts *[]models.CodeContext) error {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()
	docs := make([]any, len(*codeContexts))
	for i, cc := range *codeContexts {
		docs[i] = cc
	}
	_, err := r.Collection.InsertMany(ctx, docs)
	return err
}

func (r *CodeContextRepository) CodeContextRetriever(
	ctx context.Context,
	codeBaseId string,
	embedding []float32,
	excludeVectorIds []string,
) ([]models.CodeContext, error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	var codeContextData []models.CodeContext

	filter := bson.M{"codebase_id": codeBaseId}
	if len(excludeVectorIds) > 0 {
		filter["vector_id"] = bson.M{"$nin": excludeVectorIds}
	}

	pipeline := mongo.Pipeline{
		{{Key: "$vectorSearch", Value: bson.D{
			{Key: "index", Value: "vector_index"},
			{Key: "path", Value: "embedding"},
			{Key: "queryVector", Value: embedding},
			{Key: "numCandidates", Value: 75},
			{Key: "limit", Value: 4},
			{Key: "filter", Value: filter},
		}}},
		{{Key: "$project", Value: bson.D{
			{Key: "_id", Value: 0},
			{Key: "vector_id", Value: 1},
			{Key: "hashId", Value: 1},
			{Key: "codebase_id", Value: 1},
			{Key: "filePath", Value: 1},
			{Key: "code", Value: 1},
			{Key: "Score", Value: bson.M{"$meta": "vectorSearchScore"}},
		}}},
	}

	cursor, err := r.Collection.Aggregate(ctx, pipeline)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	if err := cursor.All(ctx, &codeContextData); err != nil {
		return nil, err
	}

	return codeContextData, nil
}
