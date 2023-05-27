// Code generated by sqlc. DO NOT EDIT.
// version: sqlc v1.18.0

namespace SAuthors

open System
open Fumble
open SAuthors.Readers

module Sql = Sqlite
type Sql = Sqlite

module Sqls =

  [<Literal>]
  let createAuthor =
    """
    INSERT INTO authors (
  name, bio
) VALUES (
  @name, @bio
)
RETURNING id, name, bio, address, date_of_birth, last_ts, savings_amt, loan_amt, disabled, married, payable
  """

  [<Literal>]
  let deleteAuthor =
    """
    DELETE FROM authors
WHERE id = @id
  """

  [<Literal>]
  let getAuthor =
    """
    SELECT id, name, bio, address, date_of_birth, last_ts, savings_amt, loan_amt, disabled, married, payable FROM authors
WHERE id = @id LIMIT 1
  """

  [<Literal>]
  let getAuthor2 =
    """
    SELECT id, name, bio FROM authors
WHERE id = @id LIMIT 1
  """

  [<Literal>]
  let listAuthors =
    """
    SELECT id, name, bio, address, date_of_birth, last_ts, savings_amt, loan_amt, disabled, married, payable FROM authors
ORDER BY name
  """

[<RequireQualifiedAccessAttribute>]
type DB(conn: string) =

  // https://www.connectionstrings.com/sqlite-net-provider

  member this.createAuthor(name: string, ?bio: string) =

    let parameters = [ ("name", Sql.string name); ("bio", Sql.stringOrNone bio) ]

    conn
    |> Sql.connect
    |> Sql.query Sqls.createAuthor
    |> Sql.parameters parameters
    |> Sql.execute authorReader

  member this.deleteAuthor(id: int) =

    let parameters = [ ("id", Sql.int id) ]

    conn
    |> Sql.connect
    |> Sql.query Sqls.deleteAuthor
    |> Sql.parameters parameters
    |> Sql.executeNonQuery

  member this.getAuthor(id: int) =

    let parameters = [ ("id", Sql.int id) ]

    conn
    |> Sql.connect
    |> Sql.query Sqls.getAuthor
    |> Sql.parameters parameters
    |> Sql.execute authorReader

  member this.getAuthor2(id: int) =

    let parameters = [ ("id", Sql.int id) ]

    conn
    |> Sql.connect
    |> Sql.query Sqls.getAuthor2
    |> Sql.parameters parameters
    |> Sql.execute getAuthor2RowReader

  member this.listAuthors() =

    conn |> Sql.connect |> Sql.query Sqls.listAuthors |> Sql.execute authorReader