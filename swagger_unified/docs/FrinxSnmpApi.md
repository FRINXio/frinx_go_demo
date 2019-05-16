# \FrinxSnmpApi

All URIs are relative to *http://172.17.0.1:8181/restconf*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteFrinxSnmpSnmp**](FrinxSnmpApi.md#DeleteFrinxSnmpSnmp) | **Delete** /config/network-topology:network-topology/network-topology:topology/unified/network-topology:node/{node-id}/yang-ext:mount/frinx-snmp:snmp/ | 
[**DeleteFrinxSnmpSnmpInterfaces**](FrinxSnmpApi.md#DeleteFrinxSnmpSnmpInterfaces) | **Delete** /config/network-topology:network-topology/network-topology:topology/unified/network-topology:node/{node-id}/yang-ext:mount/frinx-snmp:snmp/frinx-snmp:interfaces/ | 
[**DeleteFrinxSnmpSnmpInterfacesInterface**](FrinxSnmpApi.md#DeleteFrinxSnmpSnmpInterfacesInterface) | **Delete** /config/network-topology:network-topology/network-topology:topology/unified/network-topology:node/{node-id}/yang-ext:mount/frinx-snmp:snmp/frinx-snmp:interfaces/frinx-snmp:interface/{interface-id}/ | 
[**DeleteFrinxSnmpSnmpInterfacesInterfaceConfig**](FrinxSnmpApi.md#DeleteFrinxSnmpSnmpInterfacesInterfaceConfig) | **Delete** /config/network-topology:network-topology/network-topology:topology/unified/network-topology:node/{node-id}/yang-ext:mount/frinx-snmp:snmp/frinx-snmp:interfaces/frinx-snmp:interface/{interface-id}/frinx-snmp:config/ | 
[**DeleteFrinxSnmpSnmpInterfacesInterfaceConfigEnabledTrapForEvent**](FrinxSnmpApi.md#DeleteFrinxSnmpSnmpInterfacesInterfaceConfigEnabledTrapForEvent) | **Delete** /config/network-topology:network-topology/network-topology:topology/unified/network-topology:node/{node-id}/yang-ext:mount/frinx-snmp:snmp/frinx-snmp:interfaces/frinx-snmp:interface/{interface-id}/frinx-snmp:config/frinx-snmp:enabled-trap-for-event/{event-name}/ | 
[**GetFrinxSnmpSnmp**](FrinxSnmpApi.md#GetFrinxSnmpSnmp) | **Get** /config/network-topology:network-topology/network-topology:topology/unified/network-topology:node/{node-id}/yang-ext:mount/frinx-snmp:snmp/ | 
[**GetFrinxSnmpSnmpInterfaces**](FrinxSnmpApi.md#GetFrinxSnmpSnmpInterfaces) | **Get** /config/network-topology:network-topology/network-topology:topology/unified/network-topology:node/{node-id}/yang-ext:mount/frinx-snmp:snmp/frinx-snmp:interfaces/ | 
[**GetFrinxSnmpSnmpInterfacesInterface**](FrinxSnmpApi.md#GetFrinxSnmpSnmpInterfacesInterface) | **Get** /config/network-topology:network-topology/network-topology:topology/unified/network-topology:node/{node-id}/yang-ext:mount/frinx-snmp:snmp/frinx-snmp:interfaces/frinx-snmp:interface/{interface-id}/ | 
[**GetFrinxSnmpSnmpInterfacesInterfaceConfig**](FrinxSnmpApi.md#GetFrinxSnmpSnmpInterfacesInterfaceConfig) | **Get** /config/network-topology:network-topology/network-topology:topology/unified/network-topology:node/{node-id}/yang-ext:mount/frinx-snmp:snmp/frinx-snmp:interfaces/frinx-snmp:interface/{interface-id}/frinx-snmp:config/ | 
[**GetFrinxSnmpSnmpInterfacesInterfaceConfigEnabledTrapForEvent**](FrinxSnmpApi.md#GetFrinxSnmpSnmpInterfacesInterfaceConfigEnabledTrapForEvent) | **Get** /config/network-topology:network-topology/network-topology:topology/unified/network-topology:node/{node-id}/yang-ext:mount/frinx-snmp:snmp/frinx-snmp:interfaces/frinx-snmp:interface/{interface-id}/frinx-snmp:config/frinx-snmp:enabled-trap-for-event/{event-name}/ | 
[**PutFrinxSnmpSnmp**](FrinxSnmpApi.md#PutFrinxSnmpSnmp) | **Put** /config/network-topology:network-topology/network-topology:topology/unified/network-topology:node/{node-id}/yang-ext:mount/frinx-snmp:snmp/ | 
[**PutFrinxSnmpSnmpInterfaces**](FrinxSnmpApi.md#PutFrinxSnmpSnmpInterfaces) | **Put** /config/network-topology:network-topology/network-topology:topology/unified/network-topology:node/{node-id}/yang-ext:mount/frinx-snmp:snmp/frinx-snmp:interfaces/ | 
[**PutFrinxSnmpSnmpInterfacesInterface**](FrinxSnmpApi.md#PutFrinxSnmpSnmpInterfacesInterface) | **Put** /config/network-topology:network-topology/network-topology:topology/unified/network-topology:node/{node-id}/yang-ext:mount/frinx-snmp:snmp/frinx-snmp:interfaces/frinx-snmp:interface/{interface-id}/ | 
[**PutFrinxSnmpSnmpInterfacesInterfaceConfig**](FrinxSnmpApi.md#PutFrinxSnmpSnmpInterfacesInterfaceConfig) | **Put** /config/network-topology:network-topology/network-topology:topology/unified/network-topology:node/{node-id}/yang-ext:mount/frinx-snmp:snmp/frinx-snmp:interfaces/frinx-snmp:interface/{interface-id}/frinx-snmp:config/ | 
[**PutFrinxSnmpSnmpInterfacesInterfaceConfigEnabledTrapForEvent**](FrinxSnmpApi.md#PutFrinxSnmpSnmpInterfacesInterfaceConfigEnabledTrapForEvent) | **Put** /config/network-topology:network-topology/network-topology:topology/unified/network-topology:node/{node-id}/yang-ext:mount/frinx-snmp:snmp/frinx-snmp:interfaces/frinx-snmp:interface/{interface-id}/frinx-snmp:config/frinx-snmp:enabled-trap-for-event/{event-name}/ | 


# **DeleteFrinxSnmpSnmp**
> DeleteFrinxSnmpSnmp(ctx, nodeId)


removes frinx.snmp.snmptop.Snmp

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

# **DeleteFrinxSnmpSnmpInterfaces**
> DeleteFrinxSnmpSnmpInterfaces(ctx, nodeId)


removes frinx.snmp.snmpinterfacesstructural.Interfaces

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

# **DeleteFrinxSnmpSnmpInterfacesInterface**
> DeleteFrinxSnmpSnmpInterfacesInterface(ctx, interfaceId, nodeId)


removes frinx.snmp.snmpinterfacesstructural.interfaces.Interface

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **interfaceId** | **string**| Id of interface | 
  **nodeId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteFrinxSnmpSnmpInterfacesInterfaceConfig**
> DeleteFrinxSnmpSnmpInterfacesInterfaceConfig(ctx, interfaceId, nodeId)


removes frinx.snmp.snmpinterfacesstructural.interfaces.interface.Config

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **interfaceId** | **string**| Id of interface | 
  **nodeId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteFrinxSnmpSnmpInterfacesInterfaceConfigEnabledTrapForEvent**
> DeleteFrinxSnmpSnmpInterfacesInterfaceConfigEnabledTrapForEvent(ctx, interfaceId, eventName, nodeId)


removes frinx.snmp.snmpinterfaceconfig.EnabledTrapForEvent

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **interfaceId** | **string**| Id of interface | 
  **eventName** | **string**| Id of enabled-trap-for-event | 
  **nodeId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetFrinxSnmpSnmp**
> FrinxSnmpSnmptopSnmpResponse GetFrinxSnmpSnmp(ctx, nodeId)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **nodeId** | **string**|  | 

### Return type

[**FrinxSnmpSnmptopSnmpResponse**](frinx.snmp.snmptop.Snmp.response.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetFrinxSnmpSnmpInterfaces**
> FrinxSnmpSnmpinterfacesstructuralInterfacesResponse GetFrinxSnmpSnmpInterfaces(ctx, nodeId)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **nodeId** | **string**|  | 

### Return type

[**FrinxSnmpSnmpinterfacesstructuralInterfacesResponse**](frinx.snmp.snmpinterfacesstructural.Interfaces.response.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetFrinxSnmpSnmpInterfacesInterface**
> FrinxSnmpSnmpinterfacesstructuralInterfacesInterfaceResponse GetFrinxSnmpSnmpInterfacesInterface(ctx, interfaceId, nodeId)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **interfaceId** | **string**| Id of interface | 
  **nodeId** | **string**|  | 

### Return type

[**FrinxSnmpSnmpinterfacesstructuralInterfacesInterfaceResponse**](frinx.snmp.snmpinterfacesstructural.interfaces.Interface.response.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetFrinxSnmpSnmpInterfacesInterfaceConfig**
> FrinxSnmpSnmpinterfacesstructuralInterfacesInterfaceConfigResponse GetFrinxSnmpSnmpInterfacesInterfaceConfig(ctx, interfaceId, nodeId)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **interfaceId** | **string**| Id of interface | 
  **nodeId** | **string**|  | 

### Return type

[**FrinxSnmpSnmpinterfacesstructuralInterfacesInterfaceConfigResponse**](frinx.snmp.snmpinterfacesstructural.interfaces.interface.Config.response.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetFrinxSnmpSnmpInterfacesInterfaceConfigEnabledTrapForEvent**
> FrinxSnmpSnmpinterfaceconfigEnabledTrapForEventResponse GetFrinxSnmpSnmpInterfacesInterfaceConfigEnabledTrapForEvent(ctx, interfaceId, eventName, nodeId)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **interfaceId** | **string**| Id of interface | 
  **eventName** | **string**| Id of enabled-trap-for-event | 
  **nodeId** | **string**|  | 

### Return type

[**FrinxSnmpSnmpinterfaceconfigEnabledTrapForEventResponse**](frinx.snmp.snmpinterfaceconfig.EnabledTrapForEvent.response.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PutFrinxSnmpSnmp**
> PutFrinxSnmpSnmp(ctx, frinxSnmpSnmptopSnmpBodyParam, nodeId)


creates or updates frinx.snmp.snmptop.Snmp

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **frinxSnmpSnmptopSnmpBodyParam** | [**FrinxSnmpSnmptopSnmpRequest**](FrinxSnmpSnmptopSnmpRequest.md)| frinx.snmp.snmptop.Snmp to be added or updated | 
  **nodeId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PutFrinxSnmpSnmpInterfaces**
> PutFrinxSnmpSnmpInterfaces(ctx, frinxSnmpSnmpinterfacesstructuralInterfacesBodyParam, nodeId)


creates or updates frinx.snmp.snmpinterfacesstructural.Interfaces

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **frinxSnmpSnmpinterfacesstructuralInterfacesBodyParam** | [**FrinxSnmpSnmpinterfacesstructuralInterfacesRequest**](FrinxSnmpSnmpinterfacesstructuralInterfacesRequest.md)| frinx.snmp.snmpinterfacesstructural.Interfaces to be added or updated | 
  **nodeId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PutFrinxSnmpSnmpInterfacesInterface**
> PutFrinxSnmpSnmpInterfacesInterface(ctx, interfaceId, frinxSnmpSnmpinterfacesstructuralInterfacesInterfaceBodyParam, nodeId)


creates or updates frinx.snmp.snmpinterfacesstructural.interfaces.Interface

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **interfaceId** | **string**| Id of interface | 
  **frinxSnmpSnmpinterfacesstructuralInterfacesInterfaceBodyParam** | [**FrinxSnmpSnmpinterfacesstructuralInterfacesInterfaceRequest**](FrinxSnmpSnmpinterfacesstructuralInterfacesInterfaceRequest.md)| frinx.snmp.snmpinterfacesstructural.interfaces.Interface to be added or updated | 
  **nodeId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PutFrinxSnmpSnmpInterfacesInterfaceConfig**
> PutFrinxSnmpSnmpInterfacesInterfaceConfig(ctx, interfaceId, frinxSnmpSnmpinterfacesstructuralInterfacesInterfaceConfigBodyParam, nodeId)


creates or updates frinx.snmp.snmpinterfacesstructural.interfaces.interface.Config

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **interfaceId** | **string**| Id of interface | 
  **frinxSnmpSnmpinterfacesstructuralInterfacesInterfaceConfigBodyParam** | [**FrinxSnmpSnmpinterfacesstructuralInterfacesInterfaceConfigRequest**](FrinxSnmpSnmpinterfacesstructuralInterfacesInterfaceConfigRequest.md)| frinx.snmp.snmpinterfacesstructural.interfaces.interface.Config to be added or updated | 
  **nodeId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PutFrinxSnmpSnmpInterfacesInterfaceConfigEnabledTrapForEvent**
> PutFrinxSnmpSnmpInterfacesInterfaceConfigEnabledTrapForEvent(ctx, interfaceId, eventName, frinxSnmpSnmpinterfaceconfigEnabledTrapForEventBodyParam, nodeId)


creates or updates frinx.snmp.snmpinterfaceconfig.EnabledTrapForEvent

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **interfaceId** | **string**| Id of interface | 
  **eventName** | **string**| Id of enabled-trap-for-event | 
  **frinxSnmpSnmpinterfaceconfigEnabledTrapForEventBodyParam** | [**FrinxSnmpSnmpinterfaceconfigEnabledTrapForEventRequest**](FrinxSnmpSnmpinterfaceconfigEnabledTrapForEventRequest.md)| frinx.snmp.snmpinterfaceconfig.EnabledTrapForEvent to be added or updated | 
  **nodeId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

