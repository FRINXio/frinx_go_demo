# FrinxOpenconfigBgpPolicyAspathprependtopSetaspathprependConfig

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Asn** | **int64** | Optional[The AS number to prepend to the AS path. If this leaf is not specified and repeat-n is set, then the local AS number will be used for prepending.] REF:Optional.empty | [optional] [default to null]
**RepeatN** | **int32** | Optional[Number of times to prepend the value specified in the asn leaf to the AS path. If no value is specified by the asn leaf, the local AS number of the system is used. The value should be between 1 and the maximum supported by the implementation.] REF:Optional.empty | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


