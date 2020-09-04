package entity

//Item interface
type Item interface {
	ID() string
	Name() string
	Description() string
	Price() int
}

//ItemConstructor constructor
type ItemConstructor struct {
	ID          string
	Name        string
	Description string
	Price       int
}

//New item
func (c ItemConstructor) New() Item {
	return &item{
		id:          c.ID,
		name:        c.Name,
		description: c.Description,
		price:       c.Price,
	}
}
