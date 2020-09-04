package entity

type item struct {
	id          string
	name        string
	description string
	price       int
}

func (i *item) ID() string {
	return i.id
}
func (i *item) Name() string {
	return i.name
}
func (i *item) Description() string {
	return i.description
}
func (i *item) Price() int {
	return i.price
}
