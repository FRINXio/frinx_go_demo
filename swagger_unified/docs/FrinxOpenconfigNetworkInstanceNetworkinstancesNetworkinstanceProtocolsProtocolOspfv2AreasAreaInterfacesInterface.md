# FrinxOpenconfigNetworkInstanceNetworkinstancesNetworkinstanceProtocolsProtocolOspfv2AreasAreaInterfacesInterface

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FrinxBfdExtensionbfd** | [***FrinxBfdBfdtopBfd**](frinx.bfd.bfdtop.Bfd.md) | Optional[Enclosing container for BFD interface-specific data.] REF:Optional.empty | [optional] [default to null]
**Timers** | [***FrinxOpenconfigOspfv2Ospfv2areainterfacesstructureInterfacesInterfaceTimers**](frinx.openconfig.ospfv2.ospfv2areainterfacesstructure.interfaces.interface.Timers.md) | Optional[Timers relating to OSPFv2 on the interface] REF:Optional.empty | [optional] [default to null]
**FrinxBfdExtensionenableBfd** | [***FrinxOpenconfigBfdBfdenableEnableBfd**](frinx.openconfig.bfd.bfdenable.EnableBfd.md) | Optional[Enable BFD for liveliness detection to the next-hop or neighbour.] REF:Optional.empty | [optional] [default to null]
**Neighbors** | [***FrinxOpenconfigOspfv2Ospfv2areainterfacesstructureInterfacesInterfaceNeighbors**](frinx.openconfig.ospfv2.ospfv2areainterfacesstructure.interfaces.interface.Neighbors.md) | Optional[Enclosing container for the list of neighbors that an adjacency has been established with on the interface] REF:Optional.empty | [optional] [default to null]
**LsaFilter** | [***FrinxOpenconfigOspfv2Ospfv2areainterfacesstructureInterfacesInterfaceLsaFilter**](frinx.openconfig.ospfv2.ospfv2areainterfacesstructure.interfaces.interface.LsaFilter.md) | Optional[OSPFv2 parameters relating to filtering of LSAs to neighbors the specified interface.] REF:Optional.empty | [optional] [default to null]
**Mpls** | [***FrinxOpenconfigOspfv2Ospfv2areainterfacesstructureInterfacesInterfaceMpls**](frinx.openconfig.ospfv2.ospfv2areainterfacesstructure.interfaces.interface.Mpls.md) | Optional[Configuration and operational state parameters for OSPFv2 extensions related to MPLS on the interface.] REF:Optional.empty | [optional] [default to null]
**Id** | **string** | Optional[A pointer to the identifier for the interface.] REF:Optional.empty | [optional] [default to null]
**Config** | [***FrinxOpenconfigNetworkInstanceNetworkinstancesNetworkinstanceProtocolsProtocolOspfv2AreasAreaInterfacesInterfaceConfig**](frinx.openconfig.network.instance.networkinstances.networkinstance.protocols.protocol.ospfv2.areas.area.interfaces.interface.Config.md) | Optional[Configuration parameters for the interface on which OSPFv2 is enabled] REF:Optional.empty | [optional] [default to null]
**InterfaceRef** | [***FrinxOpenconfigInterfacesInterfacerefInterfaceRef**](frinx.openconfig.interfaces.interfaceref.InterfaceRef.md) | Optional[Reference to an interface or subinterface] REF:Optional.empty | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


