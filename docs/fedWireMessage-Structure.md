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
**TestProductionCode** | **string** | TestProductionCode <li>`T` - Test<li>`P` - Production  |
**MessageDuplicationCode** | **string** | MessageDuplicationCode <li>` ` - Original Message<li>`R` - Retrieval of an original message<li>`P` - Resend  | 

### TypeSubType

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TypeCode** | **string** | TypeCode: <li>`10` - Funds Transfer - A funds transfer in which the sender and/or receiver may be a bank or a third party (i.e., customer of a bank).<li>`15` - Foreign Transfer - A funds transfer to or from a foreign central bank or government or international organization with an account at the Federal Reserve Bank of New York.<li>`16` - Settlement Transfer - A funds transfer between Fedwire Funds Service participants.  | 
**SubTypeCode** | **string** | SubTypeCode: <li>`00` - Basic Funds Transfer - A basic value funds transfer.<li>`01` - Request for Reversal - A non-value request for reversal of a funds transfer originated on the current business day.<li>`02` - Reversal of Transfer - A value reversal of a funds transfer received on the current business day.  May be used in response to a subtype code ‘01’ Request for Reversal.<li>`07` - Request for Reversal of a Prior Day Transfer - A non-value request for a reversal of a funds transfer originated on a prior business day.<li>`08` - Reversal of a Prior Day Transfer - A value reversal of a funds transfer received on a prior business day.  May be used in response to a subtype code ‘07’ Request for Reversal of a Prior Day Transfer.<li>`31` - Request for Credit (Drawdown) - A non-value request for the receiver to send a funds transfer to a designated party.<li>`32` - Funds Transfer Honoring a Request for Credit (Drawdown) -  A value funds transfer honoring a subtype 31 request for credit.<li>`33` -Refusal to Honor a Request for Credit (Drawdown) - A non-value message indicating refusal to honor a subtype 31 request for credit.<li>`90` - Service Message - A non-value message used to communicate questions and information that is not covered by a specific subtype.  | 

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
**BusinessFunctionCode** | **string** | BusinessFunctionCode<li>`BTR` - Bank Transfer (Beneficiary is a bank)<li>`DRC` - Customer or Corporate Drawdown Request<li>`CKS` - Check Same Day Settlement<li>`DRW` - Drawdown Payment<li>`CTP` - Customer Transfer Plus<li>`FFR` - Fed Funds Returned<li>`CTR` - Customer Transfer (Beneficiary is a not a bank)<li>`FFS` - Fed Funds Sold<li>`DEP` - Deposit to Sender’s Account<li>`SVC` - Service Message<li>`DRB` - Bank-to-Bank Drawdown Request  | 
**TransactionTypeCode** | **string** | TransactionTypeCode If {3600} is CTR, an optional Transaction Type Code element is permitted; however, the Transaction Type Code &#39;COV&#39; is not permitted. | [optional] 

## Other Transfer Information

### SenderReference

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SenderReference** | **string** | SenderReference | [optional]

### LocalInstrument

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LocalInstrumentCode** | **string** | LocalInstrument <li>`ANSI` - ANSI X12 format<li>`COVS` - Sequence B Cover Payment Structured<li>`GXML` - General XML format<li>`IXML` - ISO 20022 XML formaT<li>`NARR` - Narrative Text<li>`PROP` - Proprietary Local Instrument Code<li>`RMTS` - Remittance Information Structured<li>`RRMT` - Related Remittance Information<li>`S820` - STP 820 format<li>`SWIF` - SWIFT field 70 (Remittance Information)<li>`UEDI` - UN/EDIFACT format  | [optional] 
**ProprietaryCode** | **string** | ProprietaryCode | [optional] 

### PaymentNotification

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PaymentNotificationIndicator** | **string** | PaymentNotificationIndicator <li>`0 - 6` - Reserved for market practice conventions.<li>`7 - 9` - Reserved for bilateral agreements between Fedwire senders and receivers.  | [optional] 
**ContactNotificationElectronicAddress** | **string** | ContactNotificationElectronicAddress | [optional] 
**ContactName** | **string** | ContactName | [optional] 
**ContactPhoneNumber** | **string** | ContactPhoneNumber | [optional] 
**ContactMobileNumber** | **string** | ContactMobileNumber | [optional] 
**FaxNumber** | **string** | FaxNumber | [optional] 
**EndToEndIdentification** | **string** | EndToEndIdentification | [optional] 

### Charges

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ChargeDetails** | **string** | ChargeDetails<li>`B` - Beneficiary<li>`S` - Shared  | [optional] 
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
**IdentificationCode** | **string** | IdentificationCode: <li>`B` - SWIFT Bank Identifier Code (BIC)<li>`C` - CHIPS Participant<li>`D` - Demand Deposit Account (DDA) Number<li>`F` - Fed Routing Number<li>`T` - SWIFT BIC or Bank Entity Identifier (BEI) and Account Number<li>`U` - CHIPS Identifier  | 
**Identifier** | **string** | Identifier | 
**Name** | **string** | Name | 
**Address** | [**Address**](address.md) |  | 

### BeneficiaryFI

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IdentificationCode** | **string** | IdentificationCode: <li>`B` - SWIFT Bank Identifier Code (BIC)<li>`C` - CHIPS Participant<li>`D` - Demand Deposit Account (DDA) Number<li>`F` - Fed Routing Number<li>`T` - SWIFT BIC or Bank Entity Identifier (BEI) and Account Number<li>`U` - CHIPS Identifier  | 
**Identifier** | **string** | Identifier | 
**Name** | **string** | Name | 
**Address** | [**Address**](address.md) |  | 

### Beneficiary

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IdentificationCode** | **string** | IdentificationCode: <li>`1` - Passport Number<li>`2` - Tax Identification Number<li>`3` - Driver’s License Number<li>`4` - Alien Registration Number<li>`5` - Corporate Identification<li>`9` - Other Identification  | 
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
**IdentificationCode** | **string** | Identification Code<li>`D` - Debit  | 
**Identifier** | **string** | Identifier | 
**Name** | **string** | Name | 
**Address** | [**Address**](address.md) |  | [optional]

## Originator Information

### Originator

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IdentificationCode** | **string** | IdentificationCode: <li>`1` - Passport Number<li>`2` - Tax Identification Number<li>`3` - Driver’s License Number<li>`4` - Alien Registration Number<li>`5` - Corporate Identification<li>`9` - Other Identification  | 
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
 **IdentificationCode** | **string** | IdentificationCode: <li>`B` - SWIFT Bank Identifier Code (BIC)<li>`C` - CHIPS Participant<li>`D` - Demand Deposit Account (DDA) Number<li>`F` - Fed Routing Number<li>`T` - SWIFT BIC or Bank Entity Identifier (BEI) and Account Number<li>`U` - CHIPS Identifier  | 
 **Identifier** | **string** | Identifier | 
 **Name** | **string** | Name | 
 **Address** | [**Address**](address.md) |  
 
### InstructingFI
  
 Name | Type | Description | Notes
 ------------ | ------------- | ------------- | -------------
 **IdentificationCode** | **string** | IdentificationCode: <li>`B` - SWIFT Bank Identifier Code (BIC)<li>`C` - CHIPS Participant<li>`D` - Demand Deposit Account (DDA) Number<li>`F` - Fed Routing Number<li>`T` - SWIFT BIC or Bank Entity Identifier (BEI) and Account Number<li>`U` - CHIPS Identifier  | 
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
**AdviceCode** | **string** | AdviceCode <li>`HLD` - Hold<li>`LTR` - Letter<li>`PHN` - Phone<li>`TLX` - Telex<li>`WRE` - Wire | [optional] 
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
**AdviceCode** | **string** | AdviceCode <li>`HLD` - Hold<li>`LTR` - Letter<li>`PHN` - Phone<li>`TLX` - Telex<li>`WRE` - Wire | [optional] 
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
**AdviceCode** | **string** | AdviceCode <li>`HLD` - Hold<li>`LTR` - Letter<li>`PHN` - Phone<li>`TLX` - Telex<li>`WRE` - Wire | [optional] 
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
**AdviceCode** | **string** | AdviceCode <li>`HLD` - Hold<li>`LTR` - Letter<li>`PHN` - Phone<li>`TLX` - Telex<li>`WRE` - Wire | [optional] 
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
**RemittanceLocationMethod** | **string** | RemittanceLocationMethod <li>`EDIC` - Electronic Data Interchange<li>`EMAL` - Email<li>`FAXI` - Fax<li>`POST` - Postal services<li>`SMS` - Short Message Service (text)<li>`URI` - Uniform Resource Identifier  | [optional] 
**RemittanceLocationElectronicAddress** | **string** | RemittanceLocationElectronicAddress (E-mail or URL address) | [optional] 
**RemittanceData** | [**RemittanceData**](remittanceData.md) |  | [optional] 

## Structured Remittance Information

#### RemittanceOriginator

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IdentificationType** | **string** | IdentificationType <li>`OI` - Organization ID<li>`PI` - Private ID  | [optional] 
**IdentificationCode** | **string** | IdentificationCode  Organization Identification Codes <li>`BANK` - Bank Party Identification<li>`CUST` - Customer Number<li>`DUNS` - Data Universal Number System (Dun &amp; Bradstreet)<li>`EMPL` - Employer Identification Number<li>`GS1G` - Global Location Number<li>`PROP` - Proprietary Identification Number<li>`SWBB` - SWIFT BIC or BEI<li>`TXID` - Tax Identification Number  Private Identification Codes <li>`ARNU` - Alien Registration Number<li>`CCPT` - Passport Number<li>`CUST` - Customer Number<li>`DPOB` - Date &amp; Place of Birth<li>`DRLC` - Driver’s License Number<li>`EMPL` - Employee Identification Number<li>`NIDN` - National Identity Number<li>`PROP` - Proprietary Identification Number<li>`SOSE` - Social Security Number<li>`TXID` - Tax Identification Number  | [optional] 
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
**IdentificationType** | **string** | IdentificationType <li>`OI` - Organization ID<li>`PI` - Private ID  | [optional] 
**IdentificationCode** | **string** | IdentificationCode  Organization Identification Codes <li>`BANK` - Bank Party Identification<li>`CUST` - Customer Number<li>`DUNS` - Data Universal Number System (Dun &amp; Bradstreet)<li>`EMPL` - Employer Identification Number<li>`GS1G` - Global Location Number<li>`PROP` - Proprietary Identification Number<li>`SWBB` - SWIFT BIC or BEI<li>`TXID` - Tax Identification Number  Private Identification Codes <li>`ARNU` - Alien Registration Number<li>`CCPT` - Passport Number<li>`CUST` - Customer Number<li>`DPOB` - Date &amp; Place of Birth<li>`DRLC` - Driver’s License Number<li>`EMPL` - Employee Identification Number<li>`NIDN` - National Identity Number<li>`PROP` - Proprietary Identification Number<li>`SOSE` - Social Security Number<li>`TXID` - Tax Identification Number  | [optional] 
**IdentificationNumber** | **string** | IdentificationNumber | [optional] 
**IdentificationNumberIssuer** | **string** | IdentificationNumberIssuer | [optional] 
**DateBirthPlace** | **string** | DateBirthPlace | [optional] 
**RemittanceData** | [**RemittanceData**](remittanceData.md) |  | [optional]

### PrimaryRemittanceDocument

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DocumentTypeCode** | **string** | DocumentTypeCode <li>`AROI` - Accounts Receivable Open Item<li>`BOLD` - Bill of Lading Shipping Notice<li>`CINV` - Commercial Invoice<li>`CMCN` - Commercial Contract<li>`CNFA` - Credit Note Related to Financial Adjustment<li>`CREN` - Credit Note<li>`DEBN` - Debit Note<li>`DISP` - Dispatch Advice<li>`DNFA` - Debit Note Related to Financial Adjustment HIRI Hire Invoice<li>`MSIN` - Metered Service Invoice<li>`PROP` - Proprietary Document Type<li>`PUOR` - Purchase Order<li>`SBIN` - Self Billed Invoice<li>`SOAC` - Statement of Account<li>`TSUT` - Trade Services Utility Transaction VCHR Voucher  | [optional] 
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
 **AdjustmentReasonCode** | **string** | Adjustment <li>`01` - Pricing Error<li>`03` - Extension Error<li>`04` - Item Not Accepted (Damaged)<li>`05` - Item Not Accepted (Quality)<li>`06` - Quantity Contested 07   Incorrect Product<li>`11` - Returns (Damaged)<li>`12` - Returns (Quality)<li>`59` - Item Not Received<li>`75` - Total Order Not Received<li>`81` - Credit as Agreed<li>`CM` - Covered by Credit Memo  | [optional] 
 **CreditDebitIndicator** | **string** | CreditDebitIndicator <li>`CRDT` - Credit<li>`DBIT` - Debit  | [optional] 
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
 **DocumentTypeCode** | **string** | SecondaryRemittanceDocument <li>`AROI` - Accounts Receivable Open Item<li>`DISP` - Dispatch Advice<li>`FXDR` - Foreign Exchange Deal Reference<li>`PROP` - Proprietary Document Type PUOR Purchase Order<li>`RADM` - Remittance Advice Message<li>`RPIN` - Related Payment Instruction<li>`SCOR1` - Structured Communication Reference VCHR Voucher  | [optional] 
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
**TestProductionCode** | **string** | TestProductionCode <li>`T` - Test<li>`P` - Production  | [optional] 
**MessageDuplicationCode** | **string** | MessageDuplicationCode <li>` ` - Original Message<li>`R` - Retrieval of an original message<li>`P` - Resend  | [optional] 
**MessageStatusIndicator** | **string** | Message Status Indicator  Outgoing Messages<li>`0` - In process or Intercepted<li>`2` - Successful with Accounting (Value)<li>`3` - Rejected due to Error Condition<li>`7` - Successful without Accounting (Non-Value)  Incoming Messages<li>`N` - Successful with Accounting (Value)<li>`S` - Successful without Accounting (Non-Value)  | [optional] 

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
**ErrorCategory** | **string** | <li>`E` - Data Error<li>`F` - Insufficient Balance<li>`H` - Accountability Error<li>`I` - In Process or Intercepted<li>`W` - Cutoff Hour Error<li>`X` - Duplicate IMAD  | [optional] 
**ErrorCode** | **string** | ErrorCode | [optional] 
**ErrorDescription** | **string** | ErrorDescription | [optional]                       