# FrinxOpenconfigRoutingPolicyPrefixtopPrefixesPrefixConfig

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MasklengthRange** | **string** | Optional[Defines a range for the masklength, or &#39;exact&#39; if the prefix has an exact length.  Example: 10.3.192.0/21 through 10.3.192.0/24 would be expressed as prefix: 10.3.192.0/21, masklength-range: 21..24.  Example: 10.3.192.0/21 would be expressed as prefix: 10.3.192.0/21, masklength-range: exact] REF:Optional.empty | [optional] [default to null]
**IpPrefix** | **string** | Optional[The prefix member in CIDR notation -- while the prefix may be either IPv4 or IPv6, most implementations require all members of the prefix set to be the same address family.  Mixing address types in the same prefix set is likely to cause an error.] REF:Optional.empty | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


