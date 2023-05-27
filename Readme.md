# Sqlc plugin for F# 
## Codegen F# from SQL
`sqlc` is a command line program that generates type-safe database access code from SQL.\
Sqlc documentation - https://sqlc.dev

**Inputs**
  - DB schema.sql file
  - File containing SQL statements
  - Configuration file. 

**Outputs**
  - Models as F# data structures
  - Queries as functions taking named-typed parameters
  - Readers to decode DB response into F# data structures 


| Target    |  Library          |    |
|-----------|-------------------|----|
|Postgres   |`Npgsql.FSharp`    |   |
|MySql      | Not supported     | Models will be generated|
|Sqlite     |`Fumble`           |   |

## Why this ?
Type safe DB access in F# is tedious with manually written data structures.\
SqlHydra is a great dotnet tool to generate F# boiler plate. Works great with ORMs.\
I found I was writing a lot of custom SQL and wanted a solution that can generate 90% of the code.
  
This is intended for devs who prefer to write SQL by hand. 

|SqlHydra  | Sqlc|
|-----------|-------------------|
|Uses a connection to the database to generate data structures| Uses schema file and SQL files|
|Postgres, Oracle, MSSql & Sqlite | Postgres & Sqlite |
|SqlHydra.Query uses Sqlkata | Handwritten Sql |
|Wraps Microsoft.Data.SqlClient. Flexible. Bring your own ADO.net wrapper| Wraps higher level F# libraries. Opinionated. Less generated code. |
|Cannot introspect queries | Wraps the pg_query Postgres SQL parser. It syntax checks the SQL & DDL statements|
|Handwritten data structures are required for custom queries| Produces exact data structures and readers for custom queries |


## How to use

- Install [Sqlc](https://docs.sqlc.dev/en/latest/overview/install.html)
- Create Schema.sql containing DDL statements. (or generate using pg_dump)
- Create Query.sql containing SQL statements with an annotation like in [docs](https://docs.sqlc.dev/en/latest/reference/query-annotations.html)
    ```sql
    -- name: ListAuthors :many
    SELECT * FROM authors ORDER BY name;
    ```
- Create sqlc.json & configure the options
  ```json
  {
    "version": "2",
    "plugins": [
      {
        "name": "fsharp",
        "wasm": {
          "url": "https://github.com/kaashyapan/sqlc-gen-fsharp/releases/download/latest/sqlc-gen-fsharp_1.0.0.wasm",
          "sha256": "85a42e4f3d70feb8eb725d7406593114723ab474ecd4c74d1f9edb4867515ea4"
        }
      }
    ],
    "sql": [
      {
        "engine": "postgresql",
        "schema": "schema.sql",
        "queries": "query.sql",
        "codegen": [
          {
            "out": <..target_folder...>,
            "plugin": "fsharp",
            "options": {
              "namespace": <...Namespace...>,
              "async": false,
              "type_affinity": true 
            }
          }
        ]
      }
    ]
  }
  ```
- ```sqlc generate```

See the test folder for a sample setup.



### fsharp config options
`namespace`: The namespace to use for the generated code.\
`out`: Output directory for generated code.\
`emit_exact_table_names`: If true, use the exact table name for generated models. Otherwise, guess a singular form. Defaults to *false*.\
`async`: If true, all query functions generated will be async. Defaults to *false*.\
`type_affinity`: If true, all DB integers (except Bigint) will be mapped to F#int. All DB floats will be mapped to F#double. Defaults to *false*.


### TODO
- Support for enumerated column types.
- Postgis type support
- Optionally generate classes instead of records
- Autogenerate basic CRUD without writing SQL
