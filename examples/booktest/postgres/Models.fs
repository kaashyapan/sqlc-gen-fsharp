
// Code generated by sqlc. DO NOT EDIT.
// version: sqlc v1.18.0

namespace Booktest

open System
open Npgsql

type Author = 
    {
        AuthorId : int
        Name : string
    }

type Book = 
    {
        BookId : int
        AuthorId : int
        Isbn : string
        BookType : book_type
        Title : string
        Year : int
        Available : DateTimeOffset
        Tags : string
    }

type BooksByTagsRow = 
    {
        BookId : int
        Title : string
        Name : string option
        Isbn : string
        Tags : List&lt;string&gt;
    }
