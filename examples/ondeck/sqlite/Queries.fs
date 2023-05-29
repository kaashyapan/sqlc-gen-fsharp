// Code generated by sqlc. DO NOT EDIT.
// version: sqlc v1.18.0
// for nuget package Fumble --version 1.0.1

namespace Ondeck

open System
open Fumble
open Ondeck.Readers

module Sqls =

  [<Literal>]
  let listCities =
    """
    SELECT slug, name
FROM city
ORDER BY name
  """

  [<Literal>]
  let getCity =
    """
    SELECT slug, name
FROM city
WHERE slug = ?
  """

  [<Literal>]
  let createCity =
    """
    INSERT INTO city (
    name,
    slug
) VALUES (
    ?,
    ? 
)
  """

  [<Literal>]
  let updateCityName =
    """
    UPDATE city
SET name = ?
WHERE slug = ?
  """

  [<Literal>]
  let listVenues =
    """
    SELECT id, status, statuses, slug, name, city, spotify_playlist, songkick_id, tags, created_at
FROM venue
WHERE city = ?
ORDER BY name
  """

  [<Literal>]
  let deleteVenue =
    """
    DELETE FROM venue
WHERE slug = ? AND slug = ?
  """

  [<Literal>]
  let getVenue =
    """
    SELECT id, status, statuses, slug, name, city, spotify_playlist, songkick_id, tags, created_at
FROM venue
WHERE slug = ? AND city = ?
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
    ?,
    ?,
    ?,
    CURRENT_TIMESTAMP,
    ?,
    ?,
    ?,
    ?
)
  """

  [<Literal>]
  let updateVenueName =
    """
    UPDATE venue
SET name = ?
WHERE slug = ?
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
type DB(conn: string) =
  // https://www.connectionstrings.com/sqlite-net-provider

  member this.listCities() =

    conn |> Sql.connect |> Sql.query Sqls.listCities |> Sql.execute cityReader

  member this.getCity(slug: string) =

    let parameters = [ ("slug", Sql.string slug) ]

    conn
    |> Sql.connect
    |> Sql.query Sqls.getCity
    |> Sql.parameters parameters
    |> Sql.execute cityReader

  member this.createCity(name: string, slug: string) =

    let parameters = [ ("name", Sql.string name); ("slug", Sql.string slug) ]

    conn
    |> Sql.connect
    |> Sql.query Sqls.createCity
    |> Sql.parameters parameters
    |> Sql.executeNonQuery

  member this.updateCityName(name: string, slug: string) =

    let parameters = [ ("name", Sql.string name); ("slug", Sql.string slug) ]

    conn
    |> Sql.connect
    |> Sql.query Sqls.updateCityName
    |> Sql.parameters parameters
    |> Sql.executeNonQuery

  member this.listVenues(city: string) =

    let parameters = [ ("city", Sql.string city) ]

    conn
    |> Sql.connect
    |> Sql.query Sqls.listVenues
    |> Sql.parameters parameters
    |> Sql.execute venueReader

  member this.deleteVenue(slug: string, slug_2: string) =

    let parameters = [ ("slug", Sql.string slug); ("slug", Sql.string slug_2) ]

    conn
    |> Sql.connect
    |> Sql.query Sqls.deleteVenue
    |> Sql.parameters parameters
    |> Sql.executeNonQuery

  member this.getVenue(slug: string, city: string) =

    let parameters = [ ("slug", Sql.string slug); ("city", Sql.string city) ]

    conn
    |> Sql.connect
    |> Sql.query Sqls.getVenue
    |> Sql.parameters parameters
    |> Sql.execute venueReader

  member this.createVenue
    (
      slug: string,
      name: string,
      city: string,
      spotifyPlaylist: string,
      status: string,
      ?statuses: string,
      ?tags: string
    ) =

    let parameters =
      [
        ("slug", Sql.string slug)
        ("name", Sql.string name)
        ("city", Sql.string city)
        ("spotify_playlist", Sql.string spotifyPlaylist)
        ("status", Sql.string status)
        ("statuses", Sql.stringOrNone statuses)
        ("tags", Sql.stringOrNone tags)
      ]

    conn
    |> Sql.connect
    |> Sql.query Sqls.createVenue
    |> Sql.parameters parameters
    |> Sql.executeNonQuery

  member this.updateVenueName(name: string, slug: string) =

    let parameters = [ ("name", Sql.string name); ("slug", Sql.string slug) ]

    conn
    |> Sql.connect
    |> Sql.query Sqls.updateVenueName
    |> Sql.parameters parameters
    |> Sql.executeNonQuery

  member this.venueCountByCity() =

    conn
    |> Sql.connect
    |> Sql.query Sqls.venueCountByCity
    |> Sql.execute venueCountByCityRowReader
