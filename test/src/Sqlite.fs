module Sqlite

open System
open Fumble
open FSharp.Data.LiteralProviders
open SAuthors
open System.IO

[<Literal>]
let initsql = TextFile<"sqlite/schema.sql">.Text

[<Literal>]
let conn = "Data Source=/tmp/sample.db;"

let initiate () =
  conn
  |> Sqlite.connect
  |> Sqlite.query initsql
  |> Sqlite.executeNonQuery
  |> printfn "%A"

let run () =
  let db = SAuthors.DB(conn)

  printfn "\n----------------------------------------------------------------- \n"
  printfn "Initiating Sqlite DB"

  ignore <| initiate ()

  db.listAuthors () |> printfn "List authors - %A"

  db.createAuthor ("Elon Musk", "CEO, CTO") |> printfn "Create authors - %A"

  db.createAuthor ("Jeff Bezos", "Chairman Amazon")
  |> function
    | Ok rows ->
      let r = List.head rows
      db.deleteAuthor (r.Id) |> printfn "Delete authors - %A"

    | Error e -> raise e

  db.listAuthors () |> printfn "List authors - %A"

  db.listAuthors () |> printfn "List authors - %A"
  db.getAuthor (1) |> printfn "Get authors - %A"

  File.Delete("/tmp/sample.db")
  ()
