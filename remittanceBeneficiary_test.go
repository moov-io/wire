package wire

import (
	"github.com/moov-io/base"
	"testing"
)

// RemittanceBeneficiary creates a RemittanceBeneficiary
func mockRemittanceBeneficiary() *RemittanceBeneficiary {
	rb := NewRemittanceBeneficiary()
	rb.RemittanceData.Name = "Name"
	rb.IdentificationType = OrganizationID
	rb.IdentificationCode = OICCustomerNumber
	rb.IdentificationNumber = "111111"
	rb.IdentificationNumberIssuer = "Bank"
	rb.RemittanceData.DateBirthPlace = ""
	rb.RemittanceData.AddressType = CompletePostalAddress
	rb.RemittanceData.Department = "Department"
	rb.RemittanceData.SubDepartment = "Sub-Department"
	rb.RemittanceData.StreetName = "Street Name"
	rb.RemittanceData.BuildingNumber = "16"
	rb.RemittanceData.PostCode = "19405"
	rb.RemittanceData.TownName = "AnyTown"
	rb.RemittanceData.CountrySubDivisionState = "PA"
	rb.RemittanceData.Country = "UA"
	rb.RemittanceData.AddressLineOne = "Address Line One"
	rb.RemittanceData.AddressLineTwo = "Address Line Two"
	rb.RemittanceData.AddressLineThree = "Address Line Three"
	rb.RemittanceData.AddressLineFour = "Address Line Four"
	rb.RemittanceData.AddressLineFive = "Address Line Five"
	rb.RemittanceData.AddressLineSix = "Address Line Six"
	rb.RemittanceData.AddressLineSeven = "Address Line Seven"
	rb.RemittanceData.CountryOfResidence = "US"
	return rb
}

// TestMockRemittanceBeneficiary validates mockRemittanceBeneficiary
func TestMockRemittanceBeneficiary(t *testing.T) {
	rb := mockRemittanceBeneficiary()
	if err := rb.Validate(); err != nil {
		t.Error("mockRemittanceBeneficiary does not validate and will break other tests")
	}
}

// TestRemittanceBeneficiaryIdentificationTypeValid validates RemittanceBeneficiary IdentificationType
func TestRemittanceBeneficiaryIdentificationTypeValid(t *testing.T) {
	rb := mockRemittanceBeneficiary()
	rb.IdentificationType = "zz"
	if err := rb.Validate(); err != nil {
		if !base.Match(err, ErrIdentificationType) {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// TestRemittanceBeneficiaryIdentificationCodeValid validates RemittanceBeneficiary IdentificationCode
func TestRemittanceBeneficiaryIdentificationCodeValid(t *testing.T) {
	rb := mockRemittanceBeneficiary()
	rb.IdentificationCode = "zz"
	if err := rb.Validate(); err != nil {
		if !base.Match(err, ErrOrganizationIdentificationCode) {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// TestRemittanceBeneficiaryIdentificationCodeValid2 validates RemittanceBeneficiary IdentificationCode
func TestRemittanceBeneficiaryIdentificationCodeValid2(t *testing.T) {
	rb := mockRemittanceBeneficiary()
	rb.IdentificationType = PrivateID
	rb.IdentificationCode = "zz"
	if err := rb.Validate(); err != nil {
		if !base.Match(err, ErrPrivateIdentificationCode) {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// TestRemittanceBeneficiaryAddressTypeValid validates RemittanceBeneficiary AddressType
func TestRemittanceBeneficiaryAddressTypeValid(t *testing.T) {
	rb := mockRemittanceBeneficiary()
	rb.RemittanceData.AddressType = "BBRB"
	if err := rb.Validate(); err != nil {
		if !base.Match(err, ErrAddressType) {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// TestRemittanceBeneficiaryNameAlphaNumeric validates RemittanceBeneficiary Name is alphanumeric
func TestRemittanceBeneficiaryNameAlphaNumeric (t *testing.T) {
	rb := mockRemittanceBeneficiary()
	rb.RemittanceData.Name = "®"
	if err := rb.Validate(); err != nil {
		if !base.Match(err, ErrNonAlphanumeric) {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// TestRemittanceBeneficiaryIdentificationNumberAlphaNumeric validates RemittanceBeneficiary IdentificationNumber is alphanumeric
func TestRemittanceBeneficiaryIdentificationNumberAlphaNumeric (t *testing.T) {
	rb := mockRemittanceBeneficiary()
	rb.IdentificationNumber = "®"
	if err := rb.Validate(); err != nil {
		if !base.Match(err, ErrNonAlphanumeric) {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// TestRemittanceBeneficiaryIdentificationNumberIssuerAlphaNumeric validates RemittanceBeneficiary IdentificationNumberIssuer is alphanumeric
func TestRemittanceBeneficiaryIdentificationNumberIssuerAlphaNumeric (t *testing.T) {
	rb := mockRemittanceBeneficiary()
	rb.IdentificationNumberIssuer = "®"
	if err := rb.Validate(); err != nil {
		if !base.Match(err, ErrNonAlphanumeric) {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// TestRemittanceBeneficiaryDepartmentAlphaNumeric validates RemittanceBeneficiary Department is alphanumeric
func TestRemittanceBeneficiaryDepartmentAlphaNumeric (t *testing.T) {
	rb := mockRemittanceBeneficiary()
	rb.RemittanceData.Department = "®"
	if err := rb.Validate(); err != nil {
		if !base.Match(err, ErrNonAlphanumeric) {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// TestRemittanceBeneficiarySubDepartmentAlphaNumeric validates RemittanceBeneficiary SubDepartment is alphanumeric
func TestRemittanceBeneficiarySubDepartmentAlphaNumeric (t *testing.T) {
	rb := mockRemittanceBeneficiary()
	rb.RemittanceData.SubDepartment = "®"
	if err := rb.Validate(); err != nil {
		if !base.Match(err, ErrNonAlphanumeric) {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// TestRemittanceBeneficiaryStreetNameAlphaNumeric validates RemittanceBeneficiary StreetName is alphanumeric
func TestRemittanceBeneficiaryStreetNameAlphaNumeric (t *testing.T) {
	rb := mockRemittanceBeneficiary()
	rb.RemittanceData.StreetName = "®"
	if err := rb.Validate(); err != nil {
		if !base.Match(err, ErrNonAlphanumeric) {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// TestRemittanceBeneficiaryBuildingNumberAlphaNumeric validates RemittanceBeneficiary BuildingNumber is alphanumeric
func TestRemittanceBeneficiaryBuildingNumberAlphaNumeric (t *testing.T) {
	rb := mockRemittanceBeneficiary()
	rb.RemittanceData.BuildingNumber = "®"
	if err := rb.Validate(); err != nil {
		if !base.Match(err, ErrNonAlphanumeric) {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// TestRemittanceBeneficiaryPostCodeAlphaNumeric validates RemittanceBeneficiary PostCode is alphanumeric
func TestRemittanceBeneficiaryPostCodeAlphaNumeric (t *testing.T) {
	rb := mockRemittanceBeneficiary()
	rb.RemittanceData.PostCode = "®"
	if err := rb.Validate(); err != nil {
		if !base.Match(err, ErrNonAlphanumeric) {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// TestRemittanceBeneficiaryTownNameAlphaNumeric validates RemittanceBeneficiary TownName is alphanumeric
func TestRemittanceBeneficiaryTownNameAlphaNumeric (t *testing.T) {
	rb := mockRemittanceBeneficiary()
	rb.RemittanceData.TownName = "®"
	if err := rb.Validate(); err != nil {
		if !base.Match(err, ErrNonAlphanumeric) {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// TestRemittanceBeneficiaryCountrySubDivisionStateAlphaNumeric validates RemittanceBeneficiary CountrySubDivisionState
// is alphanumeric
func TestRemittanceBeneficiaryCountrySubDivisionStateAlphaNumeric (t *testing.T) {
	rb := mockRemittanceBeneficiary()
	rb.RemittanceData.CountrySubDivisionState = "®"
	if err := rb.Validate(); err != nil {
		if !base.Match(err, ErrNonAlphanumeric) {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// TestRemittanceBeneficiaryCountryAlphaNumeric validates RemittanceBeneficiary Country is alphanumeric
func TestRemittanceBeneficiaryCountryAlphaNumeric (t *testing.T) {
	rb := mockRemittanceBeneficiary()
	rb.RemittanceData.Country = "®"
	if err := rb.Validate(); err != nil {
		if !base.Match(err, ErrNonAlphanumeric) {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// TestRemittanceBeneficiaryAddressLineOneAlphaNumeric validates RemittanceBeneficiary AddressLineOne is alphanumeric
func TestRemittanceBeneficiaryAddressLineOneAlphaNumeric (t *testing.T) {
	rb := mockRemittanceBeneficiary()
	rb.RemittanceData.AddressLineOne = "®"
	if err := rb.Validate(); err != nil {
		if !base.Match(err, ErrNonAlphanumeric) {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// TestRemittanceBeneficiaryAddressLineTwoAlphaNumeric validates RemittanceBeneficiary AddressLineTwo is alphanumeric
func TestRemittanceBeneficiaryAddressLineTwoAlphaNumeric (t *testing.T) {
	rb := mockRemittanceBeneficiary()
	rb.RemittanceData.AddressLineTwo = "®"
	if err := rb.Validate(); err != nil {
		if !base.Match(err, ErrNonAlphanumeric) {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// TestRemittanceBeneficiaryAddressLineThreeAlphaNumeric validates RemittanceBeneficiary AddressLineThree is alphanumeric
func TestRemittanceBeneficiaryAddressLineThreeAlphaNumeric (t *testing.T) {
	rb := mockRemittanceBeneficiary()
	rb.RemittanceData.AddressLineThree = "®"
	if err := rb.Validate(); err != nil {
		if !base.Match(err, ErrNonAlphanumeric) {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// TestRemittanceBeneficiaryAddressLineFourAlphaNumeric validates RemittanceBeneficiary AddressLineFour is alphanumeric
func TestRemittanceBeneficiaryAddressLineFourAlphaNumeric (t *testing.T) {
	rb := mockRemittanceBeneficiary()
	rb.RemittanceData.AddressLineFour = "®"
	if err := rb.Validate(); err != nil {
		if !base.Match(err, ErrNonAlphanumeric) {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// TestRemittanceBeneficiaryAddressLineFiveAlphaNumeric validates RemittanceBeneficiary AddressLineFive is alphanumeric
func TestRemittanceBeneficiaryAddressLineFiveAlphaNumeric (t *testing.T) {
	rb := mockRemittanceBeneficiary()
	rb.RemittanceData.AddressLineFive = "®"
	if err := rb.Validate(); err != nil {
		if !base.Match(err, ErrNonAlphanumeric) {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// TestRemittanceBeneficiaryAddressLineSixAlphaNumeric validates RemittanceBeneficiary AddressLineSix is alphanumeric
func TestRemittanceBeneficiaryAddressLineSixAlphaNumeric (t *testing.T) {
	rb := mockRemittanceBeneficiary()
	rb.RemittanceData.AddressLineSix = "®"
	if err := rb.Validate(); err != nil {
		if !base.Match(err, ErrNonAlphanumeric) {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// TestRemittanceBeneficiaryAddressLineSevenAlphaNumeric validates RemittanceBeneficiary AddressLineSeven is alphanumeric
func TestRemittanceBeneficiaryAddressLineSevenAlphaNumeric (t *testing.T) {
	rb := mockRemittanceBeneficiary()
	rb.RemittanceData.AddressLineSeven = "®"
	if err := rb.Validate(); err != nil {
		if !base.Match(err, ErrNonAlphanumeric) {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// TestRemittanceBeneficiaryCountryOfResidenceAlphaNumeric validates RemittanceBeneficiary CountryOfResidence is alphanumeric
func TestRemittanceBeneficiaryCountryOfResidenceAlphaNumeric (t *testing.T) {
	rb := mockRemittanceBeneficiary()
	rb.RemittanceData.CountryOfResidence = "®"
	if err := rb.Validate(); err != nil {
		if !base.Match(err, ErrNonAlphanumeric) {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// TestRemittanceBeneficiaryNameRequired validates RemittanceBeneficiary Name is required
func TestRemittanceBeneficiaryNameRequired (t *testing.T) {
	rb := mockRemittanceBeneficiary()
	rb.RemittanceData.Name = ""
	if err := rb.Validate(); err != nil {
		if !base.Match(err, ErrFieldRequired) {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// TestRemittanceBeneficiaryIdentificationNumberInvalid validates RemittanceBeneficiary IdentificationNumber
func TestRemittanceBeneficiaryIdentificationNumberInvalid(t *testing.T) {
	rb := mockRemittanceBeneficiary()
	rb.IdentificationCode = PICDateBirthPlace
	rb.IdentificationNumber = "zz"
	if err := rb.Validate(); err != nil {
		if !base.Match(err, ErrInvalidProperty) {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// TestIdentificationNumberIssuerInvalid_IdentificationNumber validates RemittanceBeneficiary IdentificationNumberIssuer
func TestIdentificationNumberIssuerInvalid_IdentificationNumber(t *testing.T) {
	rb := mockRemittanceBeneficiary()
	rb.IdentificationNumber = ""
	rb.IdentificationNumberIssuer = "zz"
	if err := rb.Validate(); err != nil {
		if !base.Match(err, ErrInvalidProperty) {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// TestIdentificationNumberIssuerInvalid_PICDateBirthPlace validates RemittanceBeneficiary IdentificationNumberIssuer
func TestIdentificationNumberIssuerInvalid_PICDateBirthPlace(t *testing.T) {
	rb := mockRemittanceBeneficiary()
	rb.IdentificationCode = PICDateBirthPlace
	rb.IdentificationNumberIssuer = "zz"
	if err := rb.Validate(); err != nil {
		if !base.Match(err, ErrInvalidProperty) {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// TestIdentificationNumberIssuerInvalid_OICSWIFTBICORBEI validates RemittanceBeneficiary IdentificationNumberIssuer
func TestIdentificationNumberIssuerInvalid_OICSWIFTBICORBEI(t *testing.T) {
	rb := mockRemittanceBeneficiary()
	rb.IdentificationCode = OICSWIFTBICORBEI
	rb.IdentificationNumberIssuer = "zz"
	if err := rb.Validate(); err != nil {
		if !base.Match(err, ErrInvalidProperty) {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// TestRemittanceBeneficiaryDateBirthPlaceInvalid validates RemittanceBeneficiary DateBirthPlace
func TestRemittanceBeneficiaryDateBirthPlaceInvalid(t *testing.T) {
	rb := mockRemittanceBeneficiary()
	rb.IdentificationCode = PICCustomerNumber
	rb.RemittanceData.DateBirthPlace = "Pottstown"
	if err := rb.Validate(); err != nil {
		if !base.Match(err, ErrInvalidProperty) {
			t.Errorf("%T: %s", err, err)
		}
	}
}