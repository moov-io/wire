// Copyright 2019 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package wire

type Personal struct {
	// IdentificationCode:  * `1` - Passport Number * `2` - Tax Identification Number * `3` - Driver’s License Number * `4` - Alien Registration Number * `5` - Corporate Identification * `9` - Other Identification
	IdentificationCode string `json:"identificationCode"`
	// Identifier
	Identifier string `json:"identifier"`
	// Name
	Name string `json:"name"`
	Address Address `json:"address"`
}
