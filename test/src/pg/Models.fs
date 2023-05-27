// Code generated by sqlc. DO NOT EDIT.
// version: sqlc v1.18.0

namespace PAuthors

open System
open Npgsql

type Author =
  {
    Id: int64
    Ssid: int64 option
    FirstName: string
    MiddleName: string option
    LastName: string option
    Avatar: byte[] option
    Dead: bool option
    Disabled: bool option
    Address: string option
    Country: string option
    Spouses: int option
    Children: int option
    Grandchildren: int option
    Bio: string option
    SavingsAcct: decimal option
    LoanAcct: decimal option
    DepositAcct: decimal option
    BookCount: int
    DateOfBirth: DateOnly option
    T1: DateTimeOffset option
    T2: TimeSpan option
    Ts1: DateTimeOffset option
    Ts2: DateTime option
    PassportId: Guid option
    Metadata: string option
    Metadatab: string option
    ColFl: double option
    ColReal: double option
    ColDbl: double option
    ColFl8: double option
  }

type ListAuthorsRow = { FirstName: string; LastName: string option; Dead: bool option }