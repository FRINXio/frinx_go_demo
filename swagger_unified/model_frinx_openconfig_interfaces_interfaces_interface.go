/*
 * unified
 *
 * API generated from yang definitions: [fake,frinx-acl-extension,frinx-bfd,frinx-bfd-extension,frinx-bgp-extension,frinx-cdp,frinx-cisco-if-extension,frinx-cisco-mpls-te-extension,frinx-cisco-ospf-extension,frinx-cisco-pf-interfaces-extension,frinx-damping,frinx-dasan-vlan-extension,frinx-event-types,frinx-evpn,frinx-evpn-types,frinx-hsrp,frinx-if-aggregate-extension,frinx-isis-extension,frinx-juniper-if-aggregate-extension,frinx-juniper-if-extension,frinx-juniper-pf-interfaces-extension,frinx-juniper-probes-extension,frinx-l3ipvlan,frinx-lacp-lag-member,frinx-logging,frinx-mpls-ldp-extension,frinx-mpls-rsvp-extension,frinx-netflow,frinx-openconfig-acl,frinx-openconfig-bfd,frinx-openconfig-bgp,frinx-openconfig-bgp-policy,frinx-openconfig-bgp-types,frinx-openconfig-extensions,frinx-openconfig-if-aggregate,frinx-openconfig-if-ethernet,frinx-openconfig-if-ip,frinx-openconfig-if-ip-ext,frinx-openconfig-inet-types,frinx-openconfig-interfaces,frinx-openconfig-isis,frinx-openconfig-isis-lsdb-types,frinx-openconfig-isis-policy,frinx-openconfig-isis-types,frinx-openconfig-lacp,frinx-openconfig-lldp,frinx-openconfig-lldp-types,frinx-openconfig-local-routing,frinx-openconfig-mpls,frinx-openconfig-mpls-ldp,frinx-openconfig-mpls-rsvp,frinx-openconfig-mpls-types,frinx-openconfig-network-instance,frinx-openconfig-network-instance-l3,frinx-openconfig-network-instance-policy,frinx-openconfig-network-instance-types,frinx-openconfig-ospf-policy,frinx-openconfig-ospf-types,frinx-openconfig-ospfv2,frinx-openconfig-packet-match,frinx-openconfig-packet-match-types,frinx-openconfig-platform,frinx-openconfig-platform-linecard,frinx-openconfig-platform-port,frinx-openconfig-platform-transceiver,frinx-openconfig-platform-types,frinx-openconfig-policy-forwarding,frinx-openconfig-policy-types,frinx-openconfig-probes,frinx-openconfig-probes-types,frinx-openconfig-qos,frinx-openconfig-qos-types,frinx-openconfig-rib-bgp,frinx-openconfig-rib-bgp-ext,frinx-openconfig-rib-bgp-types,frinx-openconfig-routing-policy,frinx-openconfig-transport-types,frinx-openconfig-types,frinx-openconfig-vlan,frinx-openconfig-vlan-types,frinx-openconfig-yang-types,frinx-ospf-extension,frinx-qos-extension,frinx-snmp,frinx-uniconfig-topology,iana-if-type,ietf-inet-types,ietf-interfaces,ietf-yang-types,network-topology,network-topology,unified-topology,yang-ext]
 *
 * API version: 4.2.0.frinx
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger_unified

// Optional[The list of named interfaces on the device.] REF:Optional.empty
type FrinxOpenconfigInterfacesInterfacesInterface struct {
	// Optional.empty REF:Optional.empty
	FrinxCiscoIfExtensionverifyUnicastSourceReachableVia *FrinxCiscoIfExtensionVerifyunicastsourcereachableviatopVerifyUnicastSourceReachableVia `json:"frinx-cisco-if-extension:verify-unicast-source-reachable-via,omitempty"`
	// Optional[Enclosing container for statistics interface-specific data.] REF:Optional.empty
	FrinxCiscoIfExtensionstatistics *FrinxCiscoIfExtensionStatisticstopStatistics `json:"frinx-cisco-if-extension:statistics,omitempty"`
	// Optional[Configuration for L3 VLAN interface] REF:Optional.empty
	FrinxL3ipvlanl3ipvlan *FrinxL3ipvlanL3ipvlaninterfacetopL3ipvlan `json:"frinx-l3ipvlan:l3ipvlan,omitempty"`
	// Optional[Top-level container for routed vlan interfaces.  These logical interfaces are also known as SVI (switched virtual interface), IRB (integrated routing and bridging), RVI (routed VLAN interface)] REF:Optional.empty
	FrinxOpenconfigVlanroutedVlan *FrinxOpenconfigVlanVlanroutedtopRoutedVlan `json:"frinx-openconfig-vlan:routed-vlan,omitempty"`
	// Optional[Enclosing container for damping interface-specific data.] REF:Optional.empty
	FrinxDampingdamping *FrinxDampingDampingtopDamping `json:"frinx-damping:damping,omitempty"`
	// Optional[References the configured name of the interface] REF:Optional.empty
	Name string `json:"name,omitempty"`
	// Optional[Options for logical interfaces representing aggregates] REF:Optional.empty
	FrinxOpenconfigIfAggregateaggregation *FrinxOpenconfigIfAggregateAggregationlogicaltopAggregation `json:"frinx-openconfig-if-aggregate:aggregation,omitempty"`
	// Optional[Top-level container for ethernet configuration and state] REF:Optional.empty
	FrinxOpenconfigIfEthernetethernet *FrinxOpenconfigIfEthernetEthernettopEthernet `json:"frinx-openconfig-if-ethernet:ethernet,omitempty"`
	// Optional[Enclosing container for the list of subinterfaces associated with a physical interface] REF:Optional.empty
	Subinterfaces *FrinxOpenconfigInterfacesInterfacesInterfaceSubinterfaces `json:"subinterfaces,omitempty"`
	// Optional[Configurable items at the global, physical interface level] REF:Optional.empty
	Config *FrinxOpenconfigInterfacesInterfacesInterfaceConfig `json:"config,omitempty"`
	// Optional[Top-level container for hold-time settings to enable dampening advertisements of interface transitions.] REF:Optional.empty
	HoldTime *FrinxOpenconfigInterfacesInterfacephysholdtimetopHoldTime `json:"hold-time,omitempty"`
}
