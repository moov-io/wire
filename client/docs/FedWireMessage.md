# FedWireMessage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ID** | **string** | Fedwire Message ID | [optional] 
**MessageDisposition** | [**MessageDisposition**](MessageDisposition.md) |  | [optional] 
**ReceiptTimeStamp** | [**ReceiptTimeStamp**](ReceiptTimeStamp.md) |  | [optional] 
**OutputMessageAccountabilityData** | [**OutputMessageAccountabilityData**](OutputMessageAccountabilityData.md) |  | [optional] 
**ErrorWire** | [**ErrorWire**](ErrorWire.md) |  | [optional] 
**SenderSupplied** | [**SenderSupplied**](SenderSupplied.md) |  | 
**TypeSubType** | [**TypeSubType**](TypeSubType.md) |  | 
**InputMessageAccountabilityData** | [**InputMessageAccountabilityData**](InputMessageAccountabilityData.md) |  | 
**Amount** | [**WireAmount**](WireAmount.md) |  | 
**SenderDepositoryInstitution** | [**SenderDepositoryInstitution**](SenderDepositoryInstitution.md) |  | 
**ReceiverDepositoryInstitution** | [**ReceiverDepositoryInstitution**](ReceiverDepositoryInstitution.md) |  | 
**BusinessFunctionCode** | [**BusinessFunctionCode**](BusinessFunctionCode.md) |  | 
**SenderReference** | [**SenderReference**](SenderReference.md) |  | [optional] 
**PreviousMessageIdentifier** | [**PreviousMessageIdentifier**](PreviousMessageIdentifier.md) |  | [optional] 
**LocalInstrument** | [**LocalInstrument**](LocalInstrument.md) |  | [optional] 
**PaymentNotification** | [**PaymentNotification**](PaymentNotification.md) |  | [optional] 
**Charges** | [**Charges**](Charges.md) |  | [optional] 
**InstructedAmount** | [**InstructedAmount**](InstructedAmount.md) |  | [optional] 
**ExchangeRate** | [**ExchangeRate**](ExchangeRate.md) |  | [optional] 
**BeneficiaryIntermediaryFI** | [**FinancialInstitution**](FinancialInstitution.md) |  | [optional] 
**BeneficiaryFI** | [**FinancialInstitution**](FinancialInstitution.md) |  | [optional] 
**Beneficiary** | [**Beneficiary**](Beneficiary.md) |  | [optional] 
**BeneficiaryReference** | [**BeneficiaryReference**](BeneficiaryReference.md) |  | [optional] 
**AccountDebitedDrawdown** | [**AccountDebitedDrawdown**](AccountDebitedDrawdown.md) |  | [optional] 
**Originator** | [**Beneficiary**](Beneficiary.md) |  | [optional] 
**OriginatorOptionF** | [**OriginatorOptionF**](OriginatorOptionF.md) |  | [optional] 
**OriginatorFI** | [**FinancialInstitution**](FinancialInstitution.md) |  | [optional] 
**InstructingFI** | [**FinancialInstitution**](FinancialInstitution.md) |  | [optional] 
**AccountCreditedDrawdown** | [**AccountCreditedDrawdown**](AccountCreditedDrawdown.md) |  | [optional] 
**OriginatorToBeneficiary** | [**OriginatorToBeneficiary**](OriginatorToBeneficiary.md) |  | [optional] 
**FiReceiverFI** | [**FiToFi**](FIToFI.md) |  | [optional] 
**FiDrawdownDebitAccountAdvice** | [**Advice**](Advice.md) |  | [optional] 
**FiIntermediaryFI** | [**FiToFi**](FIToFI.md) |  | [optional] 
**FiIntermediaryFIAdvice** | [**Advice**](Advice.md) |  | [optional] 
**FiBeneficiaryFI** | [**FiToFi**](FIToFI.md) |  | [optional] 
**FiBeneficiaryFIAdvice** | [**Advice**](Advice.md) |  | [optional] 
**FiBeneficiary** | [**FiToFi**](FIToFI.md) |  | [optional] 
**FiBeneficiaryAdvice** | [**Advice**](Advice.md) |  | [optional] 
**FiPaymentMethodToBeneficiary** | [**FiPaymentMethodToBeneficiary**](FIPaymentMethodToBeneficiary.md) |  | [optional] 
**FiAdditionalFIToFI** | [**AdditionalFiToFi**](AdditionalFIToFI.md) |  | [optional] 
**CurrencyInstructedAmount** | [**CurrencyInstructedAmount**](CurrencyInstructedAmount.md) |  | [optional] 
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
**Adjustment** | [**Adjustment**](Adjustment.md) |  | [optional] 
**DateRemittanceDocument** | [**DateRemittanceDocument**](DateRemittanceDocument.md) |  | [optional] 
**SecondaryRemittanceDocument** | [**SecondaryRemittanceDocument**](SecondaryRemittanceDocument.md) |  | [optional] 
**RemittanceFreeText** | [**RemittanceFreeText**](RemittanceFreeText.md) |  | [optional] 
**ServiceMessage** | [**ServiceMessage**](ServiceMessage.md) |  | [optional] 
**ValidateOptions** | Pointer to [**ValidateOptions**](ValidateOptions.md) |  | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


