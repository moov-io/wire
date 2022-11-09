---
layout: page
title: Go library
hide_hero: true
show_sidebar: false
menubar: docs-menu
---

# Go library

This project uses [Go Modules](https://go.dev/blog/using-go-modules) and Go v1.18 or newer. See [Golang's install instructions](https://golang.org/doc/install) for help setting up Go. You can download the source code and we offer [tagged and released versions](https://github.com/moov-io/wire/releases/latest) as well. We highly recommend you use a tagged release for production.

```
$ git@github.com:moov-io/wire.git

$ go get -u github.com/moov-io/wire

$ go doc github.com/moov-io/wire fedWireMessage
```

The package [`github.com/moov-io/wire`](https://pkg.go.dev/github.com/moov-io/wire) offers a Go-based Wire file reader and writer. To get started, check out a specific example:

### Supported business function codes

| Business Function Code | Name               | Example | Read | Write |
|----------|----------------------------------|---------|------|-------|
| DRB      | BankDrawDownRequest            | [Link](https://github.com/moov-io/wire/blob/master/examples/bankDrawDownRequest-read/bankDrawDownRequest.txt) | [Link](https://github.com/moov-io/wire/blob/master/examples/bankDrawDownRequest-read/main.go) | [Link](https://github.com/moov-io/wire/blob/master/examples/bankDrawDownRequest-write/main.go) |
| BTR      | BankTransfer                     | [Link](https://github.com/moov-io/wire/blob/master/examples/bankTransfer-read/bankTransfer.txt) | [Link](https://github.com/moov-io/wire/blob/master/examples/bankTransfer-read/main.go) | [Link](https://github.com/moov-io/wire/blob/master/examples/bankTransfer-write/main.go) |
| CKS      | CheckSameDaySettlement           | [Link](https://github.com/moov-io/wire/blob/master/examples/checkSameDaySettlement-read/checkSameDaySettlement.txt) | [Link](https://github.com/moov-io/wire/blob/master/examples/checkSameDaySettlement-read/main.go) | [Link](https://github.com/moov-io/wire/blob/master/examples/checkSameDaySettlement-write/main.go) |
| DRC      | CustomerCorporateDrawDownRequest | [Link](https://github.com/moov-io/wire/blob/master/examples/customerCorporateDrawDownRequest-read/customerCorporateDrawDownRequest.txt) | [Link](https://github.com/moov-io/wire/blob/master/examples/customerCorporateDrawDownRequest-read/main.go) | [Link](https://github.com/moov-io/wire/blob/master/examples/customerCorporateDrawDownRequest-write/main.go) |
| CTR      | CustomerTransfer                 | [Link](https://github.com/moov-io/wire/blob/master/examples/customerTransfer-read/customerTransfer.txt) | [Link](https://github.com/moov-io/wire/blob/master/examples/customerTransfer-read/main.go) | [Link](https://github.com/moov-io/wire/blob/master/examples/customerTransfer-write/main.go) |
| CTP      | CustomerTransferPlus             | [Link](https://github.com/moov-io/wire/blob/master/examples/customerTransferPlus-read/customerTransferPlus.txt) | [Link](https://github.com/moov-io/wire/blob/master/examples/customerTransferPlus-read/main.go) | [Link](https://github.com/moov-io/wire/blob/master/examples/customerTransferPlus-write/main.go) |
| CTP      | CustomerTransferPlusCOVS         | [Link](https://github.com/moov-io/wire/blob/master/examples/customerTransferPlusCOVS-read/customerTransferPlusCOVS.txt) | [Link](https://github.com/moov-io/wire/blob/master/examples/customerTransferPlusCOVS-read/main.go) | [Link](https://github.com/moov-io/wire/blob/master/examples/customerTransferPlusCOVS-write/main.go) |
| DEP      | DepositSendersAccount            | [Link](https://github.com/moov-io/wire/blob/master/examples/depositSendersAccount-read/depositSendersAccount.txt) | [Link](https://github.com/moov-io/wire/blob/master/examples/depositSendersAccount-read/main.go) | [Link](https://github.com/moov-io/wire/blob/master/examples/depositSendersAccount-write/main.go) |
| FFR      | FEDFundsReturned                 | [Link](https://github.com/moov-io/wire/blob/master/examples/fedFundsReturned-read/fedFundsReturned.txt) | [Link](https://github.com/moov-io/wire/blob/master/examples/fedFundsReturned-read/main.go) | [Link](https://github.com/moov-io/wire/blob/master/examples/fedFundsReturned-write/main.go) |
| FFS      | FEDFundsSold                     | [Link](https://github.com/moov-io/wire/blob/master/examples/fedFundsSold-read/fedFundsSold.txt) | [Link](https://github.com/moov-io/wire/blob/master/examples/fedFundsSold-read/main.go) | [Link](https://github.com/moov-io/wire/blob/master/examples/fedFundsSold-write/main.go) |
| SVC      | ServiceMessage                   | [Link](https://github.com/moov-io/wire/blob/master/examples/serviceMessage-read/serviceMessage.txt) | [Link](https://github.com/moov-io/wire/blob/master/examples/serviceMessage-read/main.go) | [Link](https://github.com/moov-io/wire/blob/master/examples/serviceMessage-write/main.go) |
