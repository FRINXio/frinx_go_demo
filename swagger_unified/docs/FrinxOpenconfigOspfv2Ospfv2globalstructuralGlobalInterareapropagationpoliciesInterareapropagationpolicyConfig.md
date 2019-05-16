# FrinxOpenconfigOspfv2Ospfv2globalstructuralGlobalInterareapropagationpoliciesInterareapropagationpolicyConfig

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DefaultImportPolicy** | [***FrinxOpenconfigRoutingPolicyDefaultPolicyType**](frinx.openconfig.routing.policy.DefaultPolicyType.md) | Optional[explicitly set a default policy if no policy definition in the import policy chain is satisfied.] REF:Optional.empty | [optional] [default to null]
**DstArea** | **string** | Optional[The destination area to which prefixes are to be imported] REF:Optional.empty | [optional] [default to null]
**SrcArea** | **string** | Optional[The area from which prefixes are to be exported.] REF:Optional.empty | [optional] [default to null]
**ImportPolicy** | **[]string** | Optional[list of policy names in sequence to be applied on receiving a routing update in the current context, e.g., for the current peer group, neighbor, address family, etc.] REF:Optional.empty | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


