// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/moov-io/wire"
)

func main() {
	f, err := os.Open(filepath.Join("examples", "bankTransfer-read", "bankTransfer.txt"))
	if err != nil {
		log.Fatalf("Could not open FEDWireMessage: %s\n", err)

	}
	defer f.Close()
	r := wire.NewReader(f)

	fwmFile, err := r.Read()
	if err != nil {
		log.Fatalf("Could not read FEDWireMessage: %s\n", err)
	}
	// ensure we have a validated file structure
	if err = fwmFile.Validate(); err != nil {
		log.Fatalf("Could not validate FEDWireMessage: %s\n", err)
	}

	fmt.Printf("Sender Supplied: %v \n", fwmFile.FEDWireMessages[0].SenderSupplied)
	fmt.Printf("Type and Subtype: %v \n", fwmFile.FEDWireMessages[0].TypeSubType)
	fmt.Printf("Input Message Accountability Data: %v \n", fwmFile.FEDWireMessages[0].InputMessageAccountabilityData)
	fmt.Printf("Amount: %v \n", fwmFile.FEDWireMessages[0].Amount)
	fmt.Printf("Sender Depository Institution: %v \n", fwmFile.FEDWireMessages[0].SenderDepositoryInstitution)
	fmt.Printf("Receiver Depository Institution: %v \n", fwmFile.FEDWireMessages[0].ReceiverDepositoryInstitution)
	fmt.Printf("Business Function Code: %v \n", fwmFile.FEDWireMessages[0].BusinessFunctionCode)
}
