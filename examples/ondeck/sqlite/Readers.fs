// Code generated by sqlc. DO NOT EDIT.
// version: sqlc v1.18.0
// for nuget package Fumble --version 1.0.1

namespace Ondeck

open System
open Fumble

module Readers =

  let cityReader (r: RowReader) : City = { City.Slug = r.string "slug"; Name = r.string "name" }

  let venueReader (r: RowReader) : Venue =
    {
      Venue.Id = r.int "id"
      Status = r.string "status"
      Statuses = r.stringOrNone "statuses"
      Slug = r.string "slug"
      Name = r.string "name"
      City = r.string "city"
      SpotifyPlaylist = r.string "spotify_playlist"
      SongkickId = r.stringOrNone "songkick_id"
      Tags = r.stringOrNone "tags"
      CreatedAt = r.dateTime "created_at"
    }

  let venueCountByCityRowReader (r: RowReader) : VenueCountByCityRow =
    { VenueCountByCityRow.City = r.string "city"; Count = r.int "count" }
