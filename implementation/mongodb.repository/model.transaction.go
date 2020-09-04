package repository

import (
	"github.com/aeramu/spektrum-server/entity"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type transaction struct {
	ID          primitive.ObjectID `bson:"_id"`
	Source      string
	Destination string
	Item        string
	Amount      int
}

//Entity convert
func (t *transaction) Entity() entity.Transaction {
	return entity.TransactionConstructor{
		ID:          t.ID.Hex(),
		Source:      t.Source,
		Destination: t.Destination,
		Item:        t.Item,
		Amount:      t.Amount,
	}.New()
}

func transactionListToEntity(transactionList []*transaction) []entity.Transaction {
	var entityList []entity.Transaction
	for _, transaction := range transactionList {
		entityList = append(entityList, transaction.Entity())
	}
	return entityList
}
