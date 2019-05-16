# \NetconfNodeTopologyApi

All URIs are relative to *http://172.17.0.1:8181/restconf*

Method | HTTP request | Description
------------- | ------------- | -------------
[**RpcNetconfNodeTopologyCreateDevice**](NetconfNodeTopologyApi.md#RpcNetconfNodeTopologyCreateDevice) | **Post** /operations/netconf-node-topology:create-device/ | 
[**RpcNetconfNodeTopologyDeleteDevice**](NetconfNodeTopologyApi.md#RpcNetconfNodeTopologyDeleteDevice) | **Post** /operations/netconf-node-topology:delete-device/ | 


# **RpcNetconfNodeTopologyCreateDevice**
> RpcNetconfNodeTopologyCreateDevice(ctx, netconfNodeTopologyCreatedeviceInputBodyParam)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **netconfNodeTopologyCreatedeviceInputBodyParam** | [**NetconfNodeTopologyCreatedeviceInputBodyparam**](NetconfNodeTopologyCreatedeviceInputBodyparam.md)|  | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RpcNetconfNodeTopologyDeleteDevice**
> RpcNetconfNodeTopologyDeleteDevice(ctx, netconfNodeTopologyDeletedeviceInputBodyParam)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **netconfNodeTopologyDeletedeviceInputBodyParam** | [**NetconfNodeTopologyDeletedeviceInputBodyparam**](NetconfNodeTopologyDeletedeviceInputBodyparam.md)|  | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

