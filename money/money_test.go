package money

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMultiplication(t *testing.T) {
	five := NewDollar(5)
	assert.Equal(t, NewDollar(10), five.Times(2))
	assert.Equal(t, NewDollar(15), five.Times(3))
}

func TestEquality(t *testing.T) {
	assert.True(t, NewDollar(5).equals(NewDollar(5)))
	assert.False(t, NewDollar(5).equals(NewDollar(6)))
	assert.False(t, NewFranc(5).equals(NewDollar(5)))
}

func TestCurrency(t *testing.T) {
	assert.Equal(t, "USD", NewDollar(1).Currency())
	assert.Equal(t, "CHF", NewFranc(1).Currency())
}

func TestDifferentCurrencyEquality(t *testing.T) {
	assert.Equal(t, newMoney(10, "CHF").equals(NewFranc(10)), true)
}

func TestSimpleAddition(t *testing.T) {
	var sum Expression
	sum = NewDollar(5).Plus(NewDollar(5))
	bank := new(Bank)
	reduced := bank.Reduce(sum, "USD")
	assert.Equal(t, NewDollar(10), reduced)
}

func TestPlusReturnsSum(t *testing.T) {
	five := NewDollar(5)
	result := five.Plus(five)
	sum := result.(*Sum)
	assert.Equal(t, five, sum.augend)
	assert.Equal(t, five, sum.addend)
}

func TestReduceSum(t *testing.T) {
	sum := &Sum{NewDollar(3), NewDollar(4)}
	bank := new(Bank)
	result := bank.Reduce(sum, "USD")
	assert.Equal(t, NewDollar(7), result)
}

func TestReduceMoney(t *testing.T) {
	b := Bank{}
	result := b.Reduce(NewDollar(1), "USD")
	assert.Equal(t, NewDollar(1), result)
}

func TestReduceMoneyDifferentCurrency(t *testing.T) {
	b := Bank{}
	b.addRate("CHF", "USD", 2)
	result := b.Reduce(NewFranc(2), "USD")
	assert.Equal(t, NewDollar(1), result)
}
