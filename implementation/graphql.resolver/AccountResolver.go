package resolver

import (
	"github.com/aeramu/spektrum-server/entity"
	"github.com/graph-gophers/graphql-go"
)

//AccountResolver interface
type AccountResolver interface {
	ID() graphql.ID
	NIM() string
	Name() string
	Money() int32
	Venture() int32
}

//AccountConnectionResolver interface
type AccountConnectionResolver interface {
	Edges() []AccountResolver
	Sum() int32
	Target() int32
}

type accountResolver struct {
	account entity.Account
}

func (r *accountResolver) ID() graphql.ID {
	return graphql.ID(r.account.ID())
}
func (r *accountResolver) NIM() string {
	return r.account.NIM()
}
func (r *accountResolver) Name() string {
	return r.account.Name()
}
func (r *accountResolver) Money() int32 {
	return int32(r.account.Money())
}
func (r *accountResolver) Venture() int32 {
	return int32(r.account.Venture())
}

type accountConnectionResolver struct {
	accountList []entity.Account
	r           *resolver
}

func (r *accountConnectionResolver) Edges() []AccountResolver {
	var accountResolverList []AccountResolver
	for _, account := range r.accountList {
		accountResolverList = append(accountResolverList, &accountResolver{account})
	}
	return accountResolverList
}

func (r *accountConnectionResolver) Sum() int32 {
	var sum int
	for _, account := range r.accountList {
		sum += account.Venture()
	}
	return int32(sum)
}

func (r *accountConnectionResolver) Target() int32 {
	target := r.r.i.VentureTarget()
	return int32(target)
}
