# \FrinxCdpApi

All URIs are relative to *http://172.17.0.1:8181/restconf*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteFrinxCdpCdp**](FrinxCdpApi.md#DeleteFrinxCdpCdp) | **Delete** /config/network-topology:network-topology/network-topology:topology/unified/network-topology:node/{node-id}/yang-ext:mount/frinx-cdp:cdp/ | 
[**DeleteFrinxCdpCdpConfig**](FrinxCdpApi.md#DeleteFrinxCdpCdpConfig) | **Delete** /config/network-topology:network-topology/network-topology:topology/unified/network-topology:node/{node-id}/yang-ext:mount/frinx-cdp:cdp/frinx-cdp:config/ | 
[**DeleteFrinxCdpCdpInterfaces**](FrinxCdpApi.md#DeleteFrinxCdpCdpInterfaces) | **Delete** /config/network-topology:network-topology/network-topology:topology/unified/network-topology:node/{node-id}/yang-ext:mount/frinx-cdp:cdp/frinx-cdp:interfaces/ | 
[**DeleteFrinxCdpCdpInterfacesInterface**](FrinxCdpApi.md#DeleteFrinxCdpCdpInterfacesInterface) | **Delete** /config/network-topology:network-topology/network-topology:topology/unified/network-topology:node/{node-id}/yang-ext:mount/frinx-cdp:cdp/frinx-cdp:interfaces/frinx-cdp:interface/{name}/ | 
[**DeleteFrinxCdpCdpInterfacesInterfaceConfig**](FrinxCdpApi.md#DeleteFrinxCdpCdpInterfacesInterfaceConfig) | **Delete** /config/network-topology:network-topology/network-topology:topology/unified/network-topology:node/{node-id}/yang-ext:mount/frinx-cdp:cdp/frinx-cdp:interfaces/frinx-cdp:interface/{name}/frinx-cdp:config/ | 
[**GetFrinxCdpCdp**](FrinxCdpApi.md#GetFrinxCdpCdp) | **Get** /config/network-topology:network-topology/network-topology:topology/unified/network-topology:node/{node-id}/yang-ext:mount/frinx-cdp:cdp/ | 
[**GetFrinxCdpCdpConfig**](FrinxCdpApi.md#GetFrinxCdpCdpConfig) | **Get** /config/network-topology:network-topology/network-topology:topology/unified/network-topology:node/{node-id}/yang-ext:mount/frinx-cdp:cdp/frinx-cdp:config/ | 
[**GetFrinxCdpCdpInterfaces**](FrinxCdpApi.md#GetFrinxCdpCdpInterfaces) | **Get** /config/network-topology:network-topology/network-topology:topology/unified/network-topology:node/{node-id}/yang-ext:mount/frinx-cdp:cdp/frinx-cdp:interfaces/ | 
[**GetFrinxCdpCdpInterfacesInterface**](FrinxCdpApi.md#GetFrinxCdpCdpInterfacesInterface) | **Get** /config/network-topology:network-topology/network-topology:topology/unified/network-topology:node/{node-id}/yang-ext:mount/frinx-cdp:cdp/frinx-cdp:interfaces/frinx-cdp:interface/{name}/ | 
[**GetFrinxCdpCdpInterfacesInterfaceConfig**](FrinxCdpApi.md#GetFrinxCdpCdpInterfacesInterfaceConfig) | **Get** /config/network-topology:network-topology/network-topology:topology/unified/network-topology:node/{node-id}/yang-ext:mount/frinx-cdp:cdp/frinx-cdp:interfaces/frinx-cdp:interface/{name}/frinx-cdp:config/ | 
[**PutFrinxCdpCdp**](FrinxCdpApi.md#PutFrinxCdpCdp) | **Put** /config/network-topology:network-topology/network-topology:topology/unified/network-topology:node/{node-id}/yang-ext:mount/frinx-cdp:cdp/ | 
[**PutFrinxCdpCdpConfig**](FrinxCdpApi.md#PutFrinxCdpCdpConfig) | **Put** /config/network-topology:network-topology/network-topology:topology/unified/network-topology:node/{node-id}/yang-ext:mount/frinx-cdp:cdp/frinx-cdp:config/ | 
[**PutFrinxCdpCdpInterfaces**](FrinxCdpApi.md#PutFrinxCdpCdpInterfaces) | **Put** /config/network-topology:network-topology/network-topology:topology/unified/network-topology:node/{node-id}/yang-ext:mount/frinx-cdp:cdp/frinx-cdp:interfaces/ | 
[**PutFrinxCdpCdpInterfacesInterface**](FrinxCdpApi.md#PutFrinxCdpCdpInterfacesInterface) | **Put** /config/network-topology:network-topology/network-topology:topology/unified/network-topology:node/{node-id}/yang-ext:mount/frinx-cdp:cdp/frinx-cdp:interfaces/frinx-cdp:interface/{name}/ | 
[**PutFrinxCdpCdpInterfacesInterfaceConfig**](FrinxCdpApi.md#PutFrinxCdpCdpInterfacesInterfaceConfig) | **Put** /config/network-topology:network-topology/network-topology:topology/unified/network-topology:node/{node-id}/yang-ext:mount/frinx-cdp:cdp/frinx-cdp:interfaces/frinx-cdp:interface/{name}/frinx-cdp:config/ | 


# **DeleteFrinxCdpCdp**
> DeleteFrinxCdpCdp(ctx, nodeId)


removes frinx.cdp.cdptop.Cdp

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

# **DeleteFrinxCdpCdpConfig**
> DeleteFrinxCdpCdpConfig(ctx, nodeId)


removes frinx.cdp.cdptop.cdp.Config

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

# **DeleteFrinxCdpCdpInterfaces**
> DeleteFrinxCdpCdpInterfaces(ctx, nodeId)


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

# **DeleteFrinxCdpCdpInterfacesInterface**
> DeleteFrinxCdpCdpInterfacesInterface(ctx, name, nodeId)


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

# **DeleteFrinxCdpCdpInterfacesInterfaceConfig**
> DeleteFrinxCdpCdpInterfacesInterfaceConfig(ctx, name, nodeId)


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

# **GetFrinxCdpCdp**
> FrinxCdpCdptopCdpResponse GetFrinxCdpCdp(ctx, nodeId)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **nodeId** | **string**|  | 

### Return type

[**FrinxCdpCdptopCdpResponse**](frinx.cdp.cdptop.Cdp.response.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetFrinxCdpCdpConfig**
> FrinxCdpCdptopCdpConfigResponse GetFrinxCdpCdpConfig(ctx, nodeId)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **nodeId** | **string**|  | 

### Return type

[**FrinxCdpCdptopCdpConfigResponse**](frinx.cdp.cdptop.cdp.Config.response.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetFrinxCdpCdpInterfaces**
> FrinxOpenconfigLldpLldpinterfacetopInterfacesResponse GetFrinxCdpCdpInterfaces(ctx, nodeId)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **nodeId** | **string**|  | 

### Return type

[**FrinxOpenconfigLldpLldpinterfacetopInterfacesResponse**](frinx.openconfig.lldp.lldpinterfacetop.Interfaces.response.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetFrinxCdpCdpInterfacesInterface**
> FrinxOpenconfigLldpLldpinterfacetopInterfacesInterfaceResponse GetFrinxCdpCdpInterfacesInterface(ctx, name, nodeId)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| Id of interface | 
  **nodeId** | **string**|  | 

### Return type

[**FrinxOpenconfigLldpLldpinterfacetopInterfacesInterfaceResponse**](frinx.openconfig.lldp.lldpinterfacetop.interfaces.Interface.response.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetFrinxCdpCdpInterfacesInterfaceConfig**
> FrinxOpenconfigLldpLldpinterfacetopInterfacesInterfaceConfigResponse GetFrinxCdpCdpInterfacesInterfaceConfig(ctx, name, nodeId)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| Id of interface | 
  **nodeId** | **string**|  | 

### Return type

[**FrinxOpenconfigLldpLldpinterfacetopInterfacesInterfaceConfigResponse**](frinx.openconfig.lldp.lldpinterfacetop.interfaces.interface.Config.response.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PutFrinxCdpCdp**
> PutFrinxCdpCdp(ctx, frinxCdpCdptopCdpBodyParam, nodeId)


creates or updates frinx.cdp.cdptop.Cdp

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **frinxCdpCdptopCdpBodyParam** | [**FrinxCdpCdptopCdpRequest**](FrinxCdpCdptopCdpRequest.md)| frinx.cdp.cdptop.Cdp to be added or updated | 
  **nodeId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PutFrinxCdpCdpConfig**
> PutFrinxCdpCdpConfig(ctx, frinxCdpCdptopCdpConfigBodyParam, nodeId)


creates or updates frinx.cdp.cdptop.cdp.Config

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **frinxCdpCdptopCdpConfigBodyParam** | [**FrinxCdpCdptopCdpConfigRequest**](FrinxCdpCdptopCdpConfigRequest.md)| frinx.cdp.cdptop.cdp.Config to be added or updated | 
  **nodeId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PutFrinxCdpCdpInterfaces**
> PutFrinxCdpCdpInterfaces(ctx, frinxOpenconfigLldpLldpinterfacetopInterfacesBodyParam, nodeId)


creates or updates frinx.openconfig.lldp.lldpinterfacetop.Interfaces

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **frinxOpenconfigLldpLldpinterfacetopInterfacesBodyParam** | [**FrinxOpenconfigLldpLldpinterfacetopInterfacesRequest**](FrinxOpenconfigLldpLldpinterfacetopInterfacesRequest.md)| frinx.openconfig.lldp.lldpinterfacetop.Interfaces to be added or updated | 
  **nodeId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PutFrinxCdpCdpInterfacesInterface**
> PutFrinxCdpCdpInterfacesInterface(ctx, name, frinxOpenconfigLldpLldpinterfacetopInterfacesInterfaceBodyParam, nodeId)


creates or updates frinx.openconfig.lldp.lldpinterfacetop.interfaces.Interface

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| Id of interface | 
  **frinxOpenconfigLldpLldpinterfacetopInterfacesInterfaceBodyParam** | [**FrinxOpenconfigLldpLldpinterfacetopInterfacesInterfaceRequest**](FrinxOpenconfigLldpLldpinterfacetopInterfacesInterfaceRequest.md)| frinx.openconfig.lldp.lldpinterfacetop.interfaces.Interface to be added or updated | 
  **nodeId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PutFrinxCdpCdpInterfacesInterfaceConfig**
> PutFrinxCdpCdpInterfacesInterfaceConfig(ctx, name, frinxOpenconfigLldpLldpinterfacetopInterfacesInterfaceConfigBodyParam, nodeId)


creates or updates frinx.openconfig.lldp.lldpinterfacetop.interfaces.interface.Config

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| Id of interface | 
  **frinxOpenconfigLldpLldpinterfacetopInterfacesInterfaceConfigBodyParam** | [**FrinxOpenconfigLldpLldpinterfacetopInterfacesInterfaceConfigRequest**](FrinxOpenconfigLldpLldpinterfacetopInterfacesInterfaceConfigRequest.md)| frinx.openconfig.lldp.lldpinterfacetop.interfaces.interface.Config to be added or updated | 
  **nodeId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

