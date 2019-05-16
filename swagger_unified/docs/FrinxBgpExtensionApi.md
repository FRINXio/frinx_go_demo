# \FrinxBgpExtensionApi

All URIs are relative to *http://172.17.0.1:8181/restconf*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteFrinxOpenconfigNetworkInstanceNetworkInstancesNetworkInstanceProtocolsProtocolBgpNeighborsNeighborAfiSafisAfiSafiConfigSoftReconfiguration**](FrinxBgpExtensionApi.md#DeleteFrinxOpenconfigNetworkInstanceNetworkInstancesNetworkInstanceProtocolsProtocolBgpNeighborsNeighborAfiSafisAfiSafiConfigSoftReconfiguration) | **Delete** /config/network-topology:network-topology/network-topology:topology/unified/network-topology:node/{node-id}/yang-ext:mount/frinx-openconfig-network-instance:network-instances/frinx-openconfig-network-instance:network-instance/{name}/frinx-openconfig-network-instance:protocols/frinx-openconfig-network-instance:protocol/{identifier}/{protocol-name}/frinx-openconfig-network-instance:bgp/frinx-openconfig-network-instance:neighbors/frinx-openconfig-network-instance:neighbor/{neighbor-address}/frinx-openconfig-network-instance:afi-safis/frinx-openconfig-network-instance:afi-safi/{afi-safi-name}/frinx-openconfig-network-instance:config/frinx-bgp-extension:soft-reconfiguration/ | 
[**GetFrinxOpenconfigNetworkInstanceNetworkInstancesNetworkInstanceProtocolsProtocolBgpNeighborsNeighborAfiSafisAfiSafiConfigSoftReconfiguration**](FrinxBgpExtensionApi.md#GetFrinxOpenconfigNetworkInstanceNetworkInstancesNetworkInstanceProtocolsProtocolBgpNeighborsNeighborAfiSafisAfiSafiConfigSoftReconfiguration) | **Get** /config/network-topology:network-topology/network-topology:topology/unified/network-topology:node/{node-id}/yang-ext:mount/frinx-openconfig-network-instance:network-instances/frinx-openconfig-network-instance:network-instance/{name}/frinx-openconfig-network-instance:protocols/frinx-openconfig-network-instance:protocol/{identifier}/{protocol-name}/frinx-openconfig-network-instance:bgp/frinx-openconfig-network-instance:neighbors/frinx-openconfig-network-instance:neighbor/{neighbor-address}/frinx-openconfig-network-instance:afi-safis/frinx-openconfig-network-instance:afi-safi/{afi-safi-name}/frinx-openconfig-network-instance:config/frinx-bgp-extension:soft-reconfiguration/ | 
[**PutFrinxOpenconfigNetworkInstanceNetworkInstancesNetworkInstanceProtocolsProtocolBgpNeighborsNeighborAfiSafisAfiSafiConfigSoftReconfiguration**](FrinxBgpExtensionApi.md#PutFrinxOpenconfigNetworkInstanceNetworkInstancesNetworkInstanceProtocolsProtocolBgpNeighborsNeighborAfiSafisAfiSafiConfigSoftReconfiguration) | **Put** /config/network-topology:network-topology/network-topology:topology/unified/network-topology:node/{node-id}/yang-ext:mount/frinx-openconfig-network-instance:network-instances/frinx-openconfig-network-instance:network-instance/{name}/frinx-openconfig-network-instance:protocols/frinx-openconfig-network-instance:protocol/{identifier}/{protocol-name}/frinx-openconfig-network-instance:bgp/frinx-openconfig-network-instance:neighbors/frinx-openconfig-network-instance:neighbor/{neighbor-address}/frinx-openconfig-network-instance:afi-safis/frinx-openconfig-network-instance:afi-safi/{afi-safi-name}/frinx-openconfig-network-instance:config/frinx-bgp-extension:soft-reconfiguration/ | 


# **DeleteFrinxOpenconfigNetworkInstanceNetworkInstancesNetworkInstanceProtocolsProtocolBgpNeighborsNeighborAfiSafisAfiSafiConfigSoftReconfiguration**
> DeleteFrinxOpenconfigNetworkInstanceNetworkInstancesNetworkInstanceProtocolsProtocolBgpNeighborsNeighborAfiSafisAfiSafiConfigSoftReconfiguration(ctx, name, identifier, protocolName, neighborAddress, afiSafiName, nodeId)


removes frinx.bgp.extension.softreconfigurationgroup.SoftReconfiguration

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| Id of network-instance | 
  **identifier** | **string**| Id of protocol | 
  **protocolName** | **string**| Id of protocol | 
  **neighborAddress** | **string**| Id of neighbor | 
  **afiSafiName** | **string**| Id of afi-safi | 
  **nodeId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetFrinxOpenconfigNetworkInstanceNetworkInstancesNetworkInstanceProtocolsProtocolBgpNeighborsNeighborAfiSafisAfiSafiConfigSoftReconfiguration**
> FrinxBgpExtensionSoftreconfigurationgroupSoftReconfigurationResponse GetFrinxOpenconfigNetworkInstanceNetworkInstancesNetworkInstanceProtocolsProtocolBgpNeighborsNeighborAfiSafisAfiSafiConfigSoftReconfiguration(ctx, name, identifier, protocolName, neighborAddress, afiSafiName, nodeId)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| Id of network-instance | 
  **identifier** | **string**| Id of protocol | 
  **protocolName** | **string**| Id of protocol | 
  **neighborAddress** | **string**| Id of neighbor | 
  **afiSafiName** | **string**| Id of afi-safi | 
  **nodeId** | **string**|  | 

### Return type

[**FrinxBgpExtensionSoftreconfigurationgroupSoftReconfigurationResponse**](frinx.bgp.extension.softreconfigurationgroup.SoftReconfiguration.response.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PutFrinxOpenconfigNetworkInstanceNetworkInstancesNetworkInstanceProtocolsProtocolBgpNeighborsNeighborAfiSafisAfiSafiConfigSoftReconfiguration**
> PutFrinxOpenconfigNetworkInstanceNetworkInstancesNetworkInstanceProtocolsProtocolBgpNeighborsNeighborAfiSafisAfiSafiConfigSoftReconfiguration(ctx, name, identifier, protocolName, neighborAddress, afiSafiName, frinxBgpExtensionSoftreconfigurationgroupSoftReconfigurationBodyParam, nodeId)


creates or updates frinx.bgp.extension.softreconfigurationgroup.SoftReconfiguration

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| Id of network-instance | 
  **identifier** | **string**| Id of protocol | 
  **protocolName** | **string**| Id of protocol | 
  **neighborAddress** | **string**| Id of neighbor | 
  **afiSafiName** | **string**| Id of afi-safi | 
  **frinxBgpExtensionSoftreconfigurationgroupSoftReconfigurationBodyParam** | [**FrinxBgpExtensionSoftreconfigurationgroupSoftReconfigurationRequest**](FrinxBgpExtensionSoftreconfigurationgroupSoftReconfigurationRequest.md)| frinx.bgp.extension.softreconfigurationgroup.SoftReconfiguration to be added or updated | 
  **nodeId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

