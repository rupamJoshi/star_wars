package repo

import (
	"context"
	"fmt"
	"time"

	"github.com/rupam_joshi/star_wars/graph/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type StarWarRepo interface {
	Save(character *model.FavoriteCharacter) error
	GetAll() ([]*model.FavoriteCharacter, error)
}

type database struct {
	client     *mongo.Client
	collection *mongo.Collection
}

func New(uri, dbName, collectionName string) (StarWarRepo, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		return nil, fmt.Errorf("failed to connect to mongo: %w", err)
	}
	if err := client.Ping(ctx, nil); err != nil {
		return nil, fmt.Errorf("failed to ping mongo: %w", err)
	}

	return &database{
		client:     client,
		collection: client.Database(dbName).Collection(collectionName),
	}, nil
}

func (db *database) Save(character *model.FavoriteCharacter) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	character.ID = primitive.NewObjectID().String()
	character.SavedAt = time.Now().Format("2006-01-02 15:04:05")

	_, err := db.collection.InsertOne(ctx, character)
	if err != nil {
		return fmt.Errorf("failed to insert character: %w", err)
	}

	return nil
}

func (db *database) GetAll() ([]*model.FavoriteCharacter, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	cursor, err := db.collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, fmt.Errorf("failed to fetch characters: %w", err)
	}
	defer cursor.Close(ctx)

	var characters []*model.FavoriteCharacter
	for cursor.Next(ctx) {
		var c model.FavoriteCharacter
		if err := cursor.Decode(&c); err != nil {
			return nil, fmt.Errorf("failed to decode character: %w", err)
		}
		characters = append(characters, &c)
	}

	return characters, nil
}
