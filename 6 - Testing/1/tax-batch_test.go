package tax

import "testing"

func TestCalculateTaxBatch(t *testing.T) {
	type calcTax struct {
		amount, expected float64
	}

	table := []calcTax{
		{500.0, 5.0},
		{1000.0, 10.0},
		{1500.0, 10.0},
		{0, 0},
	}

	for _, calc := range table {
		result := CalculateTax(calc.amount)

		if calc.expected != result {
			t.Errorf("Expected %f but got %f", calc.expected, result)
		}
	}
}

func TestCalculateTaxBatchLazy(t *testing.T) {
	type calcTax struct {
		amount, expected float64
	}

	table := []calcTax{
		{500.0, 5.0},
		{1000.0, 10.0},
		{1500.0, 10.0},
		{0, 0},
	}

	for _, calc := range table {
		result := CalculateTaxLazy(calc.amount)

		if calc.expected != result {
			t.Errorf("Expected %f but got %f", calc.expected, result)
		}
	}
}
