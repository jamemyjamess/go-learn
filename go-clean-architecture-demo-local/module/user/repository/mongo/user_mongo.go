package mongo

import (
	"context"
	"log"

	"github.com/jamemyjamess/go-clean-architecture-demo/module/user/entity/domain"
	"github.com/jamemyjamess/go-clean-architecture-demo/module/user/repository"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type userRepository struct {
	Client *mongo.Client
	DB     *mongo.Database
	Coll   *mongo.Collection
	URI    string
	DBName string
}

func NewUserRepository(mongoClient *mongo.Client) repository.UserRepository {
	dbName := "something"
	collName := "something"
	userRepoMongo := &userRepository{}
	userRepoMongo.Client = mongoClient
	userRepoMongo.DB = mongoClient.Database(dbName)
	userRepoMongo.Coll = userRepoMongo.DB.Collection(collName)
	return userRepoMongo
}

func (repo *userRepository) List(ctx context.Context) ([]domain.User, error) {
	// var filters bson.M
	// var optFilter []string
	// var opts *options.FindOptions
	// if opt != nil {
	// 	opts = repo.makePagingOpts(opt.Page, opt.PerPage)
	// 	if opt.Filters != nil && len(opt.Filters) > 0 {
	// 		optFilter = opt.Filters
	// 		filters = repo.makeFilters(opt.Filters)
	// 	}
	// 	if opt.Sorts != nil && len(opt.Sorts) > 0 {
	// 		opts.Sort = repo.makeSorts(opt.Sorts)
	// 	}
	// }

	// total, err = repo.Count(ctx, optFilter)
	// if err != nil {
	// 	return 0, nil, err
	// }

	// cursor, err := repo.Coll.Find(ctx, filters, opts)
	// if err != nil {
	// 	return 0, nil, err
	// }
	// defer func() { _ = cursor.Close(ctx) }()

	// for cursor.Next(ctx) {
	// 	item, err := repo.clone(itemType)
	// 	if err != nil {
	// 		return 0, nil, err
	// 	}
	// 	err = cursor.Decode(item)
	// 	if err != nil {
	// 		return 0, nil, err
	// 	}
	// 	items = append(items, item)
	// }

	// return total, items, nil]
	return nil, nil
}

func (repo *userRepository) Find(ctx context.Context, id string) (*domain.User, error) {

	return nil, nil
}

func (repo *userRepository) Create(ctx context.Context, user *domain.User) (err error) {
	res, err := repo.Coll.InsertOne(ctx, user)
	if err != nil {
		return err
	}
	log.Println("res.InsertedID.(primitive.ObjectID).Hex():", res.InsertedID.(primitive.ObjectID).Hex())
	// return res.InsertedID.(primitive.ObjectID).Hex(), nil
	return nil
}

func (repo *userRepository) Update(ctx context.Context, user *domain.User) (err error) {
	condition := "utils.MakeFilter()"
	_, err = repo.Coll.UpdateOne(ctx, condition, bson.M{"$set": user})
	return err
}

func (repo *userRepository) Delete(ctx context.Context, user *domain.User) (err error) {
	_, err = repo.Coll.DeleteOne(ctx, user)
	return err
}
