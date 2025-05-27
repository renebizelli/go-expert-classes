
Migration SQLC
    
    https://sqlc.dev
    - instalar.
    - comando: sqlc

    Importante: somente dá suporte pra Postgree, mysql e sqlite

    https://github.com/golang-migrate/migrate
    - necessário instalar
    comando: migrate

    Comandos:

    gerar migrations
    - migrate create -ext=sql -dir=sql/migrations -seq init  (-seq identificação sequencial)
    -> vai gerar 2 arquivos: 
        00001_init.down.sql
        00001_init.up.sql

    executar migrations
    - migrate -path=sql/migrations -database "mysql://root:root@tcp(localhost:3306)/courses" -verbose up
    - migrate -path=sql/migrations -database "mysql://root:root@tcp(localhost:3306)/courses" -verbose down (rollback)


    Docker
    - docker-compose up -d
    - docker-compose exec mysql bash
    - mysql -uroot -p courses

    Rodar comandos do Make  (Makefile)
    - make migration

Para gerar os arquivos utils:

1 - criar arquivo sqlc.yaml
2 - criar pasta sql/queries
3 - criar arquivo query.sql
4 - coloca queries com as anotations, ex : (-- name: ListCategories :many)
5 - rodar comando: sqlc generate
6 - checar a pasta internal/db com 3 arquivos
    - db.go
    - models.go
    - query.go


o sqlc converte o tipo decimal para string na geração do modelo.
para isso, é inserido um overrides no sqlc.yaml
ex:
      overrides:
      - db_type: "decimal"
        go_type: float64""

Repo do curso
https://github.com/devfullcycle/goexpert/tree/main/17-SQLC
https://github.com/devfullcycle/goexpert

SQLX golang (https://github.com/jmoiron/sqlx)
- Traz um pouco de facilidade pra conseguir utilizar SQL puro
- Não é necessário fazer os For de hidratação de classes

