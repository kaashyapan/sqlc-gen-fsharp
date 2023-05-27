-- Example queries for sqlc
CREATE TABLE authors (
  id   INTEGER PRIMARY KEY,
  name TEXT NOT NULL,
  bio  TEXT,
  address  TEXT,
  date_of_birth DATE,
  last_ts TIMESTAMP,
  savings_amt REAL,
  loan_amt NUMERIC,
  disabled BOOLEAN,
  married BOOL,
  payable DECIMAL
);
