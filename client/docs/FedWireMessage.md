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
**SenderReference** | **string** | SenderReference | [optional] 
**PreviousMessageIdentifier** | **string** | PreviousMessageIdentifier | [optional] 
**LocalInstrument** | [**LocalInstrument**](LocalInstrument.md) |  | [optional] 
**PaymentNotification** | [**PaymentNotification**](PaymentNotification.md) |  | [optional] 
**Charges** | [**Charges**](Charges.md) |  | [optional] 
**InstructedAmount** | [**InstructedAmount**](InstructedAmount.md) |  | [optional] 
**ExchangeRate** | **string** | ExchangeRate  Must contain at least one numeric character and only one decimal comma marker (e.g., an exchange rate of 1.2345 should be entered as 1,2345).  | [optional] 
**BeneficiaryIntermediaryFI** | [**FinancialInstitution**](FinancialInstitution.md) |  | [optional] 
**BeneficiaryFI** | [**FinancialInstitution**](FinancialInstitution.md) |  | [optional] 
**Beneficiary** | [**Personal**](Personal.md) |  | [optional] 
**ReferenceForBeneficiary** | **string** | ReferenceForBeneficiary {4320} | [optional] 
**AccountDebitedDrawdown** | [**AccountDebitedDrawdown**](AccountDebitedDrawdown.md) |  | [optional] 
**Originator** | [**Personal**](Personal.md) |  | [optional] 
**OriginatorOptionF** | [**map[string]interface{}**](map[string]interface{}.md) | OriginatorOptionF {5010} | [optional] 
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
**Adjustment** | [**AdjustmentEnum**](AdjustmentEnum.md) |  | [optional] 
**DateRemittanceDocument** | [**DateRemittanceDocument**](DateRemittanceDocument.md) |  | [optional] 
**SecondaryRemittanceDocument** | [**SecondaryRemittanceDocumentEnum**](SecondaryRemittanceDocumentEnum.md) |  | [optional] 
**RemittanceFreeText** | [**RemittanceFreeText**](RemittanceFreeText.md) |  | [optional] 
**ServiceMessage** | [**ServiceMessage**](ServiceMessage.md) |  | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


