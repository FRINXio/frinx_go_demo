# FrinxOpenconfigOspfv2Ospfv2globalstructuralGlobalGracefulrestartConfig

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enabled** | **bool** | Optional[When the value of this leaf is set to true, graceful restart is enabled on the local system. In this case, the system will use Grace-LSAs to signal that it is restarting to its neighbors.] REF:Optional.empty | [optional] [default to null]
**HelperOnly** | **bool** | Optional[Operate graceful-restart only in helper mode. When this leaf is set to true, the local system does not use Grace-LSAs to indicate that it is restarting, but will accept Grace-LSAs from remote systems, and suppress withdrawl of adjacencies of the system for the grace period specified] REF:Optional.empty | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


