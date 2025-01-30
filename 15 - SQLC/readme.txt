https://github.com/golang-migrate/migrate


Ir na documentation para ver com instalar.

comandos:

#criar migration
migrate create -ext=sql -dir=sql/migrations -seq init
--> -seq = sequencial
--> init = nome da migration

#colocar os comandos SQL nos arquivos gerados (ele nao gera o SQL sozinho)


#rodar migration
migrate -path=sql/migrations -database="mysql://root:root@tcp(localhost:3306)/orders" up





