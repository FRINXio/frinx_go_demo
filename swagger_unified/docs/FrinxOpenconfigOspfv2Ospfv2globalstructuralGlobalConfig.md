# FrinxOpenconfigOspfv2Ospfv2globalstructuralGlobalConfig

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IgpShortcuts** | **bool** | Optional[When this leaf is set to true, OSPFv2 will route traffic to a remote system via any LSP to the system that is marked as shortcut eligible.] REF:Optional.empty | [optional] [default to null]
**LogAdjacencyChanges** | **bool** | Optional[When this leaf is set to true, a log message will be generated when the state of an OSPFv2 neighbour changes.] REF:Optional.empty | [optional] [default to null]
**HideTransitOnlyNetworks** | **bool** | Optional[When this leaf is set to true, do not advertise prefixes into OSPFv2 that correspond to transit interfaces, as per the behaviour discussed in RFC6860.] REF:Optional[RFC6860 - Hiding Transit-Only Networks in OSPF] | [optional] [default to null]
**SummaryRouteCostMode** | [***FrinxOpenconfigOspfv2SummaryRouteCostModeEnumeration**](frinx.openconfig.ospfv2.SummaryRouteCostModeEnumeration.md) | Optional[Specify how costs for the summary routes should be specified as per the behaviour in the original OSPF specification RFC1583, or alternatively whether the revised behaviour described in RFC2328 should be utilised] REF:Optional.empty | [optional] [default to null]
**RouterId** | **string** | Optional[A 32-bit number represented as a dotted quad assigned to each router running the OSPFv2 protocol. This number should be unique within the autonomous system] REF:Optional[rfc2828] | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


