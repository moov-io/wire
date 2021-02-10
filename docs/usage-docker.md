---
layout: page
title: Docker
hide_hero: true
show_sidebar: false
menubar: docs-menu
---

# Docker

We publish a [public Docker image `moov/wire`](https://hub.docker.com/r/moov/wire/tags) on Docker Hub with every tagged release of Wire. No configuration is required to serve on `:8088` and metrics at `:9098/metrics` in Prometheus format. We also have Docker images for [OpenShift](https://quay.io/repository/moov/wire?tab=tags) published as `quay.io/moov/wire`.

Moov Wire is dependent on Docker being properly installed and running on your machine. Ensure that Docker is running. If your Docker client has issues connecting to the service, review the [Docker getting started guide](https://docs.docker.com/get-started/).

```
docker ps
```
```
CONTAINER ID        IMAGE        COMMAND        CREATED        STATUS        PORTS        NAMES
```

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