# FrinxOpenconfigRoutingPolicyPolicyconditionstopConditionsConfig

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**InstallProtocolEq** | [***FrinxOpenconfigRoutingPolicyInstallProtocolEqIdentityref**](frinx.openconfig.routing.policy.InstallProtocolEqIdentityref.md) | Optional[Condition to check the protocol / method used to install the route into the local routing table] REF:Optional.empty | [optional] [default to null]
**CallPolicy** | **string** | Optional[Applies the statements from the specified policy definition and then returns control the current policy statement. Note that the called policy may itself call other policies (subject to implementation limitations). This is intended to provide a policy &#39;subroutine&#39; capability.  The called policy should contain an explicit or a default route disposition that returns an effective true (accept-route) or false (reject-route), otherwise the behavior may be ambiguous and implementation dependent] REF:Optional.empty | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


