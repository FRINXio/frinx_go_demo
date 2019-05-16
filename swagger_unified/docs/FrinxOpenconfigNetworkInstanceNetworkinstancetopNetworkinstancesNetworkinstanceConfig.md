# FrinxOpenconfigNetworkInstanceNetworkinstancetopNetworkinstancesNetworkinstanceConfig

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RouterId** | **string** | Optional[A identifier for the local network instance - typically used within associated routing protocols or signalling routing information in another network instance] REF:Optional.empty | [optional] [default to null]
**Name** | **string** | Optional[An operator-assigned unique name for the forwarding instance] REF:Optional.empty | [optional] [default to null]
**Description** | **string** | Optional[A free-form string to be used by the network operator to describe the function of this network instance] REF:Optional.empty | [optional] [default to null]
**EnabledAddressFamilies** | [**[]FrinxOpenconfigNetworkInstanceL3EnabledAddressFamiliesIdentityref**](frinx.openconfig.network.instance.l3.EnabledAddressFamiliesIdentityref.md) | Optional[The address families that are to be enabled for this network instance.] REF:Optional.empty | [optional] [default to null]
**Type_** | [***FrinxOpenconfigNetworkInstanceTypeIdentityref**](frinx.openconfig.network.instance.TypeIdentityref.md) | Optional[The type of network instance. The value of this leaf indicates the type of forwarding entries that should be supported by this network instance] REF:Optional.empty | [optional] [default to null]
**Enabled** | **bool** | Optional[Whether the network instance should be configured to be active on the network element] REF:Optional.empty | [optional] [default to null]
**Mtu** | **int32** | Optional[The maximum frame size which should be supported for this instance for Layer 2 frames] REF:Optional.empty | [optional] [default to null]
**RouteDistinguisher** | **string** | Optional[The route distinguisher that should be used for the local VRF or VSI instance when it is signalled via BGP.] REF:Optional.empty | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


