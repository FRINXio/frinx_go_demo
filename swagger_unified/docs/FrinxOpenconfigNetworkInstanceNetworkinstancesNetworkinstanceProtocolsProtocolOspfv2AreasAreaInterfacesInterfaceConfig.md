# FrinxOpenconfigNetworkInstanceNetworkinstancesNetworkinstanceProtocolsProtocolOspfv2AreasAreaInterfacesInterfaceConfig

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MultiAreaAdjacencyPrimary** | **bool** | Optional[When the specified interface is included in more than one area&#39;s configuration, this leaf marks whether the area should be considered the primary (when the value is true). In the case that this value is false, the area is considered a secondary area.] REF:Optional.empty | [optional] [default to null]
**FrinxOspfExtensionenabled** | **bool** | Optional.empty REF:Optional.empty | [optional] [default to null]
**AuthenticationType** | **string** | Optional[The type of authentication that should be used on this interface] REF:Optional.empty | [optional] [default to null]
**Metric** | **int32** | Optional[The metric for the interface] REF:Optional.empty | [optional] [default to null]
**Id** | **string** | Optional[An operator-specified string utilised to uniquely reference this interface] REF:Optional.empty | [optional] [default to null]
**Priority** | **int32** | Optional[The local system&#39;s priority to become the designated router] REF:Optional.empty | [optional] [default to null]
**HideNetwork** | **bool** | Optional[When this leaf is set to true, the network connected to the interface should be hidden from OSPFv2 advertisements per the procedure described in RFC6860.] REF:Optional[RFC6860 - Hiding Transit-Only Networks in OSFF] | [optional] [default to null]
**Passive** | **bool** | Optional[When this leaf is set to true, the interface should be advertised within the OSPF area but OSPF adjacencies should not be established over the interface] REF:Optional.empty | [optional] [default to null]
**NetworkType** | [***FrinxOpenconfigOspfv2NetworkTypeIdentityref**](frinx.openconfig.ospfv2.NetworkTypeIdentityref.md) | Optional[The type of network that OSPFv2 should use for the specified interface.] REF:Optional.empty | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


