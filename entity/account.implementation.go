package entity

type account struct {
	id      string
	nim     string
	name    string
	code    string
	money   int
	venture int
}

func (a *account) ID() string {
	return a.id
}
func (a *account) NIM() string {
	return a.nim
}
func (a *account) Name() string {
	return a.name
}
func (a *account) IsCodeRight(code string) bool {
	return a.code == code
}
func (a *account) Money() int {
	return a.money
}
func (a *account) Venture() int {
	return a.venture
}

func (a *account) IncMoney(money int) int {
	a.money += money
	return a.Money()
}
func (a *account) DecMoney(money int) int {
	if a.money >= money {
		a.money -= money
	}
	return a.Money()
}
func (a *account) IncVenture(amount int) int {
	a.venture += amount
	return a.Venture()
}
