// Code generated by sqlc. DO NOT EDIT.
// version: sqlc v1.18.0
// for nuget package Fumble --version 1.0.1

namespace Booktest

open System
open Fumble

module Readers =

  let authorReader (r: RowReader) : Author = { Author.AuthorId = r.int "author_id"; Name = r.string "name" }

  let bookReader (r: RowReader) : Book =
    {
      Book.BookId = r.int "book_id"
      AuthorId = r.int "author_id"
      Isbn = r.string "isbn"
      BookType = r.string "book_type"
      Title = r.string "title"
      Yr = r.int "yr"
      Available = r.dateTime "available"
      Tags = r.string "tags"
    }

  let booksByTagsRowReader (r: RowReader) : BooksByTagsRow =
    {
      BooksByTagsRow.BookId = r.int "book_id"
      Title = r.string "title"
      Name = r.string "name"
      Isbn = r.string "isbn"
      Tags = r.string "tags"
    }
