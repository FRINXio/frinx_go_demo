# \FrinxOpenconfigIfAggregateApi

All URIs are relative to *http://172.17.0.1:8181/restconf*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteFrinxOpenconfigInterfacesInterfacesInterfaceAggregation**](FrinxOpenconfigIfAggregateApi.md#DeleteFrinxOpenconfigInterfacesInterfacesInterfaceAggregation) | **Delete** /config/network-topology:network-topology/network-topology:topology/unified/network-topology:node/{node-id}/yang-ext:mount/frinx-openconfig-interfaces:interfaces/frinx-openconfig-interfaces:interface/{name}/frinx-openconfig-if-aggregate:aggregation/ | 
[**DeleteFrinxOpenconfigInterfacesInterfacesInterfaceAggregationConfig**](FrinxOpenconfigIfAggregateApi.md#DeleteFrinxOpenconfigInterfacesInterfacesInterfaceAggregationConfig) | **Delete** /config/network-topology:network-topology/network-topology:topology/unified/network-topology:node/{node-id}/yang-ext:mount/frinx-openconfig-interfaces:interfaces/frinx-openconfig-interfaces:interface/{name}/frinx-openconfig-if-aggregate:aggregation/frinx-openconfig-if-aggregate:config/ | 
[**GetFrinxOpenconfigInterfacesInterfacesInterfaceAggregation**](FrinxOpenconfigIfAggregateApi.md#GetFrinxOpenconfigInterfacesInterfacesInterfaceAggregation) | **Get** /config/network-topology:network-topology/network-topology:topology/unified/network-topology:node/{node-id}/yang-ext:mount/frinx-openconfig-interfaces:interfaces/frinx-openconfig-interfaces:interface/{name}/frinx-openconfig-if-aggregate:aggregation/ | 
[**GetFrinxOpenconfigInterfacesInterfacesInterfaceAggregationConfig**](FrinxOpenconfigIfAggregateApi.md#GetFrinxOpenconfigInterfacesInterfacesInterfaceAggregationConfig) | **Get** /config/network-topology:network-topology/network-topology:topology/unified/network-topology:node/{node-id}/yang-ext:mount/frinx-openconfig-interfaces:interfaces/frinx-openconfig-interfaces:interface/{name}/frinx-openconfig-if-aggregate:aggregation/frinx-openconfig-if-aggregate:config/ | 
[**PutFrinxOpenconfigInterfacesInterfacesInterfaceAggregation**](FrinxOpenconfigIfAggregateApi.md#PutFrinxOpenconfigInterfacesInterfacesInterfaceAggregation) | **Put** /config/network-topology:network-topology/network-topology:topology/unified/network-topology:node/{node-id}/yang-ext:mount/frinx-openconfig-interfaces:interfaces/frinx-openconfig-interfaces:interface/{name}/frinx-openconfig-if-aggregate:aggregation/ | 
[**PutFrinxOpenconfigInterfacesInterfacesInterfaceAggregationConfig**](FrinxOpenconfigIfAggregateApi.md#PutFrinxOpenconfigInterfacesInterfacesInterfaceAggregationConfig) | **Put** /config/network-topology:network-topology/network-topology:topology/unified/network-topology:node/{node-id}/yang-ext:mount/frinx-openconfig-interfaces:interfaces/frinx-openconfig-interfaces:interface/{name}/frinx-openconfig-if-aggregate:aggregation/frinx-openconfig-if-aggregate:config/ | 


# **DeleteFrinxOpenconfigInterfacesInterfacesInterfaceAggregation**
> DeleteFrinxOpenconfigInterfacesInterfacesInterfaceAggregation(ctx, name, nodeId)


removes frinx.openconfig.if.aggregate.aggregationlogicaltop.Aggregation

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| Id of interface | 
  **nodeId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteFrinxOpenconfigInterfacesInterfacesInterfaceAggregationConfig**
> DeleteFrinxOpenconfigInterfacesInterfacesInterfaceAggregationConfig(ctx, name, nodeId)


removes frinx.openconfig.if.aggregate.aggregationlogicaltop.aggregation.Config

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| Id of interface | 
  **nodeId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetFrinxOpenconfigInterfacesInterfacesInterfaceAggregation**
> FrinxOpenconfigIfAggregateAggregationlogicaltopAggregationResponse GetFrinxOpenconfigInterfacesInterfacesInterfaceAggregation(ctx, name, nodeId)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| Id of interface | 
  **nodeId** | **string**|  | 

### Return type

[**FrinxOpenconfigIfAggregateAggregationlogicaltopAggregationResponse**](frinx.openconfig.if.aggregate.aggregationlogicaltop.Aggregation.response.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetFrinxOpenconfigInterfacesInterfacesInterfaceAggregationConfig**
> FrinxOpenconfigIfAggregateAggregationlogicaltopAggregationConfigResponse GetFrinxOpenconfigInterfacesInterfacesInterfaceAggregationConfig(ctx, name, nodeId)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| Id of interface | 
  **nodeId** | **string**|  | 

### Return type

[**FrinxOpenconfigIfAggregateAggregationlogicaltopAggregationConfigResponse**](frinx.openconfig.if.aggregate.aggregationlogicaltop.aggregation.Config.response.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PutFrinxOpenconfigInterfacesInterfacesInterfaceAggregation**
> PutFrinxOpenconfigInterfacesInterfacesInterfaceAggregation(ctx, name, frinxOpenconfigIfAggregateAggregationlogicaltopAggregationBodyParam, nodeId)


creates or updates frinx.openconfig.if.aggregate.aggregationlogicaltop.Aggregation

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| Id of interface | 
  **frinxOpenconfigIfAggregateAggregationlogicaltopAggregationBodyParam** | [**FrinxOpenconfigIfAggregateAggregationlogicaltopAggregationRequest**](FrinxOpenconfigIfAggregateAggregationlogicaltopAggregationRequest.md)| frinx.openconfig.if.aggregate.aggregationlogicaltop.Aggregation to be added or updated | 
  **nodeId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PutFrinxOpenconfigInterfacesInterfacesInterfaceAggregationConfig**
> PutFrinxOpenconfigInterfacesInterfacesInterfaceAggregationConfig(ctx, name, frinxOpenconfigIfAggregateAggregationlogicaltopAggregationConfigBodyParam, nodeId)


creates or updates frinx.openconfig.if.aggregate.aggregationlogicaltop.aggregation.Config

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| Id of interface | 
  **frinxOpenconfigIfAggregateAggregationlogicaltopAggregationConfigBodyParam** | [**FrinxOpenconfigIfAggregateAggregationlogicaltopAggregationConfigRequest**](FrinxOpenconfigIfAggregateAggregationlogicaltopAggregationConfigRequest.md)| frinx.openconfig.if.aggregate.aggregationlogicaltop.aggregation.Config to be added or updated | 
  **nodeId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

