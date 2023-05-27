/* name: GetAuthor :one */
/* This SQL will select a single author from the table */
SELECT * FROM authors
WHERE id = @id LIMIT 1;

/* name: ListAuthors :many */
/* This SQL will list all authors from the authors table */
SELECT * FROM authors
ORDER BY name;

/* name: CreateAuthor :execresult */
/* This SQL will insert a single author into the table */
INSERT INTO authors (
  name, bio
) VALUES (
  @name, @bio
);

/* name: DeleteAuthor :exec */
/* This SQL will delete a given author */
DELETE FROM authors
WHERE id = @id;
