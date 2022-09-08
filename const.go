// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package wire

const (
	// TagMessageDisposition is MessageDisposition
	TagMessageDisposition = "{1100}"
	// TagReceiptTimeStamp is ReceiptTimeStamp
	TagReceiptTimeStamp = "{1110}"
	// TagOutputMessageAccountabilityData is OutputMessageAccountabilityData
	TagOutputMessageAccountabilityData = "{1120}"
	// TagErrorWire  is ErrorWire
	TagErrorWire = "{1130}"

	// TagSenderSupplied is SenderSuppliedInformation
	TagSenderSupplied = "{1500}"
	// TagTypeSubType is TypeSubType
	TagTypeSubType = "{1510}"
	// TagInputMessageAccountabilityData is InputMessageAccountabilityData (IMAD)
	TagInputMessageAccountabilityData = "{1520}"
	// TagAmount is Amount
	TagAmount = "{2000}"
	// TagSenderDepositoryInstitution is SenderDepositoryInstitution
	TagSenderDepositoryInstitution = "{3100}"
	// TagReceiverDepositoryInstitution is ReceiverDepositoryInstitution
	TagReceiverDepositoryInstitution = "{3400}"
	// TagBusinessFunctionCode is BusinessFunctionCode
	TagBusinessFunctionCode = "{3600}"

	// TagSenderReference is SenderReference
	TagSenderReference = "{3320}"
	// TagPreviousMessageIdentifier is PreviousMessageIdentifier
	TagPreviousMessageIdentifier = "{3500}"
	// TagLocalInstrument is LocalInstrument
	TagLocalInstrument = "{3610}"
	// TagPaymentNotification is PaymentNotification
	TagPaymentNotification = "{3620}"
	// TagCharges is Charges
	TagCharges = "{3700}"
	// TagInstructedAmount is InstructedAmount
	TagInstructedAmount = "{3710}"
	// TagExchangeRate is ExchangeRate
	TagExchangeRate = "{3720}"

	// TagBeneficiaryIntermediaryFI is BeneficiaryIntermediaryFI
	TagBeneficiaryIntermediaryFI = "{4000}"
	// TagBeneficiaryFI is BeneficiaryFI
	TagBeneficiaryFI = "{4100}"
	// TagBeneficiary is Beneficiary
	TagBeneficiary = "{4200}"
	// TagBeneficiaryReference  is BeneficiaryReference
	TagBeneficiaryReference = "{4320}"
	// TagAccountDebitedDrawdown is AccountDebitedDrawdown
	TagAccountDebitedDrawdown = "{4400}"

	// TagOriginator is Originator
	TagOriginator = "{5000}"
	// TagOriginatorOptionF is OriginatorOptionF
	TagOriginatorOptionF = "{5010}"
	// TagOriginatorFI is OriginatorFI
	TagOriginatorFI = "{5100}"
	// TagInstructingFI is InstructingFI
	TagInstructingFI = "{5200}"
	// TagAccountCreditedDrawdown is AccountCreditedDrawdown
	TagAccountCreditedDrawdown = "{5400}"

	// TagOriginatorToBeneficiary is OriginatorToBeneficiary
	TagOriginatorToBeneficiary = "{6000}"
	// TagFIReceiverFI is FIReceiverFI
	TagFIReceiverFI = "{6100}"
	// TagFIDrawdownDebitAccountAdvice is FIDrawdownDebitAccountAdvice
	TagFIDrawdownDebitAccountAdvice = "{6110}"
	// TagFIIntermediaryFI is FIIntermediaryFI
	TagFIIntermediaryFI = "{6200}"
	// TagFIIntermediaryFIAdvice is FIIntermediaryFIAdvice
	TagFIIntermediaryFIAdvice = "{6210}"
	// TagFIBeneficiaryFI is FIBeneficiaryFI
	TagFIBeneficiaryFI = "{6300}"
	// TagFIBeneficiaryFIAdvice is FIBeneficiaryFIAdvice
	TagFIBeneficiaryFIAdvice = "{6310}"
	// TagFIBeneficiary is FIBeneficiary
	TagFIBeneficiary = "{6400}"
	// TagFIBeneficiaryAdvice is FIBeneficiaryAdvice
	TagFIBeneficiaryAdvice = "{6410}"
	// TagFIPaymentMethodToBeneficiary is FIPaymentMethodToBeneficiary
	TagFIPaymentMethodToBeneficiary = "{6420}"
	// TagFIAdditionalFIToFI is FIAdditionalFIToFI
	TagFIAdditionalFIToFI = "{6500}"

	// TagCurrencyInstructedAmount is CurrencyInstructedAmount
	TagCurrencyInstructedAmount = "{7033}"
	// TagOrderingCustomer is OrderingCustomer
	TagOrderingCustomer = "{7050}"
	// TagOrderingInstitution is OrderingInstitution
	TagOrderingInstitution = "{7052}"
	// TagIntermediaryInstitution is IntermediaryInstitution
	TagIntermediaryInstitution = "{7056}"
	// TagInstitutionAccount is InstitutionAccount
	TagInstitutionAccount = "{7057}"
	// TagBeneficiaryCustomer is BeneficiaryCustomer
	TagBeneficiaryCustomer = "{7059}"
	// TagRemittance is Remittance
	TagRemittance = "{7070}"
	// TagSenderToReceiver is SenderToReceiver
	TagSenderToReceiver = "{7072}"

	// TagUnstructuredAddenda is UnstructuredAddenda
	TagUnstructuredAddenda = "{8200}"

	// TagRelatedRemittance is RelatedRemittance
	TagRelatedRemittance = "{8250}"
	// TagRemittanceOriginator is RemittanceOriginator
	TagRemittanceOriginator = "{8300}"
	// TagRemittanceBeneficiary is RemittanceBeneficiary
	TagRemittanceBeneficiary = "{8350}"
	// TagPrimaryRemittanceDocument is PrimaryRemittanceDocument
	TagPrimaryRemittanceDocument = "{8400}"
	// TagActualAmountPaid is ActualAmountPaid
	TagActualAmountPaid = "{8450}"
	// TagGrossAmountRemittanceDocument is GrossAmountRemittanceDocument
	TagGrossAmountRemittanceDocument = "{8500}"
	// TagAmountNegotiatedDiscount is AmountNegotiatedDiscount
	TagAmountNegotiatedDiscount = "{8550}"
	// TagAdjustment is Adjustment
	TagAdjustment = "{8600}"
	// TagDateRemittanceDocument is DateRemittanceDocument
	TagDateRemittanceDocument = "{8650}"
	// TagSecondaryRemittanceDocument is SecondaryRemittanceDocument
	TagSecondaryRemittanceDocument = "{8700}"

	//TagRemittanceFreeText is RemittanceFreeText
	TagRemittanceFreeText = "{8750}"

	// TagServiceMessage is ServiceMessage
	TagServiceMessage = "{9000}"

	// FormatVersion designates the FEDWIRE message format version
	FormatVersion = "30"
	// EnvironmentTest designates a test environment
	EnvironmentTest = "T"
	// EnvironmentProduction designates a production environment
	EnvironmentProduction = "P"
	// MessageDuplicationOriginal designates an original message
	MessageDuplicationOriginal = " "
	// MessageDuplicationResend designates a resend of a message
	MessageDuplicationResend = "P"

	// TypeCode

	// FundsTransfer is SenderSuppliedInformation {1510} TypeCode which designates a funds transfer in which the
	// sender and/or receiver may be a bank or a third party (i.e., customer of a bank).
	FundsTransfer = "10"
	// ForeignTransfer is SenderSuppliedInformation {1510} TypeCode which designates a funds transfer to
	// or from a foreign central bank or government or international organization with an account at the Federal
	// Reserve Bank of New York.
	ForeignTransfer = "15"
	// SettlementTransfer is SenderSuppliedInformation {1510} TypeCode which designates a funds transfer
	// between Fedwire Funds Service participants.
	SettlementTransfer = "16"

	// SubTypeCode

	// BasicFundsTransfer is SenderSuppliedInformation {1510} SubTypeCode which designates a basic value funds transfer.
	BasicFundsTransfer = "00"
	// RequestReversal is SenderSuppliedInformation {1510} SubTypeCode which designates a non-value request for
	// reversal of a funds transfer
	// originated on the current business day.
	RequestReversal = "01"
	// ReversalTransfer is SenderSuppliedInformation {1510} SubTypeCode which designates a value reversal of a
	// funds transfer received on the current business day.  May be used in response to a subtype
	// code 01 Request for Reversal.
	ReversalTransfer = "02"
	// RequestReversalPriorDayTransfer is SenderSuppliedInformation {1510} SubTypeCode which designates a non-value
	// request for a reversal of a funds transfer originated on a prior business day.
	RequestReversalPriorDayTransfer = "07"
	// ReversalPriorDayTransfer is SenderSuppliedInformation {1510} SubTypeCode which designates a value reversal of
	// a funds transfer received on a prior business day.  May be used in response to a subtype code 07 Request for
	// Reversal of a Prior Day Transfer.
	ReversalPriorDayTransfer = "08"
	// RequestCredit is SenderSuppliedInformation {1510} SubTypeCode which designates a non-value request for the
	// receiver to send a funds transfer to a designated party.
	RequestCredit = "31"
	// FundsTransferRequestCredit is SenderSuppliedInformation {1510} SubTypeCode which designates a value funds
	// transfer honoring a subtype 31 request for credit.
	FundsTransferRequestCredit = "32"
	// RefusalRequestCredit is SenderSuppliedInformation {1510} SubTypeCode which designates a non-value message
	// indicating refusal to honor a subtype 31 request for credit.
	RefusalRequestCredit = "33"
	// SSIServiceMessage is SenderSuppliedInformation {1510} SubTypeCode which designates a non-value message used to
	// communicate questions and information that is not covered by a specific subtype.
	SSIServiceMessage = "90"

	// BusinessFunctionCode

	// BankTransfer is a bank transfer (beneficiary is a bank)
	BankTransfer = "BTR"
	// CheckSameDaySettlement is a check with same day settlement
	CheckSameDaySettlement = "CKS"
	// CustomerTransferPlus is a customer transfer plus
	CustomerTransferPlus = "CTP"
	// CustomerTransfer beneficiary is a not a bank
	CustomerTransfer = "CTR"
	// DepositSendersAccount is a deposit to a senders account
	DepositSendersAccount = "DEP"
	// BankDrawDownRequest is a bank to bank drawdown request
	BankDrawDownRequest = "DRB"
	// CustomerCorporateDrawdownRequest is a customer or corporate drawdown request
	CustomerCorporateDrawdownRequest = "DRC"
	// DrawdownResponse is a drawdown payment
	DrawdownResponse = "DRW"
	// FEDFundsReturned is FED funds returned
	FEDFundsReturned = "FFR"
	// FEDFundsSold is FED funds sold
	FEDFundsSold = "FFS"
	// BFCServiceMessage is Service Message
	BFCServiceMessage = "SVC"

	// ChargeDetails

	// CDBeneficiary is charge details beneficiary
	CDBeneficiary = "B"
	// CDShared is charge details shared
	CDShared = "S"

	// IDCode

	// SWIFTBankIdentifierCode is SWIFT Bank Identifier Code
	SWIFTBankIdentifierCode = "B"
	// CHIPSParticipant is CHIPS Participant
	CHIPSParticipant = "C"
	// DemandDepositAccountNumber is Demand Deposit Account Number
	DemandDepositAccountNumber = "D"
	// FEDRoutingNumber is FED Routing Number
	FEDRoutingNumber = "F"
	// SWIFTBICORBEIANDAccountNumber is SWIFT Bank Identifier Code(BIC) OR Bank Entity Identifier (BEI) AND AccountNumber
	SWIFTBICORBEIANDAccountNumber = "T"
	// CHIPSIdentifier is CHIPS Identifier
	CHIPSIdentifier = "U"
	// PassportNumber is Passport Number
	PassportNumber = "1"
	// TaxIdentificationNumber is Tax Identification Number
	TaxIdentificationNumber = "2"
	// DriversLicenseNumber is Drivers License Number
	DriversLicenseNumber = "3"
	// AlienRegistrationNumber is Alien Registration Number
	AlienRegistrationNumber = "4"
	// CorporateIdentification is corporate identification
	CorporateIdentification = "5"
	// OtherIdentification is other identification
	OtherIdentification = "9"

	// Drawdown Debit Account Advice Information

	// Advice Codes

	// AdviceCodeHold is an advice code for Hold
	AdviceCodeHold = "HLD"
	// AdviceCodeLetter is an advice code for LTR
	AdviceCodeLetter = "LTR"
	// AdviceCodePhone is an advice code for Phone
	AdviceCodePhone = "PHN"
	// AdviceCodeTelex is an advice code for Telex
	AdviceCodeTelex = "TLX"
	// AdviceCodeWire is an advice code for Wire
	AdviceCodeWire = "WRE"

	// RemittanceLocationMethod

	// RLMElectronicDataExchange is Remittance Location Method Electronic Data Exchange
	RLMElectronicDataExchange = "EDIC"
	// RLMEmail is Remittance Location Method Email
	RLMEmail = "EMAL"
	// RLMFax is Remittance Location Method Fax
	RLMFax = "FAXI"
	// RLMPostalService is Remittance Location Method Postal Service
	RLMPostalService = "POST"
	// RLMSMSM is Remittance Location Method Short Message Service (text)
	RLMSMSM = "SMSM"
	// RLMURI is Remittance Location Method Uniform Resource Identifier
	RLMURI = "URID"
	// is Remittance Location Method

	// AddressType

	// CompletePostalAddress is Complete Postal Address
	CompletePostalAddress = "ADDR"
	// HomeAddress is Home Address
	HomeAddress = "HOME"
	// BusinessAddress is Business Address
	BusinessAddress = "BIZZ"
	// MailAddress is Mail Address
	MailAddress = "MLTO"
	// DeliveryAddress is Delivery Address
	DeliveryAddress = "DLVY"
	// PostOfficeBox is Post Office Box
	PostOfficeBox = "PBOX"

	// Remittance IdentificationType

	// OrganizationID is Organization ID
	OrganizationID = "OI"
	// PrivateID is Private ID
	PrivateID = "PI"

	// Remittance Organization Identification Codes (OIC)

	// OICBankPartyIdentification is Bank Party Identification
	OICBankPartyIdentification = "BANK"
	// OICCustomerNumber is Customer Number
	OICCustomerNumber = "CUST"
	// OICDataUniversalNumberSystem (Dun & Bradstreet) is Data Universal Number System
	OICDataUniversalNumberSystem = "DUNS"
	// OICEmployerIdentificationNumber is Employee Identification Number
	OICEmployerIdentificationNumber = "EMPL"
	// OICGlobalLocationNumber is Global Location Number
	OICGlobalLocationNumber = "GS1G"
	// OICProprietaryIdentificationNumber is Proprietary Identification Number
	OICProprietaryIdentificationNumber = "PROP"
	// OICSWIFTBICORBEI is SWIFT BIC or BEI
	OICSWIFTBICORBEI = "SWBB"
	// OICTaxIdentificationNumber is Tax Identification Number
	OICTaxIdentificationNumber = "TXID"

	// Remittance Private Identification Codes (PIC)

	// PICAlienRegistrationNumber is Alien Registration Number
	PICAlienRegistrationNumber = "ARNU"
	// PICPassportNumber is Passport Number
	PICPassportNumber = "CCPT"
	// PICCustomerNumber is Customer Number
	PICCustomerNumber = "CUST"
	// PICDateBirthPlace is Date Birth Place
	PICDateBirthPlace = "DPOB"
	// PICEmployeeIdentificationNumber is Employer Identification Number
	PICEmployeeIdentificationNumber = "EMPL"
	// PICNationalIdentityNumber is National Identity Number
	PICNationalIdentityNumber = "NIDN"
	// PICProprietaryIdentificationNumber is Proprietary Identification Number
	PICProprietaryIdentificationNumber = "PROP"
	// PICSocialSecurityNumber is Social Security Number
	PICSocialSecurityNumber = "SOSE"
	// PICTaxIdentificationNumber is Tax Identification Number
	PICTaxIdentificationNumber = "TXID"

	// Document Type Code

	// AccountsReceivableOpenItem is accounts receivable open item
	AccountsReceivableOpenItem = "AROI"
	// BillLadingShippingNotice is bill lading shipping notice
	BillLadingShippingNotice = "BOLD"
	// CommercialInvoice is commercial invoice
	CommercialInvoice = "CINV"
	// CommercialContract is commercial contract
	CommercialContract = "CMCN"
	// CreditNoteRelatedFinancialAdjustment is credit note related financial adjustment
	CreditNoteRelatedFinancialAdjustment = "CNFA"
	// CreditNote is credit note
	CreditNote = "CREN"
	// DebitNote is debit note
	DebitNote = "DEBN"
	// DispatchAdvice is dispatch advice
	DispatchAdvice = "DISP"
	// DebitNoteRelatedFinancialAdjustment is debit note related financial adjustment
	DebitNoteRelatedFinancialAdjustment = "DNFA"
	// HireInvoice is hire invoice
	HireInvoice = "HIRI"
	// MeteredServiceInvoice is metered service invoice
	MeteredServiceInvoice = "MSIN"
	// ProprietaryDocumentType is proprietary document type
	ProprietaryDocumentType = "PROP"
	// PurchaseOrder is Purchase Order
	PurchaseOrder = "PUOR"
	// SelfBilledInvoice is self billed invoice
	SelfBilledInvoice = "SBIN"
	// StatementAccount is Statement of account
	StatementAccount = "SOAC"
	// TradeServicesUtilityTransaction is trade services utility transaction
	TradeServicesUtilityTransaction = "TSUT"
	// Voucher is Voucher
	Voucher = "VCHR"

	// Adjustment Reason Code

	// PricingError is pricing error
	PricingError = "01"
	// ExtensionError is extension error
	ExtensionError = "03"
	// ItemNotAcceptedDamaged is item not accepted damaged
	ItemNotAcceptedDamaged = "04"
	// ItemNotAcceptedQuality is item not accepted quality
	ItemNotAcceptedQuality = "05"
	// QuantityContested is quantity contested
	QuantityContested = "06"
	// IncorrectProduct is incorrect product
	IncorrectProduct = "07"
	// ReturnsDamaged is returns damaged
	ReturnsDamaged = "11"
	// ReturnsQuality is returns quality
	ReturnsQuality = "12"
	// ItemNotReceived is item not received
	ItemNotReceived = "59"
	// TotalOrderNotReceived is total order not received
	TotalOrderNotReceived = "75"
	// CreditAgreed is credit agreed
	CreditAgreed = "81"
	// CoveredCreditMemo is covered credit memo
	CoveredCreditMemo = "CM"

	// Debit / Credit Indicator

	// CreditIndicator is a credit
	CreditIndicator = "CRDT"
	// DebitIndicator is a debit
	DebitIndicator = "DBIT"

	// Local Instrument Code

	// ANSIX12format is ANSI X12 format
	ANSIX12format = "ANSI"
	// SequenceBCoverPaymentStructured is Sequence B Cover Payment Structured
	SequenceBCoverPaymentStructured = "COVS"
	// GeneralXMLformat is General XML format
	GeneralXMLformat = "GXML"
	// ISO20022XMLformat is ISO 20022 XML format
	ISO20022XMLformat = "IXML"
	// NarrativeText is Narrative Text
	NarrativeText = "NARR"
	// ProprietaryLocalInstrumentCode is Proprietary Local Instrument Code
	ProprietaryLocalInstrumentCode = "PROP"
	// RemittanceInformationStructured is Remittance Information Structured
	RemittanceInformationStructured = "RMTS"
	// RelatedRemittanceInformation is Related Remittance Information
	RelatedRemittanceInformation = "RRMT"
	// STP820format is STP 820 format
	STP820format = "S820"
	// SWIFTfield70 = SWIFT field 70
	SWIFTfield70 = "SWIF"
	// UNEDIFACTformat is UN-EDIFACT format
	UNEDIFACTformat = "UEDI"

	// PaymentMethod is the payment method to beneficiary.  'CHECK' is the only valid option
	PaymentMethod = "CHECK"

	// OriginatorOptionF PartyIdentifier

	// PartyIdentifierAlienRegistrationNumber is PartyIdentifier Alien Registration Number
	PartyIdentifierAlienRegistrationNumber = "ARNU"
	// PartyIdentifierPassportNumber is PartyIdentifier Passport Number
	PartyIdentifierPassportNumber = "CCPT"
	// PartyIdentifierCustomerIdentificationNumber is PartyIdentifier Customer Identification Number
	PartyIdentifierCustomerIdentificationNumber = "CUST"
	// PartyIdentifierDriversLicenseNumber is PartyIdentifier Driver’s License Number
	PartyIdentifierDriversLicenseNumber = "DRLC"
	// PartyIdentifierEmployerNumber is PartyIdentifier Employer Number
	PartyIdentifierEmployerNumber = "EMPL"
	// PartyIdentifierNationalIdentifyNumber is PartyIdentifier National Identify Number
	PartyIdentifierNationalIdentifyNumber = "NIDN"
	// PartyIdentifierSocialSecurityNumber is PartyIdentifier Social Security Number
	PartyIdentifierSocialSecurityNumber = "SOSE"
	// PartyIdentifierTaxIdentificationNumber is PartyIdentifier Tax Identification Number
	PartyIdentifierTaxIdentificationNumber = "TXID"

	// OriginatorOptionF LineOne, LineTwo, LineThree

	// OptionFName is Name
	OptionFName = "1"
	// OptionFAddress is Address
	OptionFAddress = "2"
	// OptionFCountryTown is Country and Town
	OptionFCountryTown = "3"
	// OptionFDOB  is Date of Birth
	OptionFDOB = "4"
	// OptionFBirthPlace is Place of Birth
	OptionFBirthPlace = "5"
	// OptionFCustomerIdentificationNumber is Customer Identification Number
	OptionFCustomerIdentificationNumber = "6"
	// OptionFNationalIdentityNumber is National Identity Number
	OptionFNationalIdentityNumber = "7"
	// OptionFAdditionalInformation is Additional Information
	OptionFAdditionalInformation = "8"
)
