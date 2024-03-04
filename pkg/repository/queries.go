package repository

import (
	"context"

	"code.stakefish.test/service/ip_validator/pkg/model"
	"github.com/rs/zerolog/log"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const collection = "queries"

//go:generate mockery --name=Queries --structname=MockQueries --outpkg=repository --output ./mocks --filename prices_mock.go
type Queries interface {
	Create(ctx context.Context, in *model.Query) error
	List(ctx context.Context, skip, limit int64) ([]model.Query, error)
}

type queries struct {
	pool *mongo.Database
}

func NewQueries(conn *mongo.Database) *queries {
	return &queries{
		pool: conn,
	}
}

func (a *queries) Create(ctx context.Context, in *model.Query) error {
	res, err := a.pool.Collection(collection).InsertOne(ctx, &in)
	if err != nil {
		return err
	}
	log.Debug().Msgf("inserted new collection with ID %d", res)
	return nil
}

func (a *queries) List(ctx context.Context, skip, limit int64) ([]model.Query, error) {
	filter := bson.D{}
	opts := options.Find().SetSort(bson.D{{"created_at", -1}}).SetLimit(limit).SetSkip(skip)

	cursor, err := a.pool.Collection(collection).Find(ctx, filter, opts)
	if err != nil {
		return nil, err
	}

	var results []model.Query
	if err = cursor.All(ctx, &results); err != nil {
		return nil, err
	}

	return results, nil
}
