# FrinxOpenconfigPolicyForwardingPfforwardingpolicystructuralPoliciesPolicyRulesRuleActionConfig

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NetworkInstance** | **string** | Optional[When this leaf is set, packets matching the match criteria for the forwarding rule should be looked up in the network-instance that is referenced rather than the network-instance with which the interface is associated. Such configuration allows policy-routing into multiple sub-topologies from a single ingress access interface, or different send and receive contexts for a particular interface (sometimes referred to as half-duplex VRF).] REF:Optional.empty | [optional] [default to null]
**DecapsulateGre** | **bool** | Optional[When this leaf is set to true, the local system should remove the GRE header from the packet matching the rule. Following the decapsulation it should subsequently forward the encapsulated packet according to the relevant lookup (e.g., if the encapsulated packet is IP, the packet should be routed according to the IP destination).] REF:Optional.empty | [optional] [default to null]
**Discard** | **bool** | Optional[When this leaf is set to true, the local system should drop packets that match the rule.] REF:Optional.empty | [optional] [default to null]
**NextHop** | **string** | Optional[When an IP next-hop is specified in the next-hop field, packets matching the match criteria for the forwarding rule should be forwarded to the next-hop IP address, bypassing any lookup on the local system.] REF:Optional.empty | [optional] [default to null]
**PathSelectionGroup** | **string** | Optional[When path-selection-group is set, packets matching the match criteria for the forwarding rule should be forwarded only via one of the paths that is specified within the referenced path-selection-group. The next-hop of the packet within the routing context should be used to determine between multiple paths that are specified within the group.] REF:Optional.empty | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


