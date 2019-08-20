## v0.3.0 (Released 2019-08-20)

BREAKING CHANGE

In our OpenAPI we've renamed fields generated as `Id` to `ID`, which is more in-line with Go's style conventions.

ADDITIONS

- cmd/server: initial set of HTTP routes for library methods
- cmd/server: bind HTTP server with TLS if HTTPS_* variables are defined
- cmd/server: record metrics when files are created/deleted

BUG FIXES

- all: check line lengths with RuneCountInString in Parse calls

BUILD

- build: upgrade openapi-generator to 4.1.0
- cmd/server: update github.com/moov-io/base to v0.10.0

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
