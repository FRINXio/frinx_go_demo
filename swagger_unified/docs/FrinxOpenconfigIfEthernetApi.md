# \FrinxOpenconfigIfEthernetApi

All URIs are relative to *http://172.17.0.1:8181/restconf*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteFrinxOpenconfigInterfacesInterfacesInterfaceEthernet**](FrinxOpenconfigIfEthernetApi.md#DeleteFrinxOpenconfigInterfacesInterfacesInterfaceEthernet) | **Delete** /config/network-topology:network-topology/network-topology:topology/unified/network-topology:node/{node-id}/yang-ext:mount/frinx-openconfig-interfaces:interfaces/frinx-openconfig-interfaces:interface/{name}/frinx-openconfig-if-ethernet:ethernet/ | 
[**DeleteFrinxOpenconfigInterfacesInterfacesInterfaceEthernetConfig**](FrinxOpenconfigIfEthernetApi.md#DeleteFrinxOpenconfigInterfacesInterfacesInterfaceEthernetConfig) | **Delete** /config/network-topology:network-topology/network-topology:topology/unified/network-topology:node/{node-id}/yang-ext:mount/frinx-openconfig-interfaces:interfaces/frinx-openconfig-interfaces:interface/{name}/frinx-openconfig-if-ethernet:ethernet/frinx-openconfig-if-ethernet:config/ | 
[**GetFrinxOpenconfigInterfacesInterfacesInterfaceEthernet**](FrinxOpenconfigIfEthernetApi.md#GetFrinxOpenconfigInterfacesInterfacesInterfaceEthernet) | **Get** /config/network-topology:network-topology/network-topology:topology/unified/network-topology:node/{node-id}/yang-ext:mount/frinx-openconfig-interfaces:interfaces/frinx-openconfig-interfaces:interface/{name}/frinx-openconfig-if-ethernet:ethernet/ | 
[**GetFrinxOpenconfigInterfacesInterfacesInterfaceEthernetConfig**](FrinxOpenconfigIfEthernetApi.md#GetFrinxOpenconfigInterfacesInterfacesInterfaceEthernetConfig) | **Get** /config/network-topology:network-topology/network-topology:topology/unified/network-topology:node/{node-id}/yang-ext:mount/frinx-openconfig-interfaces:interfaces/frinx-openconfig-interfaces:interface/{name}/frinx-openconfig-if-ethernet:ethernet/frinx-openconfig-if-ethernet:config/ | 
[**PutFrinxOpenconfigInterfacesInterfacesInterfaceEthernet**](FrinxOpenconfigIfEthernetApi.md#PutFrinxOpenconfigInterfacesInterfacesInterfaceEthernet) | **Put** /config/network-topology:network-topology/network-topology:topology/unified/network-topology:node/{node-id}/yang-ext:mount/frinx-openconfig-interfaces:interfaces/frinx-openconfig-interfaces:interface/{name}/frinx-openconfig-if-ethernet:ethernet/ | 
[**PutFrinxOpenconfigInterfacesInterfacesInterfaceEthernetConfig**](FrinxOpenconfigIfEthernetApi.md#PutFrinxOpenconfigInterfacesInterfacesInterfaceEthernetConfig) | **Put** /config/network-topology:network-topology/network-topology:topology/unified/network-topology:node/{node-id}/yang-ext:mount/frinx-openconfig-interfaces:interfaces/frinx-openconfig-interfaces:interface/{name}/frinx-openconfig-if-ethernet:ethernet/frinx-openconfig-if-ethernet:config/ | 


# **DeleteFrinxOpenconfigInterfacesInterfacesInterfaceEthernet**
> DeleteFrinxOpenconfigInterfacesInterfacesInterfaceEthernet(ctx, name, nodeId)


removes frinx.openconfig.if.ethernet.ethernettop.Ethernet

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

# **DeleteFrinxOpenconfigInterfacesInterfacesInterfaceEthernetConfig**
> DeleteFrinxOpenconfigInterfacesInterfacesInterfaceEthernetConfig(ctx, name, nodeId)


removes frinx.openconfig.if.ethernet.ethernettop.ethernet.Config

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

# **GetFrinxOpenconfigInterfacesInterfacesInterfaceEthernet**
> FrinxOpenconfigIfEthernetEthernettopEthernetResponse GetFrinxOpenconfigInterfacesInterfacesInterfaceEthernet(ctx, name, nodeId)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| Id of interface | 
  **nodeId** | **string**|  | 

### Return type

[**FrinxOpenconfigIfEthernetEthernettopEthernetResponse**](frinx.openconfig.if.ethernet.ethernettop.Ethernet.response.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetFrinxOpenconfigInterfacesInterfacesInterfaceEthernetConfig**
> FrinxOpenconfigIfEthernetEthernettopEthernetConfigResponse GetFrinxOpenconfigInterfacesInterfacesInterfaceEthernetConfig(ctx, name, nodeId)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| Id of interface | 
  **nodeId** | **string**|  | 

### Return type

[**FrinxOpenconfigIfEthernetEthernettopEthernetConfigResponse**](frinx.openconfig.if.ethernet.ethernettop.ethernet.Config.response.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PutFrinxOpenconfigInterfacesInterfacesInterfaceEthernet**
> PutFrinxOpenconfigInterfacesInterfacesInterfaceEthernet(ctx, name, frinxOpenconfigIfEthernetEthernettopEthernetBodyParam, nodeId)


creates or updates frinx.openconfig.if.ethernet.ethernettop.Ethernet

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| Id of interface | 
  **frinxOpenconfigIfEthernetEthernettopEthernetBodyParam** | [**FrinxOpenconfigIfEthernetEthernettopEthernetRequest**](FrinxOpenconfigIfEthernetEthernettopEthernetRequest.md)| frinx.openconfig.if.ethernet.ethernettop.Ethernet to be added or updated | 
  **nodeId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PutFrinxOpenconfigInterfacesInterfacesInterfaceEthernetConfig**
> PutFrinxOpenconfigInterfacesInterfacesInterfaceEthernetConfig(ctx, name, frinxOpenconfigIfEthernetEthernettopEthernetConfigBodyParam, nodeId)


creates or updates frinx.openconfig.if.ethernet.ethernettop.ethernet.Config

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| Id of interface | 
  **frinxOpenconfigIfEthernetEthernettopEthernetConfigBodyParam** | [**FrinxOpenconfigIfEthernetEthernettopEthernetConfigRequest**](FrinxOpenconfigIfEthernetEthernettopEthernetConfigRequest.md)| frinx.openconfig.if.ethernet.ethernettop.ethernet.Config to be added or updated | 
  **nodeId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

