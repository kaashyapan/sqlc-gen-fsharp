module SqlcTest

open System

[<EntryPoint>]
let main args =

  Postgres.run ()
  Sqlite.run ()
  0
