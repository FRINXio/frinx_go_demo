# \FrinxOpenconfigPlatformLinecardApi

All URIs are relative to *http://172.17.0.1:8181/restconf*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteFrinxOpenconfigPlatformComponentsComponentLinecard**](FrinxOpenconfigPlatformLinecardApi.md#DeleteFrinxOpenconfigPlatformComponentsComponentLinecard) | **Delete** /config/network-topology:network-topology/network-topology:topology/unified/network-topology:node/{node-id}/yang-ext:mount/frinx-openconfig-platform:components/frinx-openconfig-platform:component/{name}/frinx-openconfig-platform-linecard:linecard/ | 
[**DeleteFrinxOpenconfigPlatformComponentsComponentLinecardConfig**](FrinxOpenconfigPlatformLinecardApi.md#DeleteFrinxOpenconfigPlatformComponentsComponentLinecardConfig) | **Delete** /config/network-topology:network-topology/network-topology:topology/unified/network-topology:node/{node-id}/yang-ext:mount/frinx-openconfig-platform:components/frinx-openconfig-platform:component/{name}/frinx-openconfig-platform-linecard:linecard/frinx-openconfig-platform-linecard:config/ | 
[**GetFrinxOpenconfigPlatformComponentsComponentLinecard**](FrinxOpenconfigPlatformLinecardApi.md#GetFrinxOpenconfigPlatformComponentsComponentLinecard) | **Get** /config/network-topology:network-topology/network-topology:topology/unified/network-topology:node/{node-id}/yang-ext:mount/frinx-openconfig-platform:components/frinx-openconfig-platform:component/{name}/frinx-openconfig-platform-linecard:linecard/ | 
[**GetFrinxOpenconfigPlatformComponentsComponentLinecardConfig**](FrinxOpenconfigPlatformLinecardApi.md#GetFrinxOpenconfigPlatformComponentsComponentLinecardConfig) | **Get** /config/network-topology:network-topology/network-topology:topology/unified/network-topology:node/{node-id}/yang-ext:mount/frinx-openconfig-platform:components/frinx-openconfig-platform:component/{name}/frinx-openconfig-platform-linecard:linecard/frinx-openconfig-platform-linecard:config/ | 
[**PutFrinxOpenconfigPlatformComponentsComponentLinecard**](FrinxOpenconfigPlatformLinecardApi.md#PutFrinxOpenconfigPlatformComponentsComponentLinecard) | **Put** /config/network-topology:network-topology/network-topology:topology/unified/network-topology:node/{node-id}/yang-ext:mount/frinx-openconfig-platform:components/frinx-openconfig-platform:component/{name}/frinx-openconfig-platform-linecard:linecard/ | 
[**PutFrinxOpenconfigPlatformComponentsComponentLinecardConfig**](FrinxOpenconfigPlatformLinecardApi.md#PutFrinxOpenconfigPlatformComponentsComponentLinecardConfig) | **Put** /config/network-topology:network-topology/network-topology:topology/unified/network-topology:node/{node-id}/yang-ext:mount/frinx-openconfig-platform:components/frinx-openconfig-platform:component/{name}/frinx-openconfig-platform-linecard:linecard/frinx-openconfig-platform-linecard:config/ | 


# **DeleteFrinxOpenconfigPlatformComponentsComponentLinecard**
> DeleteFrinxOpenconfigPlatformComponentsComponentLinecard(ctx, name, nodeId)


removes frinx.openconfig.platform.linecard.linecardtop.Linecard

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| Id of component | 
  **nodeId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteFrinxOpenconfigPlatformComponentsComponentLinecardConfig**
> DeleteFrinxOpenconfigPlatformComponentsComponentLinecardConfig(ctx, name, nodeId)


removes frinx.openconfig.platform.linecard.linecardtop.linecard.Config

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| Id of component | 
  **nodeId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetFrinxOpenconfigPlatformComponentsComponentLinecard**
> FrinxOpenconfigPlatformLinecardLinecardtopLinecardResponse GetFrinxOpenconfigPlatformComponentsComponentLinecard(ctx, name, nodeId)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| Id of component | 
  **nodeId** | **string**|  | 

### Return type

[**FrinxOpenconfigPlatformLinecardLinecardtopLinecardResponse**](frinx.openconfig.platform.linecard.linecardtop.Linecard.response.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetFrinxOpenconfigPlatformComponentsComponentLinecardConfig**
> FrinxOpenconfigPlatformLinecardLinecardtopLinecardConfigResponse GetFrinxOpenconfigPlatformComponentsComponentLinecardConfig(ctx, name, nodeId)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| Id of component | 
  **nodeId** | **string**|  | 

### Return type

[**FrinxOpenconfigPlatformLinecardLinecardtopLinecardConfigResponse**](frinx.openconfig.platform.linecard.linecardtop.linecard.Config.response.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PutFrinxOpenconfigPlatformComponentsComponentLinecard**
> PutFrinxOpenconfigPlatformComponentsComponentLinecard(ctx, name, frinxOpenconfigPlatformLinecardLinecardtopLinecardBodyParam, nodeId)


creates or updates frinx.openconfig.platform.linecard.linecardtop.Linecard

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| Id of component | 
  **frinxOpenconfigPlatformLinecardLinecardtopLinecardBodyParam** | [**FrinxOpenconfigPlatformLinecardLinecardtopLinecardRequest**](FrinxOpenconfigPlatformLinecardLinecardtopLinecardRequest.md)| frinx.openconfig.platform.linecard.linecardtop.Linecard to be added or updated | 
  **nodeId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PutFrinxOpenconfigPlatformComponentsComponentLinecardConfig**
> PutFrinxOpenconfigPlatformComponentsComponentLinecardConfig(ctx, name, frinxOpenconfigPlatformLinecardLinecardtopLinecardConfigBodyParam, nodeId)


creates or updates frinx.openconfig.platform.linecard.linecardtop.linecard.Config

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| Id of component | 
  **frinxOpenconfigPlatformLinecardLinecardtopLinecardConfigBodyParam** | [**FrinxOpenconfigPlatformLinecardLinecardtopLinecardConfigRequest**](FrinxOpenconfigPlatformLinecardLinecardtopLinecardConfigRequest.md)| frinx.openconfig.platform.linecard.linecardtop.linecard.Config to be added or updated | 
  **nodeId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

