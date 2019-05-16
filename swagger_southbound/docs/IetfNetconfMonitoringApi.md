# \IetfNetconfMonitoringApi

All URIs are relative to *http://172.17.0.1:8181/restconf*

Method | HTTP request | Description
------------- | ------------- | -------------
[**RpcIetfNetconfMonitoringGetSchema**](IetfNetconfMonitoringApi.md#RpcIetfNetconfMonitoringGetSchema) | **Post** /operations/ietf-netconf-monitoring:get-schema/ | 


# **RpcIetfNetconfMonitoringGetSchema**
> IetfNetconfMonitoringGetSchema RpcIetfNetconfMonitoringGetSchema(ctx, ietfNetconfMonitoringGetschemaInputBodyParam)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **ietfNetconfMonitoringGetschemaInputBodyParam** | [**IetfNetconfMonitoringGetschemaInputBodyparam**](IetfNetconfMonitoringGetschemaInputBodyparam.md)|  | 

### Return type

[**IetfNetconfMonitoringGetSchema**](ietf.netconf.monitoring.GetSchema.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

