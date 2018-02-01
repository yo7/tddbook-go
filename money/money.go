package money

type Dollar struct {
	amount int
}

func NewDollar(i int) *Dollar {
	return &Dollar{amount: i}
}

func (d *Dollar) Amount() int {
	return d.amount
}

func (d *Dollar) Times(multiplier int) *Dollar {
	return NewDollar(d.Amount() * multiplier)
}

func (d *Dollar) equals(other interface{}) bool {
	mm := other.(*Dollar)
	return d.Amount() == mm.Amount()
}

type Franc struct {
	amount int
}

func NewFranc(i int) *Franc {
	return &Franc{amount: i}
}

func (f *Franc) Amount() int {
	return f.amount
}

func (f *Franc) Times(multiplier int) *Franc {
	return NewFranc(f.Amount() * multiplier)
}
