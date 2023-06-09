// Code generated by sqlc. DO NOT EDIT.
// version: sqlc v1.18.0
// for nuget package Npgsql.FSharp --version 5.7.0

namespace PAuthors

open System
open Npgsql
open Npgsql.FSharp
open PAuthors.Readers

module Sqls =

  [<Literal>]
  let getAuthor =
    """
    SELECT id, ssid, first_name, middle_name, last_name, avatar, dead, disabled, address, country, spouses, children, grandchildren, bio, savings_acct, loan_acct, deposit_acct, book_count, date_of_birth, t_1, t_2, ts_1, ts_2, passport_id, metadata, metadatab, col_fl, col_real, col_dbl, col_fl8 FROM authors
WHERE id = @id LIMIT 1
  """

  [<Literal>]
  let listAuthors =
    """
    SELECT first_name, last_name, dead FROM authors
ORDER BY first_name
  """

  [<Literal>]
  let createAuthor =
    """
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
RETURNING id, ssid, first_name, middle_name, last_name, avatar, dead, disabled, address, country, spouses, children, grandchildren, bio, savings_acct, loan_acct, deposit_acct, book_count, date_of_birth, t_1, t_2, ts_1, ts_2, passport_id, metadata, metadatab, col_fl, col_real, col_dbl, col_fl8
  """

  [<Literal>]
  let deleteAuthor =
    """
    DELETE FROM authors
WHERE id = @id
  """

  [<Literal>]
  let countAuthors =
    """
    SELECT count(*) as cnt from authors
  """

  [<Literal>]
  let totalBooks =
    """
    SELECT count(*) as cnt, sum(book_count) as total_books from authors
  """

  [<Literal>]
  let currentTime =
    """
    SELECT current_timestamp :: TIMESTAMP WITH TIME ZONE as current_time
  """

  [<Literal>]
  let dbString =
    """
    SELECT 'Hello world' as str
  """

[<RequireQualifiedAccessAttribute>]
type DB(conn: string) =
  // https://www.connectionstrings.com/npgsql

  member this.getAuthor(id: int64) =

    let parameters = [ ("id", Sql.int64 id) ]

    conn
    |> Sql.connect
    |> Sql.query Sqls.getAuthor
    |> Sql.parameters parameters
    |> Sql.executeRowAsync authorReader

  member this.listAuthors() =

    conn
    |> Sql.connect
    |> Sql.query Sqls.listAuthors
    |> Sql.executeAsync listAuthorsRowReader

  member this.createAuthor
    (
      firstName: string,
      bookCount: int,
      ?ssid: int64,
      ?middleName: string,
      ?lastName: string,
      ?avatar: byte[],
      ?dead: bool,
      ?disabled: bool,
      ?address: string,
      ?country: string,
      ?spouses: int,
      ?children: int,
      ?grandchildren: int,
      ?bio: string,
      ?savingsAcct: decimal,
      ?loanAcct: decimal,
      ?depositAcct: decimal,
      ?dateOfBirth: DateOnly,
      ?t1: DateTimeOffset,
      ?t2: TimeSpan,
      ?ts1: DateTimeOffset,
      ?ts2: DateTime,
      ?passportId: Guid,
      ?metadata: string,
      ?metadatab: string,
      ?colFl: double,
      ?colReal: double,
      ?colDbl: double,
      ?colFl8: double
    ) =

    let parameters =
      [
        ("first_name", Sql.string firstName)
        ("ssid", Sql.int64OrNone ssid)
        ("middle_name", Sql.stringOrNone middleName)
        ("last_name", Sql.stringOrNone lastName)
        ("avatar", Sql.byteaOrNone avatar)
        ("dead", Sql.boolOrNone dead)
        ("disabled", Sql.boolOrNone disabled)
        ("address", Sql.stringOrNone address)
        ("country", Sql.stringOrNone country)
        ("spouses", Sql.intOrNone spouses)
        ("children", Sql.intOrNone children)
        ("grandchildren", Sql.intOrNone grandchildren)
        ("bio", Sql.textOrNone bio)
        ("savings_acct", Sql.decimalOrNone savingsAcct)
        ("loan_acct", Sql.decimalOrNone loanAcct)
        ("deposit_acct", Sql.decimalOrNone depositAcct)
        ("book_count", Sql.int bookCount)
        ("date_of_birth", Sql.dateOrNone dateOfBirth)
        ("t_1", Sql.timestamptzOrNone t1)
        ("t_2", Sql.intervalOrNone t2)
        ("ts_1", Sql.timestamptzOrNone ts1)
        ("ts_2", Sql.timestampOrNone ts2)
        ("passport_id", Sql.uuidOrNone passportId)
        ("metadata", Sql.jsonbOrNone metadata)
        ("metadatab", Sql.jsonbOrNone metadatab)
        ("col_fl", Sql.doubleOrNone colFl)
        ("col_real", Sql.doubleOrNone colReal)
        ("col_dbl", Sql.doubleOrNone colDbl)
        ("col_fl8", Sql.doubleOrNone colFl8)
      ]

    conn
    |> Sql.connect
    |> Sql.query Sqls.createAuthor
    |> Sql.parameters parameters
    |> Sql.executeRowAsync authorReader

  member this.deleteAuthor(id: int64) =

    let parameters = [ ("id", Sql.int64 id) ]

    conn
    |> Sql.connect
    |> Sql.query Sqls.deleteAuthor
    |> Sql.parameters parameters
    |> Sql.executeNonQueryAsync

  member this.countAuthors() =

    conn
    |> Sql.connect
    |> Sql.query Sqls.countAuthors
    |> Sql.executeRowAsync (fun r -> r.int64 "cnt")

  member this.totalBooks() =

    conn
    |> Sql.connect
    |> Sql.query Sqls.totalBooks
    |> Sql.executeRowAsync totalBooksRowReader

  member this.currentTime() =

    conn
    |> Sql.connect
    |> Sql.query Sqls.currentTime
    |> Sql.executeRowAsync (fun r -> r.datetimeOffset "current_time")

  member this.dbString() =

    conn
    |> Sql.connect
    |> Sql.query Sqls.dbString
    |> Sql.executeRowAsync (fun r -> r.text "str")
