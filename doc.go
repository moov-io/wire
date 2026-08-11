// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

// Package wire implements a reader, writer, and validator for the legacy Fedwire Funds Service
// message format defined by the Fedwire Application Interface Manual (FAIM).
//
// # Fedwire ISO 20022 transition
//
// On July 14, 2025, the Federal Reserve Banks completed a single-day cutover of the Fedwire Funds
// Service from FAIM to ISO 20022. FAIM is no longer accepted for live Fedwire Funds traffic.
// This package continues to support FAIM for historical files, archival processing, testing, and
// migration tooling.
//
// For new Fedwire integrations, use:
//   - github.com/moov-io/wire20022 — read, write, and validate Fedwire ISO 20022 XML messages
//   - github.com/moov-io/fedwire20022 — generated Go types from the Fedwire ISO 20022 XSDs
//
// Official resources:
//   - https://www.frbservices.org/resources/financial-services/wires/iso-20022-implementation-center
//   - https://www.frbservices.org/resources/financial-services/wires/faq/iso-20022/overview-implementation-details
//
// # Usage
//
// Read a FAIM file:
//
//	file, err := wire.NewReader(r).Read()
//
// Create and write a FAIM file:
//
//	file := wire.NewFile()
//	// populate file.FEDWireMessage ...
//	err := wire.NewWriter(w).Write(file)
//
// See the project README and examples for business function codes and full message construction.
package wire
