# \WireFilesApi

All URIs are relative to *http://localhost:8088*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddFEDWireMessageToFile**](WireFilesApi.md#AddFEDWireMessageToFile) | **Post** /files/{fileID}/FEDWireMessage | Add Fedwire message to file
[**CreateWireFile**](WireFilesApi.md#CreateWireFile) | **Post** /files/create | Create file
[**DeleteWireFileByID**](WireFilesApi.md#DeleteWireFileByID) | **Delete** /files/{fileID} | Delete file
[**GetWireFileByID**](WireFilesApi.md#GetWireFileByID) | **Get** /files/{fileID} | Retrieve file
[**GetWireFileContents**](WireFilesApi.md#GetWireFileContents) | **Get** /files/{fileID}/contents | Get file contents
[**GetWireFiles**](WireFilesApi.md#GetWireFiles) | **Get** /files | List files
[**Ping**](WireFilesApi.md#Ping) | **Get** /ping | Ping Wire service
[**ValidateWireFile**](WireFilesApi.md#ValidateWireFile) | **Get** /files/{fileID}/validate | Validate file



## AddFEDWireMessageToFile

> AddFEDWireMessageToFile(ctx, fileID, fedWireMessage, optional)

Add Fedwire message to file

Add a Fedwire Message to the specified file.

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


 **xRequestID** | **optional.String**| Optional Request ID allows application developer to trace requests through the system&#39;s logs | 

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


## CreateWireFile

> WireFile CreateWireFile(ctx, wireFile, optional)

Create file

Upload a new Wire file, or create one from JSON. When uploading a file, query parameters can be used to configure the FedWireMessage validation options. For JSON requests, validation options are set in the  request body under fedWireMessage.validateOptions. 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**wireFile** | [**WireFile**](WireFile.md)| Content of the Wire file (in json or raw text) | 
 **optional** | ***CreateWireFileOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CreateWireFileOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xRequestID** | **optional.String**| Optional Request ID allows application developer to trace requests through the system&#39;s logs | 
 **skipMandatoryIMAD** | **optional.Bool**| Optional flag to skip mandatory IMAD validation | [default to false]
 **allowMissingSenderSupplied** | **optional.Bool**| Optional flag to allow SenderSupplied to be nil, which is generally the case in incoming files. | [default to false]

### Return type

[**WireFile**](WireFile.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json, text/plain
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteWireFileByID

> DeleteWireFileByID(ctx, fileID, optional)

Delete file

Permanently delete a File and associated message. It cannot be undone.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**fileID** | **string**| File ID | 
 **optional** | ***DeleteWireFileByIDOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a DeleteWireFileByIDOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xRequestID** | **optional.String**| Optional Request ID allows application developer to trace requests through the system&#39;s logs | 

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


## GetWireFileByID

> WireFile GetWireFileByID(ctx, fileID, optional)

Retrieve file

Get the details of an existing File using the unique File identifier that was returned upon creation.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**fileID** | **string**| File ID | 
 **optional** | ***GetWireFileByIDOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetWireFileByIDOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xRequestID** | **optional.String**| Optional Request ID allows application developer to trace requests through the system&#39;s logs | 

### Return type

[**WireFile**](WireFile.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetWireFileContents

> string GetWireFileContents(ctx, fileID, optional)

Get file contents

Assembles the existing file, computes sequence numbers and totals. Returns plaintext file. 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**fileID** | **string**| File ID | 
 **optional** | ***GetWireFileContentsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetWireFileContentsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xRequestID** | **optional.String**| Optional Request ID allows application developer to trace requests through the system&#39;s logs | 
 **format** | **optional.String**| Optional file type to get file as fixed length or variable length type | 
 **newline** | **optional.Bool**| Optional new line flag to have new line or no new line | 

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

> []WireFile GetWireFiles(ctx, optional)

List files

List all Wire files created with the Wire service. These files are not persisted through multiple runs of the service.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GetWireFilesOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetWireFilesOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xRequestID** | **optional.String**| Optional Request ID allows application developer to trace requests through the system&#39;s logs | 

### Return type

[**[]WireFile**](WireFile.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Ping

> Ping(ctx, )

Ping Wire service

Check if the Wire service is running.

### Required Parameters

This endpoint does not need any parameter.

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


## ValidateWireFile

> WireFile ValidateWireFile(ctx, fileID, optional)

Validate file

Validates the existing file. You need only supply the unique File identifier that was returned upon creation.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**fileID** | **string**| File ID | 
 **optional** | ***ValidateWireFileOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ValidateWireFileOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xRequestID** | **optional.String**| Optional Request ID allows application developer to trace requests through the system&#39;s logs | 

### Return type

[**WireFile**](WireFile.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

