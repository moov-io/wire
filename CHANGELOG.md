## v0.6.1 (Released 2020-10-19)

BUILD

- build: push moov/wire-webui image in make tasks
- build: upgrade to Go 1.15 for other docker images

## v0.6.0 (Released 2020-10-19)

ADDITIONS

- cmd/webui: initial setup for client-side file parsing to their JSON forms in a web browser

BUG FIXES

- api,client: 'beneficiary' has a Personal sub-object
- api: match openapi spec to Go library (and HTTP server) expectations
- api: update Personal identification codes
- api,client: add MessageDisposition.messageDuplicationCode " " enum value
- api: fix bug in validation of FIIntermediaryFI tag

IMPROVEMENTS

- docs: readme improvements, prioritize HTTP server / docker image

BUILD

- chore(deps): update golang docker tag to v1.15
- chore(deps): update module gorilla/mux to v1.8.0

## v0.5.1 (Released 2020-07-07)

BUILD

- build: add OpenShift [`quay.io/moov/wire`](https://quay.io/repository/moov/wire) Docker image
- build: convert to Actions from TravisCI
- chore(deps): update module prometheus/client_golang to v1.6.0
- chore(deps): upgrade github.com/gorilla/websocket to v1.4.2

## v0.5.0 (Released 2020-04-14)

IMPROVEMENTS

- build: minify Dockerfile, remove cgo
- api: remove strange "FEDWireMessage File" tag
- api,client: rename models whose names are shared across projects

BUG FIXES

- wire: log crasher file after it's parsed

BUILD

- Update module prometheus/client_golang to v1.2.1
- build: add slack notifications
- build: run sonatype-nexus-community/nancy in CI
- build: upgrade openapi-generator to 4.2.0
- chore(deps): update golang docker tag to v1.14

## v0.4.0 (Released 2019-10-18)

ADDITIONS

- cmd/server: allow creating Wire files from JSON and plain text

BUILD

- build: upgrade to Go 1.13 and Debian 10
- build: update openapi-generator to v4.1.3

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
