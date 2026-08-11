---
layout: page
title: Overview
hide_hero: true
show_sidebar: false
menubar: docs-menu
---

# Overview

![Moov Wire Logo](https://repository-images.githubusercontent.com/174186064/4a00f100-c6d9-11ea-9aee-bd8f7fce633b)

Moov's mission is to give developers an easy way to create and integrate bank processing into their own software products. Our open source projects are each focused on solving a single responsibility in financial services and designed around performance, scalability, and ease of use.

**Moov Wire** implements a reader, writer, and validator for the **legacy Fedwire Funds Service message format (FAIM)** in an HTTP server and Go library. The HTTP server is available in a Docker image and the Go package `github.com/moov-io/wire` is available.

> **Fedwire transitioned to ISO 20022 on July 14, 2025.** FAIM is no longer accepted for live Fedwire Funds traffic. For new integrations, use [moov-io/wire20022](https://github.com/moov-io/wire20022). See [What is Fedwire?](/intro/) for details on the cutover.

**[Documentation](https://moov-io.github.io/wire)** | **[Source](https://github.com/moov-io/wire)** | **[Running](https://github.com/moov-io/wire#usage)** | **[Configuration](https://github.com/moov-io/wire#configuration-settings)**

## Purpose

Moov Wire implements a reader, writer, and validator for **legacy FAIM** Fedwire messages (the proprietary tag-based format from the Fedwire Application Interface Manual). It remains useful for historical files, archival processing, testing, and migration to ISO 20022.

For **ISO 20022** Fedwire messages (required on the Fedwire Funds Service as of July 14, 2025):

- [moov-io/wire20022](https://github.com/moov-io/wire20022) — read, write, and validate Fedwire ISO 20022 XML
- [moov-io/fedwire20022](https://github.com/moov-io/fedwire20022) — generated Go types from the Fedwire ISO 20022 XSDs

## Getting help

 channel | info
 ------- | -------
 [Project Documentation](https://moov-io.github.io/wire/) | Our project documentation available online.
Twitter [@moov](https://twitter.com/moov)| You can follow Moov.io's Twitter feed to get updates on our project(s). You can also tweet us questions or just share blogs or stories.
[GitHub Issue](https://github.com/moov-io/wire/issues) | If you are able to reproduce a problem please open a GitHub Issue under the specific project that caused the error.
[moov-io slack](https://slack.moov.io/) | Join our slack channel (`#wire`) to have an interactive discussion about the development of the project.
