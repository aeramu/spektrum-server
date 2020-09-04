package repository

import (
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/aeramu/spektrum-server/entity"
	"github.com/aeramu/spektrum-server/interactor"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type repository struct {
	client *mongo.Client
}

//New mongodb repository
func New() interactor.Repository {
	client, _ := mongo.Connect(context.Background(), options.Client().ApplyURI(
		"mongodb+srv://admin:admin@qiup-wrbox.mongodb.net/",
	))
	return &repository{
		client: client,
	}
}

func (r *repository) GetDataByNIM(nim string) entity.Account {
	filter := bson.D{{"nim", nim}}
	var account account
	r.client.Database("spektrum").Collection("wallet").FindOne(context.TODO(), filter).Decode(&account)

	if account.ID.IsZero() {
		return nil
	}
	return account.Entity()
}
func (r *repository) GetDataListSortedByIndex(index string) []entity.Account {
	filter := bson.D{}
	sortOpt := bson.D{{"venture", -1}}
	option := options.Find().SetSort(sortOpt).SetLimit(int64(37))
	cursor, _ := r.client.Database("spektrum").Collection("wallet").Find(context.TODO(), filter, option)

	var accountList []*account
	cursor.All(context.TODO(), &accountList)
	return accountListToEntity(accountList)
}
func (r *repository) UpdateMoney(nim string, money int) {
	filter := bson.D{{"nim", nim}}
	update := bson.D{
		{"$set", bson.D{
			{"money", money},
		}},
	}
	r.client.Database("spektrum").Collection("wallet").UpdateOne(context.TODO(), filter, update)
}

func (r *repository) UpdateVenture(nim string, amount int) {
	filter := bson.D{{"nim", nim}}
	update := bson.D{
		{"$set", bson.D{
			{"venture", amount},
		}},
	}
	r.client.Database("spektrum").Collection("wallet").UpdateOne(context.TODO(), filter, update)
}

func (r *repository) PutTransaction(source string, destination string, item string, amount int) {
	transaction := &transaction{
		ID:          primitive.NewObjectID(),
		Source:      source,
		Destination: destination,
		Item:        item,
		Amount:      amount,
	}
	r.client.Database("spektrum").Collection("transaction").InsertOne(context.TODO(), transaction)
}

func (r *repository) GetTransactionList(nim string) []entity.Transaction {
	filter := bson.D{{
		"$or", bson.A{
			bson.D{{"source", nim}},
			bson.D{{"destination", nim}},
			bson.D{{"item", nim}},
		},
	}}
	sortOpt := bson.D{{"_id", -1}}
	option := options.Find().SetSort(sortOpt)

	cursor, _ := r.client.Database("spektrum").Collection("transaction").Find(context.TODO(), filter, option)

	var transactionList []*transaction
	cursor.All(context.TODO(), &transactionList)
	return transactionListToEntity(transactionList)
}
