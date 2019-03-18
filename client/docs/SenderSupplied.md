# SenderSupplied

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FormatVersion** | **string** | FormatVersion 30  | 
**UserRequestCorrelation** | **string** | UserRequestCorrelation | 
**TestProductionCode** | [**TestProductionCodeEnum**](TestProductionCodeEnum.md) |  | 
**MessageDuplicationCode** | **string** | MessageDuplicationCode  * &#x60; &#x60; - Original Message * &#x60;R&#x60; - Retrieval of an original message * &#x60;P&#x60; - Resend  | 
**MessageStatusIndicator** | **string** | Message Status Indicator  Outgoing Messages * &#x60;0&#x60; - In process or Intercepted * &#x60;2&#x60; - Successful with Accounting (Value) * &#x60;3&#x60; - Rejected due to Error Condition * &#x60;7&#x60; - Successful without Accounting (Non-Value)  Incoming Messages * &#x60;N&#x60; - Successful with Accounting (Value) * &#x60;S&#x60; - Successful without Accounting (Non-Value)  | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


