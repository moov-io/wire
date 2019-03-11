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
**PaymentMethodToBeneficiary** | [**map[string]interface{}**](map[string]interface{}.md) | PaymentMethodToBeneficiary | [optional] 
**AdditionalFIToFIInfo** | [**map[string]interface{}**](map[string]interface{}.md) | AdditionalFIToFIInfo | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


