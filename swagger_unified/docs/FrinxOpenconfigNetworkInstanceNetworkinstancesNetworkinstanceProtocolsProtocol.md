# FrinxOpenconfigNetworkInstanceNetworkinstancesNetworkinstanceProtocolsProtocol

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Identifier** | [***FrinxOpenconfigNetworkInstanceIdentifierIdentityref**](frinx.openconfig.network.instance.IdentifierIdentityref.md) | Optional[The protocol name for the routing or forwarding protocol to be instantiated] REF:Optional.empty | [optional] [default to null]
**Ospfv2** | [***FrinxOpenconfigNetworkInstanceNetworkinstancesNetworkinstanceProtocolsProtocolOspfv2**](frinx.openconfig.network.instance.networkinstances.networkinstance.protocols.protocol.Ospfv2.md) | Optional[Top-level configuration and operational state for Open Shortest Path First (OSPF) v2] REF:Optional.empty | [optional] [default to null]
**Isis** | [***FrinxOpenconfigNetworkInstanceNetworkinstancesNetworkinstanceProtocolsProtocolIsis**](frinx.openconfig.network.instance.networkinstances.networkinstance.protocols.protocol.Isis.md) | Optional[This container defines top-level ISIS configuration and state information.] REF:Optional.empty | [optional] [default to null]
**Name** | **string** | Optional[An operator-assigned identifier for the routing or forwarding protocol. For some processes this leaf may be system defined.] REF:Optional.empty | [optional] [default to null]
**Bgp** | [***FrinxOpenconfigNetworkInstanceNetworkinstancesNetworkinstanceProtocolsProtocolBgp**](frinx.openconfig.network.instance.networkinstances.networkinstance.protocols.protocol.Bgp.md) | Optional[Top-level configuration and state for the BGP router] REF:Optional.empty | [optional] [default to null]
**StaticRoutes** | [***FrinxOpenconfigNetworkInstanceNetworkinstancesNetworkinstanceProtocolsProtocolStaticRoutes**](frinx.openconfig.network.instance.networkinstances.networkinstance.protocols.protocol.StaticRoutes.md) | Optional[Enclosing container for the list of static routes] REF:Optional.empty | [optional] [default to null]
**LocalAggregates** | [***FrinxOpenconfigNetworkInstanceNetworkinstancesNetworkinstanceProtocolsProtocolLocalAggregates**](frinx.openconfig.network.instance.networkinstances.networkinstance.protocols.protocol.LocalAggregates.md) | Optional[Enclosing container for locally-defined aggregate routes] REF:Optional.empty | [optional] [default to null]
**Config** | [***FrinxOpenconfigNetworkInstanceNetworkinstancesNetworkinstanceProtocolsProtocolConfig**](frinx.openconfig.network.instance.networkinstances.networkinstance.protocols.protocol.Config.md) | Optional[Configuration parameters relating to the routing protocol instance] REF:Optional.empty | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


