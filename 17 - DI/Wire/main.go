package main

import "fmt"

func main() {

	s := NewServiceWire()

	r := s.Get(1)

	fmt.Println(r)

	//  go run .\main.go .\wire_gen.go .\service.go .\repositoy.go

}
