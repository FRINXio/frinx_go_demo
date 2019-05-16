# FrinxOpenconfigRoutingPolicyPrefixsettopPrefixsetsPrefixsetConfig

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Mode** | [***FrinxOpenconfigRoutingPolicyModeEnumeration**](frinx.openconfig.routing.policy.ModeEnumeration.md) | Optional[Indicates the mode of the prefix set, in terms of which address families (IPv4, IPv6, or both) are present.  The mode provides a hint, but the device must validate that all prefixes are of the indicated type, and is expected to reject the configuration if there is a discrepancy.  The MIXED mode may not be supported on devices that require prefix sets to be of only one address family.] REF:Optional.empty | [optional] [default to null]
**Name** | **string** | Optional[name / label of the prefix set -- this is used to reference the set in match conditions] REF:Optional.empty | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


