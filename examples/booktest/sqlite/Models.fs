// Code generated by sqlc. DO NOT EDIT.
// version: sqlc v1.18.0

namespace Booktest

open System
open Fumble

type Author = { AuthorId: int; Name: string }

type Book =
    { BookId: int
      AuthorId: int
      Isbn: string
      BookType: string
      Title: string
      Yr: int
      Available: DateTime
      Tags: string }

type BooksByTagsRow =
    { BookId: int
      Title: string
      Name: string
      Isbn: string
      Tags: string }