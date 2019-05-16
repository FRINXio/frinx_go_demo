# FrinxCiscoOspfExtensionMaxmetricsfieldsMaxmetrictimersMaxmetrictimerConfig

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Include** | [**[]FrinxCiscoOspfExtensionIncludeIdentityref**](frinx.cisco.ospf.extension.IncludeIdentityref.md) | Optional[By default, the maximum metric is advertised for all non-stub interfaces of a device. When identities are specified within this leaf-list, additional entities are also advertised with the maximum metric according to the values within the list.] REF:Optional.empty | [optional] [default to null]
**Trigger** | [***FrinxCiscoOspfExtensionTriggerIdentityref**](frinx.cisco.ospf.extension.TriggerIdentityref.md) | Optional.empty REF:Optional.empty | [optional] [default to null]
**Timeout** | **int32** | Optional[The delay, in seconds, after which the advertisement of entities with the maximum metric should be cleared, and the system reverts to the default, or configured, metrics.] REF:Optional.empty | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


