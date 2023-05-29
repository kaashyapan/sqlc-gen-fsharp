
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

-- name: countAuthors :exec
SELECT count(*) from authors;

-- name: countAuthorsNamed :one
SELECT count(*) as no_of_authors, sum(id) as total from authors;
