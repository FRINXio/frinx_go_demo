# FrinxOpenconfigNetworkInstanceNetworkinstancetopNetworkinstancesNetworkinstanceTableconnectionsTableconnectionConfig

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DefaultImportPolicy** | [***FrinxOpenconfigRoutingPolicyDefaultPolicyType**](frinx.openconfig.routing.policy.DefaultPolicyType.md) | Optional[explicitly set a default policy if no policy definition in the import policy chain is satisfied.] REF:Optional.empty | [optional] [default to null]
**DstProtocol** | [***FrinxOpenconfigNetworkInstanceDstProtocolIdentityref**](frinx.openconfig.network.instance.DstProtocolIdentityref.md) | Optional[The destination protocol for the table connection] REF:Optional.empty | [optional] [default to null]
**SrcProtocol** | [***FrinxOpenconfigNetworkInstanceSrcProtocolIdentityref**](frinx.openconfig.network.instance.SrcProtocolIdentityref.md) | Optional[The source protocol for the table connection] REF:Optional.empty | [optional] [default to null]
**ImportPolicy** | **[]string** | Optional[list of policy names in sequence to be applied on receiving a routing update in the current context, e.g., for the current peer group, neighbor, address family, etc.] REF:Optional.empty | [optional] [default to null]
**AddressFamily** | [***FrinxOpenconfigNetworkInstanceAddressFamilyIdentityref**](frinx.openconfig.network.instance.AddressFamilyIdentityref.md) | Optional[The address family associated with the connection. This must be defined for the source protocol. The target address family is implicitly defined by the address family specified for the source protocol.] REF:Optional.empty | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


