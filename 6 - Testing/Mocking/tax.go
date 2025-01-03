package tax

type Repository interface {
	SaveTax(tax float64) error
}

func CalculateTax(amount float64) float64 {

	if amount <= 0.0 {
		return 0.0
	}

	if amount >= 20000.0 {
		return 20.0
	}

	if amount >= 1000.0 {
		return 10.0
	}

	return 5.0
}

func Process(amount float64, repository Repository) error {

	tax := CalculateTax(amount)

	return repository.SaveTax(tax)
}
