# \FrinxOpenconfigIfIpExtApi

All URIs are relative to *http://172.17.0.1:8181/restconf*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteFrinxOpenconfigInterfacesInterfacesInterfaceSubinterfacesSubinterfaceIpv6Autoconf**](FrinxOpenconfigIfIpExtApi.md#DeleteFrinxOpenconfigInterfacesInterfacesInterfaceSubinterfacesSubinterfaceIpv6Autoconf) | **Delete** /config/network-topology:network-topology/network-topology:topology/unified/network-topology:node/{node-id}/yang-ext:mount/frinx-openconfig-interfaces:interfaces/frinx-openconfig-interfaces:interface/{name}/frinx-openconfig-interfaces:subinterfaces/frinx-openconfig-interfaces:subinterface/{index}/frinx-openconfig-if-ip:ipv6/frinx-openconfig-if-ip-ext:autoconf/ | 
[**DeleteFrinxOpenconfigInterfacesInterfacesInterfaceSubinterfacesSubinterfaceIpv6AutoconfConfig**](FrinxOpenconfigIfIpExtApi.md#DeleteFrinxOpenconfigInterfacesInterfacesInterfaceSubinterfacesSubinterfaceIpv6AutoconfConfig) | **Delete** /config/network-topology:network-topology/network-topology:topology/unified/network-topology:node/{node-id}/yang-ext:mount/frinx-openconfig-interfaces:interfaces/frinx-openconfig-interfaces:interface/{name}/frinx-openconfig-interfaces:subinterfaces/frinx-openconfig-interfaces:subinterface/{index}/frinx-openconfig-if-ip:ipv6/frinx-openconfig-if-ip-ext:autoconf/frinx-openconfig-if-ip-ext:config/ | 
[**GetFrinxOpenconfigInterfacesInterfacesInterfaceSubinterfacesSubinterfaceIpv6Autoconf**](FrinxOpenconfigIfIpExtApi.md#GetFrinxOpenconfigInterfacesInterfacesInterfaceSubinterfacesSubinterfaceIpv6Autoconf) | **Get** /config/network-topology:network-topology/network-topology:topology/unified/network-topology:node/{node-id}/yang-ext:mount/frinx-openconfig-interfaces:interfaces/frinx-openconfig-interfaces:interface/{name}/frinx-openconfig-interfaces:subinterfaces/frinx-openconfig-interfaces:subinterface/{index}/frinx-openconfig-if-ip:ipv6/frinx-openconfig-if-ip-ext:autoconf/ | 
[**GetFrinxOpenconfigInterfacesInterfacesInterfaceSubinterfacesSubinterfaceIpv6AutoconfConfig**](FrinxOpenconfigIfIpExtApi.md#GetFrinxOpenconfigInterfacesInterfacesInterfaceSubinterfacesSubinterfaceIpv6AutoconfConfig) | **Get** /config/network-topology:network-topology/network-topology:topology/unified/network-topology:node/{node-id}/yang-ext:mount/frinx-openconfig-interfaces:interfaces/frinx-openconfig-interfaces:interface/{name}/frinx-openconfig-interfaces:subinterfaces/frinx-openconfig-interfaces:subinterface/{index}/frinx-openconfig-if-ip:ipv6/frinx-openconfig-if-ip-ext:autoconf/frinx-openconfig-if-ip-ext:config/ | 
[**PutFrinxOpenconfigInterfacesInterfacesInterfaceSubinterfacesSubinterfaceIpv6Autoconf**](FrinxOpenconfigIfIpExtApi.md#PutFrinxOpenconfigInterfacesInterfacesInterfaceSubinterfacesSubinterfaceIpv6Autoconf) | **Put** /config/network-topology:network-topology/network-topology:topology/unified/network-topology:node/{node-id}/yang-ext:mount/frinx-openconfig-interfaces:interfaces/frinx-openconfig-interfaces:interface/{name}/frinx-openconfig-interfaces:subinterfaces/frinx-openconfig-interfaces:subinterface/{index}/frinx-openconfig-if-ip:ipv6/frinx-openconfig-if-ip-ext:autoconf/ | 
[**PutFrinxOpenconfigInterfacesInterfacesInterfaceSubinterfacesSubinterfaceIpv6AutoconfConfig**](FrinxOpenconfigIfIpExtApi.md#PutFrinxOpenconfigInterfacesInterfacesInterfaceSubinterfacesSubinterfaceIpv6AutoconfConfig) | **Put** /config/network-topology:network-topology/network-topology:topology/unified/network-topology:node/{node-id}/yang-ext:mount/frinx-openconfig-interfaces:interfaces/frinx-openconfig-interfaces:interface/{name}/frinx-openconfig-interfaces:subinterfaces/frinx-openconfig-interfaces:subinterface/{index}/frinx-openconfig-if-ip:ipv6/frinx-openconfig-if-ip-ext:autoconf/frinx-openconfig-if-ip-ext:config/ | 


# **DeleteFrinxOpenconfigInterfacesInterfacesInterfaceSubinterfacesSubinterfaceIpv6Autoconf**
> DeleteFrinxOpenconfigInterfacesInterfacesInterfaceSubinterfacesSubinterfaceIpv6Autoconf(ctx, name, index, nodeId)


removes frinx.openconfig.if.ip.ext.ipv6autoconftop.Autoconf

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| Id of interface | 
  **index** | **int64**| Id of subinterface | 
  **nodeId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteFrinxOpenconfigInterfacesInterfacesInterfaceSubinterfacesSubinterfaceIpv6AutoconfConfig**
> DeleteFrinxOpenconfigInterfacesInterfacesInterfaceSubinterfacesSubinterfaceIpv6AutoconfConfig(ctx, name, index, nodeId)


removes frinx.openconfig.if.ip.ext.ipv6autoconftop.autoconf.Config

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| Id of interface | 
  **index** | **int64**| Id of subinterface | 
  **nodeId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetFrinxOpenconfigInterfacesInterfacesInterfaceSubinterfacesSubinterfaceIpv6Autoconf**
> FrinxOpenconfigIfIpExtIpv6autoconftopAutoconfResponse GetFrinxOpenconfigInterfacesInterfacesInterfaceSubinterfacesSubinterfaceIpv6Autoconf(ctx, name, index, nodeId)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| Id of interface | 
  **index** | **int64**| Id of subinterface | 
  **nodeId** | **string**|  | 

### Return type

[**FrinxOpenconfigIfIpExtIpv6autoconftopAutoconfResponse**](frinx.openconfig.if.ip.ext.ipv6autoconftop.Autoconf.response.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetFrinxOpenconfigInterfacesInterfacesInterfaceSubinterfacesSubinterfaceIpv6AutoconfConfig**
> FrinxOpenconfigIfIpExtIpv6autoconftopAutoconfConfigResponse GetFrinxOpenconfigInterfacesInterfacesInterfaceSubinterfacesSubinterfaceIpv6AutoconfConfig(ctx, name, index, nodeId)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| Id of interface | 
  **index** | **int64**| Id of subinterface | 
  **nodeId** | **string**|  | 

### Return type

[**FrinxOpenconfigIfIpExtIpv6autoconftopAutoconfConfigResponse**](frinx.openconfig.if.ip.ext.ipv6autoconftop.autoconf.Config.response.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PutFrinxOpenconfigInterfacesInterfacesInterfaceSubinterfacesSubinterfaceIpv6Autoconf**
> PutFrinxOpenconfigInterfacesInterfacesInterfaceSubinterfacesSubinterfaceIpv6Autoconf(ctx, name, index, frinxOpenconfigIfIpExtIpv6autoconftopAutoconfBodyParam, nodeId)


creates or updates frinx.openconfig.if.ip.ext.ipv6autoconftop.Autoconf

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| Id of interface | 
  **index** | **int64**| Id of subinterface | 
  **frinxOpenconfigIfIpExtIpv6autoconftopAutoconfBodyParam** | [**FrinxOpenconfigIfIpExtIpv6autoconftopAutoconfRequest**](FrinxOpenconfigIfIpExtIpv6autoconftopAutoconfRequest.md)| frinx.openconfig.if.ip.ext.ipv6autoconftop.Autoconf to be added or updated | 
  **nodeId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PutFrinxOpenconfigInterfacesInterfacesInterfaceSubinterfacesSubinterfaceIpv6AutoconfConfig**
> PutFrinxOpenconfigInterfacesInterfacesInterfaceSubinterfacesSubinterfaceIpv6AutoconfConfig(ctx, name, index, frinxOpenconfigIfIpExtIpv6autoconftopAutoconfConfigBodyParam, nodeId)


creates or updates frinx.openconfig.if.ip.ext.ipv6autoconftop.autoconf.Config

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| Id of interface | 
  **index** | **int64**| Id of subinterface | 
  **frinxOpenconfigIfIpExtIpv6autoconftopAutoconfConfigBodyParam** | [**FrinxOpenconfigIfIpExtIpv6autoconftopAutoconfConfigRequest**](FrinxOpenconfigIfIpExtIpv6autoconftopAutoconfConfigRequest.md)| frinx.openconfig.if.ip.ext.ipv6autoconftop.autoconf.Config to be added or updated | 
  **nodeId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

