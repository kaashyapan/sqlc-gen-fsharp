-- name: GetEmployee :one
SELECT * FROM employees
WHERE id = @id LIMIT 1;

-- name: ListEmployees :many
SELECT * FROM employees
ORDER BY name;

-- name: CreateEmployee :exec
INSERT INTO employees (
  name, occupation, age
) VALUES (
  @name, @occupation, @age
);

-- name: DeleteEmployee :exec
DELETE FROM employees
WHERE id = @id;
