package mongo

import (
	"context"

	"github.com/jamemyjamess/go-clean-architecture-demo/module/company/entity/domain"
	"github.com/jamemyjamess/go-clean-architecture-demo/module/company/repository"
	"go.mongodb.org/mongo-driver/mongo"
)

type companyRepository struct {
	Client *mongo.Client
	DB     *mongo.Database
	Coll   *mongo.Collection
	URI    string
	DBName string
}

func NewUserRepository(mongoClient *mongo.Client) repository.CompanyRepository {
	dbName := "something"
	collName := "something"
	companyMongo := &companyRepository{}
	companyMongo.Client = mongoClient
	companyMongo.DB = mongoClient.Database(dbName)
	companyMongo.Coll = companyMongo.DB.Collection(collName)
	return companyMongo
}

func (repo *companyRepository) Find(ctx context.Context, id string) (*domain.Company, error) {
	return nil, nil
}
