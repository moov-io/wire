---
layout: page
title: Go Library
hide_hero: true
show_sidebar: false
menubar: docs-menu
---

# Go Library

This project uses [Go Modules](https://github.com/golang/go/wiki/Modules) and uses Go v1.14 or higher. See [Golang's install instructions](https://golang.org/doc/install) for help setting up Go. You can download the source code and we offer [tagged and released versions](https://github.com/moov-io/wire/releases/latest) as well. We highly recommend you use a tagged release for production.

```
$ git@github.com:moov-io/wire.git

# Pull down into the Go Module cache
$ go get -u github.com/moov-io/wire

$ go doc github.com/moov-io/wire fedWireMessage
```

The package [`github.com/moov-io/wire`](https://pkg.go.dev/github.com/moov-io/wire) offers a Go-based Wire file reader and writer. To get started, check out a specific example:

### Supported Business Function Codes

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