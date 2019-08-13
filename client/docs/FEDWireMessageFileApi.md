# \FEDWireMessageFileApi

All URIs are relative to *http://localhost:8087*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddFEDWireMessageToFile**](FEDWireMessageFileApi.md#AddFEDWireMessageToFile) | **Post** /files/{fileID}/FEDWireMessage | Add FEDWireMessage to File



## AddFEDWireMessageToFile

> AddFEDWireMessageToFile(ctx, fileID, fedWireMessage, optional)
Add FEDWireMessage to File

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**fileID** | **string**| File ID | 
**fedWireMessage** | [**FedWireMessage**](FedWireMessage.md)|  | 
 **optional** | ***AddFEDWireMessageToFileOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a AddFEDWireMessageToFileOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **xRequestID** | **optional.String**| Optional Request ID allows application developer to trace requests through the systems logs | 
 **xIdempotencyKey** | **optional.String**| Idempotent key in the header which expires after 24 hours. These strings should contain enough entropy for to not collide with each other in your requests. | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

