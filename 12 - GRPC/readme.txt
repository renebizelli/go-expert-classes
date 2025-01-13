gRPC-GO

Instalando compilador e plugins

https://grpc.io/docs/languages/go/quickstart/

Precisamos ter o compilador do protocol buffer (protoc). 

apos instalar testar com o comando:
protoc --version
-> libprotoc 3.X  (a versão deve ser >=3)

- é preciso 2 libs instaladas (comando no quick start)
 - - protoc-gen-go: para gerar entidades
 - - protoc-gen-go-gRPC: gerar todas interfaces de comunicação

Plugin VS para protobuffer
- vscode-proto3

É muito comum os package dos protos serem todos: package pb;

Comando para gerar os protofiles:
- protoc --go_out=. --go-grpc_out=. proto/course_category.proto
. é o contexto onde será guardado o arquivo
--go_out -> chamada do plugin pro arquivo Go
--go-grpc_out -> gerar arquivos e interfaces 

COm isso, vai gerar a pasta pb com os arquivos gerados.

Evans - Client de gRPC
http://github.com/ktr0731/Evans

install -> go install github.com/ktr0731/evans@latest

Para rodar :evans -r repl

É necessário "entrar" no service, então rtodar os comandos:
- package pb
- service CategoryService

Para chamar um servico:
- call CreateCategory (próximos passos será prompts para inserir os parâmetros do request)