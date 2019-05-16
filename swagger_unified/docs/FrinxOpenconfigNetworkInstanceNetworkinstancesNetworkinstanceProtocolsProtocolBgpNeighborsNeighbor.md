# FrinxOpenconfigNetworkInstanceNetworkinstancesNetworkinstanceProtocolsProtocolBgpNeighborsNeighbor

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Timers** | [***FrinxOpenconfigBgpBgpneighborbaseTimers**](frinx.openconfig.bgp.bgpneighborbase.Timers.md) | Optional[Timers related to a BGP neighbor] REF:Optional.empty | [optional] [default to null]
**ApplyPolicy** | [***FrinxOpenconfigRoutingPolicyApplypolicygroupApplyPolicy**](frinx.openconfig.routing.policy.applypolicygroup.ApplyPolicy.md) | Optional[Anchor point for routing policies in the model. Import and export policies are with respect to the local routing table, i.e., export (send) and import (receive), depending on the context.] REF:Optional.empty | [optional] [default to null]
**RouteReflector** | [***FrinxOpenconfigBgpBgpcommonstructureneighborgrouproutereflectorRouteReflector**](frinx.openconfig.bgp.bgpcommonstructureneighborgrouproutereflector.RouteReflector.md) | Optional[Route reflector parameters for the BGPgroup] REF:Optional.empty | [optional] [default to null]
**FrinxOpenconfigBfdenableBfd** | [***FrinxOpenconfigBfdBfdenableEnableBfd**](frinx.openconfig.bfd.bfdenable.EnableBfd.md) | Optional[Enable BFD for liveliness detection to the next-hop or neighbour.] REF:Optional.empty | [optional] [default to null]
**Transport** | [***FrinxOpenconfigBgpBgpneighborbaseTransport**](frinx.openconfig.bgp.bgpneighborbase.Transport.md) | Optional[Transport session parameters for the BGP neighbor] REF:Optional.empty | [optional] [default to null]
**AsPathOptions** | [***FrinxOpenconfigBgpBgpcommonstructureneighborgroupaspathoptionsAsPathOptions**](frinx.openconfig.bgp.bgpcommonstructureneighborgroupaspathoptions.AsPathOptions.md) | Optional[AS_PATH manipulation parameters for the BGP neighbor or group] REF:Optional.empty | [optional] [default to null]
**EbgpMultihop** | [***FrinxOpenconfigBgpBgpcommonstructureneighborgroupebgpmultihopEbgpMultihop**](frinx.openconfig.bgp.bgpcommonstructureneighborgroupebgpmultihop.EbgpMultihop.md) | Optional[eBGP multi-hop parameters for the BGPgroup] REF:Optional.empty | [optional] [default to null]
**ErrorHandling** | [***FrinxOpenconfigBgpBgpneighborbaseErrorHandling**](frinx.openconfig.bgp.bgpneighborbase.ErrorHandling.md) | Optional[Error handling parameters used for the BGP neighbor or group] REF:Optional.empty | [optional] [default to null]
**UseMultiplePaths** | [***FrinxOpenconfigBgpBgpneighborusemultiplepathsUseMultiplePaths**](frinx.openconfig.bgp.bgpneighborusemultiplepaths.UseMultiplePaths.md) | Optional[Parameters related to the use of multiple-paths for the same NLRI when they are received only from this neighbor] REF:Optional.empty | [optional] [default to null]
**AddPaths** | [***FrinxOpenconfigBgpBgpcommonstructureneighborgroupaddpathsAddPaths**](frinx.openconfig.bgp.bgpcommonstructureneighborgroupaddpaths.AddPaths.md) | Optional[Parameters relating to the advertisement and receipt of multiple paths for a single NLRI (add-paths)] REF:Optional.empty | [optional] [default to null]
**AfiSafis** | [***FrinxOpenconfigNetworkInstanceNetworkinstancesNetworkinstanceProtocolsProtocolBgpNeighborsNeighborAfiSafis**](frinx.openconfig.network.instance.networkinstances.networkinstance.protocols.protocol.bgp.neighbors.neighbor.AfiSafis.md) | Optional[Per-address-family configuration parameters associated with the neighbor] REF:Optional.empty | [optional] [default to null]
**GracefulRestart** | [***FrinxOpenconfigBgpBgpneighborbaseGracefulRestart**](frinx.openconfig.bgp.bgpneighborbase.GracefulRestart.md) | Optional[Parameters relating the graceful restart mechanism for BGP] REF:Optional.empty | [optional] [default to null]
**NeighborAddress** | **string** | Optional[Reference to the address of the BGP neighbor used as a key in the neighbor list] REF:Optional.empty | [optional] [default to null]
**LoggingOptions** | [***FrinxOpenconfigBgpBgpcommonstructureneighborgrouploggingoptionsLoggingOptions**](frinx.openconfig.bgp.bgpcommonstructureneighborgrouploggingoptions.LoggingOptions.md) | Optional[Logging options for events related to the BGP neighbor or group] REF:Optional.empty | [optional] [default to null]
**Config** | [***FrinxOpenconfigBgpBgpneighborbaseConfig**](frinx.openconfig.bgp.bgpneighborbase.Config.md) | Optional[Configuration parameters relating to the BGP neighbor or group] REF:Optional.empty | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


