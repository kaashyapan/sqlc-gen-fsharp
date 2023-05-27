-- name: GetAuthor :one
SELECT * FROM authors
WHERE id = @id LIMIT 1;

-- name: ListAuthors :many
SELECT first_name, last_name, dead FROM authors
ORDER BY first_name;

-- name: CreateAuthor :one
INSERT INTO authors (
  first_name, 
  ssid, 
  middle_name,
  last_name, 
  avatar ,
  dead ,
  disabled ,
  address ,
  country ,
  spouses ,
  children ,
  grandchildren ,
  bio,  
  savings_acct ,
  loan_acct ,
  deposit_acct ,
  book_count ,
  date_of_birth ,
  t_1 ,
  t_2 ,
  ts_1 ,
  ts_2 ,
  passport_id , 
  metadata , 
  metadatab,
  col_fl,
  col_real ,
  col_dbl ,
  col_fl8 
) VALUES (
  @first_name, 
  @ssid, 
  @middle_name,
  @last_name, 
  @avatar ,
  @dead ,
  @disabled ,
  @address ,
  @country ,
  @spouses , 
  @children ,
  @grandchildren ,
  @bio, 
  @savings_acct ,
  @loan_acct ,
  @deposit_acct ,
  @book_count ,
  @date_of_birth ,
  @t_1 ,
  @t_2 ,
  @ts_1 ,
  @ts_2 ,
  @passport_id,
  @metadata ,
  @metadatab , 
  @col_fl ,
  @col_real ,
  @col_dbl ,
  @col_fl8 

)
RETURNING *;

-- name: DeleteAuthor :exec
DELETE FROM authors
WHERE id = @id;
