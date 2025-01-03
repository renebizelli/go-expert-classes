package tax

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCalculateTax(t *testing.T) {

	tax, e := CalculateTax(1000)

	assert.Nil(t, e)
	assert.Equal(t, 10.0, tax)
}

func TestCalculateTaxError(t *testing.T) {

	tax, e := CalculateTax(0)

	assert.Error(t, e, "amount must be greater than 0")
	assert.Equal(t, 0.0, tax)
}
