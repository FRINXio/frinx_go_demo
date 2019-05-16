# FrinxOpenconfigNetworkInstanceNetworkinstancesNetworkinstanceProtocolsProtocolLocalaggregatesAggregateConfig

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SetTag** | **string** | Optional[Set a generic tag value on the route. This tag can be used for filtering routes that are distributed to other routing protocols.] REF:Optional.empty | [optional] [default to null]
**Discard** | **bool** | Optional[When true, install the aggregate route with a discard next-hop -- traffic destined to the aggregate will be discarded with no ICMP message generated.  When false, traffic destined to an aggregate address when no constituent routes are present will generate an ICMP unreachable message.] REF:Optional.empty | [optional] [default to null]
**FrinxBgpExtensionsummaryOnly** | **bool** | Optional.empty REF:Optional.empty | [optional] [default to null]
**FrinxBgpExtensionapplyPolicy** | **[]string** | Optional[Name of policy which needs to be applied.] REF:Optional.empty | [optional] [default to null]
**Prefix** | **string** | Optional[Aggregate prefix to be advertised] REF:Optional.empty | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


