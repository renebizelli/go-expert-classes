package tax

func CalculateTax(amount float64) float64 {

	if amount == 0.0 {
		return 0.0
	}

	// if amount >= 20000.0 {
	// 	return 20.0
	// }

	// if amount >= 1000.0 {
	// 	return 10.0
	// }

	return 5.0
}

func CalculateTaxLazy(amount float64) float64 {

	// time.Sleep(time.Microsecond)

	// if amount == 0 {
	// 	return 0
	// }

	// if amount >= 1000 {
	// 	return 10
	// }

	// return 5.0

	return 9
}
