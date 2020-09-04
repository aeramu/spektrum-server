package entity

//Transaction interface
type Transaction interface {
	ID() string
	Source() string
	Destination() string
	Item() string
	Amount() int
}

//TransactionConstructor constructor
type TransactionConstructor struct {
	ID          string
	Source      string
	Destination string
	Item        string
	Amount      int
}

//New transaction
func (tc TransactionConstructor) New() Transaction {
	return &transaction{
		id:          tc.ID,
		source:      tc.Source,
		destination: tc.Destination,
		item:        tc.Item,
		amount:      tc.Amount,
	}
}
