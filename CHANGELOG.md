## v0.15.1 (Released 2024-02-01)

BUG FIXES

- api: fix validation of IdentificationCode in InstructingFI and OriginatorFI to support valid use of empty value

BUILD

- build: print crashers after fuzzing
- fix(deps): update module golang.org/x/oauth2 to v0.16.0
- fix(deps): update module github.com/moov-io/base to v0.48.5
- fix(deps): update module github.com/prometheus/client_golang to v1.18.0
- chore(deps): update dependency jekyll-feed to v0.17.0

## v0.15.0 (Released 2023-11-28)

ADDITIONS

- cmd/server: introduced query parameters for ValidationOptions when uploading raw files

BUILD

- fix(deps): update module github.com/moov-io/base to v0.48.2
- fix(deps): update module golang.org/x/oauth2 to v0.14.0
- fix(deps): update module golang.org/x/text to v0.14.0
- fix(deps): update module github.com/gorilla/mux to v1.8.1

## v0.14.1 (Released 2023-10-19)

IMPROVEMENTS

- feat: FAIM 3.0.6 Fixed ID Code/Identifier validation
- feat: implemted new field parser functions, updated record parse, format functions
- fix: FAIM 3.0.6 Updated alphanumeric regex
- fix: parse error with max length fix error tests
- fix: wrong response data in server
- test: updates after requiring error in NewTagMaxLengthErr(..)

BUILD

- build: fix openshift image, update to ubi9
- build(deps): bump golang.org/x/net to 0.17.0
- chore(deps): update dependency jekyll-feed to v0.17.0
- fix(deps): update module github.com/moov-io/base to v0.47.0
- fix(deps): update module github.com/prometheus/client_golang to v1.17.0
- fix(deps): update module golang.org/x/oauth2 to v0.13.0
- test: switch to Go's fuzz setup

## v0.13.3 (Released 2023-04-03)

IMPROVEMENTS

- Max Length issue resolved for InstructingFI Lines if it exceeds max char Limit
- 5010(OriginatorOptionF should be mandatory only if {5000} is not present and {3600} is CTP)

## v0.13.2 (Released 2023-03-22)

IMPROVEMENTS

- fix: delimiter issue resolved for OriginatorToBeneficiary Lines if it exceeds limit of 35 chars per line and 140 chars in total (#296)

BUILD

- build: upgrade golang to 1.20
- build(deps): bump golang.org/x/net from 0.6.0 to 0.7.0
- chore(deps): update dependency bulma-clean-theme to v0.13.2
- chore(deps): update dependency github-pages to v228
- fix(deps): update module github.com/stretchr/testify to v1.8.2
- fix(deps): update module golang.org/x/oauth2 to v0.6.0
- fix(deps): update module golang.org/x/text to v0.8.0

## v0.13.0 (Released 2023-02-01)

BUG FIXES

- fedWireMessage: correct validation of OriginatorFI and OriginatorToBeneficiary (see [PR #272](https://github.com/moov-io/wire/pull/272) by [@mfdeveloper508](https://github.com/mfdeveloper508) and [Issue #217](https://github.com/moov-io/wire/issues/217) by [@vishwasbabu](https://github.com/vishwasbabu))

ADDITIONS

- feat: Introduced `ValidateOpts`, which allows users to customize file validation logic for their use cases (see [PR #283](https://github.com/moov-io/wire/pull/283) by [@mfdeveloper508](https://github.com/mfdeveloper508))
  - The first option implemented for this new feature is the ability to skip validation of IMAD, which is optional for FedLine Advantage customers (see [Issue #207](https://github.com/moov-io/wire/issues/207) by [@gpark1005](https://github.com/gpark1005))
- webui: Added buttons to convert between JSON and Wire formats, and validate files (see [PR #275](https://github.com/moov-io/wire/pull/275) by [@mfdeveloper508](https://github.com/mfdeveloper508))

BUILD

- build: update module github.com/moov-io/base to v0.39.0
- build: update module github.com/prometheus/client_golang to v1.14.0
- build: update module golang.org/x/oauth2 to v0.4.0
- build: update module golang.org/x/text to v0.6.0
- build(docs): bump activesupport to 6.0.6.1
- build(docs): bump concurrent-ruby to 1.2.0
- build(docs): bump minitest to 5.17.0
- build(docs): bump nokogiri to 1.13.10
- build(docs): bump racc to 1.6.1
- build(docs): bump zeitwerk to 2.6.6

## v0.12.2 (Released 2022-11-04)

IMPROVEMENTS

- docs: update navigation links; include (Awesome Fintech)[https://github.com/moov-io/awesome-fintech] and (Terms Dictionary)[https://github.com/moov-io/terms-dictionary]]
- receiverDepositoryInstitution: `ReceiverShortName` is optional and can be omitted without padding or delimiter from tag `{3400}` when empty (see [#267](https://github.com/moov-io/wire/issues/267))
- senderDepositoryInstitution: `SenderShortName` is optional and can be omitted without padding or delimiter from tag `{3100}` when empty (see [#267](https://github.com/moov-io/wire/issues/267))

BUILD

- build: bump nokogiri to 1.13.9 in /docs
- build: update module github.com/moov-io/base to v0.36.1

## v0.12.1 (Released 2022-09-13)

IMPROVEMENTS

- businessFunctionCode: `TransactionTypeCode` is optional and can be omitted without padding from tag `{3600}` when empty (see [#240](https://github.com/moov-io/wire/issues/240) by [@anujtewari](https://github.com/anujtewari))

BUILD

- build: update module github.com/moov-io/base to v0.35.0
- build: update module golang.org/x/oauth2 to v0.0.0-20220909003341-f21342109be1

## v0.12.0 (Released 2022-09-09)

IMPROVEMENTS

- fedWireMessage: call each tag's `Validate()` method instead of only checking for mandatory field inclusion (see [#246](https://github.com/moov-io/wire/pull/246) by [@anujtewari](https://github.com/anujtewari))
- docs: replace outdated maintainer email addresses with oss@moov.io

BUG FIXES

- amount: update validation to ensure amount is strictly numeric (no commas or decimals)
- converters: handle case where the input length is equal to max length in `parseVariableLengthField(r string, maxLen int)` (see [#252](https://github.com/moov-io/wire/pull/252) by [@bhedge](https://github.com/bhedge) and [@martinz-provisions](https://github.com/martinz-provisions))
- senderSupplied: `MessageDuplicationCode` should be `" "` instead of `""` (see [#249](https://github.com/moov-io/wire/issues/249) by [@anujtewari](https://github.com/anujtewari))

## v0.11.0 (Released 2022-09-06)

ADDITIONS

- feat: introduce conditional validation for incoming vs outgoing files (see [#244](https://github.com/moov-io/wire/pull/244) by [@mfdeveloper508](https://github.com/mfdeveloper508))

IMPROVEMENTS

- fedWireMessage: only require `SenderSupplied` for outgoing files
- senderSupplied: make `UserRequestCorrelation` optional

## v0.10.0 (Released 2022-09-01)

ADDITIONS

- feat: `format` and `newline` query params introduced to the Get File Contents endpoint to set writer `FormatOptions` (see [#239](https://github.com/moov-io/wire/issues/239) by [@anujtewari](https://github.com/anujtewari))

BUILD

- build: upgrade to Go 1.19 for docker images
- build: update module github.com/prometheus/client_golang to v1.13.0
- build: update module golang.org/x/oauth2 to v0.0.0-20220822191816-0ebed06d0094
- build: update module github.com/moov-io/base to v0.34.1

## v0.9.0 (Released 2022-07-15)

ADDITIONS

- feat: support variable length fields

IMPROVEMENTS

- cmd/server: call validate after creating Wire file via JSON

BUILD

- build: upgrade to Go 1.18 for docker images
- build: update module github.com/moov-io/base to v0.32.0
- build: update module github.com/prometheus/client_golang to v1.12.2
- build: update module github.com/stretchr/testify to v1.8.0
- build: update module golang.org/x/oauth2 to v0.0.0-20220630143837-2104d58473e0

## v0.8.0 (Released 2022-02-09)

IMPROVEMENTS

- fedWireMessage: support fed-appended tags (`{1100}`, `{1110}`, `{1120}`, and `{1130}`) in `Reader` and `Writer`

## v0.7.5 (Released 2021-09-13)

BUG FIXES

- originator: parse address line two

## v0.7.4 (Released 2021-08-09)

BUG FIXES

- api: respond with 404 instead of empty body when `fileId` is not found

## v0.7.3 (Released 2021-07-16)

BUILD

- build(deps): bump addressable from 2.7.0 to 2.8.0 in /docs
- build(deps): bump nokogiri from 1.11.1 to 1.11.5 in /docs
- fix(deps): update golang.org/x/oauth2 commit hash to a41e5a7
- fix(deps): update module github.com/go-kit/kit to v0.11.0
- fix(deps): update module github.com/moov-io/base to v0.20.1
- fix(deps): update module gotest.tools to v2.3.0
- fix: Dockerfile.webui to reduce vulnerabilities

## v0.7.2 (Released 2021-05-07)

BUG FIXES

- fedWireMessage: add `nil` check for the optional `LocalInstrument` field to prevent panics

BUILD

- build: update github.com/moov-io/base to v0.18.3
- build: update github.com/prometheus/client_golang to v1.10.0
- build: update github.com/stretchr/testify to v1.7.0
- build: update golang.org/x/text to v0.3.6

## v0.7.1 (Released 2020-12-18)

BUILD

- build: update github.com/moov-io/base to v0.15.2
- build: update github.com/moov-io/paygate to v0.9.2
## v0.7.0 (Released 2020-11-23)

BREAKING CHANGE

- fedWireMessage: remove Java-like getters and setters

IMPROVEMENTS

- cmd/webui: display detailed error message when parsing fails

## v0.6.2 (Released 2020-10-28)

BUG FIXES

- api: update validation of `AccountCreditedDrawdown.DrawdownCreditAccountNumber` (tag `{5400}`)
  - After reviewing the specification for the Account Credited in Drawdown tag we've changed
  the validation logic to ensure the value of Drawdown Credit Account Number is numeric. Our
  previous understanding was that this field could be alphanumeric. If you're aware of
  implementations or use-cases for the previous formatting please contact us and/or submit an Issue.
- all: populate record tag after JSON unmarshal (Issue [#104](https://github.com/moov-io/wire/issues/104))

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
