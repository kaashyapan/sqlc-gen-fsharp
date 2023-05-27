// Code generated by sqlc. DO NOT EDIT.
// version: sqlc v1.18.0

namespace Authors

open System
open Npgsql
open Npgsql.FSharp
open Authors.Readers

module Sqls =

  [<Literal>]
  let createAuthor =
    """
    INSERT INTO authors (
  name, bio
) VALUES (
  @name, @bio
)
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
    SELECT id, name, bio FROM authors
WHERE id = @id LIMIT 1
  """

  [<Literal>]
  let listAuthors =
    """
    SELECT id, name, bio FROM authors
ORDER BY name
  """

[<RequireQualifiedAccessAttribute>]
type DB(conn: string) =
  // https://www.connectionstrings.com/npgsql

  /// This SQL will insert a single author into the table
  member this.createAuthor(name: string, ?bio: string) =

    let parameters = [ ("name", Sql.text name); ("bio", Sql.textOrNone bio) ]

    conn
    |> Sql.connect
    |> Sql.query Sqls.createAuthor
    |> Sql.parameters parameters
    |> Sql.executeNonQuery

  /// This SQL will delete a given author
  member this.deleteAuthor(id: int64) =

    let parameters = [ ("id", Sql.int64 id) ]

    conn
    |> Sql.connect
    |> Sql.query Sqls.deleteAuthor
    |> Sql.parameters parameters
    |> Sql.executeNonQuery

  /// This SQL will select a single author from the table
  member this.getAuthor(id: int64) =

    let parameters = [ ("id", Sql.int64 id) ]

    conn
    |> Sql.connect
    |> Sql.query Sqls.getAuthor
    |> Sql.parameters parameters
    |> Sql.executeRow authorReader

  /// This SQL will list all authors from the authors table
  member this.listAuthors() =

    conn |> Sql.connect |> Sql.query Sqls.listAuthors |> Sql.execute authorReader
