package money

type IMoney interface {
	Amount() int
}

type Money struct {
	amount int
}

func (m *Money) Amount() int {
	return m.amount
}

func (m *Money) equals(other interface{}) bool {
	mm := other.(IMoney)
	return m.Amount() == mm.Amount()
}

type Dollar struct {
	Money
}

func NewDollar(i int) *Dollar {
	return &Dollar{Money{amount: i}}
}

func (d *Dollar) Times(multiplier int) *Dollar {
	return NewDollar(d.Amount() * multiplier)
}

type Franc struct {
	Money
}

func NewFranc(i int) *Franc {
	return &Franc{Money{amount: i}}
}

func (f *Franc) Times(multiplier int) *Franc {
	return NewFranc(f.Amount() * multiplier)
}
