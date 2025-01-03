package matematica

func Soma[T int | float64](a, b T) T {
	return soma(a, b)
}

func soma[T int | float64](a, b T) T {
	return a + b
}
