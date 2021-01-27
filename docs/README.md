# Moov Wire

**[GitHub](https://github.com/moov-io/wire)** | **[Running](https://github.com/moov-io/wire#usage)** | **[Configuration](https://github.com/moov-io/wire#configuration-settings)**

## Purpose

[Moov Wire](https://github.com/moov-io/wire) implements an HTTP interface to read and write Fedwire Messages.

FEDWire (formerly known as the Federal Reserve Wire Network) is a real-time gross settlement funds transfer system operated by the United States Federal Reserve Banks that allows financial institutions to electronically transfer funds between its more than 9,289 participants (as of March 19, 2009).[1] Transfers can only be initiated by the sending bank once they receive the proper wiring instructions from the receiving bank. These instructions include: the receiving bank's routing number, account number, name and dollar amount being transferred. This information is submitted to the Federal Reserve via the Fedwire system. Once the instructions are received and processed, the Fed will debit the funds from the sending bank's reserve account and credit the receiving bank's account. Wire transfers sent via Fedwire are completed in the same day, while some are completed instantly.

[Source: Wikipedia - Fedwire](https://en.wikipedia.org/wiki/Fedwire)

### How does Fedwire work?

[What are Fedwire Transfers?](https://www.americanexpress.com/us/foreign-exchange/articles/fedwire-transfers/)

## Running Moov Wire Server

Moov Wire can be deployed in multiple scenarios.

- <a href="#binary-distribution">Binary Distributions</a> are released with every versioned release. These are frequently added to the VM/AMI build script for applications needing Moov Wire.
- A <a href="#docker-container">Docker container</a> is built and added to Docker Hub with every versioned release.
- Our hosted [api.moov.io](https://api.moov.io) is updated with every versioned release. Our Kubernetes example is what Moov utilizes in our production environment.

### Binary Distribution

Download the [latest Moov Wire server release](https://github.com/moov-io/wire/releases) for your operating system and run it from a terminal.

```sh
$ ./wire-darwin-amd64
ts=2019-06-20T23:23:44.870717Z caller=main.go:75 startup="Starting wire server version v0.2.0"
ts=2019-06-20T23:23:44.871623Z caller=main.go:135 transport=HTTP addr=:8088
ts=2019-06-20T23:23:44.871692Z caller=main.go:125 admin="listening on :9098"
```

Next: [Connect to Moov Wire](#connecting-to-moov-wire)

### Docker Container

Moov Wire is dependent on Docker being properly installed and running on your machine. Ensure that Docker is running. If your Docker client has issues connecting to the service, review the [Docker getting started guide](https://docs.docker.com/get-started/).

```sh
$ docker ps
CONTAINER ID        IMAGE               COMMAND             CREATED             STATUS              PORTS               NAMES
$
```

Execute the Docker run command

```sh
$ docker run -p 8088:8088 -p 9098:9098 moov/wire:latest
ts=2019-06-21T17:03:23.782592Z caller=main.go:69 startup="Starting wire server version v0.2.0"
ts=2019-06-21T17:03:23.78314Z caller=main.go:129 transport=HTTP addr=:8088
ts=2019-06-21T17:03:23.783252Z caller=main.go:119 admin="listening on :9098"
```

Next: [Connect to Moov Wire](#connecting-to-moov-wire)

### Kubernetes

The following snippet runs the Wire Server on [Kubernetes](https://kubernetes.io/docs/tutorials/kubernetes-basics/) in the `apps` namespace. You can reach the Wire instance at the following URL from inside the cluster.

```
# Needs to be ran from inside the cluster
$ curl http://wire.apps.svc.cluster.local:8088/ping
PONG
```

Kubernetes manifest - save in a file (`wire.yaml`) and apply with `kubectl apply -f wire.yaml`.

Next: [Connect to Moov Wire](#connecting-to-moov-wire)

## Connecting to Moov Wire

The Moov Wire service will be running on port `8088` (with an admin port on `9098`).

Confirm that the service is running by issuing the following command or simply visiting [localhost:8088/ping](http://localhost:8088/ping) in your browser.

```bash
$ curl http://localhost:8088/ping
PONG

$ curl http://localhost:8088/files
null
```

### API Documentation

See our [API documentation](https://moov-io.github.io/wire/api/) for Moov Wire endpoints.

### Wire Admin Port

The port `9098` is bound by ACH for our admin service. This HTTP server has endpoints for Prometheus metrics (`GET /metrics`), readiness checks (`GET /ready`) and liveness checks (`GET /live`).

## Getting Help

 channel | info
 ------- | -------
 [Project Documentation](https://moov-io.github.io/wire/) | Our project documentation available online.
Twitter [@moov_io](https://twitter.com/moov_io)	| You can follow Moov.IO's Twitter feed to get updates on our project(s). You can also tweet us questions or just share blogs or stories.
[GitHub Issue](https://github.com/moov-io/wire/issues) | If you are able to reproduce a problem please open a GitHub Issue under the specific project that caused the error.
[moov-io slack](https://slack.moov.io/) | Join our slack channel (`#wire`) to have an interactive discussion about the development of the project.
