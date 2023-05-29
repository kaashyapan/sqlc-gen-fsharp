
-- name: GetAuthor2 :one
SELECT id, name, bio FROM authors
WHERE id = @id LIMIT 1;

-- name: GetAuthor :one
SELECT * FROM authors
WHERE id = @id LIMIT 1;

-- name: ListAuthors :many
SELECT * FROM authors
ORDER BY name;

-- name: CreateAuthor :one
INSERT INTO authors (
  name, bio
) VALUES (
  @name, @bio
)
RETURNING *;

-- name: DeleteAuthor :exec
DELETE FROM authors
WHERE id = @id;

-- name: CountAuthors :one
SELECT count(*) as cnt from authors;

-- name: TotalBooks :one
SELECT count(*) as cnt, sum(id) as total_books from authors;

-- name: DbString :one
SELECT 'Hello world' as str ;

