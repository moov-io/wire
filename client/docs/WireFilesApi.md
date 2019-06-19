# \WireFilesApi

All URIs are relative to *http://localhost:8087*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateWireFile**](WireFilesApi.md#CreateWireFile) | **Post** /files/create | Create a new File object
[**DeleteFEDWireMessageFromFile**](WireFilesApi.md#DeleteFEDWireMessageFromFile) | **Delete** /files/{file_id}/fEDWireMessage/{fEDWireMessage_id} | Delete a FEDWireMessage from a File
[**DeleteWireFileByID**](WireFilesApi.md#DeleteWireFileByID) | **Delete** /files/{file_id} | Permanently deletes a File and associated FEDWireMessage. It cannot be undone.
[**GetFEDWireMessageFromFile**](WireFilesApi.md#GetFEDWireMessageFromFile) | **Get** /files/{file_id}/fEDWireMessage/{fEDWireMessage_id} | Get a specific FEDWireMessage on a FIle
[**GetWireFileContents**](WireFilesApi.md#GetWireFileContents) | **Get** /files/{file_id}/contents | Assembles the existing file witha FEDWireMessage, Returns plaintext file.
[**GetWireFiles**](WireFilesApi.md#GetWireFiles) | **Get** /files | Gets a list of Files
[**UpdateWireFileByID**](WireFilesApi.md#UpdateWireFileByID) | **Post** /files/{file_id} | Updates the specified FEDWire Message by setting the values of the parameters passed. Any parameters not provided will be left unchanged.
[**ValidateWireFile**](WireFilesApi.md#ValidateWireFile) | **Get** /files/{file_id}/validate | Validates the existing file. You need only supply the unique File identifier that was returned upon creation.



## CreateWireFile

> File CreateWireFile(ctx, createFile, optional)
Create a new File object

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**createFile** | [**CreateFile**](CreateFile.md)| Content of the WIRE file (in json or raw text) | 
 **optional** | ***CreateWireFileOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CreateWireFileOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xRequestId** | **optional.String**| Optional Request ID allows application developer to trace requests through the systems logs | 
 **xIdempotencyKey** | **optional.String**| Idempotent key in the header which expires after 24 hours. These strings should contain enough entropy for to not collide with each other in your requests. | 

### Return type

[**File**](File.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json, text/plain
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteFEDWireMessageFromFile

> DeleteFEDWireMessageFromFile(ctx, fileId, fEDWireMessageId, optional)
Delete a FEDWireMessage from a File

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**fileId** | **string**| File ID | 
**fEDWireMessageId** | **string**| FEDWireMessage ID | 
 **optional** | ***DeleteFEDWireMessageFromFileOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a DeleteFEDWireMessageFromFileOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **xRequestId** | **optional.String**| Optional Request ID allows application developer to trace requests through the systems logs | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteWireFileByID

> DeleteWireFileByID(ctx, fileId, optional)
Permanently deletes a File and associated FEDWireMessage. It cannot be undone.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**fileId** | **string**| File ID | 
 **optional** | ***DeleteWireFileByIDOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a DeleteWireFileByIDOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xRequestId** | **optional.String**| Optional Request ID allows application developer to trace requests through the systems logs | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFEDWireMessageFromFile

> FedWireMessage GetFEDWireMessageFromFile(ctx, fileId, fEDWireMessageId, optional)
Get a specific FEDWireMessage on a FIle

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**fileId** | **string**| File ID | 
**fEDWireMessageId** | **string**| FEDWireMessage ID | 
 **optional** | ***GetFEDWireMessageFromFileOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetFEDWireMessageFromFileOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **xRequestId** | **optional.String**| Optional Request ID allows application developer to trace requests through the systems logs | 

### Return type

[**FedWireMessage**](FEDWireMessage.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetWireFileContents

> string GetWireFileContents(ctx, fileId, optional)
Assembles the existing file witha FEDWireMessage, Returns plaintext file.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**fileId** | **string**| File ID | 
 **optional** | ***GetWireFileContentsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetWireFileContentsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xRequestId** | **optional.String**| Optional Request ID allows application developer to trace requests through the systems logs | 

### Return type

**string**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetWireFiles

> []File GetWireFiles(ctx, optional)
Gets a list of Files

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GetWireFilesOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetWireFilesOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xRequestId** | **optional.String**| Optional Request ID allows application developer to trace requests through the systems logs | 

### Return type

[**[]File**](File.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateWireFileByID

> File UpdateWireFileByID(ctx, fileId, createFile, optional)
Updates the specified FEDWire Message by setting the values of the parameters passed. Any parameters not provided will be left unchanged.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**fileId** | **string**| File ID | 
**createFile** | [**CreateFile**](CreateFile.md)|  | 
 **optional** | ***UpdateWireFileByIDOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UpdateWireFileByIDOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **xRequestId** | **optional.String**| Optional Request ID allows application developer to trace requests through the systems logs | 
 **xIdempotencyKey** | **optional.String**| Idempotent key in the header which expires after 24 hours. These strings should contain enough entropy for to not collide with each other in your requests. | 

### Return type

[**File**](File.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ValidateWireFile

> File ValidateWireFile(ctx, fileId, optional)
Validates the existing file. You need only supply the unique File identifier that was returned upon creation.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**fileId** | **string**| File ID | 
 **optional** | ***ValidateWireFileOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ValidateWireFileOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xRequestId** | **optional.String**| Optional Request ID allows application developer to trace requests through the systems logs | 

### Return type

[**File**](File.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

