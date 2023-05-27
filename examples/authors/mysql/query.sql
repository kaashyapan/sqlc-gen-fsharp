/* name: GetAuthor :one */
SELECT * FROM authors
WHERE id = @id LIMIT 1;

/* name: ListAuthors :many */
SELECT * FROM authors
ORDER BY name;

/* name: CreateAuthor :execresult */
INSERT INTO authors (
  name, bio
) VALUES (
  @name, @bio
);

/* name: DeleteAuthor :exec */
DELETE FROM authors
WHERE id = @id;
