package main

import (
	"fmt"

	"github.com/valyala/fastjson"
)

func main() {

	// para quando json tem erros

	var s fastjson.Scanner

	s.Init(`{ "user"  "xxx"}`) //json com erro

	for s.Next() {
		fmt.Printf("Value: %s\n", s.Value())
	}

}
