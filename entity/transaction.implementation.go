package entity

type transaction struct {
	id          string
	source      string
	destination string
	item        string
	amount      int
}

func (t *transaction) ID() string {
	return t.id
}

func (t *transaction) Source() string {
	return t.source
}

func (t *transaction) Destination() string {
	return t.destination
}

func (t *transaction) Item() string {
	return t.item
}

func (t *transaction) Amount() int {
	return t.amount
}
