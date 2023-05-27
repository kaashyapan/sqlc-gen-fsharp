
// Code generated by sqlc. DO NOT EDIT.
// version: sqlc v1.18.0

namespace Ondeck

open System
open Npgsql
open Npgsql.FSharp
open Ondeck.Readers

module Sqls = 

  [<Literal>]
  let createCity =
    """
    INSERT INTO city (
    name,
    slug
) VALUES (
    @name,
    @slug
) RETURNING slug, name
  """

  [<Literal>]
  let createVenue =
    """
    INSERT INTO venue (
    slug,
    name,
    city,
    created_at,
    spotify_playlist,
    status,
    statuses,
    tags
) VALUES (
    @slug,
    @name,
    @city,
    NOW(),
    @spotify_playlist,
    @status,
    @statuses,
    @tags
) RETURNING id
  """

  [<Literal>]
  let deleteVenue =
    """
    DELETE FROM venue
WHERE slug = @slug AND slug = @slug
  """

  [<Literal>]
  let getCity =
    """
    SELECT slug, name
FROM city
WHERE slug = @slug
  """

  [<Literal>]
  let getVenue =
    """
    SELECT id, status, statuses, slug, name, city, spotify_playlist, songkick_id, tags, created_at
FROM venue
WHERE slug = @slug AND city = @city
  """

  [<Literal>]
  let listCities =
    """
    SELECT slug, name
FROM city
ORDER BY name
  """

  [<Literal>]
  let listVenues =
    """
    SELECT id, status, statuses, slug, name, city, spotify_playlist, songkick_id, tags, created_at
FROM venue
WHERE city = @city
ORDER BY name
  """

  [<Literal>]
  let updateCityName =
    """
    UPDATE city
SET name = @name
WHERE slug = @slug
  """

  [<Literal>]
  let updateVenueName =
    """
    UPDATE venue
SET name = @name
WHERE slug = @slug
RETURNING id
  """

  [<Literal>]
  let venueCountByCity =
    """
    SELECT
    city,
    count(*)
FROM venue
GROUP BY 1
ORDER BY 1
  """

[<RequireQualifiedAccessAttribute>]
type DB (conn: string) =
  // https://www.connectionstrings.com/npgsql

  /// Create a new city. The slug must be unique.
  /// This is the second line of the comment
  /// This is the third line
  member this.createCity  (name: string, slug: string) =

    let parameters = [ ("name", Sql.text name); ("slug", Sql.text slug) ]
    
    conn
    |> Sql.connect
    |> Sql.query Sqls.createCity
    |> Sql.parameters parameters
    |> Sql.executeRow cityReader

  member this.createVenue  (slug: string, name: string, city: string, spotifyPlaylist: string, status: status, ?statuses: List&lt;status&gt;, ?tags: List&lt;string option&gt;) =

    let parameters = [ ("slug", Sql.text slug); ("name", Sql.string name); ("city", Sql.text city); ("spotify_playlist", Sql.string spotifyPlaylist); ("status", Sql.unhandled_report_issue status); ("statuses", Sql.unhandled_report_issue statuses); ("tags", Sql.textOrNone tags) ]
    
    conn
    |> Sql.connect
    |> Sql.query Sqls.createVenue
    |> Sql.parameters parameters
    |> Sql.executeRow intReader

  member this.deleteVenue  (slug: string) =

    let parameters = [ ("slug", Sql.text slug) ]
    
    conn
    |> Sql.connect
    |> Sql.query Sqls.deleteVenue
    |> Sql.parameters parameters
    |> Sql.executeNonQuery

  member this.getCity  (slug: string) =

    let parameters = [ ("slug", Sql.text slug) ]
    
    conn
    |> Sql.connect
    |> Sql.query Sqls.getCity
    |> Sql.parameters parameters
    |> Sql.executeRow cityReader

  member this.getVenue  (slug: string, city: string) =

    let parameters = [ ("slug", Sql.text slug); ("city", Sql.text city) ]
    
    conn
    |> Sql.connect
    |> Sql.query Sqls.getVenue
    |> Sql.parameters parameters
    |> Sql.executeRow venueReader

  member this.listCities  () =

    conn
    |> Sql.connect
    |> Sql.query Sqls.listCities
    |> Sql.execute cityReader

  member this.listVenues  (city: string) =

    let parameters = [ ("city", Sql.text city) ]
    
    conn
    |> Sql.connect
    |> Sql.query Sqls.listVenues
    |> Sql.parameters parameters
    |> Sql.execute venueReader

  member this.updateCityName  (slug: string, name: string) =

    let parameters = [ ("slug", Sql.text slug); ("name", Sql.text name) ]
    
    conn
    |> Sql.connect
    |> Sql.query Sqls.updateCityName
    |> Sql.parameters parameters
    |> Sql.executeNonQuery

  member this.updateVenueName  (slug: string, name: string) =

    let parameters = [ ("slug", Sql.text slug); ("name", Sql.string name) ]
    
    conn
    |> Sql.connect
    |> Sql.query Sqls.updateVenueName
    |> Sql.parameters parameters
    |> Sql.executeRow intReader

  member this.venueCountByCity  () =

    conn
    |> Sql.connect
    |> Sql.query Sqls.venueCountByCity
    |> Sql.execute venueCountByCityRowReader
