// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package fuzzreader

import (
	"bytes"

	"github.com/moov-io/wire"
)

// Return codes (from go-fuzz docs)
//
// The function must return 1 if the fuzzer should increase priority
// of the given input during subsequent fuzzing (for example, the input is
// lexically correct and was parsed successfully); -1 if the input must not be
// added to corpus even if gives new coverage; and 0 otherwise; other values are
// reserved for future use.
func Fuzz(data []byte) int {
	f, err := wire.NewReader(bytes.NewReader(data)).Read()
	if err != nil {
		return 0
	}
	if f.ID != "" {
		return 1
	}
	return checkSenderSupplied(f.FEDWireMessage.SenderSupplied)
}

func checkSenderSupplied(ss *wire.SenderSupplied) int {
	if ss == nil {
		return -1 // depriortize
	}
	if ss.FormatVersion != "" {
		return 1
	}
	if ss.TestProductionCode != "" {
		return 1
	}
	if ss.MessageDuplicationCode != "" {
		return 1
	}
	if ss.UserRequestCorrelation != "" {
		return 1
	}
	return 0
}
