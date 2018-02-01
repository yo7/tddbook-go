package money

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMultiplication(t *testing.T) {
	five := NewMoney(5)
	assert.Equal(t, NewMoney(10), five.Times(2))
	assert.Equal(t, NewMoney(15), five.Times(3))
}

func TestEquality(t *testing.T) {
	assert.True(t, NewMoney(5).equals(NewMoney(5)))
	assert.False(t, NewMoney(5).equals(NewMoney(6)))
}
