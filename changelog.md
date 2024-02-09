# Rockup Go Standard Library Changelog

## v1.1.0 - go1.22
- added package v(alidate)
- added package form
- remove dimfield treemux
- remove a bunch of outdated dependencies
- upgrade dependencies

## v1.0.7
None - Force Upgrade

## v1.0.6
01/02/2024

- added DefaultStaticCache to package web
- removed DefaultCacheSeconds from package web

## v1.0.5
30/01/2024

- remove some outdated internals in package web (gzip)
- created a gzip middleware

## v1.0.4
- archived pgxx.
- archived redx.
- randx panics instead of returning error, no way to handle error cases.

## v1.0.3
14/12/2023

- Move package core/conf to internal/archive, prefer ardanlabs conf.
- Move package core/async to internal/archive, prefer errpool
- Small breaking changes to package core/cli:
    - removed map from implementation, making the implementation more lightweight. (not breaking)
    - remove Menu App return.
    - add "proxy" methods on App to enable setting fields on Menu.
- Move package x/fsx to internal/archive, unused and there are better solutions.

## v1.0.0
24/09/2023

- Version 1.0.0
