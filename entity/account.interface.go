package entity

//Account interface
type Account interface {
	getter
	wallet
}

type getter interface {
	ID() string
	NIM() string
	Name() string
	IsCodeRight(string) bool
	Money() int
	Venture() int
}

type wallet interface {
	IncMoney(money int) int
	DecMoney(money int) int
	IncVenture(amount int) int
}

//AccountConstructor constructor
type AccountConstructor struct {
	ID      string
	NIM     string
	Name    string
	Code    string
	Money   int
	Venture int
}

//New new account
func (ac AccountConstructor) New() Account {
	return &account{
		id:      ac.ID,
		nim:     ac.NIM,
		name:    ac.Name,
		code:    ac.Code,
		money:   ac.Money,
		venture: ac.Venture,
	}
}
