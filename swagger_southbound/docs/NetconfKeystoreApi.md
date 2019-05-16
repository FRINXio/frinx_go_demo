# \NetconfKeystoreApi

All URIs are relative to *http://172.17.0.1:8181/restconf*

Method | HTTP request | Description
------------- | ------------- | -------------
[**RpcNetconfKeystoreAddKeystoreEntry**](NetconfKeystoreApi.md#RpcNetconfKeystoreAddKeystoreEntry) | **Post** /operations/netconf-keystore:add-keystore-entry/ | 
[**RpcNetconfKeystoreAddPrivateKey**](NetconfKeystoreApi.md#RpcNetconfKeystoreAddPrivateKey) | **Post** /operations/netconf-keystore:add-private-key/ | 
[**RpcNetconfKeystoreAddTrustedCertificate**](NetconfKeystoreApi.md#RpcNetconfKeystoreAddTrustedCertificate) | **Post** /operations/netconf-keystore:add-trusted-certificate/ | 
[**RpcNetconfKeystoreRemoveKeystoreEntry**](NetconfKeystoreApi.md#RpcNetconfKeystoreRemoveKeystoreEntry) | **Post** /operations/netconf-keystore:remove-keystore-entry/ | 
[**RpcNetconfKeystoreRemovePrivateKey**](NetconfKeystoreApi.md#RpcNetconfKeystoreRemovePrivateKey) | **Post** /operations/netconf-keystore:remove-private-key/ | 
[**RpcNetconfKeystoreRemoveTrustedCertificate**](NetconfKeystoreApi.md#RpcNetconfKeystoreRemoveTrustedCertificate) | **Post** /operations/netconf-keystore:remove-trusted-certificate/ | 


# **RpcNetconfKeystoreAddKeystoreEntry**
> RpcNetconfKeystoreAddKeystoreEntry(ctx, netconfKeystoreAddkeystoreentryInputBodyParam)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **netconfKeystoreAddkeystoreentryInputBodyParam** | [**NetconfKeystoreAddkeystoreentryInputBodyparam**](NetconfKeystoreAddkeystoreentryInputBodyparam.md)|  | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RpcNetconfKeystoreAddPrivateKey**
> RpcNetconfKeystoreAddPrivateKey(ctx, netconfKeystoreAddprivatekeyInputBodyParam)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **netconfKeystoreAddprivatekeyInputBodyParam** | [**NetconfKeystoreAddprivatekeyInputBodyparam**](NetconfKeystoreAddprivatekeyInputBodyparam.md)|  | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RpcNetconfKeystoreAddTrustedCertificate**
> RpcNetconfKeystoreAddTrustedCertificate(ctx, netconfKeystoreAddtrustedcertificateInputBodyParam)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **netconfKeystoreAddtrustedcertificateInputBodyParam** | [**NetconfKeystoreAddtrustedcertificateInputBodyparam**](NetconfKeystoreAddtrustedcertificateInputBodyparam.md)|  | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RpcNetconfKeystoreRemoveKeystoreEntry**
> RpcNetconfKeystoreRemoveKeystoreEntry(ctx, netconfKeystoreRemovekeystoreentryInputBodyParam)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **netconfKeystoreRemovekeystoreentryInputBodyParam** | [**NetconfKeystoreRemovekeystoreentryInputBodyparam**](NetconfKeystoreRemovekeystoreentryInputBodyparam.md)|  | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RpcNetconfKeystoreRemovePrivateKey**
> RpcNetconfKeystoreRemovePrivateKey(ctx, netconfKeystoreRemoveprivatekeyInputBodyParam)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **netconfKeystoreRemoveprivatekeyInputBodyParam** | [**NetconfKeystoreRemoveprivatekeyInputBodyparam**](NetconfKeystoreRemoveprivatekeyInputBodyparam.md)|  | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RpcNetconfKeystoreRemoveTrustedCertificate**
> RpcNetconfKeystoreRemoveTrustedCertificate(ctx, netconfKeystoreRemovetrustedcertificateInputBodyParam)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **netconfKeystoreRemovetrustedcertificateInputBodyParam** | [**NetconfKeystoreRemovetrustedcertificateInputBodyparam**](NetconfKeystoreRemovetrustedcertificateInputBodyparam.md)|  | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

