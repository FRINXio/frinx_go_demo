# FrinxOpenconfigOspfv2Ospfv2globalstructuralGlobalTimersMaxmetricConfig

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Include** | [**[]FrinxOpenconfigOspfv2IncludeIdentityref**](frinx.openconfig.ospfv2.IncludeIdentityref.md) | Optional[By default, the maximum metric is advertised for all non-stub interfaces of a device. When identities are specified within this leaf-list, additional entities are also advertised with the maximum metric according to the values within the list.] REF:Optional.empty | [optional] [default to null]
**Set** | **bool** | Optional[When this leaf is set to true, all non-stub interfaces of the local system are advertised with the maximum metric, such that the router does not act as a transit system, (similarly to the IS-IS overload functionality).] REF:Optional[RFC3137 - OSPF Stub Router Advertisement] | [optional] [default to null]
**Trigger** | [**[]FrinxOpenconfigOspfv2TriggerIdentityref**](frinx.openconfig.ospfv2.TriggerIdentityref.md) | Optional[By default, the maximum metric is only advertised when the max-metric/set leaf is specified as true. In the case that identities are specified within this list, they provide additional triggers (e.g., system boot) that may cause the max-metric to be set. In this case, the system should still honour the timeout specified by the max-metric/timeout leaf, and clear the max-metric advertisements after the expiration of this timer.] REF:Optional.empty | [optional] [default to null]
**Timeout** | **int32** | Optional[The delay, in seconds, after which the advertisement of entities with the maximum metric should be cleared, and the system reverts to the default, or configured, metrics.] REF:Optional.empty | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


