# FrinxOpenconfigNetworkInstanceNetworkinstancesNetworkinstanceProtocolsProtocolIsisInterfacesInterfaceConfig

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FrinxIsisExtensionlevelCapability** | [***FrinxOpenconfigIsisTypesLevelType**](frinx.openconfig.isis.types.LevelType.md) | Optional[ISIS level capability(level-1, level-2,vlevel-1-2).] REF:Optional.empty | [optional] [default to null]
**CircuitType** | [***FrinxOpenconfigIsisTypesCircuitType**](frinx.openconfig.isis.types.CircuitType.md) | Optional[ISIS circuit type (p2p, broadcast).] REF:Optional.empty | [optional] [default to null]
**Passive** | **bool** | Optional[When set to true, the referenced interface is a passive interface such that it is not eligible to establish adjacencies with other systems, but is advertised into the IS-IS topology.] REF:Optional.empty | [optional] [default to null]
**InterfaceId** | **string** | Optional[Interface for which ISIS configuration is to be applied.] REF:Optional.empty | [optional] [default to null]
**Enabled** | **bool** | Optional[When set to true, the functionality within which this leaf is defined is enabled, when set to false it is explicitly disabled.] REF:Optional.empty | [optional] [default to null]
**HelloPadding** | [***FrinxOpenconfigIsisTypesHelloPaddingType**](frinx.openconfig.isis.types.HelloPaddingType.md) | Optional[This leaf controls padding type for IS-IS Hello PDUs.] REF:Optional.empty | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


