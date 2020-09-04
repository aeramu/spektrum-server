package resolver

import (
	"github.com/aeramu/spektrum-server/entity"
	"github.com/graph-gophers/graphql-go"
)

//ItemResolver interface
type ItemResolver interface {
	ID() graphql.ID
	Name() string
	Description() string
	Price() int32
}

//ItemConnectionResolver interface
type ItemConnectionResolver interface {
	Edges() []ItemResolver
}

type itemResolver struct {
	item entity.Item
}

func (r *itemResolver) ID() graphql.ID {
	return graphql.ID(r.item.ID())
}
func (r *itemResolver) Name() string {
	return r.item.Name()
}
func (r *itemResolver) Description() string {
	return r.item.Description()
}
func (r *itemResolver) Price() int32 {
	return int32(r.item.Price())
}

type itemConnectionResolver struct {
	itemList []entity.Item
}

func (r *itemConnectionResolver) Edges() []ItemResolver {
	var itemResolverList []ItemResolver
	for _, item := range r.itemList {
		itemResolverList = append(itemResolverList, &itemResolver{item})
	}
	return itemResolverList
}
