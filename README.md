# Rockup/Go Std

Rockup/Go utilities and extensions

## Overview
Rockup Go StdLib is a collection of extensions and commonly used helpers. There are three main categories
that you should be aware of:
- core: helper packages that provide useful functionality
- x: extensions for standard library packages and client drivers
- beta: packages that do not fall under the semantic versioning guarantee, they are still being evaluated

Further, there is a collection of common patterns under internal/examples. We also keep deprecated packages under internal/archive.




## Approved Third Party Packages
The following is a list of packages that have been approved for usage within Rockup projects.

- github.com/dimfeld/httptreemux (remove once [this](https://github.com/golang/go/issues/61410) approval is through)
- github.com/jackc/pgx
- github.com/neo4j/neo4j-go-driver/v5
- github.com/dgraph-io/badger
- github.com/blugelabs/bluge
- github.com/mitchellh/mapstructure
- github.com/carlmjohnson/requests
- github.com/gorilla/schema
- github.com/biter777/countries
- github.com/dustin/go-humanize
- github.com/docker/go-units
- github.com/alecthomas/chroma
- github.com/go-telegram-bot-api/telegram-bot-api/v5

## Evaluating Packages
The following is a list of packages that are under evaluation for common project usage approval.

- github.com/yuin/goldmark (I think this package is better, but will still have to play a bit)
- github.com/gomarkdown/markdown