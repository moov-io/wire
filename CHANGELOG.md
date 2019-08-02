## v0.3.0 (Unreleased)

ADDITIONS

- cmd/server: initial set of HTTP routes for library methods
- cmd/server: bind HTTP server with TLS if HTTPS_* variables are defined

## v0.2.0 (Released 2019-06-25)

BUG FIXES

- charges: handle fuzz crash
- server: Register HTTP routes

ADDITIONS

- build: push moov/wire:latest on 'make release-push'

UPGRADES

- api,client: openapi-generator 4.0.2 and embed request parameters

## v0.1.0 (Released 2019-06-19)

- Initial release
