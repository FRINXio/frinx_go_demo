# FrinxOpenconfigNetworkInstanceNetworkinstancesNetworkinstanceProtocolsProtocolConfig

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FrinxOspfExtensionexportPolicy** | **[]string** | Optional[Name of policy which needs to be exported.] REF:Optional.empty | [optional] [default to null]
**Name** | **string** | Optional[A unique name for the protocol instance] REF:Optional.empty | [optional] [default to null]
**Identifier** | [***FrinxOpenconfigNetworkInstanceIdentifierIdentityref**](frinx.openconfig.network.instance.IdentifierIdentityref.md) | Optional[The protocol identifier for the instance] REF:Optional.empty | [optional] [default to null]
**DefaultMetric** | **int64** | Optional[The default metric within the RIB for entries that are installed by this protocol instance. This value may be overridden by protocol specific configuration options. The lower the metric specified the more preferable the RIB entry is to be selected for use within the network instance. Where multiple entries have the same metric value then these equal cost paths should be treated according to the specified ECMP path selection behaviour for the instance] REF:Optional.empty | [optional] [default to null]
**Enabled** | **bool** | Optional[A boolean value indicating whether the local protocol instance is enabled.] REF:Optional.empty | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


