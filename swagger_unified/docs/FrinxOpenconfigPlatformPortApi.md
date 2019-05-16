# \FrinxOpenconfigPlatformPortApi

All URIs are relative to *http://172.17.0.1:8181/restconf*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteFrinxOpenconfigPlatformComponentsComponentBreakoutMode**](FrinxOpenconfigPlatformPortApi.md#DeleteFrinxOpenconfigPlatformComponentsComponentBreakoutMode) | **Delete** /config/network-topology:network-topology/network-topology:topology/unified/network-topology:node/{node-id}/yang-ext:mount/frinx-openconfig-platform:components/frinx-openconfig-platform:component/{name}/frinx-openconfig-platform-port:breakout-mode/ | 
[**DeleteFrinxOpenconfigPlatformComponentsComponentBreakoutModeConfig**](FrinxOpenconfigPlatformPortApi.md#DeleteFrinxOpenconfigPlatformComponentsComponentBreakoutModeConfig) | **Delete** /config/network-topology:network-topology/network-topology:topology/unified/network-topology:node/{node-id}/yang-ext:mount/frinx-openconfig-platform:components/frinx-openconfig-platform:component/{name}/frinx-openconfig-platform-port:breakout-mode/frinx-openconfig-platform-port:config/ | 
[**GetFrinxOpenconfigPlatformComponentsComponentBreakoutMode**](FrinxOpenconfigPlatformPortApi.md#GetFrinxOpenconfigPlatformComponentsComponentBreakoutMode) | **Get** /config/network-topology:network-topology/network-topology:topology/unified/network-topology:node/{node-id}/yang-ext:mount/frinx-openconfig-platform:components/frinx-openconfig-platform:component/{name}/frinx-openconfig-platform-port:breakout-mode/ | 
[**GetFrinxOpenconfigPlatformComponentsComponentBreakoutModeConfig**](FrinxOpenconfigPlatformPortApi.md#GetFrinxOpenconfigPlatformComponentsComponentBreakoutModeConfig) | **Get** /config/network-topology:network-topology/network-topology:topology/unified/network-topology:node/{node-id}/yang-ext:mount/frinx-openconfig-platform:components/frinx-openconfig-platform:component/{name}/frinx-openconfig-platform-port:breakout-mode/frinx-openconfig-platform-port:config/ | 
[**PutFrinxOpenconfigPlatformComponentsComponentBreakoutMode**](FrinxOpenconfigPlatformPortApi.md#PutFrinxOpenconfigPlatformComponentsComponentBreakoutMode) | **Put** /config/network-topology:network-topology/network-topology:topology/unified/network-topology:node/{node-id}/yang-ext:mount/frinx-openconfig-platform:components/frinx-openconfig-platform:component/{name}/frinx-openconfig-platform-port:breakout-mode/ | 
[**PutFrinxOpenconfigPlatformComponentsComponentBreakoutModeConfig**](FrinxOpenconfigPlatformPortApi.md#PutFrinxOpenconfigPlatformComponentsComponentBreakoutModeConfig) | **Put** /config/network-topology:network-topology/network-topology:topology/unified/network-topology:node/{node-id}/yang-ext:mount/frinx-openconfig-platform:components/frinx-openconfig-platform:component/{name}/frinx-openconfig-platform-port:breakout-mode/frinx-openconfig-platform-port:config/ | 


# **DeleteFrinxOpenconfigPlatformComponentsComponentBreakoutMode**
> DeleteFrinxOpenconfigPlatformComponentsComponentBreakoutMode(ctx, name, nodeId)


removes frinx.openconfig.platform.port.portbreakouttop.BreakoutMode

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

# **DeleteFrinxOpenconfigPlatformComponentsComponentBreakoutModeConfig**
> DeleteFrinxOpenconfigPlatformComponentsComponentBreakoutModeConfig(ctx, name, nodeId)


removes frinx.openconfig.platform.port.portbreakouttop.breakoutmode.Config

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

# **GetFrinxOpenconfigPlatformComponentsComponentBreakoutMode**
> FrinxOpenconfigPlatformPortPortbreakouttopBreakoutModeResponse GetFrinxOpenconfigPlatformComponentsComponentBreakoutMode(ctx, name, nodeId)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| Id of component | 
  **nodeId** | **string**|  | 

### Return type

[**FrinxOpenconfigPlatformPortPortbreakouttopBreakoutModeResponse**](frinx.openconfig.platform.port.portbreakouttop.BreakoutMode.response.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetFrinxOpenconfigPlatformComponentsComponentBreakoutModeConfig**
> FrinxOpenconfigPlatformPortPortbreakouttopBreakoutmodeConfigResponse GetFrinxOpenconfigPlatformComponentsComponentBreakoutModeConfig(ctx, name, nodeId)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| Id of component | 
  **nodeId** | **string**|  | 

### Return type

[**FrinxOpenconfigPlatformPortPortbreakouttopBreakoutmodeConfigResponse**](frinx.openconfig.platform.port.portbreakouttop.breakoutmode.Config.response.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PutFrinxOpenconfigPlatformComponentsComponentBreakoutMode**
> PutFrinxOpenconfigPlatformComponentsComponentBreakoutMode(ctx, name, frinxOpenconfigPlatformPortPortbreakouttopBreakoutModeBodyParam, nodeId)


creates or updates frinx.openconfig.platform.port.portbreakouttop.BreakoutMode

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| Id of component | 
  **frinxOpenconfigPlatformPortPortbreakouttopBreakoutModeBodyParam** | [**FrinxOpenconfigPlatformPortPortbreakouttopBreakoutModeRequest**](FrinxOpenconfigPlatformPortPortbreakouttopBreakoutModeRequest.md)| frinx.openconfig.platform.port.portbreakouttop.BreakoutMode to be added or updated | 
  **nodeId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PutFrinxOpenconfigPlatformComponentsComponentBreakoutModeConfig**
> PutFrinxOpenconfigPlatformComponentsComponentBreakoutModeConfig(ctx, name, frinxOpenconfigPlatformPortPortbreakouttopBreakoutmodeConfigBodyParam, nodeId)


creates or updates frinx.openconfig.platform.port.portbreakouttop.breakoutmode.Config

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| Id of component | 
  **frinxOpenconfigPlatformPortPortbreakouttopBreakoutmodeConfigBodyParam** | [**FrinxOpenconfigPlatformPortPortbreakouttopBreakoutmodeConfigRequest**](FrinxOpenconfigPlatformPortPortbreakouttopBreakoutmodeConfigRequest.md)| frinx.openconfig.platform.port.portbreakouttop.breakoutmode.Config to be added or updated | 
  **nodeId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

