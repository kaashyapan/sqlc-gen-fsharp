module Postgres

open System
open System.Text
open Npgsql
open Npgsql.FSharp
open PAuthors
open FSharp.Data.LiteralProviders

module Async =
  let map f workflow =
    task {
      let! res = workflow
      return f res
    }

[<Literal>]
let initsql = TextFile<"postgres/schema.sql">.Text

[<Literal>]
let conn = "Server=localhost;Port=5432;Database=postgres;User Id=sqlc;Password=example;"

let initiate () =
  let c = conn |> Sql.connect |> Sql.createConnection
  c.Open()
  let cmd = new NpgsqlCommand(initsql, c)
  printfn "%A" <| cmd.ExecuteNonQuery()

let run () =
  let db = PAuthors.DB(conn)

  printfn "\n-------------------------------------------------------------------- \n"
  printfn "Initiating postgres DB"

  ignore <| initiate ()

  task {
    do! db.listAuthors () |> Async.map (printfn "List authors - %A")

    do! db.createAuthor ("Jeff Bezos", 5) |> Async.map (printfn "Create authors - %A")

    do!
      db.createAuthor (
        firstName = "Elon",
        bookCount = 2,
        ssid = int64 868,
        middleName = "E",
        lastName = "musk",
        dead = false,
        avatar = Encoding.UTF8.GetBytes("avatar"),
        disabled = false,
        address = "California",
        country = "USA",
        spouses = 2,
        bio = "Twitter CEO",
        savingsAcct = decimal 434.23,
        children = 4,
        grandchildren = 4, //int16 4,
        loanAcct = decimal 234.57,
        depositAcct = decimal 89.0,
        dateOfBirth = DateOnly.FromDateTime(DateTime.Today),
        t2 = TimeSpan.FromTicks(34567),
        t1 = DateTimeOffset.UtcNow,
        ts1 = DateTimeOffset.UtcNow,
        ts2 = DateTime.Now,
        passportId = Guid.NewGuid(),
        metadata = """{"key" : "value"}""",
        metadatab = """{"key" : "value"}""",
        colFl = 3.1475 , // float32 3.1475,
        colReal = 3.1475, //float32 3.1475,
        colDbl = double 3.1475,
        colFl8 = double 3.1475
      )

      |> Async.map (printfn "Create authors - %A")

    //do! db.listAuthors () |> Async.map (printfn "List authors - %A")
    let! author = db.getAuthor 1
    printfn "Get authors - %A" author
    do! db.deleteAuthor (author.Id) |> Async.map (printfn "Delete authors - %A")

  }
  |> Async.AwaitTask
  |> Async.RunSynchronously

  ()
