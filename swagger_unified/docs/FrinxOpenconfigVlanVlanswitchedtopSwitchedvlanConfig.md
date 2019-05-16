# FrinxOpenconfigVlanVlanswitchedtopSwitchedvlanConfig

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NativeVlan** | **int32** | Optional[Set the native VLAN id for untagged frames arriving on a trunk interface.  Tagged frames sent on an interface configured with a native VLAN should have their tags stripped prior to transmission. This configuration is only valid on a trunk interface.] REF:Optional.empty | [optional] [default to null]
**AccessVlan** | **int32** | Optional[Assign the access vlan to the access port.] REF:Optional.empty | [optional] [default to null]
**TrunkVlans** | **[]string** | Optional[Specify VLANs, or ranges thereof, that the interface may carry when in trunk mode.  If not specified, all VLANs are allowed on the interface. Ranges are specified in the form x..y, where x&lt;y - ranges are assumed to be inclusive (such that the VLAN range is x &lt;&#x3D; range &lt;&#x3D; y.] REF:Optional.empty | [optional] [default to null]
**InterfaceMode** | [***FrinxOpenconfigVlanTypesVlanModeType**](frinx.openconfig.vlan.types.VlanModeType.md) | Optional[Set the interface to access or trunk mode for VLANs] REF:Optional.empty | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


