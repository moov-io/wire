---
layout: page
title: Kubernetes
hide_hero: true
show_sidebar: false
menubar: docs-menu
---

# Kubernetes

The following snippet runs the Wire Server on [Kubernetes](https://kubernetes.io/docs/tutorials/kubernetes-basics/) in the `apps` namespace. You can reach the Wire instance at the following URL from inside the cluster.

```
# Needs to be ran from inside the cluster
$ curl http://wire.apps.svc.cluster.local:8088/ping
PONG
```

Kubernetes manifest - save in a file (`wire.yaml`) and apply with `kubectl apply -f wire.yaml`.