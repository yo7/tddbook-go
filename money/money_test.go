package money

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMoney(t *testing.T) {
	five := NewMoney(5)
	product := five.Times(2)
	assert.Equal(t, 10, product.Amount())
	product = five.Times(3)
	assert.Equal(t, 15, product.Amount())
}

func TestEquality(t *testing.T) {
	assert.True(t, NewMoney(5).equals(NewMoney(5)))
	assert.False(t, NewMoney(5).equals(NewMoney(6)))
}
