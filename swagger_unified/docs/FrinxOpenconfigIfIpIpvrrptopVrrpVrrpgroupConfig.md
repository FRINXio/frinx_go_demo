# FrinxOpenconfigIfIpIpvrrptopVrrpVrrpgroupConfig

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PreemptDelay** | **int32** | Optional[Set the delay the higher priority router waits before preempting] REF:Optional.empty | [optional] [default to null]
**VirtualAddress** | **[]string** | Optional[Configure one or more virtual addresses for the VRRP group] REF:Optional.empty | [optional] [default to null]
**AdvertisementInterval** | **int32** | Optional[Sets the interval between successive VRRP advertisements -- RFC 5798 defines this as a 12-bit value expressed as 0.1 seconds, with default 100, i.e., 1 second.  Several implementation express this in units of seconds] REF:Optional.empty | [optional] [default to null]
**VirtualRouterId** | **int32** | Optional[Set the virtual router id for use by the VRRP group.  This usually also determines the virtual MAC address that is generated for the VRRP group] REF:Optional.empty | [optional] [default to null]
**FrinxOpenconfigIfIpvirtualLinkLocal** | **string** | Optional[For VRRP on IPv6 interfaces, sets the virtual link local address] REF:Optional.empty | [optional] [default to null]
**Priority** | **int32** | Optional[Specifies the sending VRRP interface&#39;s priority for the virtual router.  Higher values equal higher priority] REF:Optional.empty | [optional] [default to null]
**AcceptMode** | **bool** | Optional[Configure whether packets destined for virtual addresses are accepted even when the virtual address is not owned by the router interface] REF:Optional.empty | [optional] [default to null]
**Preempt** | **bool** | Optional[When set to true, enables preemption by a higher priority backup router of a lower priority master router] REF:Optional.empty | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


