# FedWireMessage

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SenderSuppliedInfo** | [**SenderSuppliedInfo**](SenderSuppliedInfo.md) |  | 
**TypeSubType** | [**TypeSubType**](TypeSubType.md) |  | 
**InputMessageAccountabilityData** | [**InputMessageAccountabilityData**](InputMessageAccountabilityData.md) |  | 
**Amount** | [**Amount**](Amount.md) |  | 
**SenderDepositoryInstitution** | [**SenderDepositoryInstitution**](SenderDepositoryInstitution.md) |  | 
**ReceiverDepositoryInstitution** | [**ReceiverDepositoryInstitution**](ReceiverDepositoryInstitution.md) |  | 
**BusinessFunctionCode** | [**BusinessFunctionCode**](BusinessFunctionCode.md) |  | 
**SenderReference** | **string** | SenderReference | [optional] 
**PreviousMessageIdentifier** | **string** | PreviousMessageIdentifier | [optional] 
**LocalInstrument** | [**LocalInstrument**](LocalInstrument.md) |  | [optional] 
**Charges** | [**Charges**](Charges.md) |  | [optional] 
**InstructedAmount** | [**InstructedAmount**](InstructedAmount.md) |  | [optional] 
**ExchangeRate** | **string** | ExchangeRate  Must contain at least one numeric character and only one decimal comma marker (e.g., an exchange rate of 1.2345 should be entered as 1,2345).  | [optional] 
**IntermediaryFinancialInstitution** | [**FinancialInstitution**](FinancialInstitution.md) |  | [optional] 
**BeneficiaryFinancialInstitution** | [**FinancialInstitution**](FinancialInstitution.md) |  | [optional] 
**Beneficiary** | [**Personal**](Personal.md) |  | [optional] 
**ReferenceForBeneficiary** | **string** | ReferenceForBeneficiary {4320} | [optional] 
**AccountDebitedDrawdown** | [**AccountDebitedDrawdown**](AccountDebitedDrawdown.md) |  | [optional] 
**Originator** | [**Personal**](Personal.md) |  | [optional] 
**OriginatorOptionF** | [**map[string]interface{}**](map[string]interface{}.md) | OriginatorOptionF {5010} | [optional] 
**OriginatorFinancialInstitution** | [**FinancialInstitution**](FinancialInstitution.md) |  | [optional] 
**InstructingFinancialInstitution** | [**FinancialInstitution**](FinancialInstitution.md) |  | [optional] 
**AccountCreditedDrawdown** | [**AccountCreditedDrawdown**](AccountCreditedDrawdown.md) |  | [optional] 
**OriginatorToBeneficiaryInfo** | [**OriginatorToBeneficiaryInfo**](OriginatorToBeneficiaryInfo.md) |  | [optional] 
**ReceiverFinancialInstitutionInfo** | [**FiToFiInfo**](FIToFIInfo.md) |  | [optional] 
**DrawdownDebitAccountAdviceInfo** | [**AdviceInfo**](AdviceInfo.md) |  | [optional] 
**IntermediaryFinancialInstitutionInfo** | [**FiToFiInfo**](FIToFIInfo.md) |  | [optional] 
**IntermediaryFinacialInstitutionAdviceInfo** | [**AdviceInfo**](AdviceInfo.md) |  | [optional] 
**BeneficiaryFinancialInstitutionInfo** | [**FiToFiInfo**](FIToFIInfo.md) |  | [optional] 
**BeneficiaryFinancialInstitutionAdviceInfo** | [**AdviceInfo**](AdviceInfo.md) |  | [optional] 
**BeneficiaryInfo** | [**FiToFiInfo**](FIToFIInfo.md) |  | [optional] 
**BeneficiaryAdviceInfo** | [**AdviceInfo**](AdviceInfo.md) |  | [optional] 
**PaymentMethodToBeneficiary** | [**PaymentMethodToBeneficiary**](PaymentMethodToBeneficiary.md) |  | [optional] 
**AdditionalFIToFIInfo** | [**AdditionalFiToFiInfo**](AdditionalFIToFIInfo.md) |  | [optional] 
**CurrencyInstructedAmount** | [**CoverPaymentInfo**](CoverPaymentInfo.md) |  | [optional] 
**OrderingCustomer** | [**CoverPaymentInfo**](CoverPaymentInfo.md) |  | [optional] 
**OrderingInstitution** | [**CoverPaymentInfo**](CoverPaymentInfo.md) |  | [optional] 
**IntermediaryInstitution** | [**CoverPaymentInfo**](CoverPaymentInfo.md) |  | [optional] 
**InstitutionAccount** | [**CoverPaymentInfo**](CoverPaymentInfo.md) |  | [optional] 
**BeneficiaryCustomer** | [**CoverPaymentInfo**](CoverPaymentInfo.md) |  | [optional] 
**RemittanceInfo** | [**CoverPaymentInfo**](CoverPaymentInfo.md) |  | [optional] 
**SenderToReceiverInfo** | [**CoverPaymentInfo**](CoverPaymentInfo.md) |  | [optional] 
**UnstructuredAddendaInfo** | [**UnstructuredAddendaInfo**](UnstructuredAddendaInfo.md) |  | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


