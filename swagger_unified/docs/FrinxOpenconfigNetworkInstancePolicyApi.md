# \FrinxOpenconfigNetworkInstancePolicyApi

All URIs are relative to *http://172.17.0.1:8181/restconf*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteFrinxOpenconfigRoutingPolicyRoutingPolicyPolicyDefinitionsPolicyDefinitionStatementsStatementConditionsMatchProtocolInstance**](FrinxOpenconfigNetworkInstancePolicyApi.md#DeleteFrinxOpenconfigRoutingPolicyRoutingPolicyPolicyDefinitionsPolicyDefinitionStatementsStatementConditionsMatchProtocolInstance) | **Delete** /config/network-topology:network-topology/network-topology:topology/unified/network-topology:node/{node-id}/yang-ext:mount/frinx-openconfig-routing-policy:routing-policy/frinx-openconfig-routing-policy:policy-definitions/frinx-openconfig-routing-policy:policy-definition/{name}/frinx-openconfig-routing-policy:statements/frinx-openconfig-routing-policy:statement/{statement-name}/frinx-openconfig-routing-policy:conditions/frinx-openconfig-network-instance-policy:match-protocol-instance/ | 
[**DeleteFrinxOpenconfigRoutingPolicyRoutingPolicyPolicyDefinitionsPolicyDefinitionStatementsStatementConditionsMatchProtocolInstanceConfig**](FrinxOpenconfigNetworkInstancePolicyApi.md#DeleteFrinxOpenconfigRoutingPolicyRoutingPolicyPolicyDefinitionsPolicyDefinitionStatementsStatementConditionsMatchProtocolInstanceConfig) | **Delete** /config/network-topology:network-topology/network-topology:topology/unified/network-topology:node/{node-id}/yang-ext:mount/frinx-openconfig-routing-policy:routing-policy/frinx-openconfig-routing-policy:policy-definitions/frinx-openconfig-routing-policy:policy-definition/{name}/frinx-openconfig-routing-policy:statements/frinx-openconfig-routing-policy:statement/{statement-name}/frinx-openconfig-routing-policy:conditions/frinx-openconfig-network-instance-policy:match-protocol-instance/frinx-openconfig-network-instance-policy:config/ | 
[**GetFrinxOpenconfigRoutingPolicyRoutingPolicyPolicyDefinitionsPolicyDefinitionStatementsStatementConditionsMatchProtocolInstance**](FrinxOpenconfigNetworkInstancePolicyApi.md#GetFrinxOpenconfigRoutingPolicyRoutingPolicyPolicyDefinitionsPolicyDefinitionStatementsStatementConditionsMatchProtocolInstance) | **Get** /config/network-topology:network-topology/network-topology:topology/unified/network-topology:node/{node-id}/yang-ext:mount/frinx-openconfig-routing-policy:routing-policy/frinx-openconfig-routing-policy:policy-definitions/frinx-openconfig-routing-policy:policy-definition/{name}/frinx-openconfig-routing-policy:statements/frinx-openconfig-routing-policy:statement/{statement-name}/frinx-openconfig-routing-policy:conditions/frinx-openconfig-network-instance-policy:match-protocol-instance/ | 
[**GetFrinxOpenconfigRoutingPolicyRoutingPolicyPolicyDefinitionsPolicyDefinitionStatementsStatementConditionsMatchProtocolInstanceConfig**](FrinxOpenconfigNetworkInstancePolicyApi.md#GetFrinxOpenconfigRoutingPolicyRoutingPolicyPolicyDefinitionsPolicyDefinitionStatementsStatementConditionsMatchProtocolInstanceConfig) | **Get** /config/network-topology:network-topology/network-topology:topology/unified/network-topology:node/{node-id}/yang-ext:mount/frinx-openconfig-routing-policy:routing-policy/frinx-openconfig-routing-policy:policy-definitions/frinx-openconfig-routing-policy:policy-definition/{name}/frinx-openconfig-routing-policy:statements/frinx-openconfig-routing-policy:statement/{statement-name}/frinx-openconfig-routing-policy:conditions/frinx-openconfig-network-instance-policy:match-protocol-instance/frinx-openconfig-network-instance-policy:config/ | 
[**PutFrinxOpenconfigRoutingPolicyRoutingPolicyPolicyDefinitionsPolicyDefinitionStatementsStatementConditionsMatchProtocolInstance**](FrinxOpenconfigNetworkInstancePolicyApi.md#PutFrinxOpenconfigRoutingPolicyRoutingPolicyPolicyDefinitionsPolicyDefinitionStatementsStatementConditionsMatchProtocolInstance) | **Put** /config/network-topology:network-topology/network-topology:topology/unified/network-topology:node/{node-id}/yang-ext:mount/frinx-openconfig-routing-policy:routing-policy/frinx-openconfig-routing-policy:policy-definitions/frinx-openconfig-routing-policy:policy-definition/{name}/frinx-openconfig-routing-policy:statements/frinx-openconfig-routing-policy:statement/{statement-name}/frinx-openconfig-routing-policy:conditions/frinx-openconfig-network-instance-policy:match-protocol-instance/ | 
[**PutFrinxOpenconfigRoutingPolicyRoutingPolicyPolicyDefinitionsPolicyDefinitionStatementsStatementConditionsMatchProtocolInstanceConfig**](FrinxOpenconfigNetworkInstancePolicyApi.md#PutFrinxOpenconfigRoutingPolicyRoutingPolicyPolicyDefinitionsPolicyDefinitionStatementsStatementConditionsMatchProtocolInstanceConfig) | **Put** /config/network-topology:network-topology/network-topology:topology/unified/network-topology:node/{node-id}/yang-ext:mount/frinx-openconfig-routing-policy:routing-policy/frinx-openconfig-routing-policy:policy-definitions/frinx-openconfig-routing-policy:policy-definition/{name}/frinx-openconfig-routing-policy:statements/frinx-openconfig-routing-policy:statement/{statement-name}/frinx-openconfig-routing-policy:conditions/frinx-openconfig-network-instance-policy:match-protocol-instance/frinx-openconfig-network-instance-policy:config/ | 


# **DeleteFrinxOpenconfigRoutingPolicyRoutingPolicyPolicyDefinitionsPolicyDefinitionStatementsStatementConditionsMatchProtocolInstance**
> DeleteFrinxOpenconfigRoutingPolicyRoutingPolicyPolicyDefinitionsPolicyDefinitionStatementsStatementConditionsMatchProtocolInstance(ctx, name, statementName, nodeId)


removes frinx.openconfig.network.instance.policy.protocolinstancepolicytop.MatchProtocolInstance

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| Id of policy-definition | 
  **statementName** | **string**| Id of statement | 
  **nodeId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteFrinxOpenconfigRoutingPolicyRoutingPolicyPolicyDefinitionsPolicyDefinitionStatementsStatementConditionsMatchProtocolInstanceConfig**
> DeleteFrinxOpenconfigRoutingPolicyRoutingPolicyPolicyDefinitionsPolicyDefinitionStatementsStatementConditionsMatchProtocolInstanceConfig(ctx, name, statementName, nodeId)


removes frinx.openconfig.network.instance.policy.protocolinstancepolicytop.matchprotocolinstance.Config

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| Id of policy-definition | 
  **statementName** | **string**| Id of statement | 
  **nodeId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetFrinxOpenconfigRoutingPolicyRoutingPolicyPolicyDefinitionsPolicyDefinitionStatementsStatementConditionsMatchProtocolInstance**
> FrinxOpenconfigNetworkInstancePolicyProtocolinstancepolicytopMatchProtocolInstanceResponse GetFrinxOpenconfigRoutingPolicyRoutingPolicyPolicyDefinitionsPolicyDefinitionStatementsStatementConditionsMatchProtocolInstance(ctx, name, statementName, nodeId)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| Id of policy-definition | 
  **statementName** | **string**| Id of statement | 
  **nodeId** | **string**|  | 

### Return type

[**FrinxOpenconfigNetworkInstancePolicyProtocolinstancepolicytopMatchProtocolInstanceResponse**](frinx.openconfig.network.instance.policy.protocolinstancepolicytop.MatchProtocolInstance.response.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetFrinxOpenconfigRoutingPolicyRoutingPolicyPolicyDefinitionsPolicyDefinitionStatementsStatementConditionsMatchProtocolInstanceConfig**
> FrinxOpenconfigNetworkInstancePolicyProtocolinstancepolicytopMatchprotocolinstanceConfigResponse GetFrinxOpenconfigRoutingPolicyRoutingPolicyPolicyDefinitionsPolicyDefinitionStatementsStatementConditionsMatchProtocolInstanceConfig(ctx, name, statementName, nodeId)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| Id of policy-definition | 
  **statementName** | **string**| Id of statement | 
  **nodeId** | **string**|  | 

### Return type

[**FrinxOpenconfigNetworkInstancePolicyProtocolinstancepolicytopMatchprotocolinstanceConfigResponse**](frinx.openconfig.network.instance.policy.protocolinstancepolicytop.matchprotocolinstance.Config.response.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PutFrinxOpenconfigRoutingPolicyRoutingPolicyPolicyDefinitionsPolicyDefinitionStatementsStatementConditionsMatchProtocolInstance**
> PutFrinxOpenconfigRoutingPolicyRoutingPolicyPolicyDefinitionsPolicyDefinitionStatementsStatementConditionsMatchProtocolInstance(ctx, name, statementName, frinxOpenconfigNetworkInstancePolicyProtocolinstancepolicytopMatchProtocolInstanceBodyParam, nodeId)


creates or updates frinx.openconfig.network.instance.policy.protocolinstancepolicytop.MatchProtocolInstance

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| Id of policy-definition | 
  **statementName** | **string**| Id of statement | 
  **frinxOpenconfigNetworkInstancePolicyProtocolinstancepolicytopMatchProtocolInstanceBodyParam** | [**FrinxOpenconfigNetworkInstancePolicyProtocolinstancepolicytopMatchProtocolInstanceRequest**](FrinxOpenconfigNetworkInstancePolicyProtocolinstancepolicytopMatchProtocolInstanceRequest.md)| frinx.openconfig.network.instance.policy.protocolinstancepolicytop.MatchProtocolInstance to be added or updated | 
  **nodeId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PutFrinxOpenconfigRoutingPolicyRoutingPolicyPolicyDefinitionsPolicyDefinitionStatementsStatementConditionsMatchProtocolInstanceConfig**
> PutFrinxOpenconfigRoutingPolicyRoutingPolicyPolicyDefinitionsPolicyDefinitionStatementsStatementConditionsMatchProtocolInstanceConfig(ctx, name, statementName, frinxOpenconfigNetworkInstancePolicyProtocolinstancepolicytopMatchprotocolinstanceConfigBodyParam, nodeId)


creates or updates frinx.openconfig.network.instance.policy.protocolinstancepolicytop.matchprotocolinstance.Config

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| Id of policy-definition | 
  **statementName** | **string**| Id of statement | 
  **frinxOpenconfigNetworkInstancePolicyProtocolinstancepolicytopMatchprotocolinstanceConfigBodyParam** | [**FrinxOpenconfigNetworkInstancePolicyProtocolinstancepolicytopMatchprotocolinstanceConfigRequest**](FrinxOpenconfigNetworkInstancePolicyProtocolinstancepolicytopMatchprotocolinstanceConfigRequest.md)| frinx.openconfig.network.instance.policy.protocolinstancepolicytop.matchprotocolinstance.Config to be added or updated | 
  **nodeId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

