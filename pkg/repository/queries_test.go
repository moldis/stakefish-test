package repository

import (
	"context"
	"fmt"
	"testing"
	"time"

	"code.stakefish.test/service/ip_validator/pkg/model"

	testdb "code.stakefish.test/service/ip_validator/pkg/repository/test_db"
	"github.com/stretchr/testify/suite"
	"go.mongodb.org/mongo-driver/mongo"
)

type QueriesRepositorySuite struct {
	suite.Suite
	repository *queries
	db         *mongo.Database
	now        time.Time
	cleanups   []func()
}

func (suite *QueriesRepositorySuite) SetupSuite() {
	db, cancel, err := testdb.NewMongoDB()
	if err != nil {
		suite.Failf("Error creating ClickHouse Client", fmt.Sprintf("%v", err))
	}
	suite.db = db
	suite.cleanups = append(suite.cleanups, cancel)
	repo := NewQueries(suite.db)
	suite.repository = repo
}

func (suite *QueriesRepositorySuite) TestCreate() {
	ctx := context.Background()
	createdAt := time.Now().UnixMilli()

	in := model.Query{
		CreatedAt: createdAt,
		ClientIP:  "192.178.0.1",
		Domain:    "google.com",
		Addresses: []model.Address{model.Address{"192.168.0.2"}},
	}
	err := suite.repository.Create(ctx, &in)
	suite.Assert().NoError(err)

	in.Domain = "yandex.ru"
	in.CreatedAt = createdAt + 1000
	err = suite.repository.Create(ctx, &in)
	suite.Assert().NoError(err)

	result, err := suite.repository.List(ctx, 0, 20)
	suite.Assert().NoError(err)
	suite.Assert().Len(result, 2)
	suite.Assert().Equal(result[0].CreatedAt, createdAt+1000)
	suite.Assert().Equal(result[1].CreatedAt, createdAt)

	// test limits
	for i := 1; i <= 40; i++ {
		in.CreatedAt = createdAt + 1000
		err = suite.repository.Create(ctx, &in)
		suite.Assert().NoError(err)
	}

	result, err = suite.repository.List(ctx, 0, 20)
	suite.Assert().NoError(err)
	suite.Assert().Len(result, 20)
}

func (suite *QueriesRepositorySuite) TearDownSuite() {
	for i := range suite.cleanups {
		suite.cleanups[i]()
	}
}

func TestPricesIntegrationSuite(t *testing.T) {
	suite.Run(t, new(QueriesRepositorySuite))
}
