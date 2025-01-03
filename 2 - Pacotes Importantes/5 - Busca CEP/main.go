package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

type ViaCEPResponse struct {
	Cep         string `json:"cep"`
	Logradouro  string `json:"logradouro"`
	Complemento string `json:"complemento"`
	Unidade     string `json:"unidade"`
	Bairro      string `json:"bairro"`
	Localidade  string `json:"localidade"`
	Uf          string `json:"uf"`
	Estado      string `json:"estado"`
	Regiao      string `json:"regiao"`
	Ibge        string `json:"ibge"`
	Gia         string `json:"gia"`
	Ddd         string `json:"ddd"`
	Siafi       string `json:"siafi"`
}

func main() {

	for _, cep := range os.Args[1:] {

		req, e1 := http.Get("http://viacep.com.br/ws/" + cep + "/json/")
		panicIfError(e1)

		defer req.Body.Close()

		response, e2 := io.ReadAll(req.Body)
		panicIfError(e2)

		var viaCEP ViaCEPResponse
		e3 := json.Unmarshal(response, &viaCEP)
		panicIfError(e3)

		file, e4 := os.Create(cep + ".txt")
		panicIfError(e4)
		defer file.Close()

		file.WriteString(string(response))

		fmt.Println(viaCEP)
	}

}

func panicIfError(err error) {
	if err != nil {
		fmt.Fprint(os.Stderr, "Ocorreu algum erro:")
		panic(err)
	}
}
