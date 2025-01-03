Inicializar módulo

go mod init NAME
- NAME pode ser qualquer valor, mas por convensão é um nome como pacote (endereço git)

ex: go mod init github.com/renebizelli/aulas-go

go mod tidy
- atualiza o arqruivo .mod
- opção -e ignora pacotes não encontrados 

go work init ./Package1 ./Package2

