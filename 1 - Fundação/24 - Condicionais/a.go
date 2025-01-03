package main

func main() {
	a := 1
	b := 2
	c := 3

	if a > b && c > a {
		println(a)
	} else {
		println(b)
	}

	if a > b || c > a {
		println(a)
	} else {
		println(b)
	}

	switch a {
	case 1:
		println("1")
	case 2:
		println("2")
	default:
		println("d")
	}

}
