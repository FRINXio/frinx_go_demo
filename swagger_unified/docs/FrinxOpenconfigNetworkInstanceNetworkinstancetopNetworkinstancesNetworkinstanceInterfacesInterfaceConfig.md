# FrinxOpenconfigNetworkInstanceNetworkinstancetopNetworkinstancesNetworkinstanceInterfacesInterfaceConfig

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Optional[A unique identifier for this interface - this is expressed as a free-text string] REF:Optional.empty | [optional] [default to null]
**Subinterface** | **int64** | Optional[Reference to a subinterface -- this requires the base interface to be specified using the interface leaf in this container.  If only a reference to a base interface is requuired, this leaf should not be set.] REF:Optional.empty | [optional] [default to null]
**Interface_** | **string** | Optional[Reference to a base interface.  If a reference to a subinterface is required, this leaf must be specified to indicate the base interface.] REF:Optional.empty | [optional] [default to null]
**AssociatedAddressFamilies** | [**[]FrinxOpenconfigNetworkInstanceAssociatedAddressFamiliesIdentityref**](frinx.openconfig.network.instance.AssociatedAddressFamiliesIdentityref.md) | Optional[The address families on the subinterface which are to be associated with this network instance. When this leaf-list is empty and the network instance requires Layer 3 information the address families for which the network instance is enabled should be imported. If the value of this leaf-list is specified then the association MUST only be made for those address families that are included in the list.] REF:Optional.empty | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


