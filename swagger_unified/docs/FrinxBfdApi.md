# \FrinxBfdApi

All URIs are relative to *http://172.17.0.1:8181/restconf*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteFrinxOpenconfigInterfacesInterfacesInterfaceAggregationBfd**](FrinxBfdApi.md#DeleteFrinxOpenconfigInterfacesInterfacesInterfaceAggregationBfd) | **Delete** /config/network-topology:network-topology/network-topology:topology/unified/network-topology:node/{node-id}/yang-ext:mount/frinx-openconfig-interfaces:interfaces/frinx-openconfig-interfaces:interface/{name}/frinx-openconfig-if-aggregate:aggregation/frinx-bfd:bfd/ | 
[**DeleteFrinxOpenconfigInterfacesInterfacesInterfaceAggregationBfdConfig**](FrinxBfdApi.md#DeleteFrinxOpenconfigInterfacesInterfacesInterfaceAggregationBfdConfig) | **Delete** /config/network-topology:network-topology/network-topology:topology/unified/network-topology:node/{node-id}/yang-ext:mount/frinx-openconfig-interfaces:interfaces/frinx-openconfig-interfaces:interface/{name}/frinx-openconfig-if-aggregate:aggregation/frinx-bfd:bfd/frinx-bfd:config/ | 
[**GetFrinxOpenconfigInterfacesInterfacesInterfaceAggregationBfd**](FrinxBfdApi.md#GetFrinxOpenconfigInterfacesInterfacesInterfaceAggregationBfd) | **Get** /config/network-topology:network-topology/network-topology:topology/unified/network-topology:node/{node-id}/yang-ext:mount/frinx-openconfig-interfaces:interfaces/frinx-openconfig-interfaces:interface/{name}/frinx-openconfig-if-aggregate:aggregation/frinx-bfd:bfd/ | 
[**GetFrinxOpenconfigInterfacesInterfacesInterfaceAggregationBfdConfig**](FrinxBfdApi.md#GetFrinxOpenconfigInterfacesInterfacesInterfaceAggregationBfdConfig) | **Get** /config/network-topology:network-topology/network-topology:topology/unified/network-topology:node/{node-id}/yang-ext:mount/frinx-openconfig-interfaces:interfaces/frinx-openconfig-interfaces:interface/{name}/frinx-openconfig-if-aggregate:aggregation/frinx-bfd:bfd/frinx-bfd:config/ | 
[**PutFrinxOpenconfigInterfacesInterfacesInterfaceAggregationBfd**](FrinxBfdApi.md#PutFrinxOpenconfigInterfacesInterfacesInterfaceAggregationBfd) | **Put** /config/network-topology:network-topology/network-topology:topology/unified/network-topology:node/{node-id}/yang-ext:mount/frinx-openconfig-interfaces:interfaces/frinx-openconfig-interfaces:interface/{name}/frinx-openconfig-if-aggregate:aggregation/frinx-bfd:bfd/ | 
[**PutFrinxOpenconfigInterfacesInterfacesInterfaceAggregationBfdConfig**](FrinxBfdApi.md#PutFrinxOpenconfigInterfacesInterfacesInterfaceAggregationBfdConfig) | **Put** /config/network-topology:network-topology/network-topology:topology/unified/network-topology:node/{node-id}/yang-ext:mount/frinx-openconfig-interfaces:interfaces/frinx-openconfig-interfaces:interface/{name}/frinx-openconfig-if-aggregate:aggregation/frinx-bfd:bfd/frinx-bfd:config/ | 


# **DeleteFrinxOpenconfigInterfacesInterfacesInterfaceAggregationBfd**
> DeleteFrinxOpenconfigInterfacesInterfacesInterfaceAggregationBfd(ctx, name, nodeId)


removes frinx.bfd.bfdtop.Bfd

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

# **DeleteFrinxOpenconfigInterfacesInterfacesInterfaceAggregationBfdConfig**
> DeleteFrinxOpenconfigInterfacesInterfacesInterfaceAggregationBfdConfig(ctx, name, nodeId)


removes frinx.bfd.bfdtop.bfd.Config

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

# **GetFrinxOpenconfigInterfacesInterfacesInterfaceAggregationBfd**
> FrinxBfdBfdtopBfdResponse GetFrinxOpenconfigInterfacesInterfacesInterfaceAggregationBfd(ctx, name, nodeId)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| Id of interface | 
  **nodeId** | **string**|  | 

### Return type

[**FrinxBfdBfdtopBfdResponse**](frinx.bfd.bfdtop.Bfd.response.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetFrinxOpenconfigInterfacesInterfacesInterfaceAggregationBfdConfig**
> FrinxBfdBfdtopBfdConfigResponse GetFrinxOpenconfigInterfacesInterfacesInterfaceAggregationBfdConfig(ctx, name, nodeId)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| Id of interface | 
  **nodeId** | **string**|  | 

### Return type

[**FrinxBfdBfdtopBfdConfigResponse**](frinx.bfd.bfdtop.bfd.Config.response.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PutFrinxOpenconfigInterfacesInterfacesInterfaceAggregationBfd**
> PutFrinxOpenconfigInterfacesInterfacesInterfaceAggregationBfd(ctx, name, frinxBfdBfdtopBfdBodyParam, nodeId)


creates or updates frinx.bfd.bfdtop.Bfd

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| Id of interface | 
  **frinxBfdBfdtopBfdBodyParam** | [**FrinxBfdBfdtopBfdRequest**](FrinxBfdBfdtopBfdRequest.md)| frinx.bfd.bfdtop.Bfd to be added or updated | 
  **nodeId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PutFrinxOpenconfigInterfacesInterfacesInterfaceAggregationBfdConfig**
> PutFrinxOpenconfigInterfacesInterfacesInterfaceAggregationBfdConfig(ctx, name, frinxBfdBfdtopBfdConfigBodyParam, nodeId)


creates or updates frinx.bfd.bfdtop.bfd.Config

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| Id of interface | 
  **frinxBfdBfdtopBfdConfigBodyParam** | [**FrinxBfdBfdtopBfdConfigRequest**](FrinxBfdBfdtopBfdConfigRequest.md)| frinx.bfd.bfdtop.bfd.Config to be added or updated | 
  **nodeId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

