# \FrinxL3ipvlanApi

All URIs are relative to *http://172.17.0.1:8181/restconf*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteFrinxOpenconfigInterfacesInterfacesInterfaceL3ipvlan**](FrinxL3ipvlanApi.md#DeleteFrinxOpenconfigInterfacesInterfacesInterfaceL3ipvlan) | **Delete** /config/network-topology:network-topology/network-topology:topology/unified/network-topology:node/{node-id}/yang-ext:mount/frinx-openconfig-interfaces:interfaces/frinx-openconfig-interfaces:interface/{name}/frinx-l3ipvlan:l3ipvlan/ | 
[**DeleteFrinxOpenconfigInterfacesInterfacesInterfaceL3ipvlanConfig**](FrinxL3ipvlanApi.md#DeleteFrinxOpenconfigInterfacesInterfacesInterfaceL3ipvlanConfig) | **Delete** /config/network-topology:network-topology/network-topology:topology/unified/network-topology:node/{node-id}/yang-ext:mount/frinx-openconfig-interfaces:interfaces/frinx-openconfig-interfaces:interface/{name}/frinx-l3ipvlan:l3ipvlan/frinx-l3ipvlan:config/ | 
[**GetFrinxOpenconfigInterfacesInterfacesInterfaceL3ipvlan**](FrinxL3ipvlanApi.md#GetFrinxOpenconfigInterfacesInterfacesInterfaceL3ipvlan) | **Get** /config/network-topology:network-topology/network-topology:topology/unified/network-topology:node/{node-id}/yang-ext:mount/frinx-openconfig-interfaces:interfaces/frinx-openconfig-interfaces:interface/{name}/frinx-l3ipvlan:l3ipvlan/ | 
[**GetFrinxOpenconfigInterfacesInterfacesInterfaceL3ipvlanConfig**](FrinxL3ipvlanApi.md#GetFrinxOpenconfigInterfacesInterfacesInterfaceL3ipvlanConfig) | **Get** /config/network-topology:network-topology/network-topology:topology/unified/network-topology:node/{node-id}/yang-ext:mount/frinx-openconfig-interfaces:interfaces/frinx-openconfig-interfaces:interface/{name}/frinx-l3ipvlan:l3ipvlan/frinx-l3ipvlan:config/ | 
[**PutFrinxOpenconfigInterfacesInterfacesInterfaceL3ipvlan**](FrinxL3ipvlanApi.md#PutFrinxOpenconfigInterfacesInterfacesInterfaceL3ipvlan) | **Put** /config/network-topology:network-topology/network-topology:topology/unified/network-topology:node/{node-id}/yang-ext:mount/frinx-openconfig-interfaces:interfaces/frinx-openconfig-interfaces:interface/{name}/frinx-l3ipvlan:l3ipvlan/ | 
[**PutFrinxOpenconfigInterfacesInterfacesInterfaceL3ipvlanConfig**](FrinxL3ipvlanApi.md#PutFrinxOpenconfigInterfacesInterfacesInterfaceL3ipvlanConfig) | **Put** /config/network-topology:network-topology/network-topology:topology/unified/network-topology:node/{node-id}/yang-ext:mount/frinx-openconfig-interfaces:interfaces/frinx-openconfig-interfaces:interface/{name}/frinx-l3ipvlan:l3ipvlan/frinx-l3ipvlan:config/ | 


# **DeleteFrinxOpenconfigInterfacesInterfacesInterfaceL3ipvlan**
> DeleteFrinxOpenconfigInterfacesInterfacesInterfaceL3ipvlan(ctx, name, nodeId)


removes frinx.l3ipvlan.l3ipvlaninterfacetop.L3ipvlan

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

# **DeleteFrinxOpenconfigInterfacesInterfacesInterfaceL3ipvlanConfig**
> DeleteFrinxOpenconfigInterfacesInterfacesInterfaceL3ipvlanConfig(ctx, name, nodeId)


removes frinx.l3ipvlan.l3ipvlaninterfacetop.l3ipvlan.Config

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

# **GetFrinxOpenconfigInterfacesInterfacesInterfaceL3ipvlan**
> FrinxL3ipvlanL3ipvlaninterfacetopL3ipvlanResponse GetFrinxOpenconfigInterfacesInterfacesInterfaceL3ipvlan(ctx, name, nodeId)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| Id of interface | 
  **nodeId** | **string**|  | 

### Return type

[**FrinxL3ipvlanL3ipvlaninterfacetopL3ipvlanResponse**](frinx.l3ipvlan.l3ipvlaninterfacetop.L3ipvlan.response.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetFrinxOpenconfigInterfacesInterfacesInterfaceL3ipvlanConfig**
> FrinxL3ipvlanL3ipvlaninterfacetopL3ipvlanConfigResponse GetFrinxOpenconfigInterfacesInterfacesInterfaceL3ipvlanConfig(ctx, name, nodeId)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| Id of interface | 
  **nodeId** | **string**|  | 

### Return type

[**FrinxL3ipvlanL3ipvlaninterfacetopL3ipvlanConfigResponse**](frinx.l3ipvlan.l3ipvlaninterfacetop.l3ipvlan.Config.response.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PutFrinxOpenconfigInterfacesInterfacesInterfaceL3ipvlan**
> PutFrinxOpenconfigInterfacesInterfacesInterfaceL3ipvlan(ctx, name, frinxL3ipvlanL3ipvlaninterfacetopL3ipvlanBodyParam, nodeId)


creates or updates frinx.l3ipvlan.l3ipvlaninterfacetop.L3ipvlan

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| Id of interface | 
  **frinxL3ipvlanL3ipvlaninterfacetopL3ipvlanBodyParam** | [**FrinxL3ipvlanL3ipvlaninterfacetopL3ipvlanRequest**](FrinxL3ipvlanL3ipvlaninterfacetopL3ipvlanRequest.md)| frinx.l3ipvlan.l3ipvlaninterfacetop.L3ipvlan to be added or updated | 
  **nodeId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PutFrinxOpenconfigInterfacesInterfacesInterfaceL3ipvlanConfig**
> PutFrinxOpenconfigInterfacesInterfacesInterfaceL3ipvlanConfig(ctx, name, frinxL3ipvlanL3ipvlaninterfacetopL3ipvlanConfigBodyParam, nodeId)


creates or updates frinx.l3ipvlan.l3ipvlaninterfacetop.l3ipvlan.Config

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| Id of interface | 
  **frinxL3ipvlanL3ipvlaninterfacetopL3ipvlanConfigBodyParam** | [**FrinxL3ipvlanL3ipvlaninterfacetopL3ipvlanConfigRequest**](FrinxL3ipvlanL3ipvlaninterfacetopL3ipvlanConfigRequest.md)| frinx.l3ipvlan.l3ipvlaninterfacetop.l3ipvlan.Config to be added or updated | 
  **nodeId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

