# \FrinxDampingApi

All URIs are relative to *http://172.17.0.1:8181/restconf*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteFrinxOpenconfigInterfacesInterfacesInterfaceDamping**](FrinxDampingApi.md#DeleteFrinxOpenconfigInterfacesInterfacesInterfaceDamping) | **Delete** /config/network-topology:network-topology/network-topology:topology/unified/network-topology:node/{node-id}/yang-ext:mount/frinx-openconfig-interfaces:interfaces/frinx-openconfig-interfaces:interface/{name}/frinx-damping:damping/ | 
[**DeleteFrinxOpenconfigInterfacesInterfacesInterfaceDampingConfig**](FrinxDampingApi.md#DeleteFrinxOpenconfigInterfacesInterfacesInterfaceDampingConfig) | **Delete** /config/network-topology:network-topology/network-topology:topology/unified/network-topology:node/{node-id}/yang-ext:mount/frinx-openconfig-interfaces:interfaces/frinx-openconfig-interfaces:interface/{name}/frinx-damping:damping/frinx-damping:config/ | 
[**GetFrinxOpenconfigInterfacesInterfacesInterfaceDamping**](FrinxDampingApi.md#GetFrinxOpenconfigInterfacesInterfacesInterfaceDamping) | **Get** /config/network-topology:network-topology/network-topology:topology/unified/network-topology:node/{node-id}/yang-ext:mount/frinx-openconfig-interfaces:interfaces/frinx-openconfig-interfaces:interface/{name}/frinx-damping:damping/ | 
[**GetFrinxOpenconfigInterfacesInterfacesInterfaceDampingConfig**](FrinxDampingApi.md#GetFrinxOpenconfigInterfacesInterfacesInterfaceDampingConfig) | **Get** /config/network-topology:network-topology/network-topology:topology/unified/network-topology:node/{node-id}/yang-ext:mount/frinx-openconfig-interfaces:interfaces/frinx-openconfig-interfaces:interface/{name}/frinx-damping:damping/frinx-damping:config/ | 
[**PutFrinxOpenconfigInterfacesInterfacesInterfaceDamping**](FrinxDampingApi.md#PutFrinxOpenconfigInterfacesInterfacesInterfaceDamping) | **Put** /config/network-topology:network-topology/network-topology:topology/unified/network-topology:node/{node-id}/yang-ext:mount/frinx-openconfig-interfaces:interfaces/frinx-openconfig-interfaces:interface/{name}/frinx-damping:damping/ | 
[**PutFrinxOpenconfigInterfacesInterfacesInterfaceDampingConfig**](FrinxDampingApi.md#PutFrinxOpenconfigInterfacesInterfacesInterfaceDampingConfig) | **Put** /config/network-topology:network-topology/network-topology:topology/unified/network-topology:node/{node-id}/yang-ext:mount/frinx-openconfig-interfaces:interfaces/frinx-openconfig-interfaces:interface/{name}/frinx-damping:damping/frinx-damping:config/ | 


# **DeleteFrinxOpenconfigInterfacesInterfacesInterfaceDamping**
> DeleteFrinxOpenconfigInterfacesInterfacesInterfaceDamping(ctx, name, nodeId)


removes frinx.damping.dampingtop.Damping

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

# **DeleteFrinxOpenconfigInterfacesInterfacesInterfaceDampingConfig**
> DeleteFrinxOpenconfigInterfacesInterfacesInterfaceDampingConfig(ctx, name, nodeId)


removes frinx.damping.dampingtop.damping.Config

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

# **GetFrinxOpenconfigInterfacesInterfacesInterfaceDamping**
> FrinxDampingDampingtopDampingResponse GetFrinxOpenconfigInterfacesInterfacesInterfaceDamping(ctx, name, nodeId)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| Id of interface | 
  **nodeId** | **string**|  | 

### Return type

[**FrinxDampingDampingtopDampingResponse**](frinx.damping.dampingtop.Damping.response.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetFrinxOpenconfigInterfacesInterfacesInterfaceDampingConfig**
> FrinxDampingDampingtopDampingConfigResponse GetFrinxOpenconfigInterfacesInterfacesInterfaceDampingConfig(ctx, name, nodeId)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| Id of interface | 
  **nodeId** | **string**|  | 

### Return type

[**FrinxDampingDampingtopDampingConfigResponse**](frinx.damping.dampingtop.damping.Config.response.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PutFrinxOpenconfigInterfacesInterfacesInterfaceDamping**
> PutFrinxOpenconfigInterfacesInterfacesInterfaceDamping(ctx, name, frinxDampingDampingtopDampingBodyParam, nodeId)


creates or updates frinx.damping.dampingtop.Damping

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| Id of interface | 
  **frinxDampingDampingtopDampingBodyParam** | [**FrinxDampingDampingtopDampingRequest**](FrinxDampingDampingtopDampingRequest.md)| frinx.damping.dampingtop.Damping to be added or updated | 
  **nodeId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PutFrinxOpenconfigInterfacesInterfacesInterfaceDampingConfig**
> PutFrinxOpenconfigInterfacesInterfacesInterfaceDampingConfig(ctx, name, frinxDampingDampingtopDampingConfigBodyParam, nodeId)


creates or updates frinx.damping.dampingtop.damping.Config

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| Id of interface | 
  **frinxDampingDampingtopDampingConfigBodyParam** | [**FrinxDampingDampingtopDampingConfigRequest**](FrinxDampingDampingtopDampingConfigRequest.md)| frinx.damping.dampingtop.damping.Config to be added or updated | 
  **nodeId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

