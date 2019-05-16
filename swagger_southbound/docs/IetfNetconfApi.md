# \IetfNetconfApi

All URIs are relative to *http://172.17.0.1:8181/restconf*

Method | HTTP request | Description
------------- | ------------- | -------------
[**RpcIetfNetconfCancelCommit**](IetfNetconfApi.md#RpcIetfNetconfCancelCommit) | **Post** /operations/ietf-netconf:cancel-commit/ | 
[**RpcIetfNetconfCloseSession**](IetfNetconfApi.md#RpcIetfNetconfCloseSession) | **Post** /operations/ietf-netconf:close-session/ | 
[**RpcIetfNetconfCommit**](IetfNetconfApi.md#RpcIetfNetconfCommit) | **Post** /operations/ietf-netconf:commit/ | 
[**RpcIetfNetconfCopyConfig**](IetfNetconfApi.md#RpcIetfNetconfCopyConfig) | **Post** /operations/ietf-netconf:copy-config/ | 
[**RpcIetfNetconfDeleteConfig**](IetfNetconfApi.md#RpcIetfNetconfDeleteConfig) | **Post** /operations/ietf-netconf:delete-config/ | 
[**RpcIetfNetconfDiscardChanges**](IetfNetconfApi.md#RpcIetfNetconfDiscardChanges) | **Post** /operations/ietf-netconf:discard-changes/ | 
[**RpcIetfNetconfEditConfig**](IetfNetconfApi.md#RpcIetfNetconfEditConfig) | **Post** /operations/ietf-netconf:edit-config/ | 
[**RpcIetfNetconfGet**](IetfNetconfApi.md#RpcIetfNetconfGet) | **Post** /operations/ietf-netconf:get/ | 
[**RpcIetfNetconfGetConfig**](IetfNetconfApi.md#RpcIetfNetconfGetConfig) | **Post** /operations/ietf-netconf:get-config/ | 
[**RpcIetfNetconfKillSession**](IetfNetconfApi.md#RpcIetfNetconfKillSession) | **Post** /operations/ietf-netconf:kill-session/ | 
[**RpcIetfNetconfLock**](IetfNetconfApi.md#RpcIetfNetconfLock) | **Post** /operations/ietf-netconf:lock/ | 
[**RpcIetfNetconfUnlock**](IetfNetconfApi.md#RpcIetfNetconfUnlock) | **Post** /operations/ietf-netconf:unlock/ | 
[**RpcIetfNetconfValidate**](IetfNetconfApi.md#RpcIetfNetconfValidate) | **Post** /operations/ietf-netconf:validate/ | 


# **RpcIetfNetconfCancelCommit**
> RpcIetfNetconfCancelCommit(ctx, ietfNetconfCancelcommitInputBodyParam)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **ietfNetconfCancelcommitInputBodyParam** | [**IetfNetconfCancelcommitInputBodyparam**](IetfNetconfCancelcommitInputBodyparam.md)|  | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RpcIetfNetconfCloseSession**
> RpcIetfNetconfCloseSession(ctx, )


### Required Parameters
This endpoint does not need any parameter.

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RpcIetfNetconfCommit**
> RpcIetfNetconfCommit(ctx, ietfNetconfCommitInputBodyParam)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **ietfNetconfCommitInputBodyParam** | [**IetfNetconfCommitInputBodyparam**](IetfNetconfCommitInputBodyparam.md)|  | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RpcIetfNetconfCopyConfig**
> RpcIetfNetconfCopyConfig(ctx, ietfNetconfCopyconfigInputBodyParam)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **ietfNetconfCopyconfigInputBodyParam** | [**IetfNetconfCopyconfigInputBodyparam**](IetfNetconfCopyconfigInputBodyparam.md)|  | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RpcIetfNetconfDeleteConfig**
> RpcIetfNetconfDeleteConfig(ctx, ietfNetconfDeleteconfigInputBodyParam)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **ietfNetconfDeleteconfigInputBodyParam** | [**IetfNetconfDeleteconfigInputBodyparam**](IetfNetconfDeleteconfigInputBodyparam.md)|  | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RpcIetfNetconfDiscardChanges**
> RpcIetfNetconfDiscardChanges(ctx, )


### Required Parameters
This endpoint does not need any parameter.

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RpcIetfNetconfEditConfig**
> RpcIetfNetconfEditConfig(ctx, ietfNetconfEditconfigInputBodyParam)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **ietfNetconfEditconfigInputBodyParam** | [**IetfNetconfEditconfigInputBodyparam**](IetfNetconfEditconfigInputBodyparam.md)|  | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RpcIetfNetconfGet**
> IetfNetconfGet RpcIetfNetconfGet(ctx, ietfNetconfGetInputBodyParam)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **ietfNetconfGetInputBodyParam** | [**IetfNetconfGetInputBodyparam**](IetfNetconfGetInputBodyparam.md)|  | 

### Return type

[**IetfNetconfGet**](ietf.netconf.Get.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RpcIetfNetconfGetConfig**
> IetfNetconfGetConfig RpcIetfNetconfGetConfig(ctx, ietfNetconfGetconfigInputBodyParam)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **ietfNetconfGetconfigInputBodyParam** | [**IetfNetconfGetconfigInputBodyparam**](IetfNetconfGetconfigInputBodyparam.md)|  | 

### Return type

[**IetfNetconfGetConfig**](ietf.netconf.GetConfig.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RpcIetfNetconfKillSession**
> RpcIetfNetconfKillSession(ctx, ietfNetconfKillsessionInputBodyParam)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **ietfNetconfKillsessionInputBodyParam** | [**IetfNetconfKillsessionInputBodyparam**](IetfNetconfKillsessionInputBodyparam.md)|  | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RpcIetfNetconfLock**
> RpcIetfNetconfLock(ctx, ietfNetconfLockInputBodyParam)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **ietfNetconfLockInputBodyParam** | [**IetfNetconfLockInputBodyparam**](IetfNetconfLockInputBodyparam.md)|  | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RpcIetfNetconfUnlock**
> RpcIetfNetconfUnlock(ctx, ietfNetconfUnlockInputBodyParam)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **ietfNetconfUnlockInputBodyParam** | [**IetfNetconfUnlockInputBodyparam**](IetfNetconfUnlockInputBodyparam.md)|  | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RpcIetfNetconfValidate**
> RpcIetfNetconfValidate(ctx, ietfNetconfValidateInputBodyParam)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **ietfNetconfValidateInputBodyParam** | [**IetfNetconfValidateInputBodyparam**](IetfNetconfValidateInputBodyparam.md)|  | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

