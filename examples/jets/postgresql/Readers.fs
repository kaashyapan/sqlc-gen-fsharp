// Code generated by sqlc. DO NOT EDIT.
// version: sqlc v1.18.0
// for nuget package Npgsql.FSharp --version 5.7.0

namespace Jets

open System
open Npgsql
open Npgsql.FSharp

module Readers =

  let pilotReader (r: RowReader) : Pilot = { Pilot.Id = r.int "id"; Name = r.text "name" }
