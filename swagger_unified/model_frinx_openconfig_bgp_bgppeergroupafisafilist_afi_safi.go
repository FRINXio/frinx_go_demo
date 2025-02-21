/*
 * unified
 *
 * API generated from yang definitions: [fake,frinx-acl-extension,frinx-bfd,frinx-bfd-extension,frinx-bgp-extension,frinx-cdp,frinx-cisco-if-extension,frinx-cisco-mpls-te-extension,frinx-cisco-ospf-extension,frinx-cisco-pf-interfaces-extension,frinx-damping,frinx-dasan-vlan-extension,frinx-event-types,frinx-evpn,frinx-evpn-types,frinx-hsrp,frinx-if-aggregate-extension,frinx-isis-extension,frinx-juniper-if-aggregate-extension,frinx-juniper-if-extension,frinx-juniper-pf-interfaces-extension,frinx-juniper-probes-extension,frinx-l3ipvlan,frinx-lacp-lag-member,frinx-logging,frinx-mpls-ldp-extension,frinx-mpls-rsvp-extension,frinx-netflow,frinx-openconfig-acl,frinx-openconfig-bfd,frinx-openconfig-bgp,frinx-openconfig-bgp-policy,frinx-openconfig-bgp-types,frinx-openconfig-extensions,frinx-openconfig-if-aggregate,frinx-openconfig-if-ethernet,frinx-openconfig-if-ip,frinx-openconfig-if-ip-ext,frinx-openconfig-inet-types,frinx-openconfig-interfaces,frinx-openconfig-isis,frinx-openconfig-isis-lsdb-types,frinx-openconfig-isis-policy,frinx-openconfig-isis-types,frinx-openconfig-lacp,frinx-openconfig-lldp,frinx-openconfig-lldp-types,frinx-openconfig-local-routing,frinx-openconfig-mpls,frinx-openconfig-mpls-ldp,frinx-openconfig-mpls-rsvp,frinx-openconfig-mpls-types,frinx-openconfig-network-instance,frinx-openconfig-network-instance-l3,frinx-openconfig-network-instance-policy,frinx-openconfig-network-instance-types,frinx-openconfig-ospf-policy,frinx-openconfig-ospf-types,frinx-openconfig-ospfv2,frinx-openconfig-packet-match,frinx-openconfig-packet-match-types,frinx-openconfig-platform,frinx-openconfig-platform-linecard,frinx-openconfig-platform-port,frinx-openconfig-platform-transceiver,frinx-openconfig-platform-types,frinx-openconfig-policy-forwarding,frinx-openconfig-policy-types,frinx-openconfig-probes,frinx-openconfig-probes-types,frinx-openconfig-qos,frinx-openconfig-qos-types,frinx-openconfig-rib-bgp,frinx-openconfig-rib-bgp-ext,frinx-openconfig-rib-bgp-types,frinx-openconfig-routing-policy,frinx-openconfig-transport-types,frinx-openconfig-types,frinx-openconfig-vlan,frinx-openconfig-vlan-types,frinx-openconfig-yang-types,frinx-ospf-extension,frinx-qos-extension,frinx-snmp,frinx-uniconfig-topology,iana-if-type,ietf-inet-types,ietf-interfaces,ietf-yang-types,network-topology,network-topology,unified-topology,yang-ext]
 *
 * API version: 4.2.0.frinx
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger_unified

// Optional[AFI,SAFI configuration available for the neighbour or group] REF:Optional.empty
type FrinxOpenconfigBgpBgppeergroupafisafilistAfiSafi struct {
	// Optional[IPv4 unicast configuration options] REF:Optional.empty
	Ipv4Unicast *FrinxOpenconfigBgpBgpcommonmpipv4unicastgroupIpv4Unicast `json:"ipv4-unicast,omitempty"`
	// Optional[Anchor point for routing policies in the model. Import and export policies are with respect to the local routing table, i.e., export (send) and import (receive), depending on the context.] REF:Optional.empty
	ApplyPolicy *FrinxOpenconfigRoutingPolicyApplypolicygroupApplyPolicy `json:"apply-policy,omitempty"`
	// Optional[BGP-signalled VPLS configuration options] REF:Optional.empty
	L2vpnVpls *FrinxOpenconfigBgpBgpcommonmpl2vpnvplsgroupL2vpnVpls `json:"l2vpn-vpls,omitempty"`
	// Optional[IPv6 unicast configuration options] REF:Optional.empty
	Ipv6Unicast *FrinxOpenconfigBgpBgpcommonmpipv6unicastgroupIpv6Unicast `json:"ipv6-unicast,omitempty"`
	// Optional[IPv6 Labeled Unicast configuration options] REF:Optional.empty
	Ipv6LabeledUnicast *FrinxOpenconfigBgpBgpcommonmpipv6labeledunicastgroupIpv6LabeledUnicast `json:"ipv6-labeled-unicast,omitempty"`
	// Optional[Unicast IPv4 L3VPN configuration options] REF:Optional.empty
	L3vpnIpv4Unicast *FrinxOpenconfigBgpBgpcommonmpl3vpnipv4unicastgroupL3vpnIpv4Unicast `json:"l3vpn-ipv4-unicast,omitempty"`
	// Optional[BGP EVPN configuration options] REF:Optional.empty
	L2vpnEvpn *FrinxOpenconfigBgpBgpcommonmpl2vpnevpngroupL2vpnEvpn `json:"l2vpn-evpn,omitempty"`
	// Optional[Parameters related to the use of multiple paths for the same NLRI] REF:Optional.empty
	UseMultiplePaths *FrinxOpenconfigBgpBgpcommonglobalgroupusemultiplepathsUseMultiplePaths `json:"use-multiple-paths,omitempty"`
	// Optional[Parameters relating to options for route selection] REF:Optional.empty
	RouteSelectionOptions *FrinxOpenconfigBgpBgpcommonrouteselectionoptionsRouteSelectionOptions `json:"route-selection-options,omitempty"`
	// Optional[Multicast IPv4 L3VPN configuration options] REF:Optional.empty
	L3vpnIpv4Multicast *FrinxOpenconfigBgpBgpcommonmpl3vpnipv4multicastgroupL3vpnIpv4Multicast `json:"l3vpn-ipv4-multicast,omitempty"`
	// Optional[Unicast IPv6 L3VPN configuration options] REF:Optional.empty
	L3vpnIpv6Unicast *FrinxOpenconfigBgpBgpcommonmpl3vpnipv6unicastgroupL3vpnIpv6Unicast `json:"l3vpn-ipv6-unicast,omitempty"`
	// Optional[Reference to the AFI-SAFI name used as a key for the AFI-SAFI list] REF:Optional.empty
	AfiSafiName string `json:"afi-safi-name,omitempty"`
	// Optional[Parameters relating to BGP graceful-restart] REF:Optional.empty
	GracefulRestart *FrinxOpenconfigBgpBgppeergroupafisafilistAfisafiGracefulRestart `json:"graceful-restart,omitempty"`
	// Optional[IPv4 Labeled Unicast configuration options] REF:Optional.empty
	Ipv4LabeledUnicast *FrinxOpenconfigBgpBgpcommonmpipv4labeledunicastgroupIpv4LabeledUnicast `json:"ipv4-labeled-unicast,omitempty"`
	// Optional[Multicast IPv6 L3VPN configuration options] REF:Optional.empty
	L3vpnIpv6Multicast *FrinxOpenconfigBgpBgpcommonmpl3vpnipv6multicastgroupL3vpnIpv6Multicast `json:"l3vpn-ipv6-multicast,omitempty"`
	// Optional[Configuration parameters for the AFI-SAFI] REF:Optional.empty
	Config *FrinxOpenconfigBgpBgppeergroupafisafilistAfisafiConfig `json:"config,omitempty"`
}
