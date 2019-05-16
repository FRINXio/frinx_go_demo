# FrinxOpenconfigLocalRoutingLocalstatictopStaticroutesStaticNexthopsNexthopConfig

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Index** | **string** | Optional[An user-specified identifier utilised to uniquely reference the next-hop entry in the next-hop list. The value of this index has no semantic meaning other than for referencing the entry.] REF:Optional.empty | [optional] [default to null]
**NextHop** | **string** | Optional[The next-hop that is to be used for the static route - this may be specified as an IP address, an interface or a pre-defined next-hop type - for instance, DROP or LOCAL_LINK. When this leaf is not set, and the interface-ref value is specified for the next-hop, then the system should treat the prefix as though it is directly connected to the interface.] REF:Optional.empty | [optional] [default to null]
**Metric** | **int64** | Optional[A metric which is utilised to specify the preference of the next-hop entry when it is injected into the RIB. The lower the metric, the more preferable the prefix is. When this value is not specified the metric is inherited from the default metric utilised for static routes within the network instance that the static routes are being instantiated. When multiple next-hops are specified for a static route, the metric is utilised to determine which of the next-hops is to be installed in the RIB. When multiple next-hops have the same metric (be it specified, or simply the default) then these next-hops should all be installed in the RIB] REF:Optional.empty | [optional] [default to null]
**Recurse** | **bool** | Optional[Determines whether the next-hop should be allowed to be looked up recursively - i.e., via a RIB entry which has been installed by a routing protocol, or another static route - rather than needing to be connected directly to an interface of the local system within the current network instance. When the interface reference specified within the next-hop entry is set (i.e., is not null) then forwarding is restricted to being via the interface specified - and recursion is hence disabled.] REF:Optional.empty | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


