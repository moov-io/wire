---
layout: page
title: Intro
hide_hero: true
show_sidebar: false
menubar: docs-menu
---

## What is Fedwire?

Fedwire (formerly known as the Federal Reserve Wire Network) is a real-time gross settlement funds transfer system operated by the United States Federal Reserve Banks. It allows financial institutions to electronically transfer funds between participants. Transfers are initiated by the sending bank with routing instructions (receiving bank routing number, account number, name, and amount). Once processed, the Fed debits the sending bank's reserve account and credits the receiving bank's account. Fedwire transfers settle with finality on the same business day.

[Source: Wikipedia - Fedwire](https://en.wikipedia.org/wiki/Fedwire)

## ISO 20022 transition (July 2025)

On **July 14, 2025**, the Fedwire Funds Service completed a [single-day cutover](https://www.frbservices.org/resources/financial-services/wires/iso-20022-implementation-center) from the proprietary **FAIM** (Fedwire Application Interface Manual) format to **ISO 20022** XML messaging.

| | FAIM (legacy) | ISO 20022 (current) |
|---|---------------|---------------------|
| **Status on Fedwire Funds** | Deprecated — last accepted July 11, 2025 | Required as of July 14, 2025 |
| **Encoding** | Proprietary tag-based text (e.g. `{1500}`, `{2000}`) | XML (pacs, camt, pain, admi, …) |
| **Moov library** | [moov-io/wire](https://github.com/moov-io/wire) (this project) | [moov-io/wire20022](https://github.com/moov-io/wire20022), [moov-io/fedwire20022](https://github.com/moov-io/fedwire20022) |

This documentation and library describe the **legacy FAIM** format only. They remain useful for historical files, testing, and migration tooling.

Official Federal Reserve resources:

- [ISO 20022 Implementation Center](https://www.frbservices.org/resources/financial-services/wires/iso-20022-implementation-center)
- [Overview and implementation details](https://www.frbservices.org/resources/financial-services/wires/faq/iso-20022/overview-implementation-details)
- [July 2025 ISO 20022 newsletter](https://www.frbservices.org/resources/financial-services/wires/iso-20022-implementation-center/on-the-wire-iso-20022-newsletter/july-2025-iso-20022-newsletter)

## How does Fedwire work?

[What are Fedwire Transfers?](https://www.americanexpress.com/us/foreign-exchange/articles/fedwire-transfers/)
