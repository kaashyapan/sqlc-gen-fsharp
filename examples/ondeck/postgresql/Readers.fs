// Code generated by sqlc. DO NOT EDIT.
// version: sqlc v1.18.0
// for nuget package Npgsql.FSharp --version 5.7.0

namespace Ondeck

open System
open Npgsql
open Npgsql.FSharp

module Readers =

  let cityReader (r: RowReader) : City = { City.Slug = r.text "slug"; Name = r.text "name" }

  let venueReader (r: RowReader) : Venue =
    {
      Venue.Id = r.int "id"
      Status = r.status_unhandled_report_issue "status"
      Statuses = r.status_unhandled_report_issue "statuses"
      Slug = r.text "slug"
      Name = r.string "name"
      City = r.text "city"
      SpotifyPlaylist = r.string "spotify_playlist"
      SongkickId = r.textOrNone "songkick_id"
      Tags = r.textOrNone "tags"
      CreatedAt = r.timestamp "created_at"
    }

  let venueCountByCityRowReader (r: RowReader) : VenueCountByCityRow =
    { VenueCountByCityRow.City = r.text "city"; Count = r.int64 "count" }
