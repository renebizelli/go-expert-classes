package tax

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestProcess1000(t *testing.T) {

	repository := &TaxRepositoryMock{}

	repository.On("SaveTax", 10.0).Return(nil)

	e := Process(1000.0, repository)

	assert.Nil(t, e)
}

func TestProcess0(t *testing.T) {

	repository := &TaxRepositoryMock{}

	repository.On("SaveTax", 0.0).Return(errors.New("erro..."))

	e := Process(0.0, repository)

	assert.Error(t, e, "erro...")

	repository.AssertNumberOfCalls(t, "SaveTax", 1)
}
