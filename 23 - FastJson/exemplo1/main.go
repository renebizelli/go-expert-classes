package main

import (
	"fmt"

	"github.com/valyala/fastjson"
)

func main() {

	var p fastjson.Parser

	jsonData := `{"name":"John","age":30,"city":"New York", "arr": [1,2,3]}`

	v, err := p.Parse(jsonData)

	if err != nil {
		panic(err)
	}

	fmt.Printf("Name: %s\n", v.GetStringBytes("name"))
	fmt.Printf("Age: %d\n", v.GetInt("age"))

	arr := v.GetArray("arr")

	for i, va := range arr {
		fmt.Printf("Arr[%d]: %d\n", i, va.GetInt())
	}

}
