exemplo, pacote público:
URL: "https://github.com/devfullcycle/fcutils/tree/main/pkg/events"
obs: sempre remover o "tree/main"
-> import "github.com/devfullcycle/fcutils/pkg/events"

go env | grep PRIVATE -
Exibe as variáveis privadas, uma delas é a GOPRIVATE
Na GOPRIVATE é inseridos os repos separados por vírgulas

ex ->  go env -w GOPRIVATE=github.com/renebizelli/fc-utils-secret,...,...

Para autenticar

1 - modo token

criar arquivo na raiz do diretório do user chamado .netrc
machine github.com
login renebizelli@gmail.com
password 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  -> token gerado no github (settings -> dev settings ->)

2 - modo ssh

editar arquivo .git/config
adicionar
[url "ssh://git@github.com/"]
        insteadOf = https://github.com/


go mod vendor
 - baixa todos os pacotes utilizados para o projeto, para garantir que os pacotes estarão disponíveis caso ocorra algum problema com o pacote online.
 