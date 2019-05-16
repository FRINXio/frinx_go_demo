# FrinxOpenconfigBgpBgpcommonmpallafisaficommonPrefixlimitConfig

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RestartTimer** | **int32** | Optional[Time interval in seconds after which the BGP session is re-established after being torn down due to exceeding the max-prefix limit.] REF:Optional.empty | [optional] [default to null]
**ShutdownThresholdPct** | **int32** | Optional[Threshold on number of prefixes that can be received from a neighbour before generation of warning messages or log entries. Expressed as a percentage of max-prefixes] REF:Optional.empty | [optional] [default to null]
**MaxPrefixes** | **int64** | Optional[Maximum number of prefixes that will be accepted from the neighbour] REF:Optional.empty | [optional] [default to null]
**PreventTeardown** | **bool** | Optional[Do not tear down the BGP session when the maximum prefix limit is exceeded, but rather only log a warning. The default of this leaf is false, such that when it is not specified, the session is torn down.] REF:Optional.empty | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


