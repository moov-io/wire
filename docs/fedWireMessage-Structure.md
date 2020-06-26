## The following Tags are mandatory for all transfers:


Name |
------------ | 
SenderSupplied |
TypeSubType |
InputMessageAccountabilityData |
Amount | 
SenderDepositoryInstitution |
ReceiverDepositoryInstitution |
BusinessFunctionCode |

### SenderSupplied

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FormatVersion** | **string** | FormatVersion 30  | 
**UserRequestCorrelation** | **string** | UserRequestCorrelation | 
**TestProductionCode** | **string** | TestProductionCode  * &#x60;T&#x60; - Test * &#x60;P&#x60; - Production  |
**MessageDuplicationCode** | **string** | MessageDuplicationCode  * &#x60; &#x60; - Original Message * &#x60;R&#x60; - Retrieval of an original message * &#x60;P&#x60; - Resend  | 

### TypeSubType

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TypeCode** | **string** | TypeCode:  * &#x60;10&#x60; - Funds Transfer - A funds transfer in which the sender and/or receiver may be a bank or a third party (i.e., customer of a bank). * &#x60;15&#x60; - Foreign Transfer - A funds transfer to or from a foreign central bank or government or international organization with an account at the Federal Reserve Bank of New York. * &#x60;16&#x60; - Settlement Transfer - A funds transfer between Fedwire Funds Service participants.  | 
**SubTypeCode** | **string** | SubTypeCode:  * &#x60;00&#x60; - Basic Funds Transfer - A basic value funds transfer. * &#x60;01&#x60; - Request for Reversal - A non-value request for reversal of a funds transfer originated on the current business day. * &#x60;02&#x60; - Reversal of Transfer - A value reversal of a funds transfer received on the current business day.  May be used in response to a subtype code ‘01’ Request for Reversal. * &#x60;07&#x60; - Request for Reversal of a Prior Day Transfer - A non-value request for a reversal of a funds transfer originated on a prior business day. * &#x60;08&#x60; - Reversal of a Prior Day Transfer - A value reversal of a funds transfer received on a prior business day.  May be used in response to a subtype code ‘07’ Request for Reversal of a Prior Day Transfer. * &#x60;31&#x60; - Request for Credit (Drawdown) - A non-value request for the receiver to send a funds transfer to a designated party. * &#x60;32&#x60; - Funds Transfer Honoring a Request for Credit (Drawdown) -  A value funds transfer honoring a subtype 31 request for credit. * &#x60;33&#x60; -Refusal to Honor a Request for Credit (Drawdown) - A non-value message indicating refusal to honor a subtype 31 request for credit. * &#x60;90&#x60; - Service Message - A non-value message used to communicate questions and information that is not covered by a specific subtype.  | 

### InputMessageAccountabilityData (IMAD)

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**InputCycleDate** | **string** | InputCycleDate CCYYMMDD  | 
**InputSource** | **string** | InputSource | 
**InputSequenceNumber** | **string** | InputSequenceNumber | 

### Amount

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Amount** | **string** | Amount 12 numeric, right-justified with leading zeros, an implied decimal point and no commas; e.g., $12,345.67 becomes 000001234567 Can be all zeros for subtype 90  | 

### SenderDepositoryInstitution

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SenderABANumber** | **string** | SenderABANumber | 
**SenderShortName** | **string** | SenderShortName | 

### ReceiverDepositoryInstitution

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ReceiverABANumber** | **string** | ReceiverABANumber | 
**ReceiverShortName** | **string** | ReceiverShortName | 

### BusinessFunctionCode

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BusinessFunctionCode** | **string** | BusinessFunctionCode * &#x60;BTR&#x60; - Bank Transfer (Beneficiary is a bank) * &#x60;DRC&#x60; - Customer or Corporate Drawdown Request * &#x60;CKS&#x60; - Check Same Day Settlement * &#x60;DRW&#x60; - Drawdown Payment * &#x60;CTP&#x60; - Customer Transfer Plus * &#x60;FFR&#x60; - Fed Funds Returned * &#x60;CTR&#x60; - Customer Transfer (Beneficiary is a not a bank) * &#x60;FFS&#x60; - Fed Funds Sold * &#x60;DEP&#x60; - Deposit to Sender’s Account * &#x60;SVC&#x60; - Service Message * &#x60;DRB&#x60; - Bank-to-Bank Drawdown Request  | 
**TransactionTypeCode** | **string** | TransactionTypeCode If {3600} is CTR, an optional Transaction Type Code element is permitted; however, the Transaction Type Code &#39;COV&#39; is not permitted. | [optional] 

## Other Transfer Information

### SenderReference

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SenderReference** | **string** | SenderReference | [optional]

### LocalInstrument

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LocalInstrumentCode** | **string** | LocalInstrument  * &#x60;ANSI&#x60; - ANSI X12 format * &#x60;COVS&#x60; - Sequence B Cover Payment Structured * &#x60;GXML&#x60; - General XML format * &#x60;IXML&#x60; - ISO 20022 XML formaT * &#x60;NARR&#x60; - Narrative Text * &#x60;PROP&#x60; - Proprietary Local Instrument Code * &#x60;RMTS&#x60; - Remittance Information Structured * &#x60;RRMT&#x60; - Related Remittance Information * &#x60;S820&#x60; - STP 820 format * &#x60;SWIF&#x60; - SWIFT field 70 (Remittance Information) * &#x60;UEDI&#x60; - UN/EDIFACT format  | [optional] 
**ProprietaryCode** | **string** | ProprietaryCode | [optional] 

### PaymentNotification

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PaymentNotificationIndicator** | **string** | PaymentNotificationIndicator  * &#x60;0 - 6&#x60; - Reserved for market practice conventions. * &#x60;7 - 9&#x60; - Reserved for bilateral agreements between Fedwire senders and receivers.  | [optional] 
**ContactNotificationElectronicAddress** | **string** | ContactNotificationElectronicAddress | [optional] 
**ContactName** | **string** | ContactName | [optional] 
**ContactPhoneNumber** | **string** | ContactPhoneNumber | [optional] 
**ContactMobileNumber** | **string** | ContactMobileNumber | [optional] 
**FaxNumber** | **string** | FaxNumber | [optional] 
**EndToEndIdentification** | **string** | EndToEndIdentification | [optional] 

### Charges

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ChargeDetails** | **string** | ChargeDetails * &#x60;B&#x60; - Beneficiary * &#x60;S&#x60; - Shared  | [optional] 
**SendersChargesOne** | **string** | SendersChargesOne  The first three characters must contain an alpha currency code (e.g., USD).  The remaining characters for the amount must begin with at least one numeric character (0-9) and only one decimal comma marker.  $1,234.56 should be entered as USD1234,56 and $0.99 should be entered as USD0,99.  | [optional] 
**SendersChargesTwo** | **string** | SendersChargesTwo  The first three characters must contain an alpha currency code (e.g., USD).  The remaining characters for the amount must begin with at least one numeric character (0-9) and only one decimal comma marker.  $1,234.56 should be entered as USD1234,56 and $0.99 should be entered as USD0,99.  | [optional] 
**SendersChargesThree** | **string** | SendersChargesThree  The first three characters must contain an alpha currency code (e.g., USD).  The remaining characters for the amount must begin with at least one numeric character (0-9) and only one decimal comma marker.  $1,234.56 should be entered as USD1234,56 and $0.99 should be entered as USD0,99.  | [optional] 
**SendersChargesFour** | **string** | SendersChargesFour  The first three characters must contain an alpha currency code (e.g., USD).  The remaining characters for the amount must begin with at least one numeric character (0-9) and only one decimal comma marker.  $1,234.56 should be entered as USD1234,56 and $0.99 should be entered as USD0,99.  | [optional] 

### InstructedAmount

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CurrencyCode** | **string** | CurrencyCode | [optional] 
**Amount** | **string** | Amount  Must begin with at least one numeric character (0-9) and contain only one decimal comma marker (e.g., $1,234.56 should be entered as 1234,56 and $0.99 should be entered as  | [optional] 

### ExchangeRate

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExchangeRate** | **string** | ExchangeRate is the exchange rate  Must contain at least one numeric character and only one decimal comma marker (e.g., an exchange rate of 1.2345 should be entered as 1,2345).  | [optional] 

## Beneficiary Information

### IntermediaryFI

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IdentificationCode** | **string** | IdentificationCode:  * &#x60;B&#x60; - SWIFT Bank Identifier Code (BIC) * &#x60;C&#x60; - CHIPS Participant * &#x60;D&#x60; - Demand Deposit Account (DDA) Number * &#x60;F&#x60; - Fed Routing Number * &#x60;T&#x60; - SWIFT BIC or Bank Entity Identifier (BEI) and Account Number * &#x60;U&#x60; - CHIPS Identifier  | 
**Identifier** | **string** | Identifier | 
**Name** | **string** | Name | 
**Address** | [**Address**](address.md) |  | 

### BeneficiaryFI

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IdentificationCode** | **string** | IdentificationCode:  * &#x60;B&#x60; - SWIFT Bank Identifier Code (BIC) * &#x60;C&#x60; - CHIPS Participant * &#x60;D&#x60; - Demand Deposit Account (DDA) Number * &#x60;F&#x60; - Fed Routing Number * &#x60;T&#x60; - SWIFT BIC or Bank Entity Identifier (BEI) and Account Number * &#x60;U&#x60; - CHIPS Identifier  | 
**Identifier** | **string** | Identifier | 
**Name** | **string** | Name | 
**Address** | [**Address**](address.md) |  | 

### Beneficiary

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IdentificationCode** | **string** | IdentificationCode:  * &#x60;1&#x60; - Passport Number * &#x60;2&#x60; - Tax Identification Number * &#x60;3&#x60; - Driver’s License Number * &#x60;4&#x60; - Alien Registration Number * &#x60;5&#x60; - Corporate Identification * &#x60;9&#x60; - Other Identification  | 
**Identifier** | **string** | Identifier | 
**Name** | **string** | Name | 
**Address** | [**Address**](address.md) |  | 

### BeneficiaryReference

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BeneficiaryReference** | **string** | BeneficiaryReference | [optional]

### AccountDebitedDrawdown

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IdentificationCode** | **string** | Identification Code * &#x60;D&#x60; - Debit  | 
**Identifier** | **string** | Identifier | 
**Name** | **string** | Name | 
**Address** | [**Address**](address.md) |  | [optional]

## Originator Information

### Originator

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IdentificationCode** | **string** | IdentificationCode:  * &#x60;1&#x60; - Passport Number * &#x60;2&#x60; - Tax Identification Number * &#x60;3&#x60; - Driver’s License Number * &#x60;4&#x60; - Alien Registration Number * &#x60;5&#x60; - Corporate Identification * &#x60;9&#x60; - Other Identification  | 
**Identifier** | **string** | Identifier | 
**Name** | **string** | Name | 
**Address** | [**Address**](address.md) |  | 

### OriginatorOptionF

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PartyIdentifier** | **string** | PartyIdentifier  Must be one of the following two formats: 1. /Account Number (slash followed by at least one valid non-space character:  e.g., /123456)  2. Unique Identifier/ (4 character code followed by a slash and at least one valid non-space character:    e.g., SOSE/123-456-789) ARNU: Alien Registration Number CCPT: Passport Number CUST: Customer Identification Number  DRLC/    Driver’s License Number  EMPL/    Employer Number NIDN: National Identify Number  SOSE/    Social Security Number TXID: Tax Identification Number  | [optional] 
**Name** | **string** | Name  Format:  Must begin with Line Code 1 followed by a slash and at least one valid non-space character: e.g., 1/SMITH JOHN.  | [optional] 
**LineOne** | **string** | LineOne  Format: Must begin with one of the following Line Codes followed by a slash and at least one valid non-space character. 1 Name 2 Address 3 Country and Town 4 Date of Birth 5 Place of Birth 6 Customer Identification Number 7 National Identity Number 8 Additional Information  For example: 2/123 MAIN STREET 3/US/NEW YORK, NY 10000 7/111-22-3456  | [optional] 
**LineTwo** | **string** | LineTwo  Format: Must begin with one of the following Line Codes followed by a slash and at least one valid non-space character. 1 Name 2 Address 3 Country and Town 4 Date of Birth 5 Place of Birth 6 Customer Identification Number 7 National Identity Number 8 Additional Information  For example: 2/123 MAIN STREET 3/US/NEW YORK, NY 10000 7/111-22-3456  | [optional] 
**LineThree** | **string** | LineThree  Format: Must begin with one of the following Line Codes followed by a slash and at least one valid non-space character. 1 Name 2 Address 3 Country and Town 4 Date of Birth 5 Place of Birth 6 Customer Identification Number 7 National Identity Number 8 Additional Information  For example: 2/123 MAIN STREET 3/US/NEW YORK, NY 10000 7/111-22-3456  | [optional] 
 
### OriginatorFI
 
 Name | Type | Description | Notes
 ------------ | ------------- | ------------- | -------------
 **IdentificationCode** | **string** | IdentificationCode:  * &#x60;B&#x60; - SWIFT Bank Identifier Code (BIC) * &#x60;C&#x60; - CHIPS Participant * &#x60;D&#x60; - Demand Deposit Account (DDA) Number * &#x60;F&#x60; - Fed Routing Number * &#x60;T&#x60; - SWIFT BIC or Bank Entity Identifier (BEI) and Account Number * &#x60;U&#x60; - CHIPS Identifier  | 
 **Identifier** | **string** | Identifier | 
 **Name** | **string** | Name | 
 **Address** | [**Address**](address.md) |  
 
### InstructingFI
  
 Name | Type | Description | Notes
 ------------ | ------------- | ------------- | -------------
 **IdentificationCode** | **string** | IdentificationCode:  * &#x60;B&#x60; - SWIFT Bank Identifier Code (BIC) * &#x60;C&#x60; - CHIPS Participant * &#x60;D&#x60; - Demand Deposit Account (DDA) Number * &#x60;F&#x60; - Fed Routing Number * &#x60;T&#x60; - SWIFT BIC or Bank Entity Identifier (BEI) and Account Number * &#x60;U&#x60; - CHIPS Identifier  | 
 **Identifier** | **string** | Identifier | 
 **Name** | **string** | Name | 
 **Address** | [**Address**](address.md) |
 
### AccountCreditedDrawdown
 
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DrawdownCreditAccountNumber** | **string** | DrawdownCreditAccountNumber  9 character ABA  | [optional]

### OriginatorToBeneficiary

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LineOne** | **string** | LineOne | [optional] 
**LineTwo** | **string** | LineTwo | [optional] 
**LineThree** | **string** | LineThree | [optional] 
**LineFour** | **string** | LineFour | [optional]

## Financial Institution to Financial Institution Information

### ReceiverFI

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LineOne** | **string** | LineOne | [optional] 
**LineTwo** | **string** | LineTwo | [optional] 
**LineThree** | **string** | LineThree | [optional] 
**LineFour** | **string** | LineFour | [optional] 
**LineFive** | **string** | LineFive | [optional] 
**LineSix** | **string** | LineSix | [optional]

### DrawdownDebitAccountAdvice

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AdviceCode** | **string** | AdviceCode  * &#x60;HLD - Hold * &#x60;LTR&#x60; - Letter * &#x60;PHN&#x60; - Phone * &#x60;TLX&#x60; - Telex * &#x60;WRE&#x60; - Wire  *  | [optional] 
**LineOne** | **string** | LineOne | [optional] 
**LineTwo** | **string** | LineTwo | [optional] 
**LineThree** | **string** | LineThree | [optional] 
**LineFour** | **string** | LineFour | [optional] 
**LineFive** | **string** | LineFive | [optional] 
**LineSix** | **string** | LineSix | [optional] 

### IntermediaryFI

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LineOne** | **string** | LineOne | [optional] 
**LineTwo** | **string** | LineTwo | [optional] 
**LineThree** | **string** | LineThree | [optional] 
**LineFour** | **string** | LineFour | [optional] 
**LineFive** | **string** | LineFive | [optional] 
**LineSix** | **string** | LineSix | [optional]

### IntermediaryFIAdvice

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AdviceCode** | **string** | AdviceCode  * &#x60;HLD - Hold * &#x60;LTR&#x60; - Letter * &#x60;PHN&#x60; - Phone * &#x60;TLX&#x60; - Telex * &#x60;WRE&#x60; - Wire  *  | [optional] 
**LineOne** | **string** | LineOne | [optional] 
**LineTwo** | **string** | LineTwo | [optional] 
**LineThree** | **string** | LineThree | [optional] 
**LineFour** | **string** | LineFour | [optional] 
**LineFive** | **string** | LineFive | [optional] 
**LineSix** | **string** | LineSix | [optional]

### BeneficiaryFI

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LineOne** | **string** | LineOne | [optional] 
**LineTwo** | **string** | LineTwo | [optional] 
**LineThree** | **string** | LineThree | [optional] 
**LineFour** | **string** | LineFour | [optional] 
**LineFive** | **string** | LineFive | [optional] 
**LineSix** | **string** | LineSix | [optional]

### BeneficiaryFIAdvice

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AdviceCode** | **string** | AdviceCode  * &#x60;HLD - Hold * &#x60;LTR&#x60; - Letter * &#x60;PHN&#x60; - Phone * &#x60;TLX&#x60; - Telex * &#x60;WRE&#x60; - Wire  *  | [optional] 
**LineOne** | **string** | LineOne | [optional] 
**LineTwo** | **string** | LineTwo | [optional] 
**LineThree** | **string** | LineThree | [optional] 
**LineFour** | **string** | LineFour | [optional] 
**LineFive** | **string** | LineFive | [optional] 
**LineSix** | **string** | LineSix | [optional]

### Beneficiary

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LineOne** | **string** | LineOne | [optional] 
**LineTwo** | **string** | LineTwo | [optional] 
**LineThree** | **string** | LineThree | [optional] 
**LineFour** | **string** | LineFour | [optional] 
**LineFive** | **string** | LineFive | [optional] 
**LineSix** | **string** | LineSix | [optional]

### Beneficiary Advice

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AdviceCode** | **string** | AdviceCode  * &#x60;HLD - Hold * &#x60;LTR&#x60; - Letter * &#x60;PHN&#x60; - Phone * &#x60;TLX&#x60; - Telex * &#x60;WRE&#x60; - Wire  *  | [optional] 
**LineOne** | **string** | LineOne | [optional] 
**LineTwo** | **string** | LineTwo | [optional] 
**LineThree** | **string** | LineThree | [optional] 
**LineFour** | **string** | LineFour | [optional] 
**LineFive** | **string** | LineFive | [optional] 
**LineSix** | **string** | LineSix | [optional]

### PaymentMethodToBeneficiary

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PaymentMethod** | **string** | PaymentMethod | [optional] 
**Additional** | **string** |  | [optional]

### AdditionalFiToFi

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LineOne** | **string** | LineOne | [optional] 
**LineTwo** | **string** | LineTwo | [optional] 
**LineThree** | **string** | LineThree | [optional] 
**LineFour** | **string** | LineFour | [optional] 
**LineFive** | **string** | LineFive | [optional] 
**LineSix** | **string** | LineSix | [optional] 

## Cover Payment Information

### CurrencyInstructedAmount

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SwiftFieldTag** | **string** | SwiftFieldTag | [optional] 
**Amount** | **string** | Amount | [optional]

### OrderingCustomer

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SwiftFieldTag** | **string** | SwiftFieldTag | [optional] 
**SwiftLineOne** | **string** | SwiftLineOne | [optional] 
**SwiftLineTwo** | **string** | SwiftLineTwo | [optional] 
**SwiftLineThree** | **string** | SwiftLineThree | [optional] 
**SwiftLineFour** | **string** | SwiftLineFour | [optional] 
**SwiftLineFive** | **string** | SwiftLineFive | [optional]

### OrderingInstitution

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SwiftFieldTag** | **string** | SwiftFieldTag | [optional] 
**SwiftLineOne** | **string** | SwiftLineOne | [optional] 
**SwiftLineTwo** | **string** | SwiftLineTwo | [optional] 
**SwiftLineThree** | **string** | SwiftLineThree | [optional] 
**SwiftLineFour** | **string** | SwiftLineFour | [optional] 
**SwiftLineFive** | **string** | SwiftLineFive | [optional]

### IntermediaryInstitution

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SwiftFieldTag** | **string** | SwiftFieldTag | [optional] 
**SwiftLineOne** | **string** | SwiftLineOne | [optional] 
**SwiftLineTwo** | **string** | SwiftLineTwo | [optional] 
**SwiftLineThree** | **string** | SwiftLineThree | [optional] 
**SwiftLineFour** | **string** | SwiftLineFour | [optional] 
**SwiftLineFive** | **string** | SwiftLineFive | [optional]

### InstitutionAccount

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SwiftFieldTag** | **string** | SwiftFieldTag | [optional] 
**SwiftLineOne** | **string** | SwiftLineOne | [optional] 
**SwiftLineTwo** | **string** | SwiftLineTwo | [optional] 
**SwiftLineThree** | **string** | SwiftLineThree | [optional] 
**SwiftLineFour** | **string** | SwiftLineFour | [optional] 
**SwiftLineFive** | **string** | SwiftLineFive | [optional]

### BeneficiaryCustomer

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SwiftFieldTag** | **string** | SwiftFieldTag | [optional] 
**SwiftLineOne** | **string** | SwiftLineOne | [optional] 
**SwiftLineTwo** | **string** | SwiftLineTwo | [optional] 
**SwiftLineThree** | **string** | SwiftLineThree | [optional] 
**SwiftLineFour** | **string** | SwiftLineFour | [optional] 
**SwiftLineFive** | **string** | SwiftLineFive | [optional]

### Remittance

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SwiftFieldTag** | **string** | SwiftFieldTag | [optional] 
**SwiftLineOne** | **string** | SwiftLineOne | [optional] 
**SwiftLineTwo** | **string** | SwiftLineTwo | [optional] 
**SwiftLineThree** | **string** | SwiftLineThree | [optional] 
**SwiftLineFour** | **string** | SwiftLineFour | [optional]

### Sender to Receiver

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SwiftFieldTag** | **string** | SwiftFieldTag | [optional] 
**SwiftLineOne** | **string** | SwiftLineOne | [optional] 
**SwiftLineTwo** | **string** | SwiftLineTwo | [optional] 
**SwiftLineThree** | **string** | SwiftLineThree | [optional] 
**SwiftLineFour** | **string** | SwiftLineFour | [optional] 
**SwiftLineFive** | **string** | SwiftLineFive | [optional] 
**SwiftLineSix** | **string** | SwiftLineSix | [optional]

###  Unstructured Addenda

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AddendaLength** | **string** | AddendaLength  Addenda Length must be numeric, padded with leading zeros if less than four characters and must equal length of content in Addenda Information (e.g., if content of Addenda Information is 987 characters, Addenda Length must be 0987).  | [optional] 
**Addenda** | **string** | Addenda | [optional]

### RelatedRemittance

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RemittanceIdentification** | **string** | RemittanceIdentification | [optional] 
**RemittanceLocationMethod** | **string** | RemittanceLocationMethod  * &#x60;EDIC&#x60; - Electronic Data Interchange * &#x60;EMAL&#x60; - Email * &#x60;FAXI&#x60; - Fax * &#x60;POST&#x60; - Postal services * &#x60;SMS&#x60; - Short Message Service (text) * &#x60;URI&#x60; - Uniform Resource Identifier  | [optional] 
**RemittanceLocationElectronicAddress** | **string** | RemittanceLocationElectronicAddress (E-mail or URL address) | [optional] 
**RemittanceData** | [**RemittanceData**](remittanceData.md) |  | [optional] 

## Structured Remittance Information

#### RemittanceOriginator

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IdentificationType** | **string** | IdentificationType  * &#x60;OI&#x60; - Organization ID * &#x60;PI&#x60; - Private ID  | [optional] 
**IdentificationCode** | **string** | IdentificationCode  Organization Identification Codes  * &#x60;BANK&#x60; - Bank Party Identification * &#x60;CUST&#x60; - Customer Number * &#x60;DUNS&#x60; - Data Universal Number System (Dun &amp; Bradstreet) * &#x60;EMPL&#x60; - Employer Identification Number * &#x60;GS1G&#x60; - Global Location Number * &#x60;PROP&#x60; - Proprietary Identification Number * &#x60;SWBB&#x60; - SWIFT BIC or BEI * &#x60;TXID&#x60; - Tax Identification Number  Private Identification Codes  * &#x60;ARNU&#x60; - Alien Registration Number * &#x60;CCPT&#x60; - Passport Number * &#x60;CUST&#x60; - Customer Number * &#x60;DPOB&#x60; - Date &amp; Place of Birth * &#x60;DRLC&#x60; - Driver’s License Number * &#x60;EMPL&#x60; - Employee Identification Number * &#x60;NIDN&#x60; - National Identity Number * &#x60;PROP&#x60; - Proprietary Identification Number * &#x60;SOSE&#x60; - Social Security Number * &#x60;TXID&#x60; - Tax Identification Number  | [optional] 
**IdentificationNumber** | **string** | IdentificationNumber | [optional] 
**IdentificationNumberIssuer** | **string** | IdentificationNumberIssuer | [optional] 
**DateBirthPlace** | **string** | DateBirthPlace | [optional] 
**RemittanceData** | [**RemittanceData**](remittanceData.md) |  | [optional] 
**CountryOfResidence** | **string** | CountryOfResidence | [optional] 
**ContactName** | **string** | ContactName | [optional] 
**ContactPhoneNumber** | **string** | ContactPhoneNumber | [optional] 
**ContactMobileNumber** | **string** | ContactMobileNumber | [optional] 
**ContactFaxNumber** | **string** | ContactFaxNumber | [optional] 
**ContactElectronicAddress** | **string** | ContactElectronicAddress ( i.e., E-mail or URL address) | [optional] 
**ContactOther** | **string** | ContactOther | [optional] 

### RemittanceBeneficiary

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IdentificationType** | **string** | IdentificationType  * &#x60;OI&#x60; - Organization ID * &#x60;PI&#x60; - Private ID  | [optional] 
**IdentificationCode** | **string** | IdentificationCode  Organization Identification Codes  * &#x60;BANK&#x60; - Bank Party Identification * &#x60;CUST&#x60; - Customer Number * &#x60;DUNS&#x60; - Data Universal Number System (Dun &amp; Bradstreet) * &#x60;EMPL&#x60; - Employer Identification Number * &#x60;GS1G&#x60; - Global Location Number * &#x60;PROP&#x60; - Proprietary Identification Number * &#x60;SWBB&#x60; - SWIFT BIC or BEI * &#x60;TXID&#x60; - Tax Identification Number  Private Identification Codes  * &#x60;ARNU&#x60; - Alien Registration Number * &#x60;CCPT&#x60; - Passport Number * &#x60;CUST&#x60; - Customer Number * &#x60;DPOB&#x60; - Date &amp; Place of Birth * &#x60;DRLC&#x60; - Driver’s License Number * &#x60;EMPL&#x60; - Employee Identification Number * &#x60;NIDN&#x60; - National Identity Number * &#x60;PROP&#x60; - Proprietary Identification Number * &#x60;SOSE&#x60; - Social Security Number * &#x60;TXID&#x60; - Tax Identification Number  | [optional] 
**IdentificationNumber** | **string** | IdentificationNumber | [optional] 
**IdentificationNumberIssuer** | **string** | IdentificationNumberIssuer | [optional] 
**DateBirthPlace** | **string** | DateBirthPlace | [optional] 
**RemittanceData** | [**RemittanceData**](remittanceData.md) |  | [optional]

### PrimaryRemittanceDocument

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DocumentTypeCode** | **string** | DocumentTypeCode  * &#x60;AROI&#x60; - Accounts Receivable Open Item * &#x60;BOLD&#x60; - Bill of Lading Shipping Notice * &#x60;CINV&#x60; - Commercial Invoice * &#x60;CMCN&#x60; - Commercial Contract * &#x60;CNFA&#x60; - Credit Note Related to Financial Adjustment * &#x60;CREN&#x60; - Credit Note * &#x60;DEBN&#x60; - Debit Note * &#x60;DISP&#x60; - Dispatch Advice * &#x60;DNFA&#x60; - Debit Note Related to Financial Adjustment HIRI Hire Invoice * &#x60;MSIN&#x60; - Metered Service Invoice * &#x60;PROP&#x60; - Proprietary Document Type * &#x60;PUOR&#x60; - Purchase Order * &#x60;SBIN&#x60; - Self Billed Invoice * &#x60;SOAC&#x60; - Statement of Account * &#x60;TSUT&#x60; - Trade Services Utility Transaction VCHR Voucher  | [optional] 
**ProprietaryDocumentTypeCode** | **string** | ProprietaryDocumentTypeCode | [optional] 
**DocumentIdentificationNumber** | **string** | DocumentIdentificationNumber | [optional] 
**Issuer** | **string** | Issuer | [optional]

### ActualAmountPaid

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CurrencyCode** | **string** | CurrencyCode | [optional] 
**Amount** | **string** | Amount Must contain at least one numeric character and only one decimal period marker (e.g., $1,234.56 should be entered as 1234.56). Can have up to 5 numeric characters following the decimal period marker (e.g., 1234.56789). Amount must be greater than zero (i.e., at least .01).  | [optional] 

### GrossAmountOfRemittanceDocument

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CurrencyCode** | **string** | CurrencyCode | [optional] 
**Amount** | **string** | Amount Must contain at least one numeric character and only one decimal period marker (e.g., $1,234.56 should be entered as 1234.56). Can have up to 5 numeric characters following the decimal period marker (e.g., 1234.56789). Amount must be greater than zero (i.e., at least .01).  | [optional] 
 
### AmountNegotiatedDiscount

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CurrencyCode** | **string** | CurrencyCode | [optional] 
**Amount** | **string** | Amount Must contain at least one numeric character and only one decimal period marker (e.g., $1,234.56 should be entered as 1234.56). Can have up to 5 numeric characters following the decimal period marker (e.g., 1234.56789). Amount must be greater than zero (i.e., at least .01).  | [optional] 
 
### Adjustment
 
 Name | Type | Description | Notes
 ------------ | ------------- | ------------- | -------------
 **AdjustmentReasonCode** | **string** | Adjustment  * &#x60;01&#x60; - Pricing Error * &#x60;03&#x60; - Extension Error * &#x60;04&#x60; - Item Not Accepted (Damaged) * &#x60;05&#x60; - Item Not Accepted (Quality) * &#x60;06&#x60; - Quantity Contested 07   Incorrect Product * &#x60;11&#x60; - Returns (Damaged) * &#x60;12&#x60; - Returns (Quality) * &#x60;59&#x60; - Item Not Received * &#x60;75&#x60; - Total Order Not Received * &#x60;81&#x60; - Credit as Agreed * &#x60;CM&#x60; - Covered by Credit Memo  | [optional] 
 **CreditDebitIndicator** | **string** | CreditDebitIndicator  * &#x60;CRDT&#x60; - Credit * &#x60;DBIT&#x60; - Debit  | [optional] 
 **CurrencyCode** | **string** | CurrencyCode | [optional] 
 **Amount** | **string** | Amount Must contain at least one numeric character and only one decimal period marker (e.g., $1,234.56 should be entered as 1234.56). Can have up to 5 numeric characters following the decimal period marker (e.g., 1234.56789). Amount must be greater than zero (i.e., at least .01).  | [optional] 
 **AdditionalInfo** | **string** | AdditionalInfo | [optional]
 
### DateRemittanceDocument

 Name | Type | Description | Notes
 ------------ | ------------- | ------------- | -------------
 **DateRemittanceDocument** | **string** | DateRemittanceDocument CCYYMMDD | [optional]
 
### SecondaryRemittanceDocument
 
 Name | Type | Description | Notes
 ------------ | ------------- | ------------- | -------------
 **DocumentTypeCode** | **string** | SecondaryRemittanceDocument  * &#x60;AROI&#x60; - Accounts Receivable Open Item * &#x60;DISP&#x60; - Dispatch Advice * &#x60;FXDR&#x60; - Foreign Exchange Deal Reference * &#x60;PROP&#x60; - Proprietary Document Type PUOR Purchase Order * &#x60;RADM&#x60; - Remittance Advice Message * &#x60;RPIN&#x60; - Related Payment Instruction * &#x60;SCOR1&#x60; - Structured Communication Reference VCHR Voucher  | [optional] 
 **ProprietaryDocumentTypeCode** | **string** | proprietaryDocumentTypeCode | [optional] 
 **DocumentIdentificationNumber** | **string** | documentIdentificationNumber | [optional] 
 **Issuer** | **string** | Issuer | [optional]
 
### RemittanceFreeText

 Name | Type | Description | Notes
 ------------ | ------------- | ------------- | -------------
 **LineOne** | **string** | LineOne | [optional] 
 **LineTwo** | **string** | LineTwo | [optional] 
 **LineThree** | **string** | LineThree | [optional]
 
 ## Information Appended by the FEDWire Funds Service
 
### MessageDisposition
 
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FormatVersion** | **string** | FormatVersion 30  | [optional] 
**TestProductionCode** | **string** | TestProductionCode  * &#x60;T&#x60; - Test * &#x60;P&#x60; - Production  | [optional] 
**MessageDuplicationCode** | **string** | MessageDuplicationCode  * &#x60; &#x60; - Original Message * &#x60;R&#x60; - Retrieval of an original message * &#x60;P&#x60; - Resend  | [optional] 
**MessageStatusIndicator** | **string** | Message Status Indicator  Outgoing Messages * &#x60;0&#x60; - In process or Intercepted * &#x60;2&#x60; - Successful with Accounting (Value) * &#x60;3&#x60; - Rejected due to Error Condition * &#x60;7&#x60; - Successful without Accounting (Non-Value)  Incoming Messages * &#x60;N&#x60; - Successful with Accounting (Value) * &#x60;S&#x60; - Successful without Accounting (Non-Value)  | [optional] 

### ReceiptTimeStamp
 
 Name | Type | Description | Notes
 ------------ | ------------- | ------------- | -------------
 **ReceiptDate** | [**Date**](DateMMDD.md) | Date | [optional] 
 **ReceiptTime** | [**Time**](TimeHHMM.md) | Time | [optional] 
 **ReceiptApplicationIdentification** | **string** | ApplicationIdentification | [optional]
 
### OutputMessageAccountabilityData

 Name | Type | Description | Notes
 ------------ | ------------- | ------------- | -------------
 **OutputCycleDate** | **string** | OutputCycleDate (CCYYMMDD) | [optional] 
 **OutputDestinationID** | **string** | OutputDestinationID | [optional] 
 **OutputSequenceNumber** | **string** | outputSequenceNumber | [optional] 
 **OutputDate** | [**Date**](DateMMDD.md) | Date | [optional] 
 **OutputTime** | [**Time**](TimeHHMM.md) | Time | [optional] 
 **OutputFRBApplicationIdentification** | **string** | OutputFRBApplicationIdentification | [optional]  

### ErrorWire

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ErrorCategory** | **string** |  * &#x60;E&#x60; - Data Error * &#x60;F&#x60; - Insufficient Balance * &#x60;H&#x60; - Accountability Error * &#x60;I&#x60; - In Process or Intercepted * &#x60;W&#x60; - Cutoff Hour Error * &#x60;X&#x60; - Duplicate IMAD  | [optional] 
**ErrorCode** | **string** | ErrorCode | [optional] 
**ErrorDescription** | **string** | ErrorDescription | [optional]                       