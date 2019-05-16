# FrinxOpenconfigBgpPolicyBgpconditionstopBgpconditionsConfig

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MedEq** | **int64** | Optional[Condition to check if the received MED value is equal to the specified value] REF:Optional.empty | [optional] [default to null]
**LocalPrefEq** | **int64** | Optional[Condition to check if the local pref attribute is equal to the specified value] REF:Optional.empty | [optional] [default to null]
**NextHopIn** | **[]string** | Optional[List of next hop addresses to check for in the route update] REF:Optional.empty | [optional] [default to null]
**RouteType** | [***FrinxOpenconfigBgpPolicyRouteTypeEnumeration**](frinx.openconfig.bgp.policy.RouteTypeEnumeration.md) | Optional[Condition to check the route type in the route update] REF:Optional.empty | [optional] [default to null]
**AfiSafiIn** | [**[]FrinxOpenconfigBgpPolicyAfiSafiInIdentityref**](frinx.openconfig.bgp.policy.AfiSafiInIdentityref.md) | Optional[List of address families which the NLRI may be within] REF:Optional.empty | [optional] [default to null]
**OriginEq** | [***FrinxOpenconfigBgpTypesBgpOriginAttrType**](frinx.openconfig.bgp.types.BgpOriginAttrType.md) | Optional[Condition to check if the route origin is equal to the specified value] REF:Optional.empty | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


