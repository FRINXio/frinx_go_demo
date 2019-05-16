# FrinxOpenconfigRoutingPolicyRoutingpolicyPolicydefinitionsPolicydefinitionStatementsStatementConditions

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FrinxOpenconfigNetworkInstancePolicymatchProtocolInstance** | [***FrinxOpenconfigNetworkInstancePolicyProtocolinstancepolicytopMatchProtocolInstance**](frinx.openconfig.network.instance.policy.protocolinstancepolicytop.MatchProtocolInstance.md) | Optional[Top-level container for protocol instance match condition in policy statements.  The protocol instance is referenced by an identifier and name] REF:Optional.empty | [optional] [default to null]
**FrinxOpenconfigOspfPolicyospfConditions** | [***FrinxOpenconfigOspfPolicyOspfmatchconditionsOspfConditions**](frinx.openconfig.ospf.policy.ospfmatchconditions.OspfConditions.md) | Optional[Match conditions specific to OSPF] REF:Optional.empty | [optional] [default to null]
**MatchNeighborSet** | [***FrinxOpenconfigRoutingPolicyNeighborsetconditiontopMatchNeighborSet**](frinx.openconfig.routing.policy.neighborsetconditiontop.MatchNeighborSet.md) | Optional[Match a referenced neighbor set according to the logic defined in the match-set-options-leaf] REF:Optional.empty | [optional] [default to null]
**FrinxOpenconfigIsisPolicyisisConditions** | [***FrinxOpenconfigIsisPolicyIsismatchconditionsIsisConditions**](frinx.openconfig.isis.policy.isismatchconditions.IsisConditions.md) | Optional[Match conditions relating to the IS-IS protocol] REF:Optional.empty | [optional] [default to null]
**MatchTagSet** | [***FrinxOpenconfigRoutingPolicyTagsetconditiontopMatchTagSet**](frinx.openconfig.routing.policy.tagsetconditiontop.MatchTagSet.md) | Optional[Match a referenced tag set according to the logic defined in the match-options-set leaf] REF:Optional.empty | [optional] [default to null]
**FrinxOpenconfigBgpPolicybgpConditions** | [***FrinxOpenconfigBgpPolicyBgpconditionstopBgpConditions**](frinx.openconfig.bgp.policy.bgpconditionstop.BgpConditions.md) | Optional[Top-level container ] REF:Optional.empty | [optional] [default to null]
**MatchPrefixSet** | [***FrinxOpenconfigRoutingPolicyPrefixsetconditiontopMatchPrefixSet**](frinx.openconfig.routing.policy.prefixsetconditiontop.MatchPrefixSet.md) | Optional[Match a referenced prefix-set according to the logic defined in the match-set-options leaf] REF:Optional.empty | [optional] [default to null]
**Config** | [***FrinxOpenconfigRoutingPolicyPolicyconditionstopConditionsConfig**](frinx.openconfig.routing.policy.policyconditionstop.conditions.Config.md) | Optional[Configuration data for policy conditions] REF:Optional.empty | [optional] [default to null]
**MatchInterface** | [***FrinxOpenconfigRoutingPolicyMatchinterfaceconditiontopMatchInterface**](frinx.openconfig.routing.policy.matchinterfaceconditiontop.MatchInterface.md) | Optional[Top-level container for interface match conditions] REF:Optional.empty | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


