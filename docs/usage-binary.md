---
layout: page
title: Binary distribution
hide_hero: true
show_sidebar: false
menubar: docs-menu
---

# Binary distribution

Download the [latest Moov Wire server release](https://github.com/moov-io/wire/releases) for your operating system and run it from a terminal.

```sh
$ ./wire-darwin-amd64
ts=2019-06-20T23:23:44.870717Z caller=main.go:75 startup="Starting wire server version v0.2.0"
ts=2019-06-20T23:23:44.871623Z caller=main.go:135 transport=HTTP addr=:8088
ts=2019-06-20T23:23:44.871692Z caller=main.go:125 admin="listening on :9098"
```

## Connecting to Moov Wire

The Moov Wire service will be running on port `8088` (with an admin port on `9098`).

Confirm that the service is running by issuing the following command or simply visiting [localhost:8088/ping](http://localhost:8088/ping) in your browser.

```bash
$ curl http://localhost:8088/ping
PONG

$ curl http://localhost:8088/files
null
```
