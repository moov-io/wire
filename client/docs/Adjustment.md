# Adjustment

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AdjustmentReasonCode** | **string** | Adjustment  * &#x60;01&#x60; - Pricing Error * &#x60;03&#x60; - Extension Error * &#x60;04&#x60; - Item Not Accepted (Damaged) * &#x60;05&#x60; - Item Not Accepted (Quality) * &#x60;06&#x60; - Quantity Contested 07   Incorrect Product * &#x60;11&#x60; - Returns (Damaged) * &#x60;12&#x60; - Returns (Quality) * &#x60;59&#x60; - Item Not Received * &#x60;75&#x60; - Total Order Not Received * &#x60;81&#x60; - Credit as Agreed * &#x60;CM&#x60; - Covered by Credit Memo  | [optional] 
**CreditDebitIndicator** | **string** | CreditDebitIndicator  * &#x60;CRDT&#x60; - Credit * &#x60;DBIT&#x60; - Debit  | [optional] 
**CurrencyCode** | **string** | CurrencyCode | [optional] 
**Amount** | **string** | Amount Must contain at least one numeric character and only one decimal period marker (e.g., $1,234.56 should be entered as 1234.56). Can have up to 5 numeric characters following the decimal period marker (e.g., 1234.56789). Amount must be greater than zero (i.e., at least .01).  | [optional] 
**AdditionalInfo** | **string** | AdditionalInfo | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


