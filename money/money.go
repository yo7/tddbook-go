package money

type IMoney interface {
	Amount() int
	Currency() string
}

type Expression interface {
	Reduce(to string) *Money
}

type Sum struct {
	augend, addend *Money
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

func (m *Money) Plus(addend *Money) Expression {
	return &Sum{m, addend}
}

func (m *Money) equals(other interface{}) bool {
	mm := other.(IMoney)
	return m.Amount() == mm.Amount() &&
		m.Currency() == mm.Currency()
}

func (m *Money) Reduce(to string) *Money {
	var rate int
	if m.Currency() == "CHF" && to == "USD" {
		rate = 2
	} else {
		rate = 1
	}
	return newMoney(m.Amount()/rate, to)
}

func NewDollar(i int) *Money {
	return newMoney(i, "USD")
}

func NewFranc(i int) *Money {
	return newMoney(i, "CHF")
}

func (s *Sum) Reduce(to string) *Money {
	amount := s.augend.amount + s.addend.amount
	return newMoney(amount, to)
}

type Bank struct{}

func (b *Bank) Reduce(source Expression, to string) *Money {
	return source.Reduce(to)
}

func (b *Bank) addRate(from string, to string, rate int) {

}
