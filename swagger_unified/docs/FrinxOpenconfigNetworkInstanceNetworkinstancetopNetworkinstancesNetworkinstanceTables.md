# FrinxOpenconfigNetworkInstanceNetworkinstancetopNetworkinstancesNetworkinstanceTables

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Table** | [**[]FrinxOpenconfigNetworkInstanceNetworkinstancetopNetworkinstancesNetworkinstanceTablesTable**](frinx.openconfig.network.instance.networkinstancetop.networkinstances.networkinstance.tables.Table.md) | Optional[A network instance manages one or more forwarding or routing tables. These may reflect a Layer 2 forwarding information base, a Layer 3 routing table, or an MPLS LFIB.  The table populated by a protocol within an instance is identified by the protocol identifier (e.g., BGP, IS-IS) and the address family (e.g., IPv4, IPv6) supported by that protocol. Multiple instances of the same protocol populate a single table -- such that a single IS-IS or OSPF IPv4 table exists per network instance.  An implementation is expected to create entries within this list when the relevant protocol context is enabled. i.e., when a BGP instance is created with IPv4 and IPv6 address families enabled, the protocol&#x3D;BGP, address-family&#x3D;IPv4 table is created by the system.] REF:Optional.empty | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


