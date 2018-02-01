package money

type Money struct {
	amount int
}

func NewMoney(i int) *Money {
	return &Money{amount: i}
}

func (m *Money) Amount() int {
	return m.amount
}

func (m *Money) Times(multiplier int) *Money {
	return NewMoney(m.Amount() * multiplier)
}

func (m *Money) equals(other interface{}) bool {
	mm := other.(*Money)
	return m.Amount() == mm.Amount()
}
