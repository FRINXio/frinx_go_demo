/*
 * unified
 *
 * API generated from yang definitions: [fake,frinx-acl-extension,frinx-bfd,frinx-bfd-extension,frinx-bgp-extension,frinx-cdp,frinx-cisco-if-extension,frinx-cisco-mpls-te-extension,frinx-cisco-ospf-extension,frinx-cisco-pf-interfaces-extension,frinx-damping,frinx-dasan-vlan-extension,frinx-event-types,frinx-evpn,frinx-evpn-types,frinx-hsrp,frinx-if-aggregate-extension,frinx-isis-extension,frinx-juniper-if-aggregate-extension,frinx-juniper-if-extension,frinx-juniper-pf-interfaces-extension,frinx-juniper-probes-extension,frinx-l3ipvlan,frinx-lacp-lag-member,frinx-logging,frinx-mpls-ldp-extension,frinx-mpls-rsvp-extension,frinx-netflow,frinx-openconfig-acl,frinx-openconfig-bfd,frinx-openconfig-bgp,frinx-openconfig-bgp-policy,frinx-openconfig-bgp-types,frinx-openconfig-extensions,frinx-openconfig-if-aggregate,frinx-openconfig-if-ethernet,frinx-openconfig-if-ip,frinx-openconfig-if-ip-ext,frinx-openconfig-inet-types,frinx-openconfig-interfaces,frinx-openconfig-isis,frinx-openconfig-isis-lsdb-types,frinx-openconfig-isis-policy,frinx-openconfig-isis-types,frinx-openconfig-lacp,frinx-openconfig-lldp,frinx-openconfig-lldp-types,frinx-openconfig-local-routing,frinx-openconfig-mpls,frinx-openconfig-mpls-ldp,frinx-openconfig-mpls-rsvp,frinx-openconfig-mpls-types,frinx-openconfig-network-instance,frinx-openconfig-network-instance-l3,frinx-openconfig-network-instance-policy,frinx-openconfig-network-instance-types,frinx-openconfig-ospf-policy,frinx-openconfig-ospf-types,frinx-openconfig-ospfv2,frinx-openconfig-packet-match,frinx-openconfig-packet-match-types,frinx-openconfig-platform,frinx-openconfig-platform-linecard,frinx-openconfig-platform-port,frinx-openconfig-platform-transceiver,frinx-openconfig-platform-types,frinx-openconfig-policy-forwarding,frinx-openconfig-policy-types,frinx-openconfig-probes,frinx-openconfig-probes-types,frinx-openconfig-qos,frinx-openconfig-qos-types,frinx-openconfig-rib-bgp,frinx-openconfig-rib-bgp-ext,frinx-openconfig-rib-bgp-types,frinx-openconfig-routing-policy,frinx-openconfig-transport-types,frinx-openconfig-types,frinx-openconfig-vlan,frinx-openconfig-vlan-types,frinx-openconfig-yang-types,frinx-ospf-extension,frinx-qos-extension,frinx-snmp,frinx-uniconfig-topology,iana-if-type,ietf-inet-types,ietf-interfaces,ietf-yang-types,network-topology,network-topology,unified-topology,yang-ext]
 *
 * API version: 4.2.0.frinx
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger_unified

// Optional[Configuration parameters relating to an interface to policy forwarding rule binding.] REF:Optional.empty
type FrinxOpenconfigNetworkInstanceNetworkinstancesNetworkinstancePolicyforwardingInterfacesInterfaceConfig struct {
	// Optional[The policy to be applied on the interface. Packets ingress on the referenced interface should be compared to the match criteria within the specified policy, and in the case that these criteria are met, the forwarding actions specified applied. These policies should be applied following quality of service classification, and ACL actions if such entities are referenced by the corresponding interface.] REF:Optional.empty
	ApplyForwardingPolicy string `json:"apply-forwarding-policy,omitempty"`
	// Optional[Classify incoming packets based on code point value] REF:Optional.empty
	FrinxJuniperPfInterfacesExtensionclassifiers *FrinxJuniperPfInterfacesExtensionJuniperpfinterfaceextensionconfigClassifiers `json:"frinx-juniper-pf-interfaces-extension:classifiers,omitempty"`
	// Optional[Output scheduler map] REF:Optional.empty
	FrinxJuniperPfInterfacesExtensionschedulerMap string `json:"frinx-juniper-pf-interfaces-extension:scheduler-map,omitempty"`
	// Optional[Service policy which is applied on packets in ingress direction] REF:Optional.empty
	FrinxCiscoPfInterfacesExtensioninputServicePolicy string `json:"frinx-cisco-pf-interfaces-extension:input-service-policy,omitempty"`
	// Optional[Service policy which is applied on packets in egress direction] REF:Optional.empty
	FrinxCiscoPfInterfacesExtensionoutputServicePolicy string `json:"frinx-cisco-pf-interfaces-extension:output-service-policy,omitempty"`
	// Optional[A unique identifier for the interface.] REF:Optional.empty
	InterfaceId string `json:"interface-id,omitempty"`
}
