package wire

import (
	"github.com/moov-io/base"
	"testing"
)

// RemittanceOriginator creates a RemittanceOriginator
func mockRemittanceOriginator() *RemittanceOriginator {
	ro := NewRemittanceOriginator()
	ro.IdentificationType = OrganizationID
	ro.IdentificationCode = OICCustomerNumber
	ro.IdentificationNumber = "111111"
	ro.IdentificationNumberIssuer = "Bank"
	//ro.RemittanceData.DateBirthPlace = "12072008 AnyTown"
	ro.RemittanceData.Name = "Name"
	ro.RemittanceData.AddressType = CompletePostalAddress
	ro.RemittanceData.Department = "Department"
	ro.RemittanceData.SubDepartment = "Sub-Department"
	ro.RemittanceData.StreetName = "Street Name"
	ro.RemittanceData.BuildingNumber = "16"
	ro.RemittanceData.PostCode = "19405"
	ro.RemittanceData.TownName = "AnyTown"
	ro.RemittanceData.CountrySubDivisionState = "PA"
	ro.RemittanceData.Country = "UA"
	ro.RemittanceData.AddressLineOne = "Address Line One"
	ro.RemittanceData.AddressLineTwo = "Address Line Two"
	ro.RemittanceData.AddressLineThree = "Address Line Three"
	ro.RemittanceData.AddressLineFour = "Address Line Four"
	ro.RemittanceData.AddressLineFive = "Address Line Five"
	ro.RemittanceData.AddressLineSix = "Address Line Six"
	ro.RemittanceData.AddressLineSeven = "Address Line Seven"
	ro.RemittanceData.CountryOfResidence = "US"

	ro.ContactName = "Contact Name"
	ro.ContactPhoneNumber = "5551231212"
	ro.ContactMobileNumber = "5551231212"
	ro.ContactFaxNumber = "5551231212"
	ro.ContactElectronicAddress = "http://www.moov.io"
	ro.ContactOther = "Contact Other"
	return ro
}

// TestMockRemittanceOriginator validates mockRemittanceOriginator
func TestMockRemittanceOriginator(t *testing.T) {
	ro := mockRemittanceOriginator()
	if err := ro.Validate(); err != nil {
		t.Error("mockRemittanceOriginator does not validate and will break other tests")
	}
}

// TestRemittanceOriginatorNameAlphaNumeric validates RemittanceOriginator Name is alphanumeric
func TestRemittanceOriginatorNameAlphaNumeric (t *testing.T) {
	ro := mockRemittanceOriginator()
	ro.RemittanceData.Name = "®"
	if err := ro.Validate(); err != nil {
		if !base.Match(err, ErrNonAlphanumeric) {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// TestRemittanceOriginatorIdentificationNumberAlphaNumeric validates RemittanceOriginator IdentificationNumber is alphanumeric
func TestRemittanceOriginatorIdentificationNumberAlphaNumeric (t *testing.T) {
	ro := mockRemittanceOriginator()
	ro.IdentificationNumber = "®"
	if err := ro.Validate(); err != nil {
		if !base.Match(err, ErrNonAlphanumeric) {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// TestRemittanceOriginatorIdentificationNumberIssuerAlphaNumeric validates RemittanceOriginator IdentificationNumberIssuer is alphanumeric
func TestRemittanceOriginatorIdentificationNumberIssuerAlphaNumeric (t *testing.T) {
	ro := mockRemittanceOriginator()
	ro.IdentificationNumberIssuer = "®"
	if err := ro.Validate(); err != nil {
		if !base.Match(err, ErrNonAlphanumeric) {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// TestRemittanceOriginatorDepartmentAlphaNumeric validates RemittanceOriginator Department is alphanumeric
func TestRemittanceOriginatorDepartmentAlphaNumeric (t *testing.T) {
	ro := mockRemittanceOriginator()
	ro.RemittanceData.Department = "®"
	if err := ro.Validate(); err != nil {
		if !base.Match(err, ErrNonAlphanumeric) {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// TestRemittanceOriginatorSubDepartmentAlphaNumeric validates RemittanceOriginator SubDepartment is alphanumeric
func TestRemittanceOriginatorSubDepartmentAlphaNumeric (t *testing.T) {
	ro := mockRemittanceOriginator()
	ro.RemittanceData.SubDepartment = "®"
	if err := ro.Validate(); err != nil {
		if !base.Match(err, ErrNonAlphanumeric) {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// TestRemittanceOriginatorStreetNameAlphaNumeric validates RemittanceOriginator StreetName is alphanumeric
func TestRemittanceOriginatorStreetNameAlphaNumeric (t *testing.T) {
	ro := mockRemittanceOriginator()
	ro.RemittanceData.StreetName = "®"
	if err := ro.Validate(); err != nil {
		if !base.Match(err, ErrNonAlphanumeric) {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// TestRemittanceOriginatorBuildingNumberAlphaNumeric validates RemittanceOriginator BuildingNumber is alphanumeric
func TestRemittanceOriginatorBuildingNumberAlphaNumeric (t *testing.T) {
	ro := mockRemittanceOriginator()
	ro.RemittanceData.BuildingNumber = "®"
	if err := ro.Validate(); err != nil {
		if !base.Match(err, ErrNonAlphanumeric) {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// TestRemittanceOriginatorPostCodeAlphaNumeric validates RemittanceOriginator PostCode is alphanumeric
func TestRemittanceOriginatorPostCodeAlphaNumeric (t *testing.T) {
	ro := mockRemittanceOriginator()
	ro.RemittanceData.PostCode = "®"
	if err := ro.Validate(); err != nil {
		if !base.Match(err, ErrNonAlphanumeric) {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// TestRemittanceOriginatorTownNameAlphaNumeric validates RemittanceOriginator TownName is alphanumeric
func TestRemittanceOriginatorTownNameAlphaNumeric (t *testing.T) {
	ro := mockRemittanceOriginator()
	ro.RemittanceData.TownName = "®"
	if err := ro.Validate(); err != nil {
		if !base.Match(err, ErrNonAlphanumeric) {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// TestRemittanceOriginatorCountrySubDivisionStateAlphaNumeric validates RemittanceOriginator CountrySubDivisionState
// is alphanumeric
func TestRemittanceOriginatorCountrySubDivisionStateAlphaNumeric (t *testing.T) {
	ro := mockRemittanceOriginator()
	ro.RemittanceData.CountrySubDivisionState = "®"
	if err := ro.Validate(); err != nil {
		if !base.Match(err, ErrNonAlphanumeric) {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// TestRemittanceOriginatorCountryAlphaNumeric validates RemittanceOriginator Country is alphanumeric
func TestRemittanceOriginatorCountryAlphaNumeric (t *testing.T) {
	ro := mockRemittanceOriginator()
	ro.RemittanceData.Country = "®"
	if err := ro.Validate(); err != nil {
		if !base.Match(err, ErrNonAlphanumeric) {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// TestRemittanceOriginatorAddressLineOneAlphaNumeric validates RemittanceOriginator AddressLineOne is alphanumeric
func TestRemittanceOriginatorAddressLineOneAlphaNumeric (t *testing.T) {
	ro := mockRemittanceOriginator()
	ro.RemittanceData.AddressLineOne = "®"
	if err := ro.Validate(); err != nil {
		if !base.Match(err, ErrNonAlphanumeric) {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// TestRemittanceOriginatorAddressLineTwoAlphaNumeric validates RemittanceOriginator AddressLineTwo is alphanumeric
func TestRemittanceOriginatorAddressLineTwoAlphaNumeric (t *testing.T) {
	ro := mockRemittanceOriginator()
	ro.RemittanceData.AddressLineTwo = "®"
	if err := ro.Validate(); err != nil {
		if !base.Match(err, ErrNonAlphanumeric) {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// TestRemittanceOriginatorAddressLineThreeAlphaNumeric validates RemittanceOriginator AddressLineThree is alphanumeric
func TestRemittanceOriginatorAddressLineThreeAlphaNumeric (t *testing.T) {
	ro := mockRemittanceOriginator()
	ro.RemittanceData.AddressLineThree = "®"
	if err := ro.Validate(); err != nil {
		if !base.Match(err, ErrNonAlphanumeric) {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// TestRemittanceOriginatorAddressLineFourAlphaNumeric validates RemittanceOriginator AddressLineFour is alphanumeric
func TestRemittanceOriginatorAddressLineFourAlphaNumeric (t *testing.T) {
	ro := mockRemittanceOriginator()
	ro.RemittanceData.AddressLineFour = "®"
	if err := ro.Validate(); err != nil {
		if !base.Match(err, ErrNonAlphanumeric) {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// TestRemittanceOriginatorAddressLineFiveAlphaNumeric validates RemittanceOriginator AddressLineFive is alphanumeric
func TestRemittanceOriginatorAddressLineFiveAlphaNumeric (t *testing.T) {
	ro := mockRemittanceOriginator()
	ro.RemittanceData.AddressLineFive = "®"
	if err := ro.Validate(); err != nil {
		if !base.Match(err, ErrNonAlphanumeric) {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// TestRemittanceOriginatorAddressLineSixAlphaNumeric validates RemittanceOriginator AddressLineSix is alphanumeric
func TestRemittanceOriginatorAddressLineSixAlphaNumeric (t *testing.T) {
	ro := mockRemittanceOriginator()
	ro.RemittanceData.AddressLineSix = "®"
	if err := ro.Validate(); err != nil {
		if !base.Match(err, ErrNonAlphanumeric) {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// TestRemittanceOriginatorAddressLineSevenAlphaNumeric validates RemittanceOriginator AddressLineSeven is alphanumeric
func TestRemittanceOriginatorAddressLineSevenAlphaNumeric (t *testing.T) {
	ro := mockRemittanceOriginator()
	ro.RemittanceData.AddressLineSeven = "®"
	if err := ro.Validate(); err != nil {
		if !base.Match(err, ErrNonAlphanumeric) {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// TestRemittanceOriginatorCountryOfResidenceAlphaNumeric validates RemittanceOriginator CountryOfResidence is alphanumeric
func TestRemittanceOriginatorCountryOfResidenceAlphaNumeric (t *testing.T) {
	ro := mockRemittanceOriginator()
	ro.RemittanceData.CountryOfResidence = "®"
	if err := ro.Validate(); err != nil {
		if !base.Match(err, ErrNonAlphanumeric) {
			t.Errorf("%T: %s", err, err)
		}
	}
}