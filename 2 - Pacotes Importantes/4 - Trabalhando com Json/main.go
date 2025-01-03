package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Person struct {
	Name     string  `json:"n"`
	Nickname string  `json:"nick"`
	Value    float64 `json:"v"`
	Age      int     `json:"-"`
}

/// GORM -> ORM do GO

func main() {

	createJson()
	encoding()
	deserializer()

}

func deserializer() {

	fmt.Println("\ndeserializer")

	var p Person

	jsonString := []byte(`{ "n": "René", "nick": "Dê", "a": 46, "v": 10.2 }`)

	json.Unmarshal(jsonString, &p)

	fmt.Println(p)
}

func encoding() {

	fmt.Println("\nencoding")

	rene := Person{Name: "René", Nickname: "Dê", Age: 46, Value: 10.0}

	encoder := json.NewEncoder(os.Stdout)

	err := encoder.Encode(rene)

	panicIfError(err)

}

func createJson() {

	fmt.Println("\ncreateJson")

	rene := Person{Name: "René", Nickname: "Dê", Age: 46, Value: 10.4}

	json, err := json.Marshal(rene)

	panicIfError(err)

	fmt.Printf("\n %s", json)
}

func panicIfError(err error) {
	if err != nil {
		fmt.Println("Ocorreu algum erro:")
		panic(err)
	}
}
