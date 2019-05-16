# FrinxOpenconfigPolicyForwardingPfforwardingpolicystructuralPoliciesPolicyRulesRule

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ipv4** | [***FrinxOpenconfigPacketMatchIpv4protocolfieldstopIpv4**](frinx.openconfig.packet.match.ipv4protocolfieldstop.Ipv4.md) | Optional[Top level container for IPv4 match field data] REF:Optional.empty | [optional] [default to null]
**Ipv6** | [***FrinxOpenconfigPacketMatchIpv6protocolfieldstopIpv6**](frinx.openconfig.packet.match.ipv6protocolfieldstop.Ipv6.md) | Optional[Top-level container for IPv6 match field data] REF:Optional.empty | [optional] [default to null]
**L2** | [***FrinxOpenconfigPacketMatchEthernetheadertopL2**](frinx.openconfig.packet.match.ethernetheadertop.L2.md) | Optional[Ethernet header fields] REF:Optional.empty | [optional] [default to null]
**SequenceId** | **int64** | Optional[A unique sequence identifier for the match rule.] REF:Optional.empty | [optional] [default to null]
**Action** | [***FrinxOpenconfigPolicyForwardingPfforwardingpolicystructuralPoliciesPolicyRulesRuleAction**](frinx.openconfig.policy.forwarding.pfforwardingpolicystructural.policies.policy.rules.rule.Action.md) | Optional[The forwarding policy action to be applied for packets matching the rule.] REF:Optional.empty | [optional] [default to null]
**Transport** | [***FrinxOpenconfigPacketMatchTransportfieldstopTransport**](frinx.openconfig.packet.match.transportfieldstop.Transport.md) | Optional[Transport fields container] REF:Optional.empty | [optional] [default to null]
**Config** | [***FrinxOpenconfigPolicyForwardingPfforwardingpolicystructuralPoliciesPolicyRulesRuleConfig**](frinx.openconfig.policy.forwarding.pfforwardingpolicystructural.policies.policy.rules.rule.Config.md) | Optional[Configuration parameters relating to the match rule.] REF:Optional.empty | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


