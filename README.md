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
[![Coverage Status](https://codecov.io/gh/moov-io/wire/branch/master/graph/badge.svg)](https://codecov.io/gh/moov-io/wire)
[![Go Report Card](https://goreportcard.com/badge/github.com/moov-io/wire)](https://goreportcard.com/report/github.com/moov-io/wire)
[![Repo Size](https://img.shields.io/github/languages/code-size/moov-io/wire?label=project%20size)](https://github.com/moov-io/wire)
[![Apache 2 License](https://img.shields.io/badge/license-Apache2-blue.svg)](https://raw.githubusercontent.com/moov-io/wire/master/LICENSE)
[![Slack Channel](https://slack.moov.io/badge.svg?bg=e01563&fgColor=fffff)](https://slack.moov.io/)
[![Docker Pulls](https://img.shields.io/docker/pulls/moov/wire)](https://hub.docker.com/r/moov/wire)
[![GitHub Stars](https://img.shields.io/github/stars/moov-io/wire)](https://github.com/moov-io/wire)
[![Twitter](https://img.shields.io/twitter/follow/moov?style=social)](https://twitter.com/moov?lang=en)

# moov-io/wire

Moov's mission is to give developers an easy way to create and integrate bank processing into their own software products. Our open source projects are each focused on solving a single responsibility in financial services and designed around performance, scalability, and ease of use.

Wire implements a reader, writer, and validator for FED Wire Messages ([FEDWire](https://en.wikipedia.org/wiki/Fedwire)) in an HTTP server and Go library. The HTTP server is available in a [Docker image](#docker) and the Go package `github.com/moov-io/wire` is available.

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

Moov Wire is actively used in multiple production environments. Please star the project if you are interested in its progress. If you have layers above Wire to simplify tasks, perform business operations, or found bugs we would appreciate an issue or pull request. Thanks!

## Usage
The Wire project implements an HTTP server and [Go library](https://pkg.go.dev/github.com/moov-io/wire) for creating and modifying Wire files. We also have some [examples](https://pkg.go.dev/github.com/moov-io/wire/examples) of the reader and writer.

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

This project uses [Go Modules](https://github.com/golang/go/wiki/Modules) and uses Go v1.14 or higher. See [Golang's install instructions](https://golang.org/doc/install) for help setting up Go. You can download the source code and we offer [tagged and released versions](https://github.com/moov-io/wire/releases/latest) as well. We highly recommend you use a tagged release for production.

```
$ git@github.com:moov-io/wire.git

# Pull down into the Go Module cache
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
Using our [in-browser utility](http://oss.moov.io/wire/), you can instantly convert Wire files into JSON. Either paste in Wire file content directly or choose a file from your local machine. This tool is particulary useful if you're handling sensitive PII or want perform some quick tests, as operations are fully client-side with nothing stored in memory. We plan to support bidirectional conversion in the future.

## Learn about Fedwire
- [Intro to Fedwire](https://www.americanexpress.com/us/foreign-exchange/articles/fedwire-transfers/)
- [FedWire Message Structure](./docs/fedWireMessage-Structure.md)
- [Sending or Receiving International Wires via the Fedwire Funds Service](https://www.youtube.com/watch?v=GSd2gZ8-bzQ)

## FAQ
<details open="true">
<summary ><b>Is there an in-browser tool for converting Wire files into JSON?</b></summary>
Yes! You can find our browser utility at http://oss.moov.io/wire/.
</details>
<details open="true">
<summary><b>Is my data being saved somewhere?</b></summary>
No, we do not save any data related to files or message details. All processing is done in-memory.
</details>
<details open="true">
<summary><b>What Fedwire message types are supported?</b></summary>
We support generating and parsing all Business Function codes.
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

This project uses [Go Modules](https://github.com/golang/go/wiki/Modules) and uses Go v1.14 or higher. See [Golang's install instructions](https://golang.org/doc/install) for help setting up Go. You can download the source code and we offer [tagged and released versions](https://github.com/moov-io/wire/releases/latest) as well. We highly recommend you use a tagged release for production.

### Releasing

To make a release of wire simply open a pull request with `CHANGELOG.md` and `version.go` updated with the next version number and details. You'll also need to push the tag (i.e. `git push origin v1.0.0`) to origin in order for CI to make the release.

### Testing

We maintain a comprehensive suite of unit tests and recommend table-driven testing when a particular function warrants several very similar test cases. To run all test files in the current directory, use `go test`. Current overall coverage can be found on [Codecov](https://app.codecov.io/gh/moov-io/wire/).

### Fuzzing

We currently run fuzzing over wire in the form of a [`moov/wirefuzz`](https://hub.docker.com/r/moov/wirefuzz) Docker image. You can [read more](./test/fuzz-reader/README.md) or run the image and report crasher examples to [`security@moov.io`](mailto:security@moov.io). Thanks!

## Related projects

As part of Moov's initiative to offer open source fintech infrastructure, we have a large collection of active projects you may find useful:

- [Moov Watchman](https://github.com/moov-io/watchman) offers search functions over numerous trade sanction lists from the United States and European Union.

- [Moov Fed](https://github.com/moov-io/fed) implements utility services for searching the United States Federal Reserve System such as ABA routing numbers, financial institution name lookup, and FedACH and Fedwire routing information.

- [Moov Image Cash Letter](https://github.com/moov-io/imagecashletter) implements Image Cash Letter (ICL) files used for Check21, X.9 or check truncation files for exchange and remote deposit in the U.S.

- [Moov ACH](https://github.com/moov-io/ach) provides ACH file generation and parsing, supporting all Standard Entry Codes for the primary method of money movement throughout the United States.

- [Moov Metro 2](https://github.com/moov-io/metro2) provides a way to easily read, create, and validate Metro 2 format, which is used for consumer credit history reporting by the United States credit bureaus.

## License

Apache License 2.0 - See [LICENSE](LICENSE) for details.
