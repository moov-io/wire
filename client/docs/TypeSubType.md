# TypeSubType

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TypeCode** | **string** | TypeCode:  * &#x60;10&#x60; - Funds Transfer - A funds transfer in which the sender and/or receiver may be a bank or a third party (i.e., customer of a bank). * &#x60;15&#x60; - Foreign Transfer - A funds transfer to or from a foreign central bank or government or international organization with an account at the Federal Reserve Bank of New York. * &#x60;16&#x60; - Settlement Transfer - A funds transfer between Fedwire Funds Service participants.  | 
**SubTypeCode** | **string** | SubTypeCode:  * &#x60;00&#x60; - Basic Funds Transfer - A basic value funds transfer. * &#x60;01&#x60; - Request for Reversal - A non-value request for reversal of a funds transfer originated on the current business day. * &#x60;02&#x60; - Reversal of Transfer - A value reversal of a funds transfer received on the current business day.  May be used in response to a subtype code ‘01’ Request for Reversal. * &#x60;07&#x60; - Request for Reversal of a Prior Day Transfer - A non-value request for a reversal of a funds transfer originated on a prior business day. * &#x60;08&#x60; - Reversal of a Prior Day Transfer - A value reversal of a funds transfer received on a prior business day.  May be used in response to a subtype code ‘07’ Request for Reversal of a Prior Day Transfer. * &#x60;31&#x60; - Request for Credit (Drawdown) - A non-value request for the receiver to send a funds transfer to a designated party. * &#x60;32&#x60; - Funds Transfer Honoring a Request for Credit (Drawdown) -  A value funds transfer honoring a subtype 31 request for credit. * &#x60;33&#x60; -Refusal to Honor a Request for Credit (Drawdown) - A non-value message indicating refusal to honor a subtype 31 request for credit. * &#x60;90&#x60; - Service Message - A non-value message used to communicate questions and information that is not covered by a specific subtype.  | 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


