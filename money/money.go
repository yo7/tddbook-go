package money

type IMoney interface {
	Amount() int
	Currency() string
}

type Money struct {
	amount   int
	currency string
}

func (m *Money) Amount() int {
	return m.amount
}

func (m *Money) Currency() string {
	return m.currency
}

func newMoney(i int, s string) *Money {
	return &Money{
		amount:   i,
		currency: s,
	}
}

func (m *Money) Times(multiplier int) *Money {
	return newMoney(m.amount*multiplier, m.currency)
}

func (m *Money) equals(other interface{}) bool {
	mm := other.(IMoney)
	return m.Amount() == mm.Amount() &&
		m.Currency() == mm.Currency()
}

func NewDollar(i int) *Money {
	return &Money{
		amount:   i,
		currency: "USD",
	}
}

func NewFranc(i int) *Money {
	return &Money{
		amount:   i,
		currency: "CHF",
	}
}
