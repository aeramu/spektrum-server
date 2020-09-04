package repository

import (
	"github.com/aeramu/spektrum-server/entity"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

//Account model
type account struct {
	ID      primitive.ObjectID `bson:"_id"`
	NIM     string
	Name    string
	Code    string
	Money   int
	Venture int
}

//Entity convert
func (a *account) Entity() entity.Account {
	return entity.AccountConstructor{
		ID:      a.ID.Hex(),
		NIM:     a.NIM,
		Name:    a.Name,
		Code:    a.Code,
		Money:   a.Money,
		Venture: a.Venture,
	}.New()
}

func accountListToEntity(accountList []*account) []entity.Account {
	var entityList []entity.Account
	for _, account := range accountList {
		entityList = append(entityList, account.Entity())
	}
	return entityList
}
