# \WireFilesApi

All URIs are relative to *http://localhost:8087*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateFile**](WireFilesApi.md#CreateFile) | **Post** /files/create | Create a new File object
[**DeleteFile**](WireFilesApi.md#DeleteFile) | **Delete** /files/{file_id} | Permanently deletes a File and associated FEDWireMessage. It cannot be undone.
[**DeleteFileFEDWireMessage**](WireFilesApi.md#DeleteFileFEDWireMessage) | **Delete** /files/{file_id}/fEDWireMessage/{fEDWireMessage_id} | Delete a FEDWireMessage from a File
[**GetFileContents**](WireFilesApi.md#GetFileContents) | **Get** /files/{file_id}/contents | Assembles the existing file witha FEDWireMessage, Returns plaintext file.
[**GetFileFEDWireMessage**](WireFilesApi.md#GetFileFEDWireMessage) | **Get** /files/{file_id}/fEDWireMessage/{fEDWireMessage_id} | Get a specific FEDWireMessage on a FIle
[**GetFiles**](WireFilesApi.md#GetFiles) | **Get** /files | Gets a list of Files
[**UpdateFile**](WireFilesApi.md#UpdateFile) | **Post** /files/{file_id} | Updates the specified FEDWire Message by setting the values of the parameters passed. Any parameters not provided will be left unchanged.
[**ValidateFile**](WireFilesApi.md#ValidateFile) | **Get** /files/{file_id}/validate | Validates the existing file. You need only supply the unique File identifier that was returned upon creation.


# **CreateFile**
> File CreateFile(ctx, createFile, optional)
Create a new File object

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **createFile** | [**CreateFile**](CreateFile.md)| Content of the WIRE file (in json or raw text) | 
 **optional** | ***CreateFileOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CreateFileOpts struct

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

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteFile**
> DeleteFile(ctx, fileId, optional)
Permanently deletes a File and associated FEDWireMessage. It cannot be undone.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **fileId** | **string**| File ID | 
 **optional** | ***DeleteFileOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DeleteFileOpts struct

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

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteFileFEDWireMessage**
> DeleteFileFEDWireMessage(ctx, fileId, fEDWireMessageId, optional)
Delete a FEDWireMessage from a File

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **fileId** | **string**| File ID | 
  **fEDWireMessageId** | **string**| FEDWireMessage ID | 
 **optional** | ***DeleteFileFEDWireMessageOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DeleteFileFEDWireMessageOpts struct

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

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetFileContents**
> string GetFileContents(ctx, fileId, optional)
Assembles the existing file witha FEDWireMessage, Returns plaintext file.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **fileId** | **string**| File ID | 
 **optional** | ***GetFileContentsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetFileContentsOpts struct

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

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetFileFEDWireMessage**
> FedWireMessage GetFileFEDWireMessage(ctx, fileId, fEDWireMessageId, optional)
Get a specific FEDWireMessage on a FIle

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **fileId** | **string**| File ID | 
  **fEDWireMessageId** | **string**| FEDWireMessage ID | 
 **optional** | ***GetFileFEDWireMessageOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetFileFEDWireMessageOpts struct

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

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetFiles**
> []File GetFiles(ctx, optional)
Gets a list of Files

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GetFilesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetFilesOpts struct

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

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateFile**
> File UpdateFile(ctx, fileId, createFile, optional)
Updates the specified FEDWire Message by setting the values of the parameters passed. Any parameters not provided will be left unchanged.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **fileId** | **string**| File ID | 
  **createFile** | [**CreateFile**](CreateFile.md)|  | 
 **optional** | ***UpdateFileOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UpdateFileOpts struct

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

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ValidateFile**
> File ValidateFile(ctx, fileId, optional)
Validates the existing file. You need only supply the unique File identifier that was returned upon creation.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **fileId** | **string**| File ID | 
 **optional** | ***ValidateFileOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ValidateFileOpts struct

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

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

