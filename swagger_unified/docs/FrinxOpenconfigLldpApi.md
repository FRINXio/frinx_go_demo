# \FrinxOpenconfigLldpApi

All URIs are relative to *http://172.17.0.1:8181/restconf*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteFrinxOpenconfigLldpLldp**](FrinxOpenconfigLldpApi.md#DeleteFrinxOpenconfigLldpLldp) | **Delete** /config/network-topology:network-topology/network-topology:topology/unified/network-topology:node/{node-id}/yang-ext:mount/frinx-openconfig-lldp:lldp/ | 
[**DeleteFrinxOpenconfigLldpLldpConfig**](FrinxOpenconfigLldpApi.md#DeleteFrinxOpenconfigLldpLldpConfig) | **Delete** /config/network-topology:network-topology/network-topology:topology/unified/network-topology:node/{node-id}/yang-ext:mount/frinx-openconfig-lldp:lldp/frinx-openconfig-lldp:config/ | 
[**DeleteFrinxOpenconfigLldpLldpInterfaces**](FrinxOpenconfigLldpApi.md#DeleteFrinxOpenconfigLldpLldpInterfaces) | **Delete** /config/network-topology:network-topology/network-topology:topology/unified/network-topology:node/{node-id}/yang-ext:mount/frinx-openconfig-lldp:lldp/frinx-openconfig-lldp:interfaces/ | 
[**DeleteFrinxOpenconfigLldpLldpInterfacesInterface**](FrinxOpenconfigLldpApi.md#DeleteFrinxOpenconfigLldpLldpInterfacesInterface) | **Delete** /config/network-topology:network-topology/network-topology:topology/unified/network-topology:node/{node-id}/yang-ext:mount/frinx-openconfig-lldp:lldp/frinx-openconfig-lldp:interfaces/frinx-openconfig-lldp:interface/{name}/ | 
[**DeleteFrinxOpenconfigLldpLldpInterfacesInterfaceConfig**](FrinxOpenconfigLldpApi.md#DeleteFrinxOpenconfigLldpLldpInterfacesInterfaceConfig) | **Delete** /config/network-topology:network-topology/network-topology:topology/unified/network-topology:node/{node-id}/yang-ext:mount/frinx-openconfig-lldp:lldp/frinx-openconfig-lldp:interfaces/frinx-openconfig-lldp:interface/{name}/frinx-openconfig-lldp:config/ | 
[**GetFrinxOpenconfigLldpLldp**](FrinxOpenconfigLldpApi.md#GetFrinxOpenconfigLldpLldp) | **Get** /config/network-topology:network-topology/network-topology:topology/unified/network-topology:node/{node-id}/yang-ext:mount/frinx-openconfig-lldp:lldp/ | 
[**GetFrinxOpenconfigLldpLldpConfig**](FrinxOpenconfigLldpApi.md#GetFrinxOpenconfigLldpLldpConfig) | **Get** /config/network-topology:network-topology/network-topology:topology/unified/network-topology:node/{node-id}/yang-ext:mount/frinx-openconfig-lldp:lldp/frinx-openconfig-lldp:config/ | 
[**GetFrinxOpenconfigLldpLldpInterfaces**](FrinxOpenconfigLldpApi.md#GetFrinxOpenconfigLldpLldpInterfaces) | **Get** /config/network-topology:network-topology/network-topology:topology/unified/network-topology:node/{node-id}/yang-ext:mount/frinx-openconfig-lldp:lldp/frinx-openconfig-lldp:interfaces/ | 
[**GetFrinxOpenconfigLldpLldpInterfacesInterface**](FrinxOpenconfigLldpApi.md#GetFrinxOpenconfigLldpLldpInterfacesInterface) | **Get** /config/network-topology:network-topology/network-topology:topology/unified/network-topology:node/{node-id}/yang-ext:mount/frinx-openconfig-lldp:lldp/frinx-openconfig-lldp:interfaces/frinx-openconfig-lldp:interface/{name}/ | 
[**GetFrinxOpenconfigLldpLldpInterfacesInterfaceConfig**](FrinxOpenconfigLldpApi.md#GetFrinxOpenconfigLldpLldpInterfacesInterfaceConfig) | **Get** /config/network-topology:network-topology/network-topology:topology/unified/network-topology:node/{node-id}/yang-ext:mount/frinx-openconfig-lldp:lldp/frinx-openconfig-lldp:interfaces/frinx-openconfig-lldp:interface/{name}/frinx-openconfig-lldp:config/ | 
[**PutFrinxOpenconfigLldpLldp**](FrinxOpenconfigLldpApi.md#PutFrinxOpenconfigLldpLldp) | **Put** /config/network-topology:network-topology/network-topology:topology/unified/network-topology:node/{node-id}/yang-ext:mount/frinx-openconfig-lldp:lldp/ | 
[**PutFrinxOpenconfigLldpLldpConfig**](FrinxOpenconfigLldpApi.md#PutFrinxOpenconfigLldpLldpConfig) | **Put** /config/network-topology:network-topology/network-topology:topology/unified/network-topology:node/{node-id}/yang-ext:mount/frinx-openconfig-lldp:lldp/frinx-openconfig-lldp:config/ | 
[**PutFrinxOpenconfigLldpLldpInterfaces**](FrinxOpenconfigLldpApi.md#PutFrinxOpenconfigLldpLldpInterfaces) | **Put** /config/network-topology:network-topology/network-topology:topology/unified/network-topology:node/{node-id}/yang-ext:mount/frinx-openconfig-lldp:lldp/frinx-openconfig-lldp:interfaces/ | 
[**PutFrinxOpenconfigLldpLldpInterfacesInterface**](FrinxOpenconfigLldpApi.md#PutFrinxOpenconfigLldpLldpInterfacesInterface) | **Put** /config/network-topology:network-topology/network-topology:topology/unified/network-topology:node/{node-id}/yang-ext:mount/frinx-openconfig-lldp:lldp/frinx-openconfig-lldp:interfaces/frinx-openconfig-lldp:interface/{name}/ | 
[**PutFrinxOpenconfigLldpLldpInterfacesInterfaceConfig**](FrinxOpenconfigLldpApi.md#PutFrinxOpenconfigLldpLldpInterfacesInterfaceConfig) | **Put** /config/network-topology:network-topology/network-topology:topology/unified/network-topology:node/{node-id}/yang-ext:mount/frinx-openconfig-lldp:lldp/frinx-openconfig-lldp:interfaces/frinx-openconfig-lldp:interface/{name}/frinx-openconfig-lldp:config/ | 


# **DeleteFrinxOpenconfigLldpLldp**
> DeleteFrinxOpenconfigLldpLldp(ctx, nodeId)


removes frinx.openconfig.lldp.lldptop.Lldp

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **nodeId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteFrinxOpenconfigLldpLldpConfig**
> DeleteFrinxOpenconfigLldpLldpConfig(ctx, nodeId)


removes frinx.openconfig.lldp.lldptop.lldp.Config

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **nodeId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteFrinxOpenconfigLldpLldpInterfaces**
> DeleteFrinxOpenconfigLldpLldpInterfaces(ctx, nodeId)


removes frinx.openconfig.lldp.lldpinterfacetop.Interfaces

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **nodeId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteFrinxOpenconfigLldpLldpInterfacesInterface**
> DeleteFrinxOpenconfigLldpLldpInterfacesInterface(ctx, name, nodeId)


removes frinx.openconfig.lldp.lldpinterfacetop.interfaces.Interface

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

# **DeleteFrinxOpenconfigLldpLldpInterfacesInterfaceConfig**
> DeleteFrinxOpenconfigLldpLldpInterfacesInterfaceConfig(ctx, name, nodeId)


removes frinx.openconfig.lldp.lldpinterfacetop.interfaces.interface.Config

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

# **GetFrinxOpenconfigLldpLldp**
> FrinxOpenconfigLldpLldptopLldpResponse GetFrinxOpenconfigLldpLldp(ctx, nodeId)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **nodeId** | **string**|  | 

### Return type

[**FrinxOpenconfigLldpLldptopLldpResponse**](frinx.openconfig.lldp.lldptop.Lldp.response.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetFrinxOpenconfigLldpLldpConfig**
> FrinxOpenconfigLldpLldptopLldpConfigResponse GetFrinxOpenconfigLldpLldpConfig(ctx, nodeId)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **nodeId** | **string**|  | 

### Return type

[**FrinxOpenconfigLldpLldptopLldpConfigResponse**](frinx.openconfig.lldp.lldptop.lldp.Config.response.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetFrinxOpenconfigLldpLldpInterfaces**
> FrinxOpenconfigLldpLldpinterfacetopInterfacesResponse1 GetFrinxOpenconfigLldpLldpInterfaces(ctx, nodeId)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **nodeId** | **string**|  | 

### Return type

[**FrinxOpenconfigLldpLldpinterfacetopInterfacesResponse1**](frinx.openconfig.lldp.lldpinterfacetop.Interfaces.response_1.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetFrinxOpenconfigLldpLldpInterfacesInterface**
> FrinxOpenconfigLldpLldpinterfacetopInterfacesInterfaceResponse1 GetFrinxOpenconfigLldpLldpInterfacesInterface(ctx, name, nodeId)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| Id of interface | 
  **nodeId** | **string**|  | 

### Return type

[**FrinxOpenconfigLldpLldpinterfacetopInterfacesInterfaceResponse1**](frinx.openconfig.lldp.lldpinterfacetop.interfaces.Interface.response_1.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetFrinxOpenconfigLldpLldpInterfacesInterfaceConfig**
> FrinxOpenconfigLldpLldpinterfacetopInterfacesInterfaceConfigResponse1 GetFrinxOpenconfigLldpLldpInterfacesInterfaceConfig(ctx, name, nodeId)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| Id of interface | 
  **nodeId** | **string**|  | 

### Return type

[**FrinxOpenconfigLldpLldpinterfacetopInterfacesInterfaceConfigResponse1**](frinx.openconfig.lldp.lldpinterfacetop.interfaces.interface.Config.response_1.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PutFrinxOpenconfigLldpLldp**
> PutFrinxOpenconfigLldpLldp(ctx, frinxOpenconfigLldpLldptopLldpBodyParam, nodeId)


creates or updates frinx.openconfig.lldp.lldptop.Lldp

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **frinxOpenconfigLldpLldptopLldpBodyParam** | [**FrinxOpenconfigLldpLldptopLldpRequest**](FrinxOpenconfigLldpLldptopLldpRequest.md)| frinx.openconfig.lldp.lldptop.Lldp to be added or updated | 
  **nodeId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PutFrinxOpenconfigLldpLldpConfig**
> PutFrinxOpenconfigLldpLldpConfig(ctx, frinxOpenconfigLldpLldptopLldpConfigBodyParam, nodeId)


creates or updates frinx.openconfig.lldp.lldptop.lldp.Config

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **frinxOpenconfigLldpLldptopLldpConfigBodyParam** | [**FrinxOpenconfigLldpLldptopLldpConfigRequest**](FrinxOpenconfigLldpLldptopLldpConfigRequest.md)| frinx.openconfig.lldp.lldptop.lldp.Config to be added or updated | 
  **nodeId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PutFrinxOpenconfigLldpLldpInterfaces**
> PutFrinxOpenconfigLldpLldpInterfaces(ctx, frinxOpenconfigLldpLldpinterfacetopInterfacesBodyParam, nodeId)


creates or updates frinx.openconfig.lldp.lldpinterfacetop.Interfaces

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **frinxOpenconfigLldpLldpinterfacetopInterfacesBodyParam** | [**FrinxOpenconfigLldpLldpinterfacetopInterfacesRequest1**](FrinxOpenconfigLldpLldpinterfacetopInterfacesRequest1.md)| frinx.openconfig.lldp.lldpinterfacetop.Interfaces to be added or updated | 
  **nodeId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PutFrinxOpenconfigLldpLldpInterfacesInterface**
> PutFrinxOpenconfigLldpLldpInterfacesInterface(ctx, name, frinxOpenconfigLldpLldpinterfacetopInterfacesInterfaceBodyParam, nodeId)


creates or updates frinx.openconfig.lldp.lldpinterfacetop.interfaces.Interface

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| Id of interface | 
  **frinxOpenconfigLldpLldpinterfacetopInterfacesInterfaceBodyParam** | [**FrinxOpenconfigLldpLldpinterfacetopInterfacesInterfaceRequest1**](FrinxOpenconfigLldpLldpinterfacetopInterfacesInterfaceRequest1.md)| frinx.openconfig.lldp.lldpinterfacetop.interfaces.Interface to be added or updated | 
  **nodeId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PutFrinxOpenconfigLldpLldpInterfacesInterfaceConfig**
> PutFrinxOpenconfigLldpLldpInterfacesInterfaceConfig(ctx, name, frinxOpenconfigLldpLldpinterfacetopInterfacesInterfaceConfigBodyParam, nodeId)


creates or updates frinx.openconfig.lldp.lldpinterfacetop.interfaces.interface.Config

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| Id of interface | 
  **frinxOpenconfigLldpLldpinterfacetopInterfacesInterfaceConfigBodyParam** | [**FrinxOpenconfigLldpLldpinterfacetopInterfacesInterfaceConfigRequest1**](FrinxOpenconfigLldpLldpinterfacetopInterfacesInterfaceConfigRequest1.md)| frinx.openconfig.lldp.lldpinterfacetop.interfaces.interface.Config to be added or updated | 
  **nodeId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

