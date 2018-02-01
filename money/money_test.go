package money

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMoney(t *testing.T) {
	five := NewDollar(5)
	five.Times(2)
	assert.Equal(t, 10, five.Amount())
}
