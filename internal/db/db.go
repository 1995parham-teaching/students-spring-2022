package db

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func New(cfg Config) (*mongo.Database, error) {
	options := new(options.ClientOptions)
	options.ApplyURI(cfg.URL)

	client, err := mongo.NewClient(options)
	if err != nil {
		return nil, fmt.Errorf("cannot create mongodb client %w", err)
	}

	ctx := context.Background()
	ctx, done := context.WithTimeout(ctx, cfg.ConnectionTimeout)
	defer done()

	if err := client.Connect(ctx); err != nil {
		return nil, fmt.Errorf("cannot connect to database %w", err)
	}

	{
		ctx := context.Background()
		ctx, done := context.WithTimeout(ctx, cfg.ConnectionTimeout)
		defer done()

		if err := client.Ping(ctx, readpref.Primary()); err != nil {
			return nil, fmt.Errorf("cannot ping database %w", err)
		}
	}

	return client.Database("students"), nil
}
