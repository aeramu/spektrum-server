package resolver

import (
	"context"

	"github.com/aeramu/spektrum-server/interactor"
)

//Schema graphql
var Schema = `
	schema{
		query: Query
		mutation: Mutation
	}
	type Query{
		myAccount: Account
		scoreboard: AccountConnection!
		transactionList: TransactionConnection!
		itemList: ItemConnection!
	}
	type Mutation{
		signIn(nim: String!, code: String!): String
		transfer(receiverNIM: String!, money: Int!): Account
		buyItem(item: String!, amount: Int!): Account
	}
	type Account{
		id: ID!
		nim: String!
		name: String!
		money: Int!
		venture: Int!
	}
	type AccountConnection{
		edges: [Account!]!
		sum: Int!
		target: Int!
	}
	type Transaction{
		id: ID!
		source: String!
		destination: String!
		item: String!
		amount: Int!
	}
	type TransactionConnection{
		edges: [Transaction!]!
	}
	type Item{
		id: ID!
		name: String!
		description: String!
		price: Int!
	}
	type ItemConnection{
		edges: [Item!]!
	}
`

//Resolver interface
type Resolver interface {
	MyAccount() AccountResolver
	Scoreboard() AccountConnectionResolver
	SignIn(args struct {
		NIM  string
		Code string
	}) *string
	Transfer(args struct {
		ReceiverNIM string
		Money       int32
	}) AccountResolver
	ItemList() ItemConnectionResolver
}

//Constructor Reoslver
type Constructor struct {
	Interactor interactor.Interactor
	Context    context.Context
}

//New Resolver
func (c Constructor) New() Resolver {
	return &resolver{
		i: c.Interactor,
		c: c.Context,
	}
}
