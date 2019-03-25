// Copyright 2019 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package wire

// CoverPayment is cover payment data
type CoverPayment struct {
	// SwiftFieldTag
	SwiftFieldTag string `json:"swiftFieldTag,omitempty"`
	// SwiftLineOne
	SwiftLineOne string `json:"swiftLineOne,omitempty"`
	// SwiftLineTwo
	SwiftLineTwo string `json:"swiftLineTwo,omitempty"`
	// SwiftLineThree
	SwiftLineThree string `json:"swiftLineThree,omitempty"`
	// SwiftLineFour
	SwiftLineFour string `json:"swiftLineFour,omitempty"`
	// SwiftLineFive
	SwiftLineFive string `json:"swiftLineFive,omitempty"`
	// SwiftLineSix
	SwiftLineSix string `json:"swiftLineSix,omitempty"`
}