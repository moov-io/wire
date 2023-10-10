// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package wire

import (
	"regexp"
	"strings"
	"unicode/utf8"

	"golang.org/x/text/currency"
)

var (
	// upperAlphanumericRegex = regexp.MustCompile(`[^ A-Z0-9!"#$%&'()*+,-.\\/:;<>=?@\[\]^_{}|~]+`)
	// alphanumericRegex = regexp.MustCompile(`[^ \w!"#$%&'()*+,-.\\/:;<>=?@\[\]^_{}|~]+`)

	// Alpha-Numeric including spaces and special characters as defined by FAIM 3.0.6
	alphanumericRegex = regexp.MustCompile(`[^ \w.?!,;:_@&/\\'"\x60~()<>$#%+-=]+`)

	numericRegex = regexp.MustCompile(`[^0-9]`)
	amountRegex  = regexp.MustCompile("[^0-9,.]")
)

// validator is common validation and formatting of golang types to WIRE type strings
type validator struct{}

// isAlphanumeric checks if a string only contains ASCII alphanumeric characters
func (v *validator) isAlphanumeric(s string) error {
	if alphanumericRegex.MatchString(s) {
		// ^[ A-Za-z0-9_@./#&+-]*$/
		return ErrNonAlphanumeric
	}
	return nil
}

// isNumeric checks if a string only contains ASCII numeric (0-9) characters
func (v *validator) isNumeric(s string) error {
	if numericRegex.MatchString(s) {
		// [^ 0-9]
		return ErrNonNumeric
	}
	return nil
}

// ToDo: Amount Decimal and AmountComma (only 1 per each) ?

// isAmount checks if a string only contains one comma and ASCII numeric (0-9) characters
func (v *validator) isAmount(s string) error {
	str := strings.Trim(s, ",")
	if amountRegex.MatchString(str) {
		// [^ [0-9],.]
		return ErrNonAmount
	}
	return nil
}

// isAmountImplied checks if a string contains only ASCII numeric (0-9) characters, decimal precision is
// implied (2), and no commas
func (v *validator) isAmountImplied(s string) error {
	if numericRegex.MatchString(s) {
		// [^ 0-9]
		return ErrNonAmount
	}
	return nil
}

// ToDo Add 5 decimal precision?

// isTypeCode ensures tag {1510} TypeCode is valid
func (v *validator) isTypeCode(code string) error {
	switch code {
	case
		FundsTransfer,
		ForeignTransfer,
		SettlementTransfer:
		return nil
	}
	return ErrTypeCode
}

// isSubTypeCode ensures tag {1510} SubTypeCode is valid
func (v *validator) isSubTypeCode(code string) error {
	switch code {
	case
		BasicFundsTransfer,
		RequestReversal,
		ReversalTransfer,
		RequestReversalPriorDayTransfer,
		ReversalPriorDayTransfer,
		RequestCredit,
		FundsTransferRequestCredit,
		RefusalRequestCredit,
		SSIServiceMessage:
		return nil
	}
	return ErrSubTypeCode
}

func (v *validator) isLocalInstrumentCode(code string) error {
	switch code {
	case
		ANSIX12format,
		SequenceBCoverPaymentStructured,
		GeneralXMLformat,
		ISO20022XMLformat,
		NarrativeText,
		ProprietaryLocalInstrumentCode,
		RemittanceInformationStructured,
		RelatedRemittanceInformation,
		STP820format,
		SWIFTfield70,
		UNEDIFACTformat:
		return nil
	}
	return ErrLocalInstrumentCode
}

func (v *validator) isTestProductionCode(code string) error {
	switch code {
	case
		EnvironmentTest,
		EnvironmentProduction:
		return nil
	}
	return ErrTestProductionCode
}

func (v *validator) isMessageDuplicationCode(code string) error {
	switch code {
	case
		MessageDuplicationOriginal,
		MessageDuplicationResend:
		return nil
	}
	return ErrMessageDuplicationCode
}

func (v *validator) isBusinessFunctionCode(code string) error {
	switch code {
	case
		BankTransfer,
		CheckSameDaySettlement,
		CustomerTransferPlus,
		CustomerTransfer,
		DepositSendersAccount,
		BankDrawDownRequest,
		CustomerCorporateDrawdownRequest,
		DrawdownResponse,
		FEDFundsReturned,
		FEDFundsSold,
		BFCServiceMessage:
		return nil
	}
	return ErrBusinessFunctionCode
}

func (v *validator) isChargeDetails(code string) error {
	switch code {
	case
		CDBeneficiary,
		CDShared:
		return nil
	}
	return ErrChargeDetails
}

func (v *validator) isTransactionTypeCode(code string) error {
	switch code {
	case
		"   ", "COV", "":
		return nil
	}
	return ErrTransactionTypeCode
}

func (v *validator) isIdentificationCode(code string) error {
	switch code {
	case
		SWIFTBankIdentifierCode,
		CHIPSParticipant,
		DemandDepositAccountNumber,
		FEDRoutingNumber,
		SWIFTBICORBEIANDAccountNumber,
		CHIPSIdentifier,
		PassportNumber,
		TaxIdentificationNumber,
		DriversLicenseNumber,
		AlienRegistrationNumber,
		CorporateIdentification,
		OtherIdentification:
		return nil
	}
	return ErrIdentificationCode
}

func (v *validator) isAdviceCode(code string) error {
	switch code {
	case
		AdviceCodeHold,
		AdviceCodeLetter,
		AdviceCodePhone,
		AdviceCodeTelex,
		AdviceCodeWire:
		return nil
	}
	return ErrAdviceCode
}

func (v *validator) isAddressType(code string) error {
	switch code {
	case
		CompletePostalAddress,
		HomeAddress,
		BusinessAddress,
		MailAddress,
		DeliveryAddress,
		PostOfficeBox:
		return nil
	}
	return ErrAddressType
}

func (v *validator) isRemittanceLocationMethod(code string) error {
	switch code {
	case
		RLMElectronicDataExchange,
		RLMEmail,
		RLMFax,
		RLMPostalService,
		RLMSMSM,
		RLMURI:
		return nil
	}
	return ErrRemittanceLocationMethod
}

func (v *validator) isIdentificationType(code string) error {
	switch code {
	case
		OrganizationID,
		PrivateID:
		return nil
	}
	return ErrIdentificationType
}

func (v *validator) isOrganizationIdentificationCode(code string) error {
	switch code {
	case
		OICBankPartyIdentification,
		OICCustomerNumber,
		OICDataUniversalNumberSystem,
		OICEmployerIdentificationNumber,
		OICGlobalLocationNumber,
		OICProprietaryIdentificationNumber,
		OICSWIFTBICORBEI,
		OICTaxIdentificationNumber:
		return nil
	}
	return ErrOrganizationIdentificationCode
}

func (v *validator) isPrivateIdentificationCode(code string) error {
	switch code {
	case
		PICAlienRegistrationNumber,
		PICPassportNumber,
		PICCustomerNumber,
		PICDateBirthPlace,
		PICEmployeeIdentificationNumber,
		PICNationalIdentityNumber,
		PICProprietaryIdentificationNumber,
		PICSocialSecurityNumber,
		PICTaxIdentificationNumber:
		return nil
	}
	return ErrPrivateIdentificationCode
}

func (v *validator) isDocumentTypeCode(code string) error {
	switch code {
	case
		AccountsReceivableOpenItem,
		BillLadingShippingNotice,
		CommercialInvoice,
		CommercialContract,
		CreditNoteRelatedFinancialAdjustment,
		CreditNote,
		DebitNote,
		DispatchAdvice,
		DebitNoteRelatedFinancialAdjustment,
		HireInvoice,
		MeteredServiceInvoice,
		ProprietaryDocumentType,
		PurchaseOrder,
		SelfBilledInvoice,
		StatementAccount,
		TradeServicesUtilityTransaction,
		Voucher:
		return nil
	}
	return ErrDocumentTypeCode
}

func (v *validator) isCreditDebitIndicator(code string) error {
	switch code {
	case
		CreditIndicator,
		DebitIndicator:
		return nil
	}
	return ErrCreditDebitIndicator
}

func (v *validator) isAdjustmentReasonCode(code string) error {
	switch code {
	case
		PricingError,
		ExtensionError,
		ItemNotAcceptedDamaged,
		ItemNotAcceptedQuality,
		QuantityContested,
		IncorrectProduct,
		ReturnsDamaged,
		ReturnsQuality,
		ItemNotReceived,
		TotalOrderNotReceived,
		CreditAgreed,
		CoveredCreditMemo:
		return nil
	}
	return ErrAdjustmentReasonCode
}

func (v *validator) isCurrencyCode(code string) error {
	_, err := currency.ParseISO(code)
	if err != nil {
		return ErrNonCurrencyCode
	}
	return nil
}

// isCentury validates a 2 digit century 20-29
func (v *validator) isCentury(s string) error {
	if s < "20" || s > "29" {
		return ErrValidCentury
	}
	return nil
}

// isYear validates a 2 digit year 00-99
func (v *validator) isYear(s string) error {
	if s < "00" || s > "99" {
		return ErrValidYear
	}
	return nil
}

// isMonth validates a 2 digit month 01-12
func (v *validator) isMonth(s string) error {
	switch s {
	case
		"01", "02", "03", "04", "05", "06",
		"07", "08", "09", "10", "11", "12":
		return nil
	}
	return ErrValidMonth
}

// isDay validates a 2 digit day based on a 2 digit month
// months are 01-12, days are 01-29, 01-30, or 01-31
func (v *validator) isDay(m string, d string) error {
	switch m {
	// February
	case "02":
		switch d {
		case
			"01", "02", "03", "04", "05", "06",
			"07", "08", "09", "10", "11", "12",
			"13", "14", "15", "16", "17", "18",
			"19", "20", "21", "22", "23", "24",
			"25", "26", "27", "28", "29":
			return nil
		}
	// April, June, September, November
	case "04", "06", "09", "11":
		switch d {
		case
			"01", "02", "03", "04", "05", "06",
			"07", "08", "09", "10", "11", "12",
			"13", "14", "15", "16", "17", "18",
			"19", "20", "21", "22", "23", "24",
			"25", "26", "27", "28", "29", "30":
			return nil
		}
	// January, March, May, July, August, October, December
	case "01", "03", "05", "07", "08", "10", "12":
		switch d {
		case
			"01", "02", "03", "04", "05", "06",
			"07", "08", "09", "10", "11", "12",
			"13", "14", "15", "16", "17", "18",
			"19", "20", "21", "22", "23", "24",
			"25", "26", "27", "28", "29", "30", "31":
			return nil
		}
	}
	return ErrValidDay
}

// validateDate will return the incoming string only if it matches a valid CCYYMMDD
// date format. (C=Century, Y=Year, M=Month, D=Day)
func (v *validator) validateDate(s string) error {
	if length := utf8.RuneCountInString(s); length != 8 {
		return NewTagWrongLengthErr(8, len(s))
	}
	cc, yy, mm, dd := s[:2], s[2:4], s[4:6], s[6:8]

	if err := v.isCentury(cc); err != nil {
		return ErrValidDate
	}
	if err := v.isYear(yy); err != nil {
		return ErrValidDate
	}
	if err := v.isMonth(mm); err != nil {
		return ErrValidDate
	}
	if err := v.isDay(mm, dd); err != nil {
		return ErrValidDate
	}
	return nil
}

// validatePartyIdentifier validates that PartyIdentifier must be one of the following two formats:
// 1. /Account Number (slash followed by at least one valid non-space character:  e.g., /123456)
func (v *validator) validatePartyIdentifier(s string) error {
	if s == "" {
		return ErrPartyIdentifier
	}
	if utf8.RuneCountInString(s) < 2 {
		return ErrPartyIdentifier
	}

	if s[:1] == "/" {
		if strings.TrimSpace(s[1:2]) == "" {
			return ErrPartyIdentifier
		}
		an := s[2:]
		if alphanumericRegex.MatchString(an) {
			return ErrPartyIdentifier
		}
	} else {
		if err := v.validateUIDPartyIdentifier(s); err != nil {
			return err
		}
	}
	return nil
}

// uidPartyIdentifier validates a unique identifier for PartyIdentifier which is not an account number
// 2. Unique Identifier/ (4 character code followed by a slash and at least one valid non-space character:
// e.g., SOSE/123-456-789)
//
// ARNU: Alien Registration Number
// CCPT: Passport Number
// CUST: Customer Identification Number
// DRLC: Driver’s License Number
// EMPL: Employer Number
// NIDN: National Identify Number
// SOSE: Social Security Number
// TXID: Tax Identification Number
func (v *validator) validateUIDPartyIdentifier(s string) error {
	if utf8.RuneCountInString(s) < 7 {
		return ErrPartyIdentifier
	}
	uid := s[:4]
	switch uid {
	case
		PartyIdentifierAlienRegistrationNumber,
		PartyIdentifierPassportNumber,
		PartyIdentifierCustomerIdentificationNumber,
		PartyIdentifierDriversLicenseNumber,
		PartyIdentifierEmployerNumber,
		PartyIdentifierNationalIdentifyNumber,
		PartyIdentifierSocialSecurityNumber,
		PartyIdentifierTaxIdentificationNumber:
	default:
		return ErrPartyIdentifier
	}
	if s[4:5] != "/" {
		return ErrPartyIdentifier
	}
	if strings.TrimSpace(s[5:6]) == "" {
		return ErrPartyIdentifier
	}
	an := s[5:]
	if alphanumericRegex.MatchString(an) {
		return ErrPartyIdentifier
	}
	return nil
}

// validateOptionFLine validates OriginatorOptionF LineOne, LineTwo, LineThree
// Format: Must begin with one of the following Line Codes followed by a slash and at least one
// valid non-space character.
// 1 Name
// 2 Address
// 3 Country and Town
// 4 Date of Birth
// 5 Place of Birth
// 6 Customer Identification Number
// 7 National Identity Number
// 8 Additional Information
// For example:
// 2/123 MAIN STREET
// 3/US/NEW YORK, NY 10000
// 7/111-22-3456
func (v *validator) validateOptionFLine(s string) error {
	if s == "" {
		return nil
	}
	// Can be "" without an error, but if not it has to be at least 3.
	if utf8.RuneCountInString(s) < 3 {
		return ErrOptionFLine
	}
	switch s[:1] {
	case
		OptionFName,
		OptionFAddress,
		OptionFCountryTown,
		OptionFDOB,
		OptionFBirthPlace,
		OptionFCustomerIdentificationNumber,
		OptionFNationalIdentityNumber,
		OptionFAdditionalInformation:
	default:
		return ErrOptionFLine
	}
	if s[1:2] != "/" {
		return ErrOptionFLine
	}
	if strings.TrimSpace(s[2:3]) == "" {
		return ErrOptionFLine
	}
	an := strings.TrimSpace(s[2:])
	if alphanumericRegex.MatchString(an) {
		return ErrOptionFLine
	}
	return nil
}

// validateOptionFName validates OriginatorOptionF
// Name  Format:
//
//	Must begin with Line Code 1 followed by a slash and at least one valid non-space character:
//
// e.g., 1/SMITH JOHN.
func (v *validator) validateOptionFName(s string) error {
	if utf8.RuneCountInString(s) < 3 {
		return ErrOptionFName
	}
	if s[:1] != OptionFName {
		return ErrOptionFName
	}
	if s[1:2] != "/" {
		return ErrOptionFName
	}
	if strings.TrimSpace(s[2:3]) == "" {
		return ErrOptionFName
	}
	an := strings.TrimSpace(s[2:])
	if alphanumericRegex.MatchString(an) {
		return ErrOptionFName
	}
	return nil
}
