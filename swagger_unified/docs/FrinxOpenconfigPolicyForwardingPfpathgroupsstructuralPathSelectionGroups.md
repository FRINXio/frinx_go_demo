# FrinxOpenconfigPolicyForwardingPfpathgroupsstructuralPathSelectionGroups

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PathSelectionGroup** | [**[]FrinxOpenconfigPolicyForwardingPfpathgroupsstructuralPathselectiongroupsPathSelectionGroup**](frinx.openconfig.policy.forwarding.pfpathgroupsstructural.pathselectiongroups.PathSelectionGroup.md) | Optional[A path selection group is a set of forwarding resources, which are grouped as eligible paths for a particular policy-based forwarding rule. A policy rule may select a path-selection-group as the egress for a particular type of traffic (e.g., DSCP value). The system then utilises its standard forwarding lookup mechanism to select from the paths that are specified within the group - for IP packets, the destination IP address is used such that the packet is routed to the entity within the path-selection-group that corresponds to the next-hop for the destination IP address of the packet; for L2 packets, the selection is based on the destination MAC address.] REF:Optional.empty | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


