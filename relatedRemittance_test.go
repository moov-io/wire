package wire

import (
	"github.com/moov-io/base"
	"strings"
	"testing"
)

// RelatedRemittance creates a RelatedRemittance
func mockRelatedRemittance() *RelatedRemittance {
	rr := NewRelatedRemittance()
	rr.RemittanceIdentification = "Remittance Identification"
	rr.RemittanceLocationMethod = RLMElectronicDataExchange
	rr.RemittanceLocationElectronicAddress = "http://moov.io"
	rr.RemittanceData.Name = "Name"
	rr.RemittanceData.AddressType = CompletePostalAddress
	rr.RemittanceData.Department = "Department"
	rr.RemittanceData.SubDepartment = "Sub-Department"
	rr.RemittanceData.BuildingNumber = "16"
	rr.RemittanceData.PostCode = "19405"
	rr.RemittanceData.TownName = "AnyTown"
	rr.RemittanceData.CountrySubDivisionState = "PA"
	rr.RemittanceData.Country = "UA"
	rr.RemittanceData.AddressLineOne = "Address Line One"
	rr.RemittanceData.AddressLineTwo = "Address Line Two"
	rr.RemittanceData.AddressLineThree = "Address Line Three"
	rr.RemittanceData.AddressLineFour = "Address Line Four"
	rr.RemittanceData.AddressLineFive = "Address Line Five"
	rr.RemittanceData.AddressLineSix = "Address Line Six"
	rr.RemittanceData.AddressLineSeven = "Address Line Seven"
	return rr
}

// TestMockRelatedRemittance validates mockRelatedRemittance
func TestMockRelatedRemittance(t *testing.T) {
	rr := mockRelatedRemittance()
	if err := rr.Validate(); err != nil {
		t.Error("mockRelatedRemittance does not validate and will break other tests")
	}
}

// TestRelatedRemittanceLocationMethodValid validates RelatedRemittance RemittanceLocationMethod
func TestRelatedRemittanceLocationMethodValid(t *testing.T) {
	rr := mockRelatedRemittance()
	rr.RemittanceLocationMethod = "BBRB"
	if err := rr.Validate(); err != nil {
		if !base.Match(err, ErrRemittanceLocationMethod) {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// TestRelatedRemittanceLocationElectronicAddressAlphaNumeric validates RelatedRemittance
// RemittanceLocationElectronicAddressAlphaNumeric
func TestRelatedRemittanceLocationElectronicAddressAlphaNumeric(t *testing.T) {
	rr := mockRelatedRemittance()
	rr.RemittanceLocationElectronicAddress = "®"
	if err := rr.Validate(); err != nil {
		if !base.Match(err, ErrNonAlphanumeric) {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// TestRelatedRemittanceAddressTypeValid validates RelatedRemittance AddressType
func TestRelatedRemittanceAddressTypeValid(t *testing.T) {
	rr := mockRelatedRemittance()
	rr.RemittanceData.AddressType = "BBRB"
	if err := rr.Validate(); err != nil {
		if !base.Match(err, ErrAddressType) {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// TestRelatedRemittanceNameAlphaNumeric validates RelatedRemittance Name is alphanumeric
func TestRelatedRemittanceNameAlphaNumeric(t *testing.T) {
	rr := mockRelatedRemittance()
	rr.RemittanceData.Name = "®"
	if err := rr.Validate(); err != nil {
		if !base.Match(err, ErrNonAlphanumeric) {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// TestRelatedRemittanceDepartmentAlphaNumeric validates RelatedRemittance Department is alphanumeric
func TestRelatedRemittanceDepartmentAlphaNumeric(t *testing.T) {
	rr := mockRelatedRemittance()
	rr.RemittanceData.Department = "®"
	if err := rr.Validate(); err != nil {
		if !base.Match(err, ErrNonAlphanumeric) {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// TestRelatedRemittanceSubDepartmentAlphaNumeric validates RelatedRemittance SubDepartment is alphanumeric
func TestRelatedRemittanceSubDepartmentAlphaNumeric(t *testing.T) {
	rr := mockRelatedRemittance()
	rr.RemittanceData.SubDepartment = "®"
	if err := rr.Validate(); err != nil {
		if !base.Match(err, ErrNonAlphanumeric) {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// TestRelatedRemittanceStreetNameAlphaNumeric validates RelatedRemittance StreetName is alphanumeric
func TestRelatedRemittanceStreetNameAlphaNumeric(t *testing.T) {
	rr := mockRelatedRemittance()
	rr.RemittanceData.StreetName = "®"
	if err := rr.Validate(); err != nil {
		if !base.Match(err, ErrNonAlphanumeric) {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// TestRelatedRemittanceBuildingNumberAlphaNumeric validates RelatedRemittance BuildingNumber is alphanumeric
func TestRelatedRemittanceBuildingNumberAlphaNumeric(t *testing.T) {
	rr := mockRelatedRemittance()
	rr.RemittanceData.BuildingNumber = "®"
	if err := rr.Validate(); err != nil {
		if !base.Match(err, ErrNonAlphanumeric) {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// TestRelatedRemittancePostCodeAlphaNumeric validates RelatedRemittance PostCode is alphanumeric
func TestRelatedRemittancePostCodeAlphaNumeric(t *testing.T) {
	rr := mockRelatedRemittance()
	rr.RemittanceData.PostCode = "®"
	if err := rr.Validate(); err != nil {
		if !base.Match(err, ErrNonAlphanumeric) {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// TestRelatedRemittanceTownNameAlphaNumeric validates RelatedRemittance TownName is alphanumeric
func TestRelatedRemittanceTownNameAlphaNumeric(t *testing.T) {
	rr := mockRelatedRemittance()
	rr.RemittanceData.TownName = "®"
	if err := rr.Validate(); err != nil {
		if !base.Match(err, ErrNonAlphanumeric) {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// TestRelatedRemittanceCountrySubDivisionStateAlphaNumeric validates RelatedRemittance CountrySubDivisionState
// is alphanumeric
func TestRelatedRemittanceCountrySubDivisionStateAlphaNumeric(t *testing.T) {
	rr := mockRelatedRemittance()
	rr.RemittanceData.CountrySubDivisionState = "®"
	if err := rr.Validate(); err != nil {
		if !base.Match(err, ErrNonAlphanumeric) {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// TestRelatedRemittanceCountryAlphaNumeric validates RelatedRemittance Country is alphanumeric
func TestRelatedRemittanceCountryAlphaNumeric(t *testing.T) {
	rr := mockRelatedRemittance()
	rr.RemittanceData.Country = "®"
	if err := rr.Validate(); err != nil {
		if !base.Match(err, ErrNonAlphanumeric) {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// TestRelatedRemittanceAddressLineOneAlphaNumeric validates RelatedRemittance AddressLineOne is alphanumeric
func TestRelatedRemittanceAddressLineOneAlphaNumeric(t *testing.T) {
	rr := mockRelatedRemittance()
	rr.RemittanceData.AddressLineOne = "®"
	if err := rr.Validate(); err != nil {
		if !base.Match(err, ErrNonAlphanumeric) {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// TestRelatedRemittanceAddressLineTwoAlphaNumeric validates RelatedRemittance AddressLineTwo is alphanumeric
func TestRelatedRemittanceAddressLineTwoAlphaNumeric(t *testing.T) {
	rr := mockRelatedRemittance()
	rr.RemittanceData.AddressLineTwo = "®"
	if err := rr.Validate(); err != nil {
		if !base.Match(err, ErrNonAlphanumeric) {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// TestRelatedRemittanceAddressLineThreeAlphaNumeric validates RelatedRemittance AddressLineThree is alphanumeric
func TestRelatedRemittanceAddressLineThreeAlphaNumeric(t *testing.T) {
	rr := mockRelatedRemittance()
	rr.RemittanceData.AddressLineThree = "®"
	if err := rr.Validate(); err != nil {
		if !base.Match(err, ErrNonAlphanumeric) {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// TestRelatedRemittanceAddressLineFourAlphaNumeric validates RelatedRemittance AddressLineFour is alphanumeric
func TestRelatedRemittanceAddressLineFourAlphaNumeric(t *testing.T) {
	rr := mockRelatedRemittance()
	rr.RemittanceData.AddressLineFour = "®"
	if err := rr.Validate(); err != nil {
		if !base.Match(err, ErrNonAlphanumeric) {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// TestRelatedRemittanceAddressLineFiveAlphaNumeric validates RelatedRemittance AddressLineFive is alphanumeric
func TestRelatedRemittanceAddressLineFiveAlphaNumeric(t *testing.T) {
	rr := mockRelatedRemittance()
	rr.RemittanceData.AddressLineFive = "®"
	if err := rr.Validate(); err != nil {
		if !base.Match(err, ErrNonAlphanumeric) {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// TestRelatedRemittanceAddressLineSixAlphaNumeric validates RelatedRemittance AddressLineSix is alphanumeric
func TestRelatedRemittanceAddressLineSixAlphaNumeric(t *testing.T) {
	rr := mockRelatedRemittance()
	rr.RemittanceData.AddressLineSix = "®"
	if err := rr.Validate(); err != nil {
		if !base.Match(err, ErrNonAlphanumeric) {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// TestRelatedRemittanceAddressLineSevenAlphaNumeric validates RelatedRemittance AddressLineSeven is alphanumeric
func TestRelatedRemittanceAddressLineSevenAlphaNumeric(t *testing.T) {
	rr := mockRelatedRemittance()
	rr.RemittanceData.AddressLineSeven = "®"
	if err := rr.Validate(); err != nil {
		if !base.Match(err, ErrNonAlphanumeric) {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// TestRelatedRemittanceCountryOfResidenceAlphaNumeric validates RelatedRemittance CountryOfResidence is alphanumeric
func TestRelatedRemittanceCountryOfResidenceAlphaNumeric(t *testing.T) {
	rr := mockRelatedRemittance()
	rr.RemittanceData.CountryOfResidence = "®"
	if err := rr.Validate(); err != nil {
		if !base.Match(err, ErrNonAlphanumeric) {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// TestRelatedRemittanceNameRequired validates RelatedRemittance Name is required
func TestRelatedRemittanceNameRequired(t *testing.T) {
	rr := mockRelatedRemittance()
	rr.RemittanceData.Name = ""
	if err := rr.Validate(); err != nil {
		if !base.Match(err, ErrFieldRequired) {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// TestParseRelatedRemittanceWrongLength parses a wrong RelatedRemittance record length
func TestParseRelatedRemittanceWrongLength(t *testing.T) {
	var line = "{8250}Remittance Identification          EDIChttp://moov.io                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                  Name                                                                                                                                        ADDRDepartment                                                            Sub-Department                                                                                                                              16              19405           AnyTown                            PA                                 UAAddress Line One                                                      Address Line Two                                                      Address Line Three                                                    Address Line Four                                                     Address Line Five                                                     Address Line Six                                                      Address Line Seven                                                  "
	r := NewReader(strings.NewReader(line))
	r.line = line
	fwm := new(FEDWireMessage)
	rr := mockRelatedRemittance()
	fwm.SetRelatedRemittance(rr)
	err := r.parseRelatedRemittance()
	if err != nil {
		if !base.Match(err, NewTagWrongLengthErr(3041, len(r.line))) {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// TestParseRelatedRemittanceReaderParseError parses a wrong RelatedRemittance reader parse error
func TestParseRelatedRemittanceReaderParseError(t *testing.T) {
	var line = "{8250}Remittance ®dentification          EDIChttp://moov.io                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                  Name                                                                                                                                        ADDRDepartment                                                            Sub-Department                                                                                                                              16              19405           AnyTown                            PA                                 UAAddress Line One                                                      Address Line Two                                                      Address Line Three                                                    Address Line Four                                                     Address Line Five                                                     Address Line Six                                                      Address Line Seven                                                    "
	r := NewReader(strings.NewReader(line))
	r.line = line
	fwm := new(FEDWireMessage)
	rr := mockRelatedRemittance()
	fwm.SetRelatedRemittance(rr)
	err := r.parseRelatedRemittance()
	if err != nil {
		if !base.Match(err, ErrNonAlphanumeric) {
			t.Errorf("%T: %s", err, err)
		}
	}
	_, err = r.Read()
	if err != nil {
		if !base.Has(err, ErrNonAlphanumeric) {
			t.Errorf("%T: %s", err, err)
		}
	}
}
