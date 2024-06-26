package repmongo

import (
	"context"
	"fmt"

	"log/slog"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func NewConnect() (*mongo.Database, error) {
	uri := "mongodb://root:12344321@localhost:27017/?maxPoolSize=10"

	if uri == "" {
		slog.Error("You must set your 'MONGODB_URI' environment variable. See\n\t https://www.mongodb.com/docs/drivers/go/current/usage-examples/#environment-variable")
	}

	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))

	if err != nil {
		return nil, fmt.Errorf("error to create connection with database\n %w", err)
	}

	return client.Database("library"), nil
}
