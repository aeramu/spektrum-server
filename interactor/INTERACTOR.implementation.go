package interactor

import "github.com/aeramu/spektrum-server/entity"

type interactor struct {
	repo Repository
}

//Repository interface
type Repository interface {
	GetDataByNIM(nim string) entity.Account
	UpdateMoney(nim string, money int)
	GetDataListSortedByIndex(index string) []entity.Account
	PutTransaction(source string, destination string, item string, amount int)
	GetTransactionList(nim string) []entity.Transaction
	UpdateVenture(nim string, amount int)
}

func (i *interactor) SignIn(nim string, code string) entity.Account {
	account := i.repo.GetDataByNIM(nim)
	if account == nil {
		return nil
	}

	if !account.IsCodeRight(code) {
		return nil
	}

	return account
}

func (i *interactor) Transfer(senderNIM string, receiverNIM string, money int) entity.Account {
	if senderNIM == receiverNIM {
		return nil
	}

	sender := i.repo.GetDataByNIM(senderNIM)
	if sender == nil {
		return nil
	}

	if money < 0 || sender.Money() < money {
		return sender
	}

	if receiverNIM == "18119000" {
		i.repo.UpdateMoney(senderNIM, sender.DecMoney(money))
		i.repo.UpdateVenture(senderNIM, sender.IncVenture(money))
		i.repo.PutTransaction(senderNIM, receiverNIM, "", money)
		return sender
	}

	receiver := i.repo.GetDataByNIM(receiverNIM)
	if receiver == nil {
		return sender
	}

	i.repo.UpdateMoney(senderNIM, sender.DecMoney(money))
	i.repo.UpdateMoney(receiverNIM, receiver.IncMoney(money))
	i.repo.PutTransaction(senderNIM, receiverNIM, "", money)

	return sender
}

func (i *interactor) Account(nim string) entity.Account {
	account := i.repo.GetDataByNIM(nim)
	if account == nil {
		return nil
	}
	return account
}

func (i *interactor) Scoreboard() []entity.Account {
	return i.repo.GetDataListSortedByIndex("money")
}

func (i *interactor) VentureTarget() int {
	return 1337000
}

func (i *interactor) BuyItem(nim string, item string, amount int) entity.Account {
	account := i.repo.GetDataByNIM(nim)
	if account == nil {
		return nil
	}

	if account.Money() >= amount {
		i.repo.UpdateMoney(nim, account.DecMoney(amount))
		i.repo.PutTransaction(nim, "", item, amount)
	}

	return account
}

func (i *interactor) TransactionList(nim string) []entity.Transaction {
	return i.repo.GetTransactionList(nim)
}

func (i *interactor) ItemList() []entity.Item {
	list := []entity.ItemConstructor{}
	var itemList []entity.Item
	for _, item := range list {
		itemList = append(itemList, item.New())
	}
	return itemList
}
