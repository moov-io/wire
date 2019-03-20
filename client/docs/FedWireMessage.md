# FedWireMessage

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MessageDisposition** | [**MessageDisposition**](MessageDisposition.md) |  | [optional] 
**ReceiptTimeStamp** | [**ReceiptTimeStamp**](ReceiptTimeStamp.md) |  | [optional] 
**OutputMessageAccountabilityData** | [**OutputMessageAccountabilityData**](OutputMessageAccountabilityData.md) |  | [optional] 
**Error** | [**FedWireError**](FEDWireError.md) |  | [optional] 
**SenderSupplied** | [**SenderSupplied**](SenderSupplied.md) |  | 
**TypeSubType** | [**TypeSubType**](TypeSubType.md) |  | 
**InputMessageAccountabilityData** | [**InputMessageAccountabilityData**](InputMessageAccountabilityData.md) |  | 
**Amount** | [**Amount**](Amount.md) |  | 
**SenderDepositoryInstitution** | [**SenderDepositoryInstitution**](SenderDepositoryInstitution.md) |  | 
**ReceiverDepositoryInstitution** | [**ReceiverDepositoryInstitution**](ReceiverDepositoryInstitution.md) |  | 
**BusinessFunctionCode** | [**BusinessFunctionCode**](BusinessFunctionCode.md) |  | 
**SenderReference** | [**SenderReference**](SenderReference.md) |  | [optional] 
**PreviousMessageIdentifier** | [**PreviousMessageIdentifier**](PreviousMessageIdentifier.md) |  | [optional] 
**LocalInstrument** | [**LocalInstrument**](LocalInstrument.md) |  | [optional] 
**PaymentNotification** | [**PaymentNotification**](PaymentNotification.md) |  | [optional] 
**Charges** | [**Charges**](Charges.md) |  | [optional] 
**InstructedAmount** | [**InstructedAmount**](InstructedAmount.md) |  | [optional] 
**ExchangeRate** | **string** | ExchangeRate  Must contain at least one numeric character and only one decimal comma marker (e.g., an exchange rate of 1.2345 should be entered as 1,2345).  | [optional] 
**BeneficiaryIntermediaryFI** | [**FinancialInstitution**](FinancialInstitution.md) |  | [optional] 
**BeneficiaryFI** | [**FinancialInstitution**](FinancialInstitution.md) |  | [optional] 
**Beneficiary** | [**Personal**](Personal.md) |  | [optional] 
**BeneficiaryReference** | [**BeneficiaryReference**](BeneficiaryReference.md) |  | [optional] 
**AccountDebitedDrawdown** | [**AccountDebitedDrawdown**](AccountDebitedDrawdown.md) |  | [optional] 
**Originator** | [**Personal**](Personal.md) |  | [optional] 
**OriginatorOptionF** | [**OriginatorOptionF**](OriginatorOptionF.md) |  | [optional] 
**OriginatorFI** | [**FinancialInstitution**](FinancialInstitution.md) |  | [optional] 
**InstructingFI** | [**FinancialInstitution**](FinancialInstitution.md) |  | [optional] 
**AccountCreditedDrawdown** | [**AccountCreditedDrawdown**](AccountCreditedDrawdown.md) |  | [optional] 
**OriginatorToBeneficiary** | [**OriginatorToBeneficiary**](OriginatorToBeneficiary.md) |  | [optional] 
**ReceiverFI** | [**FiToFi**](FIToFI.md) |  | [optional] 
**DrawdownDebitAccountAdvice** | [**Advice**](Advice.md) |  | [optional] 
**IntermediaryFI** | [**FiToFi**](FIToFI.md) |  | [optional] 
**IntermediaryFinacialInstitutionAdvice** | [**Advice**](Advice.md) |  | [optional] 
**OriginatorBeneficiaryFinancialInstitution** | [**FiToFi**](FIToFI.md) |  | [optional] 
**OriginatorBeneficiaryFinancialInstitutionAdvice** | [**Advice**](Advice.md) |  | [optional] 
**OriginatorBeneficiary** | [**FiToFi**](FIToFI.md) |  | [optional] 
**BeneficiaryAdvice** | [**Advice**](Advice.md) |  | [optional] 
**PaymentMethodToBeneficiary** | [**PaymentMethodToBeneficiary**](PaymentMethodToBeneficiary.md) |  | [optional] 
**AdditionalFIToFI** | [**AdditionalFiToFi**](AdditionalFIToFI.md) |  | [optional] 
**CurrencyInstructedAmount** | [**CoverPayment**](CoverPayment.md) |  | [optional] 
**OrderingCustomer** | [**CoverPayment**](CoverPayment.md) |  | [optional] 
**OrderingInstitution** | [**CoverPayment**](CoverPayment.md) |  | [optional] 
**IntermediaryInstitution** | [**CoverPayment**](CoverPayment.md) |  | [optional] 
**InstitutionAccount** | [**CoverPayment**](CoverPayment.md) |  | [optional] 
**BeneficiaryCustomer** | [**CoverPayment**](CoverPayment.md) |  | [optional] 
**Remittance** | [**CoverPayment**](CoverPayment.md) |  | [optional] 
**SenderToReceiver** | [**CoverPayment**](CoverPayment.md) |  | [optional] 
**UnstructuredAddenda** | [**UnstructuredAddenda**](UnstructuredAddenda.md) |  | [optional] 
**RelatedRemittance** | [**RelatedRemittance**](RelatedRemittance.md) |  | [optional] 
**RemittanceOriginator** | [**RemittanceOriginator**](RemittanceOriginator.md) |  | [optional] 
**RemittanceBeneficiary** | [**RemittanceBeneficiary**](RemittanceBeneficiary.md) |  | [optional] 
**PrimaryRemittanceDocument** | [**PrimaryRemittanceDocument**](PrimaryRemittanceDocument.md) |  | [optional] 
**ActualAmountPaid** | [**RemittanceAmount**](RemittanceAmount.md) |  | [optional] 
**GrossAmountRemittanceDocument** | [**RemittanceAmount**](RemittanceAmount.md) |  | [optional] 
**AmountNegotiatedDiscount** | [**RemittanceAmount**](RemittanceAmount.md) |  | [optional] 
**Adjustment** | **string** | Adjustment  * &#x60;01&#x60; - Pricing Error * &#x60;03&#x60; - Extension Error * &#x60;04&#x60; - Item Not Accepted (Damaged) * &#x60;05&#x60; - Item Not Accepted (Quality) * &#x60;06&#x60; - Quantity Contested 07   Incorrect Product * &#x60;11&#x60; - Returns (Damaged) * &#x60;12&#x60; - Returns (Quality) * &#x60;59&#x60; - Item Not Received * &#x60;75&#x60; - Total Order Not Received * &#x60;81&#x60; - Credit as Agreed * &#x60;CM&#x60; - Covered by Credit Memo  | [optional] 
**DateRemittanceDocument** | [**DateRemittanceDocument**](DateRemittanceDocument.md) |  | [optional] 
**SecondaryRemittanceDocument** | **string** | SecondaryRemittanceDocument  * &#x60;AROI&#x60; - Accounts Receivable Open Item * &#x60;DISP&#x60; - Dispatch Advice * &#x60;FXDR&#x60; - Foreign Exchange Deal Reference * &#x60;PROP&#x60; - Proprietary Document Type PUOR Purchase Order * &#x60;RADM&#x60; - Remittance Advice Message * &#x60;RPIN&#x60; - Related Payment Instruction * &#x60;SCOR1&#x60; - Structured Communication Reference VCHR Voucher  | [optional] 
**RemittanceFreeText** | [**RemittanceFreeText**](RemittanceFreeText.md) |  | [optional] 
**ServiceMessage** | [**ServiceMessage**](ServiceMessage.md) |  | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


