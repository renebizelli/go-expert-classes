go test .
- executa todos os testes do diretório
- Opção -v traz detalhes da execução do teste
- Opção -coverprofile=coverage.out: exibe a porcentagem de cobertura de código
- go tool cover -html=coverage: mostra um HTML indicando partes não cobertas pelo código (utiliza o arquivo "coverage.out" do comando acima)

go help test: help

Arquivo de teste deve possuir o sufixo _test.go (não funciona sem isso)
- Arquivo math.go -> math_test.go;
- package deve ser iguais entre o arquivo e o teste;

Iniciar métodos com o "Test" seguido do nome do método:
- Método a ser testado Calculate -> TestCalculate;

Benchmark
- go test -bench=.

Fuzz
- go test -fuzz=.


Pacote testify

- Testify
	"github.com/stretchr/testify/assert"

- Mocking
	"github.com/stretchr/testify/mock"
