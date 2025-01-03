package main

import "fmt"

type Human struct {
	Name  string
	Genre string
}

type Animal struct {
	Name  string
	Genre string
}

type Movements interface {
	Walking()
	Running()
}

func (h Human) Walking() {
	fmt.Println(h.Name + " andando")

}

func (a Animal) Walking() {
	fmt.Println(a.Name + " andando")
}

func (h Human) Running() {

}

func (a Animal) Running() {

}

func Walking(m Movements) {
	m.Walking()
}

func main() {

	human := Human{Name: "René", Genre: "M"}
	animal := Animal{Name: "Guiguinho", Genre: "M"}

	Walking(human)
	Walking(animal)

}
