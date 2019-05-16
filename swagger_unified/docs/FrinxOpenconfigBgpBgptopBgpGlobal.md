# FrinxOpenconfigBgpBgptopBgpGlobal

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UseMultiplePaths** | [***FrinxOpenconfigBgpBgpcommonglobalgroupusemultiplepathsUseMultiplePaths**](frinx.openconfig.bgp.bgpcommonglobalgroupusemultiplepaths.UseMultiplePaths.md) | Optional[Parameters related to the use of multiple paths for the same NLRI] REF:Optional.empty | [optional] [default to null]
**DynamicNeighborPrefixes** | [***FrinxOpenconfigBgpBgpglobaldynamicneighborsDynamicNeighborPrefixes**](frinx.openconfig.bgp.bgpglobaldynamicneighbors.DynamicNeighborPrefixes.md) | Optional[A list of IP prefixes from which the system should:  - Accept connections to the BGP daemon  - Dynamically configure a BGP neighbor corresponding to the    source address of the remote system, using the parameters    of the specified peer-group. For such neighbors, an entry within the neighbor list should be created, indicating that the peer was dynamically configured, and referencing the peer-group from which the configuration was derived.] REF:Optional.empty | [optional] [default to null]
**RouteSelectionOptions** | [***FrinxOpenconfigBgpBgpcommonrouteselectionoptionsRouteSelectionOptions**](frinx.openconfig.bgp.bgpcommonrouteselectionoptions.RouteSelectionOptions.md) | Optional[Parameters relating to options for route selection] REF:Optional.empty | [optional] [default to null]
**AfiSafis** | [***FrinxOpenconfigBgpBgpglobalbaseAfiSafis**](frinx.openconfig.bgp.bgpglobalbase.AfiSafis.md) | Optional[Address family specific configuration] REF:Optional.empty | [optional] [default to null]
**GracefulRestart** | [***FrinxOpenconfigBgpBgpglobalbaseGracefulRestart**](frinx.openconfig.bgp.bgpglobalbase.GracefulRestart.md) | Optional[Parameters relating the graceful restart mechanism for BGP] REF:Optional.empty | [optional] [default to null]
**DefaultRouteDistance** | [***FrinxOpenconfigBgpBgpglobalbaseDefaultRouteDistance**](frinx.openconfig.bgp.bgpglobalbase.DefaultRouteDistance.md) | Optional[Administrative distance (or preference) assigned to routes received from different sources (external, internal, and local).] REF:Optional.empty | [optional] [default to null]
**Config** | [***FrinxOpenconfigBgpBgpglobalbaseConfig**](frinx.openconfig.bgp.bgpglobalbase.Config.md) | Optional[Configuration parameters relating to the global BGP router] REF:Optional.empty | [optional] [default to null]
**Confederation** | [***FrinxOpenconfigBgpBgpglobalbaseConfederation**](frinx.openconfig.bgp.bgpglobalbase.Confederation.md) | Optional[Parameters indicating whether the local system acts as part of a BGP confederation] REF:Optional.empty | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


