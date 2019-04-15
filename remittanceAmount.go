// Copyright 2019 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package wire

// RemittanceAmount is remittance amount
type RemittanceAmount struct {
	// CurrencyCode
	CurrencyCode string `json:"currencyCode,omitempty"`
	// Amount Must contain at least one numeric character and only one decimal period marker (e.g., $1,234.56 should be entered as 1234.56). Can have up to 5 numeric characters following the decimal period marker (e.g., 1234.56789). Amount must be greater than zero (i.e., at least .01).
	Amount string `json:"amount,omitempty"`
}
