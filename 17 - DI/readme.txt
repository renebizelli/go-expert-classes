Libs de DI
- uber-go/tx - https://uber-go.github.io/fx/get-started 
usa reflection

- google/wire - https://github.com/google/wire
não usa reflection, gera código.

Notations necessários no arquivo wire.go
// go:build wireinject
// +build wireinject
Essa notations serão lidas pelo comando do wire.

Comando: wire
com isso, gera um arquivo chamado wire_gen.go
