package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

type HttpError struct {
	StatusCode int
	Err        string
}

func (e *HttpError) Error() string {
	return fmt.Sprintf("%v", e.Err)
}

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

	http.HandleFunc("/cep/{cep}", buscaCEPHandler)

	http.ListenAndServe(":8080", nil)

}

func buscaCEPHandler(w http.ResponseWriter, r *http.Request) {

	cep := r.PathValue("cep")

	if cep == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	fmt.Println("\n1 >> " + cep)

	viaCEP, err := buscaCEP(cep)

	fmt.Print(viaCEP)

	if err != nil {
		w.WriteHeader(err.StatusCode)
		w.Write([]byte(err.Error()))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	json.NewEncoder(w).Encode(viaCEP)
}

func buscaCEP(cep string) (*ViaCEPResponse, *HttpError) {

	req, err := http.Get("http://viacep.com.br/ws/" + cep + "/json/")

	if err != nil {
		return nil, &HttpError{StatusCode: 500, Err: err.Error()}
	}

	if req.StatusCode != 200 {
		fmt.Printf("\n3 >>>> %v", req.StatusCode)
		return nil, &HttpError{StatusCode: req.StatusCode, Err: req.Status}
	}

	defer req.Body.Close()

	response, errBody := io.ReadAll(req.Body)

	if errBody != nil {
		return nil, &HttpError{StatusCode: 500, Err: errBody.Error()}
	}

	var viaCEP ViaCEPResponse
	errJson := json.Unmarshal(response, &viaCEP)

	if errJson != nil {
		return nil, &HttpError{StatusCode: 500, Err: errJson.Error()}
	}

	return &viaCEP, nil
}

func panicIfError(err error) {
	if err != nil {
		fmt.Fprint(os.Stderr, "Ocorreu algum erro:")
		panic(err)
	}
}
