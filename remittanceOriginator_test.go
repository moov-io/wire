package wire

import (
	"github.com/moov-io/base"
	"strings"
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

// TestRemittanceOriginatorIdentificationTypeValid validates RemittanceOriginator IdentificationType
func TestRemittanceOriginatorIdentificationTypeValid(t *testing.T) {
	ro := mockRemittanceOriginator()
	ro.IdentificationType = "zz"
	if err := ro.Validate(); err != nil {
		if !base.Match(err, ErrIdentificationType) {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// TestRemittanceOriginatorIdentificationCodeValid validates RemittanceOriginator IdentificationCode
func TestRemittanceOriginatorIdentificationCodeValid(t *testing.T) {
	ro := mockRemittanceOriginator()
	ro.IdentificationCode = "zz"
	if err := ro.Validate(); err != nil {
		if !base.Match(err, ErrOrganizationIdentificationCode) {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// TestRemittanceOriginatorIdentificationCodeValid2 validates RemittanceOriginator IdentificationCode
func TestRemittanceOriginatorIdentificationCodeValid2(t *testing.T) {
	ro := mockRemittanceOriginator()
	ro.IdentificationType = PrivateID
	ro.IdentificationCode = "zz"
	if err := ro.Validate(); err != nil {
		if !base.Match(err, ErrPrivateIdentificationCode) {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// TestRemittanceOriginatorAddressTypeValid validates RemittanceOriginator AddressType
func TestRemittanceOriginatorAddressTypeValid(t *testing.T) {
	ro := mockRemittanceOriginator()
	ro.RemittanceData.AddressType = "BBRB"
	if err := ro.Validate(); err != nil {
		if !base.Match(err, ErrAddressType) {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// TestRemittanceOriginatorNameAlphaNumeric validates RemittanceOriginator Name is alphanumeric
func TestRemittanceOriginatorNameAlphaNumeric(t *testing.T) {
	ro := mockRemittanceOriginator()
	ro.RemittanceData.Name = "®"
	if err := ro.Validate(); err != nil {
		if !base.Match(err, ErrNonAlphanumeric) {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// TestRemittanceOriginatorIdentificationNumberAlphaNumeric validates RemittanceOriginator IdentificationNumber is alphanumeric
func TestRemittanceOriginatorIdentificationNumberAlphaNumeric(t *testing.T) {
	ro := mockRemittanceOriginator()
	ro.IdentificationNumber = "®"
	if err := ro.Validate(); err != nil {
		if !base.Match(err, ErrNonAlphanumeric) {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// TestRemittanceOriginatorIdentificationNumberIssuerAlphaNumeric validates RemittanceOriginator IdentificationNumberIssuer is alphanumeric
func TestRemittanceOriginatorIdentificationNumberIssuerAlphaNumeric(t *testing.T) {
	ro := mockRemittanceOriginator()
	ro.IdentificationNumberIssuer = "®"
	if err := ro.Validate(); err != nil {
		if !base.Match(err, ErrNonAlphanumeric) {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// TestRemittanceOriginatorDepartmentAlphaNumeric validates RemittanceOriginator Department is alphanumeric
func TestRemittanceOriginatorDepartmentAlphaNumeric(t *testing.T) {
	ro := mockRemittanceOriginator()
	ro.RemittanceData.Department = "®"
	if err := ro.Validate(); err != nil {
		if !base.Match(err, ErrNonAlphanumeric) {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// TestRemittanceOriginatorSubDepartmentAlphaNumeric validates RemittanceOriginator SubDepartment is alphanumeric
func TestRemittanceOriginatorSubDepartmentAlphaNumeric(t *testing.T) {
	ro := mockRemittanceOriginator()
	ro.RemittanceData.SubDepartment = "®"
	if err := ro.Validate(); err != nil {
		if !base.Match(err, ErrNonAlphanumeric) {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// TestRemittanceOriginatorStreetNameAlphaNumeric validates RemittanceOriginator StreetName is alphanumeric
func TestRemittanceOriginatorStreetNameAlphaNumeric(t *testing.T) {
	ro := mockRemittanceOriginator()
	ro.RemittanceData.StreetName = "®"
	if err := ro.Validate(); err != nil {
		if !base.Match(err, ErrNonAlphanumeric) {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// TestRemittanceOriginatorBuildingNumberAlphaNumeric validates RemittanceOriginator BuildingNumber is alphanumeric
func TestRemittanceOriginatorBuildingNumberAlphaNumeric(t *testing.T) {
	ro := mockRemittanceOriginator()
	ro.RemittanceData.BuildingNumber = "®"
	if err := ro.Validate(); err != nil {
		if !base.Match(err, ErrNonAlphanumeric) {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// TestRemittanceOriginatorPostCodeAlphaNumeric validates RemittanceOriginator PostCode is alphanumeric
func TestRemittanceOriginatorPostCodeAlphaNumeric(t *testing.T) {
	ro := mockRemittanceOriginator()
	ro.RemittanceData.PostCode = "®"
	if err := ro.Validate(); err != nil {
		if !base.Match(err, ErrNonAlphanumeric) {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// TestRemittanceOriginatorTownNameAlphaNumeric validates RemittanceOriginator TownName is alphanumeric
func TestRemittanceOriginatorTownNameAlphaNumeric(t *testing.T) {
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
func TestRemittanceOriginatorCountrySubDivisionStateAlphaNumeric(t *testing.T) {
	ro := mockRemittanceOriginator()
	ro.RemittanceData.CountrySubDivisionState = "®"
	if err := ro.Validate(); err != nil {
		if !base.Match(err, ErrNonAlphanumeric) {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// TestRemittanceOriginatorCountryAlphaNumeric validates RemittanceOriginator Country is alphanumeric
func TestRemittanceOriginatorCountryAlphaNumeric(t *testing.T) {
	ro := mockRemittanceOriginator()
	ro.RemittanceData.Country = "®"
	if err := ro.Validate(); err != nil {
		if !base.Match(err, ErrNonAlphanumeric) {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// TestRemittanceOriginatorAddressLineOneAlphaNumeric validates RemittanceOriginator AddressLineOne is alphanumeric
func TestRemittanceOriginatorAddressLineOneAlphaNumeric(t *testing.T) {
	ro := mockRemittanceOriginator()
	ro.RemittanceData.AddressLineOne = "®"
	if err := ro.Validate(); err != nil {
		if !base.Match(err, ErrNonAlphanumeric) {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// TestRemittanceOriginatorAddressLineTwoAlphaNumeric validates RemittanceOriginator AddressLineTwo is alphanumeric
func TestRemittanceOriginatorAddressLineTwoAlphaNumeric(t *testing.T) {
	ro := mockRemittanceOriginator()
	ro.RemittanceData.AddressLineTwo = "®"
	if err := ro.Validate(); err != nil {
		if !base.Match(err, ErrNonAlphanumeric) {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// TestRemittanceOriginatorAddressLineThreeAlphaNumeric validates RemittanceOriginator AddressLineThree is alphanumeric
func TestRemittanceOriginatorAddressLineThreeAlphaNumeric(t *testing.T) {
	ro := mockRemittanceOriginator()
	ro.RemittanceData.AddressLineThree = "®"
	if err := ro.Validate(); err != nil {
		if !base.Match(err, ErrNonAlphanumeric) {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// TestRemittanceOriginatorAddressLineFourAlphaNumeric validates RemittanceOriginator AddressLineFour is alphanumeric
func TestRemittanceOriginatorAddressLineFourAlphaNumeric(t *testing.T) {
	ro := mockRemittanceOriginator()
	ro.RemittanceData.AddressLineFour = "®"
	if err := ro.Validate(); err != nil {
		if !base.Match(err, ErrNonAlphanumeric) {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// TestRemittanceOriginatorAddressLineFiveAlphaNumeric validates RemittanceOriginator AddressLineFive is alphanumeric
func TestRemittanceOriginatorAddressLineFiveAlphaNumeric(t *testing.T) {
	ro := mockRemittanceOriginator()
	ro.RemittanceData.AddressLineFive = "®"
	if err := ro.Validate(); err != nil {
		if !base.Match(err, ErrNonAlphanumeric) {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// TestRemittanceOriginatorAddressLineSixAlphaNumeric validates RemittanceOriginator AddressLineSix is alphanumeric
func TestRemittanceOriginatorAddressLineSixAlphaNumeric(t *testing.T) {
	ro := mockRemittanceOriginator()
	ro.RemittanceData.AddressLineSix = "®"
	if err := ro.Validate(); err != nil {
		if !base.Match(err, ErrNonAlphanumeric) {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// TestRemittanceOriginatorAddressLineSevenAlphaNumeric validates RemittanceOriginator AddressLineSeven is alphanumeric
func TestRemittanceOriginatorAddressLineSevenAlphaNumeric(t *testing.T) {
	ro := mockRemittanceOriginator()
	ro.RemittanceData.AddressLineSeven = "®"
	if err := ro.Validate(); err != nil {
		if !base.Match(err, ErrNonAlphanumeric) {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// TestRemittanceOriginatorCountryOfResidenceAlphaNumeric validates RemittanceOriginator CountryOfResidence is alphanumeric
func TestRemittanceOriginatorCountryOfResidenceAlphaNumeric(t *testing.T) {
	ro := mockRemittanceOriginator()
	ro.RemittanceData.CountryOfResidence = "®"
	if err := ro.Validate(); err != nil {
		if !base.Match(err, ErrNonAlphanumeric) {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// TestRemittanceOriginatorContactNameAlphaNumeric validates RemittanceOriginator ContactName is alphanumeric
func TestRemittanceOriginatorContactNameAlphaNumeric(t *testing.T) {
	ro := mockRemittanceOriginator()
	ro.ContactName = "®"
	if err := ro.Validate(); err != nil {
		if !base.Match(err, ErrNonAlphanumeric) {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// TestRemittanceOriginatorContactPhoneNumberAlphaNumeric validates RemittanceOriginator ContactPhoneNumber is alphanumeric
func TestRemittanceOriginatorContactPhoneNumberAlphaNumeric(t *testing.T) {
	ro := mockRemittanceOriginator()
	ro.ContactPhoneNumber = "®"
	if err := ro.Validate(); err != nil {
		if !base.Match(err, ErrNonAlphanumeric) {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// TestRemittanceOriginatorContactMobileNumberAlphaNumeric validates RemittanceOriginator ContactMobileNumber is alphanumeric
func TestRemittanceOriginatorContactMobileNumberAlphaNumeric(t *testing.T) {
	ro := mockRemittanceOriginator()
	ro.ContactMobileNumber = "®"
	if err := ro.Validate(); err != nil {
		if !base.Match(err, ErrNonAlphanumeric) {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// TestRemittanceOriginatorContactFaxNumberAlphaNumeric validates RemittanceOriginator ContactFaxNumber is alphanumeric
func TestRemittanceOriginatorContactFaxNumberAlphaNumeric(t *testing.T) {
	ro := mockRemittanceOriginator()
	ro.ContactFaxNumber = "®"
	if err := ro.Validate(); err != nil {
		if !base.Match(err, ErrNonAlphanumeric) {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// TestRemittanceOriginatorContactElectronicAddressAlphaNumeric validates RemittanceOriginator ContactElectronicAddress
// is alphanumeric
func TestRemittanceOriginatorContactElectronicAddressAlphaNumeric(t *testing.T) {
	ro := mockRemittanceOriginator()
	ro.ContactElectronicAddress = "®"
	if err := ro.Validate(); err != nil {
		if !base.Match(err, ErrNonAlphanumeric) {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// TestRemittanceOriginatorContactOtherAlphaNumeric validates RemittanceOriginator ContactOther
// is alphanumeric
func TestRemittanceOriginatorContactOtherAlphaNumeric(t *testing.T) {
	ro := mockRemittanceOriginator()
	ro.ContactOther = "®"
	if err := ro.Validate(); err != nil {
		if !base.Match(err, ErrNonAlphanumeric) {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// TestRemittanceOriginatorNameRequired validates RemittanceOriginator Name is required
func TestRemittanceOriginatorNameRequired(t *testing.T) {
	ro := mockRemittanceOriginator()
	ro.RemittanceData.Name = ""
	if err := ro.Validate(); err != nil {
		if !base.Match(err, ErrFieldRequired) {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// TestRemittanceOriginatorIdentificationNumberInvalid validates RemittanceOriginator IdentificationNumber
func TestRemittanceOriginatorIdentificationNumberInvalid(t *testing.T) {
	ro := mockRemittanceOriginator()
	ro.IdentificationCode = PICDateBirthPlace
	ro.IdentificationNumber = "zz"
	if err := ro.Validate(); err != nil {
		if !base.Match(err, ErrInvalidProperty) {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// TestRemittanceOriginatorIdentificationNumberIssuerInvalid_IdentificationNumber validates RemittanceOriginator IdentificationNumberIssuer
func TestRemittanceOriginatorIdentificationNumberIssuerInvalid_IdentificationNumber(t *testing.T) {
	ro := mockRemittanceOriginator()
	ro.IdentificationNumber = ""
	ro.IdentificationNumberIssuer = "zz"
	if err := ro.Validate(); err != nil {
		if !base.Match(err, ErrInvalidProperty) {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// TestRemittanceOriginatorIdentificationNumberIssuerInvalid_PICDateBirthPlace validates RemittanceOriginator IdentificationNumberIssuer
func TestRemittanceOriginatorIdentificationNumberIssuerInvalid_PICDateBirthPlace(t *testing.T) {
	ro := mockRemittanceOriginator()
	ro.IdentificationCode = PICDateBirthPlace
	ro.IdentificationNumberIssuer = "zz"
	if err := ro.Validate(); err != nil {
		if !base.Match(err, ErrInvalidProperty) {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// TestRemittanceOriginatorIdentificationNumberIssuerInvalid_OICSWIFTBICORBEI validates RemittanceOriginator IdentificationNumberIssuer
func TestRemittanceOriginatorIdentificationNumberIssuerInvalid_OICSWIFTBICORBEI(t *testing.T) {
	ro := mockRemittanceOriginator()
	ro.IdentificationCode = OICSWIFTBICORBEI
	ro.IdentificationNumberIssuer = "zz"
	if err := ro.Validate(); err != nil {
		if !base.Match(err, ErrInvalidProperty) {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// TestRemittanceOriginatorDateBirthPlaceInvalid validates RemittanceOriginator DateBirthPlace
func TestRemittanceOriginatorDateBirthPlaceInvalid(t *testing.T) {
	ro := mockRemittanceOriginator()
	ro.IdentificationCode = PICCustomerNumber
	ro.RemittanceData.DateBirthPlace = "Pottstown"
	if err := ro.Validate(); err != nil {
		if !base.Match(err, ErrInvalidProperty) {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// TestParseRemittanceOriginatorWrongLength parses a wrong RemittanceOriginator record length
func TestParseRemittanceOriginatorWrongLength(t *testing.T) {
	var line = "{8300}OICUSTName                                                                                                                                        111111                             Bank                                                                                                                 ADDRDepartment                                                            Sub-Department                                                        Street Name                                                           16              19405           AnyTown                            PA                                 UAAddress Line One                                                      Address Line Two                                                      Address Line Three                                                    Address Line Four                                                     Address Line Five                                                     Address Line Six                                                      Address Line Seven                                                    USContact Name                                                                                                                                5551231212                         5551231212                         5551231212                         http://www.moov.io                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                              Contact Other                    "
	r := NewReader(strings.NewReader(line))
	r.line = line
	fwm := new(FEDWireMessage)
	ro := mockRemittanceOriginator()
	fwm.SetRemittanceOriginator(ro)
	err := r.parseRemittanceOriginator()
	if err != nil {
		if !base.Match(err, NewTagWrongLengthErr(3442, len(r.line))) {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// TestParseRemittanceOriginatorReaderParseError parses a wrong RemittanceOriginator reader parse error
func TestParseRemittanceOriginatorReaderParseError(t *testing.T) {
	var line = "{8300}OICUSTName                                                                                                                                        111111                             Bank                                                                                                                 ADDRDepartment                                                            Sub-Department                                                        Street Name                                                           16              19405           AnyTown                            PA                                 UA®ddress Line One                                                      Address Line Two                                                      Address Line Three                                                    Address Line Four                                                     Address Line Five                                                     Address Line Six                                                      Address Line Seven                                                    USContact Name                                                                                                                                5551231212                         5551231212                         5551231212                         http://www.moov.io                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                              Contact Other                      "
	r := NewReader(strings.NewReader(line))
	r.line = line
	fwm := new(FEDWireMessage)
	ro := mockRemittanceOriginator()
	fwm.SetRemittanceOriginator(ro)
	err := r.parseRemittanceOriginator()
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

// TestRemittanceOriginatorTagError validates a RemittanceOriginator tag
func TestRemittanceOriginatorTagError(t *testing.T) {
	ro := mockRemittanceOriginator()
	ro.tag = "{9999}"
	if err := ro.Validate(); err != nil {
		if !base.Match(err, ErrValidTagForType) {
			t.Errorf("%T: %s", err, err)
		}
	}
}
