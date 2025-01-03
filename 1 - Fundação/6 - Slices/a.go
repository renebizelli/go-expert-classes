package main

import "fmt"

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

}
