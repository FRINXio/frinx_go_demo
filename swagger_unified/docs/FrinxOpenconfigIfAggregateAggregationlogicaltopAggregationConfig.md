# FrinxOpenconfigIfAggregateAggregationlogicaltopAggregationConfig

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FrinxJuniperIfAggregateExtensionlinkSpeed** | [***FrinxJuniperIfAggregateExtensionLinkSpeedEnumeration**](frinx.juniper.if.aggregate.extension.LinkSpeedEnumeration.md) | Optional[Link speed of individual interface that joins the AE] REF:Optional.empty | [optional] [default to null]
**FrinxIfAggregateExtensionsystemIdMac** | **string** | Optional[The MAC address portion of the node&#39;s System ID. This is combined with the system priority to construct the 8-octet system-id] REF:Optional.empty | [optional] [default to null]
**LagType** | [***FrinxOpenconfigIfAggregateAggregationType**](frinx.openconfig.if.aggregate.AggregationType.md) | Optional[Sets the type of LAG, i.e., how it is configured / maintained] REF:Optional.empty | [optional] [default to null]
**FrinxIfAggregateExtensionmacAddress** | **string** | Optional[Assigns a MAC address to the Ethernet interface.  If not specified, the corresponding operational state leaf is expected to show the system-assigned MAC address.] REF:Optional.empty | [optional] [default to null]
**MinLinks** | **int32** | Optional[Specifies the mininum number of member interfaces that must be active for the aggregate interface to be available] REF:Optional.empty | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


