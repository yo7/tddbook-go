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

func (m *Money) equals(other interface{}) bool {
	mm := other.(IMoney)
	return m.Amount() == mm.Amount()
}

type Dollar struct {
	Money
}

func NewDollar(i int) *Dollar {
	return &Dollar{
		Money{
			amount:   i,
			currency: "USD",
		},
	}
}

func (d *Dollar) Times(multiplier int) *Dollar {
	return NewDollar(d.Amount() * multiplier)
}

type Franc struct {
	Money
}

func NewFranc(i int) *Franc {
	return &Franc{
		Money{
			amount:   i,
			currency: "CHF",
		},
	}
}

func (f *Franc) Times(multiplier int) *Franc {
	return NewFranc(f.Amount() * multiplier)
}
