package main

import (
	"fmt"
	"regexp"
	"strings"
)

func main() {

	s := []int{1, 2, 3, 4, 5}

	fmt.Printf("\nLen: %d | cap: %d | s: %v", len(s), cap(s), s)
	fmt.Printf("\nLen: %d | cap: %d | s: %v", len(s[:1]), cap(s[:0]), s[:2])
	fmt.Printf("\nLen: %d | cap: %d | s: %v", len(s[1:]), cap(s[0:]), s[3:])

	s = append(s, 6)
	fmt.Printf("\nLen: %d | cap: %d | s: %v", len(s), cap(s), s)

	s = append(s, 7)
	fmt.Printf("\nLen: %d | cap: %d | s: %v", len(s), cap(s), s)

	s = append(s, 8)
	fmt.Printf("\nLen: %d | cap: %d | s: %v", len(s), cap(s), s)

	s = append(s, 9)
	fmt.Printf("\nLen: %d | cap: %d | s: %v", len(s), cap(s), s)

	s = append(s, 10)
	fmt.Printf("\nLen: %d | cap: %d | s: %v", len(s), cap(s), s)

	s = append(s, 11)
	fmt.Printf("\nLen: %d | cap: %d | s: %v", len(s), cap(s), s)
	fmt.Printf("\nLen: %d | cap: %d | s: %v", len(s), cap(s[5:]), s[5:])

	sliceJoin()

	remove()

}

func sliceJoin() {
	fmt.Println("\n--------------------------------------------------------------------")
	re := regexp.MustCompile("[0-9]+")
	fmt.Println(strings.Join(re.FindAllString("1abc123def987asdf", -1)[:], ""))
}

func remove() {

	fmt.Println("\n--------------------------------------------------------------------")

	events := []string{"Evento1", "Evento2", "Evento3", "Evento4"}

	fmt.Printf("values %v \n", events)
	fmt.Printf("[:2] %v \n", events[:2])
	fmt.Printf("[2:] %v \n", events[2:])
	fmt.Printf("[:0] %v \n", events[:0])
	fmt.Printf("[1:] %v \n", events[1:])

	events = append(events[:0], events[1:]...) //	events = events[1:]

	fmt.Printf("new values %v \n", events)

}
