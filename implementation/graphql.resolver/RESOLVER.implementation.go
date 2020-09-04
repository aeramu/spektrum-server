package resolver

import (
	"context"

	"github.com/aeramu/spektrum-server/interactor"
)

type resolver struct {
	i interactor.Interactor
	c context.Context
}

func (r *resolver) MyAccount() AccountResolver {
	token := r.c.Value("request").(map[string]string)["token"]
	nim := token
	account := r.i.Account(nim)
	if account == nil {
		return nil
	}
	return &accountResolver{account}
}

func (r *resolver) Scoreboard() AccountConnectionResolver {
	accountList := r.i.Scoreboard()
	return &accountConnectionResolver{accountList, r}
}

func (r *resolver) SignIn(args struct {
	NIM  string
	Code string
}) *string {
	account := r.i.SignIn(args.NIM, args.Code)
	if account == nil {
		return nil
	}
	token := account.NIM()
	return &token
}

func (r *resolver) Transfer(args struct {
	ReceiverNIM string
	Money       int32
}) AccountResolver {
	token := r.c.Value("request").(map[string]string)["token"]
	nim := token
	account := r.i.Transfer(nim, args.ReceiverNIM, int(args.Money))
	if account == nil {
		return nil
	}
	return &accountResolver{account}
}

func (r *resolver) TransactionList() TransactionConnectionResolver {
	// token := r.c.Value("request").(map[string]string)["token"]
	// nim := token
	transactionList := r.i.TransactionList("")
	return &transactionConnectionResolver{transactionList}
}

func (r *resolver) BuyItem(args struct {
	Item   string
	Amount int32
}) AccountResolver {
	token := r.c.Value("request").(map[string]string)["token"]
	nim := token
	account := r.i.BuyItem(nim, args.Item, int(args.Amount))
	return &accountResolver{account}
}

func (r *resolver) ItemList() ItemConnectionResolver {
	itemList := r.i.ItemList()
	return &itemConnectionResolver{itemList}
}
