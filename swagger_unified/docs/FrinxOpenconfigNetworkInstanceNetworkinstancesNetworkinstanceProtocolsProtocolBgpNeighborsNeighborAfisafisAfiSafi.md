# FrinxOpenconfigNetworkInstanceNetworkinstancesNetworkinstanceProtocolsProtocolBgpNeighborsNeighborAfisafisAfiSafi

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ipv4Unicast** | [***FrinxOpenconfigBgpBgpcommonmpipv4unicastgroupIpv4Unicast**](frinx.openconfig.bgp.bgpcommonmpipv4unicastgroup.Ipv4Unicast.md) | Optional[IPv4 unicast configuration options] REF:Optional.empty | [optional] [default to null]
**ApplyPolicy** | [***FrinxOpenconfigRoutingPolicyApplypolicygroupApplyPolicy**](frinx.openconfig.routing.policy.applypolicygroup.ApplyPolicy.md) | Optional[Anchor point for routing policies in the model. Import and export policies are with respect to the local routing table, i.e., export (send) and import (receive), depending on the context.] REF:Optional.empty | [optional] [default to null]
**L2vpnVpls** | [***FrinxOpenconfigBgpBgpcommonmpl2vpnvplsgroupL2vpnVpls**](frinx.openconfig.bgp.bgpcommonmpl2vpnvplsgroup.L2vpnVpls.md) | Optional[BGP-signalled VPLS configuration options] REF:Optional.empty | [optional] [default to null]
**Ipv6Unicast** | [***FrinxOpenconfigBgpBgpcommonmpipv6unicastgroupIpv6Unicast**](frinx.openconfig.bgp.bgpcommonmpipv6unicastgroup.Ipv6Unicast.md) | Optional[IPv6 unicast configuration options] REF:Optional.empty | [optional] [default to null]
**Ipv6LabeledUnicast** | [***FrinxOpenconfigBgpBgpcommonmpipv6labeledunicastgroupIpv6LabeledUnicast**](frinx.openconfig.bgp.bgpcommonmpipv6labeledunicastgroup.Ipv6LabeledUnicast.md) | Optional[IPv6 Labeled Unicast configuration options] REF:Optional.empty | [optional] [default to null]
**L3vpnIpv4Unicast** | [***FrinxOpenconfigBgpBgpcommonmpl3vpnipv4unicastgroupL3vpnIpv4Unicast**](frinx.openconfig.bgp.bgpcommonmpl3vpnipv4unicastgroup.L3vpnIpv4Unicast.md) | Optional[Unicast IPv4 L3VPN configuration options] REF:Optional.empty | [optional] [default to null]
**L2vpnEvpn** | [***FrinxOpenconfigBgpBgpcommonmpl2vpnevpngroupL2vpnEvpn**](frinx.openconfig.bgp.bgpcommonmpl2vpnevpngroup.L2vpnEvpn.md) | Optional[BGP EVPN configuration options] REF:Optional.empty | [optional] [default to null]
**UseMultiplePaths** | [***FrinxOpenconfigBgpBgpneighborusemultiplepathsUseMultiplePaths**](frinx.openconfig.bgp.bgpneighborusemultiplepaths.UseMultiplePaths.md) | Optional[Parameters related to the use of multiple-paths for the same NLRI when they are received only from this neighbor] REF:Optional.empty | [optional] [default to null]
**L3vpnIpv4Multicast** | [***FrinxOpenconfigBgpBgpcommonmpl3vpnipv4multicastgroupL3vpnIpv4Multicast**](frinx.openconfig.bgp.bgpcommonmpl3vpnipv4multicastgroup.L3vpnIpv4Multicast.md) | Optional[Multicast IPv4 L3VPN configuration options] REF:Optional.empty | [optional] [default to null]
**L3vpnIpv6Unicast** | [***FrinxOpenconfigBgpBgpcommonmpl3vpnipv6unicastgroupL3vpnIpv6Unicast**](frinx.openconfig.bgp.bgpcommonmpl3vpnipv6unicastgroup.L3vpnIpv6Unicast.md) | Optional[Unicast IPv6 L3VPN configuration options] REF:Optional.empty | [optional] [default to null]
**AfiSafiName** | **string** | Optional[Reference to the AFI-SAFI name used as a key for the AFI-SAFI list] REF:Optional.empty | [optional] [default to null]
**GracefulRestart** | [***FrinxOpenconfigBgpBgpneighborafisafilistAfisafiGracefulRestart**](frinx.openconfig.bgp.bgpneighborafisafilist.afisafi.GracefulRestart.md) | Optional[Parameters relating to BGP graceful-restart] REF:Optional.empty | [optional] [default to null]
**Ipv4LabeledUnicast** | [***FrinxOpenconfigBgpBgpcommonmpipv4labeledunicastgroupIpv4LabeledUnicast**](frinx.openconfig.bgp.bgpcommonmpipv4labeledunicastgroup.Ipv4LabeledUnicast.md) | Optional[IPv4 Labeled Unicast configuration options] REF:Optional.empty | [optional] [default to null]
**L3vpnIpv6Multicast** | [***FrinxOpenconfigBgpBgpcommonmpl3vpnipv6multicastgroupL3vpnIpv6Multicast**](frinx.openconfig.bgp.bgpcommonmpl3vpnipv6multicastgroup.L3vpnIpv6Multicast.md) | Optional[Multicast IPv6 L3VPN configuration options] REF:Optional.empty | [optional] [default to null]
**Config** | [***FrinxOpenconfigNetworkInstanceNetworkinstancesNetworkinstanceProtocolsProtocolBgpNeighborsNeighborAfisafisAfisafiConfig**](frinx.openconfig.network.instance.networkinstances.networkinstance.protocols.protocol.bgp.neighbors.neighbor.afisafis.afisafi.Config.md) | Optional[Configuration parameters for the AFI-SAFI] REF:Optional.empty | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


