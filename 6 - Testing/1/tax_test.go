package tax

import "testing"

func TestCalculateTax(t *testing.T) {
	amount := 500.0
	expected := 5.0

	result := CalculateTax(amount)

	if result != expected {
		t.Errorf("Expected %f but got %f", expected, result)
	}
}

func TestCalculateTaxLazy(t *testing.T) {
	amount := 500.0
	expected := 5.0

	result := CalculateTaxLazy(amount)

	if result != expected {
		t.Errorf("Expected %f but got %f", expected, result)
	}
}

func FuzzCalculateTax(f *testing.F) {

	seed := []float64{-1.0, -2.0, 0, 100.0, 500.00, 2000.00, 20000.0}
	for _, s := range seed {
		f.Add(s)
	}

	f.Fuzz(func(t *testing.T, amount float64) {

		result := CalculateTax(amount)

		// if amount == 0 && result != 0 {
		// 	t.Errorf("Received %f but expected 0", result)
		// } else if amount >= 20000.0 && result != 20.0 {
		// 	t.Errorf("Received %f but expected 20", result)
		// } else if amount >= 1000.0 && result != 10.0 {
		// 	t.Errorf("Received %f but expected 10", result)
		// } else if amount < 1000.0 && result != 5.0 {
		// 	t.Errorf("Received %f but expected 5", result)
		// }

		if amount == 0 && result != 0 {
			t.Errorf("Received %f but expected 0", result)
		} else if amount != 0 && result != 5.0 {
			t.Errorf("Received %f but expected 5", result)
		}

	})

}
