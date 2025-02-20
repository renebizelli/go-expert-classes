package main

import (
	"fmt"

	"github.com/valyala/fastjson"
)

func main() {

	var p fastjson.Parser

	jsonData := `{"user":{"name":"John","age":30,"city":"New York"}, "arr": [1,2,3]}`

	v, err := p.Parse(jsonData)

	if err != nil {
		panic(err)
	}

	user := v.Get("user")

	name := user.GetStringBytes("name")

	age := user.GetInt("age")

	fmt.Printf("Name: %s\n", name)
	fmt.Printf("Age: %d\n", age)

	oUser := v.GetObject("user")

	fmt.Printf("Name: %v\n", oUser.Get("name"))
	fmt.Printf("Age: %s\n", oUser.Get("age"))

	userJson := oUser.String()

	fmt.Printf("User: %s\n", userJson)

}
