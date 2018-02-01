package money

type Money struct {
	amount int
}

func NewDollar(i int) *Money {
	return &Money{amount: i}
}

func (m *Money) Amount() int {
	return m.amount
}

func (m *Money) Times(multiplier int) *Money {
	a := m.amount * multiplier
	return &Money{amount: a}
}
