GOLANG-STANDARDS

- Repositório com modelos de organização de projetos 
> github.com/golang-standards/project-layout

*****************************************************************

PASTAS

/internal
É a pasta com códigos específicos da aplicação.

/pkg
Libraries para sem compartilhadas.

/cmd
Onde é gerado o executável, onde fica o main.go.

/configs
Configurações do sistema e boot da aplicação.

/test
Arquivos adicionais de teste, auxiliares. Não necessariamente arquivos .go. 

/api
Guarda especificações da API

*****************************************************************

go-chi 
- é um roteador

*****************************************************************

Swagger (doc: github.com/swaggo/swag)

go install github.com/swaggo/swag/cmd/swag@

// @title GO Expert API Example
// @version 1
// @description This is a sample server Petstore server.
// @host localhost:8080
// @BasePath /
// @schemes http
// @contact.email rene.bizelli@gmail.com
// @contact.name René Bizelli
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization

Gerar a pasta /docs
swag init -g cmd/app/main.go

