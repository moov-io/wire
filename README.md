[![Moov Banner Logo](https://user-images.githubusercontent.com/20115216/104214617-885b3c80-53ec-11eb-8ce0-9fc745fb5bfc.png)](https://github.com/moov-io)

<p align="center">
  <a href="https://moov-io.github.io/wire/">Project Documentation</a>
  ·
  <a href="https://moov-io.github.io/wire/api/#overview">API Endpoints</a>
  ·
  <a href="https://moov.io/blog/education/wire-api-guide/">API Guide</a>
  ·
  <a href="https://slack.moov.io/">Community</a>
  ·
  <a href="https://moov.io/blog/">Blog</a>
  <br>
  <br>
</p>

[![GoDoc](https://godoc.org/github.com/moov-io/wire?status.svg)](https://godoc.org/github.com/moov-io/wire)
[![Build Status](https://github.com/moov-io/wire/workflows/Go/badge.svg)](https://github.com/moov-io/wire/actions)
[![Go Report Card](https://goreportcard.com/badge/github.com/moov-io/wire)](https://goreportcard.com/report/github.com/moov-io/wire)
[![Repo Size](https://img.shields.io/github/languages/code-size/moov-io/wire?label=project%20size)](https://github.com/moov-io/wire)
[![Apache 2 License](https://img.shields.io/badge/license-Apache2-blue.svg)](https://raw.githubusercontent.com/moov-io/wire/master/LICENSE)
[![Slack Channel](https://slack.moov.io/badge.svg?bg=e01563&fgColor=fffff)](https://slack.moov.io/)
[![Docker Pulls](https://img.shields.io/docker/pulls/moov/wire)](https://hub.docker.com/r/moov/wire)
[![GitHub Stars](https://img.shields.io/github/stars/moov-io/wire)](https://github.com/moov-io/wire)
[![Twitter](https://img.shields.io/twitter/follow/moov?style=social)](https://twitter.com/moov?lang=en)

# moov-io/wire

Moov's mission is to give developers an easy way to create and integrate bank processing into their own software products. Our open source projects are each focused on solving a single responsibility in financial services and designed around performance, scalability, and ease of use.

**Moov Wire** implements a reader, writer, and validator for the **legacy Fedwire Funds Service message format (FAIM)** — the proprietary tag-based format described in the Fedwire Application Interface Manual. The HTTP server is available in a [Docker image](#docker) and the Go package `github.com/moov-io/wire` is available.

> [!IMPORTANT]
> **Fedwire Funds Service transitioned from FAIM to ISO 20022 on July 14, 2025.**
>
> The Federal Reserve Banks completed a [single-day cutover](https://www.frbservices.org/resources/financial-services/wires/iso-20022-implementation-center) to ISO 20022. FAIM is **no longer accepted** for live Fedwire Funds traffic. This repository remains available for historical FAIM files, archival processing, testing, and migration work.
>
> For **new Fedwire integrations**, use Moov's ISO 20022 libraries:
>
> | Project | Purpose |
> |---------|---------|
> | [`moov-io/wire20022`](https://github.com/moov-io/wire20022) | Read, write, and validate Fedwire ISO 20022 XML messages (recommended) |
> | [`moov-io/fedwire20022`](https://github.com/moov-io/fedwire20022) | Generated Go types from the Fedwire ISO 20022 XSDs |
>
> Official resources: [ISO 20022 Implementation Center](https://www.frbservices.org/resources/financial-services/wires/iso-20022-implementation-center) · [Overview & implementation details](https://www.frbservices.org/resources/financial-services/wires/faq/iso-20022/overview-implementation-details)

## Table of contents

- [Project Status](#project-status)
- [Usage](#usage)
  - As an API
    - [Docker](#docker) ([Config](#configuration-settings))
    - [Google Cloud](#google-cloud-run) ([Config](#configuration-settings))
    - [Data Persistence](#data-persistence)
  - [As a Go Module](#go-library)
  - [As an In-Browser Parser](#in-browser-wire-file-parser)
- [Learn About Wire](#learn-about-wire)
- [FAQ](#faq)
- [Getting Help](#getting-help)
- [Supported and Tested Platforms](#supported-and-tested-platforms)
- [Contributing](#contributing)
- [Related Projects](#related-projects)

## Project status

**Moov Wire is in maintenance mode** for the legacy **FAIM** (Fedwire Application Interface Manual) format.

| Format | Fedwire Funds support | Moov project |
|--------|----------------------|--------------|
| **ISO 20022** (XML) | Required as of **July 14, 2025** | [`wire20022`](https://github.com/moov-io/wire20022), [`fedwire20022`](https://github.com/moov-io/fedwire20022) |
| **FAIM** (tag-based, format version 30) | **Deprecated** — not accepted for live traffic after July 11, 2025 | This repository (`moov-io/wire`) |

This library continues to parse, write, and validate FAIM messages for existing integrations, historical files, and migration tooling. We accept bug fixes and security updates; new FAIM feature work is unlikely. Please star the project if you still rely on it, and open an issue or pull request for bugs you find. Thanks!

## Usage
The Wire project implements an HTTP server and [Go library](https://pkg.go.dev/github.com/moov-io/wire) for creating and modifying **legacy FAIM** Fedwire files. We also have some [examples](https://pkg.go.dev/github.com/moov-io/wire/examples) of the reader and writer.

### Docker

We publish a [public Docker image `moov/wire`](https://hub.docker.com/r/moov/wire/tags) on Docker Hub with every tagged release of Wire. No configuration is required to serve on `:8088` and metrics at `:9098/metrics` in Prometheus format. We also have Docker images for [OpenShift](https://quay.io/repository/moov/wire?tab=tags) published as `quay.io/moov/wire`.

Pull & start the Docker image:
```
docker pull moov/wire:latest
docker run -p 8088:8088 -p 9098:9098 moov/wire:latest
```

List files stored in-memory:
```
curl localhost:8088/files
```
```
null
```

Create a file on the HTTP server:
```
curl -X POST --data-binary "@./test/testdata/fedWireMessage-CustomerTransfer.txt" http://localhost:8088/files/create
```
```
{"id":"<YOUR-UNIQUE-FILE-ID>","fedWireMessage":{"id":"","senderSupplied":{"formatVersion":"30", .....
```

Get the file in its original format:
```
curl http://localhost:8088/files/<YOUR-UNIQUE-FILE-ID>/contents
```
```
{1500}30User ReqT
{1510}1000
{1520}20190410Source08000001
...
```

### Google Cloud Run

To get started in a hosted environment you can deploy this project to the Google Cloud Platform.

From your [Google Cloud dashboard](https://console.cloud.google.com/home/dashboard) create a new project and call it:
```
moov-wire-demo
```

Enable the [Container Registry](https://cloud.google.com/container-registry) API for your project and associate a [billing account](https://cloud.google.com/billing/docs/how-to/manage-billing-account) if needed. Then, open the Cloud Shell terminal and run the following Docker commands, substituting your unique project ID:

```
docker pull moov/wire
docker tag moov/wire gcr.io/<PROJECT-ID>/wire
docker push gcr.io/<PROJECT-ID>/wire
```

Deploy the container to Cloud Run:
```
gcloud run deploy --image gcr.io/<PROJECT-ID>/wire --port 8088
```

Select your target platform to `1`, service name to `wire`, and region to the one closest to you (enable Google API service if a prompt appears). Upon a successful build you will be given a URL where the API has been deployed:

```
https://YOUR-WIRE-APP-URL.a.run.app
```

Now you can list files stored in-memory:
```
curl https://YOUR-WIRE-APP-URL.a.run.app/files
```
You should get this response:
```
null
```

### Configuration settings

The following environmental variables can be set to configure behavior in Wire.

| Environmental Variable | Description | Default |
|-----|-----|-----|
| `HTTPS_CERT_FILE` | Filepath containing a certificate (or intermediate chain) to be served by the HTTP server. Requires all traffic be over secure HTTP. | Empty |
| `HTTPS_KEY_FILE`  | Filepath of a private key matching the leaf certificate from `HTTPS_CERT_FILE`. | Empty |
| `WIRE_FILE_TTL` | Time to live (TTL) for `*wire.File` objects stored in the in-memory repository. | 0 = No TTL / Never delete files (Example: `240m`) |

### Data persistence

By design, Wire  **does not persist** (save) any data about the files or entry details created. The only storage occurs in memory of the process and upon restart Wire will have no files or data saved. Also, no in-memory encryption of the data is performed.

### Go library

This project uses [Go Modules](https://go.dev/blog/using-go-modules) and Go v1.18 or newer. See [Golang's install instructions](https://golang.org/doc/install) for help setting up Go. You can download the source code and we offer [tagged and released versions](https://github.com/moov-io/wire/releases/latest) as well. We highly recommend you use a tagged release for production.

```
$ git@github.com:moov-io/wire.git

$ go get -u github.com/moov-io/wire

$ go doc github.com/moov-io/wire fedWireMessage
```

The package [`github.com/moov-io/wire`](https://pkg.go.dev/github.com/moov-io/wire) offers a Go-based Wire file reader and writer. To get started, check out a specific example:

<details>
<summary>Supported Business Function Codes</summary>

| Business Function Code | Name               | Example | Read | Write |
|----------|----------------------------------|---------|------|-------|
| DRB      | BankDrawDownRequest            | [Link](examples/bankDrawDownRequest-read/bankDrawDownRequest.txt) | [Link](examples/bankDrawDownRequest-read/main.go) | [Link](examples/bankDrawDownRequest-write/main.go) |
| BTR      | BankTransfer                     | [Link](examples/bankTransfer-read/bankTransfer.txt) | [Link](examples/bankTransfer-read/main.go) | [Link](examples/bankTransfer-write/main.go) |
| CKS      | CheckSameDaySettlement           | [Link](examples/checkSameDaySettlement-read/checkSameDaySettlement.txt) | [Link](examples/checkSameDaySettlement-read/main.go) | [Link](examples/checkSameDaySettlement-write/main.go) |
| DRC      | CustomerCorporateDrawDownRequest | [Link](examples/customerCorporateDrawDownRequest-read/customerCorporateDrawDownRequest.txt) | [Link](examples/customerCorporateDrawDownRequest-read/main.go) | [Link](examples/customerCorporateDrawDownRequest-write/main.go) |
| CTR      | CustomerTransfer                 | [Link](examples/customerTransfer-read/customerTransfer.txt) | [Link](examples/customerTransfer-read/main.go) | [Link](examples/customerTransfer-write/main.go) |
| CTP      | CustomerTransferPlus             | [Link](examples/customerTransferPlus-read/customerTransferPlus.txt) | [Link](examples/customerTransferPlus-read/main.go) | [Link](examples/customerTransferPlus-write/main.go) |
| CTP      | CustomerTransferPlusCOVS         | [Link](examples/customerTransferPlusCOVS-read/customerTransferPlusCOVS.txt) | [Link](examples/customerTransferPlusCOVS-read/main.go) | [Link](examples/customerTransferPlusCOVS-write/main.go) |
| DEP      | DepositSendersAccount            | [Link](examples/depositSendersAccount-read/depositSendersAccount.txt) | [Link](examples/depositSendersAccount-read/main.go) | [Link](examples/depositSendersAccount-write/main.go) |
| FFR      | FEDFundsReturned                 | [Link](examples/fedFundsReturned-read/fedFundsReturned.txt) | [Link](examples/fedFundsReturned-read/main.go) | [Link](examples/fedFundsReturned-write/main.go) |
| FFS      | FEDFundsSold                     | [Link](examples/fedFundsSold-read/fedFundsSold.txt) | [Link](examples/fedFundsSold-read/main.go) | [Link](examples/fedFundsSold-write/main.go) |
| SVC      | ServiceMessage                   | [Link](examples/serviceMessage-read/serviceMessage.txt) | [Link](examples/serviceMessage-read/main.go) | [Link](examples/serviceMessage-write/main.go) |
</details>

### In-browser Wire file parser
Using our [in-browser utility](http://oss.moov.io/wire/), you can instantly convert Wire files into JSON. Either paste in Wire file content directly or choose a file from your local machine. This tool is particularly useful if you're handling sensitive PII or want perform some quick tests, as operations are fully client-side with nothing stored in memory. We plan to support bidirectional conversion in the future.

## Learn about Fedwire
- [Intro to Fedwire](./docs/intro.md) (includes the ISO 20022 transition)
- [FAIM message structure](./docs/message-structure.md) (legacy format implemented by this project)
- [Fedwire Funds Service ISO 20022 Implementation Center](https://www.frbservices.org/resources/financial-services/wires/iso-20022-implementation-center)
- [Intro to Fedwire](https://www.americanexpress.com/us/foreign-exchange/articles/fedwire-transfers/)
- [Sending or Receiving International Wires via the Fedwire Funds Service](https://www.youtube.com/watch?v=GSd2gZ8-bzQ)

## FAQ
<details open="true">
<summary><b>Does this library support Fedwire ISO 20022?</b></summary>
No. This project implements the legacy <strong>FAIM</strong> format only. For ISO 20022 Fedwire messages, use <a href="https://github.com/moov-io/wire20022">moov-io/wire20022</a> (and optionally <a href="https://github.com/moov-io/fedwire20022">moov-io/fedwire20022</a> for generated XSD types).
</details>
<details open="true">
<summary><b>Is FAIM still valid on the Fedwire Funds Service?</b></summary>
No. After the Federal Reserve's July 14, 2025 cutover, live Fedwire Funds traffic must use ISO 20022. FAIM is retained here for historical files, testing, and migration.
</details>
<details open="true">
<summary ><b>Is there an in-browser tool for converting Wire files into JSON?</b></summary>
Yes! You can find our browser utility at http://oss.moov.io/wire/. It parses legacy FAIM files client-side.
</details>
<details open="true">
<summary><b>Is my data being saved somewhere?</b></summary>
No, we do not save any data related to files or message details. All processing is done in-memory.
</details>
<details open="true">
<summary><b>What Fedwire message types are supported?</b></summary>
For FAIM, we support generating and parsing all Business Function codes listed below. ISO 20022 message types are covered by <a href="https://github.com/moov-io/wire20022">wire20022</a>.
</details>

## Getting help

 channel | info
 ------- | -------
[Project Documentation](https://moov-io.github.io/wire/) | Our project documentation available online.
Twitter [@moov](https://twitter.com/moov)	| You can follow Moov.io's Twitter feed to get updates on our project(s). You can also tweet us questions or just share blogs or stories.
[GitHub Issue](https://github.com/moov-io/wire/issues) | If you are able to reproduce a problem please open a GitHub Issue under the specific project that caused the error.
[moov-io slack](https://slack.moov.io/) | Join our slack channel to have an interactive discussion about the development of the project.

## Supported and tested platforms

- 64-bit Linux (Ubuntu, Debian), macOS, and Windows

Note: 32-bit platforms have known issues and are not supported.

## Contributing

Yes please! Please review our [Contributing guide](CONTRIBUTING.md) and [Code of Conduct](CODE_OF_CONDUCT.md) to get started!

This project uses [Go Modules](https://go.dev/blog/using-go-modules) and Go v1.18 or newer. See [Golang's install instructions](https://golang.org/doc/install) for help setting up Go. You can download the source code and we offer [tagged and released versions](https://github.com/moov-io/wire/releases/latest) as well. We highly recommend you use a tagged release for production.

### Releasing

To make a release of wire simply open a pull request with `CHANGELOG.md` and `version.go` updated with the next version number and details. You'll also need to push the tag (i.e. `git push origin v1.0.0`) to origin in order for CI to make the release.

### Testing

We maintain a comprehensive suite of unit tests and recommend table-driven testing when a particular function warrants several very similar test cases. To run all test files in the current directory, use `go test`.

### Fuzzing

We currently run fuzzing over Wire in the form of a [Github Action](https://github.com/moov-io/wire/actions/workflows/fuzz.yml). Please report crashes examples to [`oss@moov.io`](mailto:oss@moov.io). Thanks!

## Related projects

As part of Moov's initiative to offer open source fintech infrastructure, we have a large collection of active projects you may find useful:

- [Moov Wire20022](https://github.com/moov-io/wire20022) reads, writes, and validates **Fedwire ISO 20022** XML messages (successor format to FAIM).

- [Moov Fedwire20022](https://github.com/moov-io/fedwire20022) provides generated Go types from the Fedwire Funds ISO 20022 XSDs.

- [Moov Watchman](https://github.com/moov-io/watchman) offers search functions over numerous trade sanction lists from the United States and European Union.

- [Moov Fed](https://github.com/moov-io/fed) implements utility services for searching the United States Federal Reserve System such as ABA routing numbers, financial institution name lookup, and FedACH and Fedwire routing information.

- [Moov Image Cash Letter](https://github.com/moov-io/imagecashletter) implements Image Cash Letter (ICL) files used for Check21, X.9 or check truncation files for exchange and remote deposit in the U.S.

- [Moov ACH](https://github.com/moov-io/ach) provides ACH file generation and parsing, supporting all Standard Entry Codes for the primary method of money movement throughout the United States.

- [Moov Metro 2](https://github.com/moov-io/metro2) provides a way to easily read, create, and validate Metro 2 format, which is used for consumer credit history reporting by the United States credit bureaus.

## License

Apache License 2.0 - See [LICENSE](LICENSE) for details.
