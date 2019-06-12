moov-io/wire
===
[![GoDoc](https://godoc.org/github.com/moov-io/wire?status.svg)](https://godoc.org/github.com/moov-io/wire)
[![Build Status](https://travis-ci.com/moov-io/wire.svg?branch=master)](https://travis-ci.com/moov-io/wire)
[![Coverage Status](https://codecov.io/gh/moov-io/wire/branch/master/graph/badge.svg)](https://codecov.io/gh/moov-io/wire)
[![Go Report Card](https://goreportcard.com/badge/github.com/moov-io/wire)](https://goreportcard.com/report/github.com/moov-io/wire)
[![Apache 2 licensed](https://img.shields.io/badge/license-Apache2-blue.svg)](https://raw.githubusercontent.com/moov-io/ach/master/LICENSE)

Package `github.com/moov-io/wire` implements a reader and writer written in Go  for creating, parsing and validating FED Wire Messages ([FEDWire](https://en.wikipedia.org/wiki/Fedwire)) 

Docs: [docs.moov.io](https://docs.moov.io/en/latest/) | [api docs](https://api.moov.io/apps/wire/)

## Project Status

Moov WIRE is under active development and should not currently be used for production. Please star the project if you are interested in its progress.

## Usage

### Go library

`github.com/moov-io/wire` offers a Go based ACH file reader and writer. To get started checkout a specific example:

<details>
<summary>Supported Business Function Codes</summary>

| Business Function Code | Name               | Example | Read | Write |
|----------|----------------------------------|---------|------|-------|
| DRB      | Bank DrawDown Request            | [Link](examples/bankDrawDownRequest-read/bankDrawDownRequest.txt) | [Link](examples/bankDrawDownRequest-read/main.go) | [Link](examples/bankDrawDownRequest-write/main.go) | 
| BTR      | BankTransfer                     | [Link](examples/bankTransfer-read/bankTransfer.txt) | [Link](examples/bankTransfer-read/main.go) | [Link](examples/bankTransfer-write/main.go) |
| CKS      | CheckSameDaySettlement           | [Link](examples/checkSameDaySettlement-read/checkSameDaySettlement.txt) | [Link](examples/checkSameDaySettlement-read/main.go) | [Link](examples/checkSameDaySettlement-write/main.go) |
| DRC      | CustomerCorporateDrawdownRequest | [Link](examples/customerCorporateDrawDownRequest-read/customerCorporateDrawDownRequest.txt) | [Link](examples/customerCorporateDrawDownRequest-read/main.go) | [Link](examples/customerCorporateDrawDownRequest-write/main.go) |
| CTR      | CustomerTransfer                 | [Link](examples/customerTransfer-read/customerTransfer.txt) | [Link](examples/customerTransfer-read/main.go) | [Link](examples/customerTransfer-write/main.go) |
| CTP      | CustomerTransferPlus             | [Link](examples/customerTransferPlus-read/customerTransferPlus.txt) | [Link](examples/customerTransferPlus-read/main.go) | [Link](examples/customerTransferPlus-write/main.go) |
| CTP      | CustomerTransferPlusCOVS         | [Link](examples/customerTransferPlusCOVS-read/customerTransferPlusCOVS.txt) | [Link](examples/customerTransferPlusCOVS-read/main.go) | [Link](examples/customerTransferPlusCOVS-write/main.go) |
| DEP      | DepositSendersAccount            | [Link](examples/depositSendersAccount-read/depositSendersAccount.txt) | [Link](examples/depositSendersAccount-read/main.go) | [Link](examples/depositSendersAccount-write/main.go) |
| FFR      | FEDFundsReturned                 | [Link](examples/fedFundsReturned-read/fedFundsReturned.txt) | [Link](examples/fedFundsReturned-read/main.go) | [Link](examples/fedFundsReturned-write/main.go) |
| FFS      | FEDFundsSold                     | [Link](examples/fedFundsSold-read/fedFundsSold.txt) | [Link](examples/fedFundsSold-read/main.go) | [Link](examples/fedFundsSold-write/main.go) |
| SVC      | ServiceMessage                   | [Link](examples/serviceMessage-read/serviceMessage.txt) | [Link](examples/serviceMessage-read/main.go) | [Link](examples/serviceMessage-write/main.go) |
</details>

### From Source

This project uses [Go Modules](https://github.com/golang/go/wiki/Modules) and thus requires Go 1.11+. You can download the source code and we offer [tagged and released versions](https://github.com/moov-io) as well. We highly recommend you use a tagged release for production.

```
$ git@github.com:moov-io/wire.git

# Pull down into the Go Module cache
$ go get -u github.com/moov-io/wire

$ go doc github.com/moov-io/wire fedWireMessage
```

## Getting Help

 channel | info
 ------- | -------
[Project Documentation](https://docs.moov.io/en/latest/) | Our project documentation available online.
Google Group [moov-users](https://groups.google.com/forum/#!forum/moov-users)| The Moov users Google group is for contributors other people contributing to the Moov project. You can join them without a google account by sending an email to [moov-users+subscribe@googlegroups.com](mailto:moov-users+subscribe@googlegroups.com). After receiving the join-request message, you can simply reply to that to confirm the subscription.
Twitter [@moov_io](https://twitter.com/moov_io)	| You can follow Moov.IO's Twitter feed to get updates on our project(s). You can also tweet us questions or just share blogs or stories.
[GitHub Issue](https://github.com/moov-io) | If you are able to reproduce an problem please open a GitHub Issue under the specific project that caused the error.
[moov-io slack](http://moov-io.slack.com/) | Join our slack channel to have an interactive discussion about the development of the project. [Request an invite to the slack channel](https://join.slack.com/t/moov-io/shared_invite/enQtNDE5NzIwNTYxODEwLTRkYTcyZDI5ZTlkZWRjMzlhMWVhMGZlOTZiOTk4MmM3MmRhZDY4OTJiMDVjOTE2MGEyNWYzYzY1MGMyMThiZjg)

## Supported and Tested Platforms

- 64-bit Linux (Ubuntu, Debian), macOS, and Windows

Note: 32-bit platforms have known issues and are not supported.

## Contributing

Yes please! Please review our [Contributing guide](CONTRIBUTING.md) and [Code of Conduct](CODE_OF_CONDUCT.md) to get started!

Note: This project uses Go Modules, which requires Go 1.11 or higher.

### Releasing

To make a release of wire simply open a pull request with `CHANGELOG.md` and `version.go` updated with the next version number and details. You'll also need to push the tag (i.e. `git push origin v1.0.0`) to origin in order for CI to make the release.

## License

Apache License 2.0 See [LICENSE](LICENSE) for details.
