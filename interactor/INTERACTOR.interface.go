package interactor

import "github.com/aeramu/spektrum-server/entity"

//Interactor interface
type Interactor interface {
	account
}

type account interface {
	SignIn(nim string, code string) entity.Account
	Transfer(senderNIM string, receiverNIM string, money int) entity.Account
	BuyItem(nim string, item string, amount int) entity.Account
	Scoreboard() []entity.Account
	VentureTarget() int
	TransactionList(nim string) []entity.Transaction
	Account(nim string) entity.Account
	ItemList() []entity.Item
}

//Constructor Interactor
type Constructor struct {
	Repository Repository
}

//New Interactor
func (c Constructor) New() Interactor {
	return &interactor{
		repo: c.Repository,
	}
}
