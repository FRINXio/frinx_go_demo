# FrinxOpenconfigLacpLacpinterfacestopInterfacesInterfaceConfig

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SystemIdMac** | **string** | Optional[The MAC address portion of the node&#39;s System ID. This is combined with the system priority to construct the 8-octet system-id] REF:Optional.empty | [optional] [default to null]
**Name** | **string** | Optional[Reference to the interface on which LACP should be configured.   The type of the target interface must be ieee8023adLag] REF:Optional.empty | [optional] [default to null]
**Interval** | [***FrinxOpenconfigLacpLacpPeriodType**](frinx.openconfig.lacp.LacpPeriodType.md) | Optional[Set the period between LACP messages -- uses the lacp-period-type enumeration.] REF:Optional.empty | [optional] [default to null]
**LacpMode** | [***FrinxOpenconfigLacpLacpActivityType**](frinx.openconfig.lacp.LacpActivityType.md) | Optional[Indicates participant is active or passive] REF:Optional.empty | [optional] [default to null]
**SystemPriority** | **int32** | Optional[Sytem priority used by the node on this LAG interface. Lower value is higher priority for determining which node is the controlling system.] REF:Optional.empty | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


