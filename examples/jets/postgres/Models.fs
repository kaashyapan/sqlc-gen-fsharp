// Code generated by sqlc. DO NOT EDIT.
// version: sqlc v1.18.0

namespace Jets

open System
open Npgsql

type Pilot = { Id: int; Name: string }

type Jet =
    { Id: int
      PilotId: int
      Age: int
      Name: string
      Color: string }

type Language = { Id: int; Language: string }

type PilotLanguage = { PilotId: int; LanguageId: int }