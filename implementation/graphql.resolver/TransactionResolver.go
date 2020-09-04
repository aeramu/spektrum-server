package resolver

import (
	"github.com/aeramu/spektrum-server/entity"
	"github.com/graph-gophers/graphql-go"
)

//TransactionResolver interface
type TransactionResolver interface {
	ID() graphql.ID
	Source() string
	Destination() string
	Item() string
	Amount() int32
}

//TransactionConnectionResolver interface
type TransactionConnectionResolver interface {
	Edges() []TransactionResolver
}

type transactionResolver struct {
	transaction entity.Transaction
}

func (r *transactionResolver) ID() graphql.ID {
	return graphql.ID(r.transaction.ID())
}
func (r *transactionResolver) Source() string {
	return r.transaction.Source()
}
func (r *transactionResolver) Destination() string {
	return r.transaction.Destination()
}
func (r *transactionResolver) Item() string {
	return r.transaction.Item()
}
func (r *transactionResolver) Amount() int32 {
	return int32(r.transaction.Amount())
}

type transactionConnectionResolver struct {
	transactionList []entity.Transaction
}

func (r *transactionConnectionResolver) Edges() []TransactionResolver {
	var transactionResolverList []TransactionResolver
	for _, transaction := range r.transactionList {
		transactionResolverList = append(transactionResolverList, &transactionResolver{transaction})
	}
	return transactionResolverList
}
