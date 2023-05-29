// Code generated by sqlc. DO NOT EDIT.
// version: sqlc v1.18.0

namespace SAuthors

open System
open Fumble

type Author =
  {
    Id: int
    Name: string
    Bio: string option
    Address: string option
    DateOfBirth: DateTime option
    LastTs: DateTimeOffset option
    SavingsAmt: double option
    LoanAmt: decimal option
    Disabled: bool option
    Married: bool option
    Payable: decimal option
  }

type GetAuthor2Row = { Id: int; Name: string; Bio: string option }

type CountAuthorsNamedRow = { NoOfAuthors: int; Total: double option }
