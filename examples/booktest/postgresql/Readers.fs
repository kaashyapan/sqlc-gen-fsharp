// Code generated by sqlc. DO NOT EDIT.
// version: sqlc v1.18.0

namespace Booktest

open System
open Npgsql
open Npgsql.FSharp

module Readers =

  let booksByTagsRowReader (r: RowReader) : BooksByTagsRow =
    {
      BooksByTagsRow.BookId = r.int "book_id"
      Title = r.text "title"
      Name = r.textOrNone "name"
      Isbn = r.text "isbn"
      Tags = r.string "tags"
    }

  let bookReader (r: RowReader) : Book =
    {
      Book.BookId = r.int "book_id"
      AuthorId = r.int "author_id"
      Isbn = r.text "isbn"
      BookType = r.unhandled_report_issue "book_type"
      Title = r.text "title"
      Year = r.int "year"
      Available = r.datetimeOffset "available"
      Tags = r.string "tags"
    }

  let authorReader (r: RowReader) : Author = { Author.AuthorId = r.int "author_id"; Name = r.text "name" }