-- name: ListCategories :many
SELECT * FROM categories;

--name: GetCategory :one
SELECT * FROM categories WHERE id = ?;

--name: GetCategory :exe (retorna apenas erro)
--name: GetCategory :exeresult (retora o resulta do da execução)
INSERT INTO categories (id, name, description) VALUES (?,?,?)