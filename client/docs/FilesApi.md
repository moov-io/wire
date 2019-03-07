# \FilesApi

All URIs are relative to *http://localhost:8087*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateFile**](FilesApi.md#CreateFile) | **Post** /files/create | Create a new File object
[**DeleteWIREFile**](FilesApi.md#DeleteWIREFile) | **Delete** /files/{file_id} | Permanently deletes a File and associated Batches. It cannot be undone.
[**GetFileByID**](FilesApi.md#GetFileByID) | **Get** /files/{file_id} | Retrieves the details of an existing File. You need only supply the unique File identifier that was returned upon creation.
[**GetFiles**](FilesApi.md#GetFiles) | **Get** /files | Gets a list of Files
[**Ping**](FilesApi.md#Ping) | **Get** /ping | Ping the Wire service to check if running
[**UpdateFile**](FilesApi.md#UpdateFile) | **Post** /files/{file_id} | Updates the specified FedWire Message by setting the values of the parameters passed. Any parameters not provided will be left unchanged.


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

# **DeleteWIREFile**
> DeleteWIREFile(ctx, fileId, optional)
Permanently deletes a File and associated Batches. It cannot be undone.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **fileId** | **string**| File ID | 
 **optional** | ***DeleteWIREFileOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DeleteWIREFileOpts struct

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

# **GetFileByID**
> File GetFileByID(ctx, fileId, optional)
Retrieves the details of an existing File. You need only supply the unique File identifier that was returned upon creation.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **fileId** | **string**| File ID | 
 **optional** | ***GetFileByIDOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetFileByIDOpts struct

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

# **GetFiles**
> Files GetFiles(ctx, optional)
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

[**Files**](Files.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Ping**
> Ping(ctx, )
Ping the Wire service to check if running

### Required Parameters
This endpoint does not need any parameter.

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateFile**
> File UpdateFile(ctx, fileId, createFile, optional)
Updates the specified FedWire Message by setting the values of the parameters passed. Any parameters not provided will be left unchanged.

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

