// Code generated by sqlc. DO NOT EDIT.
// version: sqlc v1.18.0

namespace Authors

open System

type Author =
    { Id: int64
      Name: string
      Bio: string option }