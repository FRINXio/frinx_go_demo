# FrinxDampingDampingtopDampingConfig

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MaxSuppress** | **int64** | Optional[Max suppress time (in minutes)] REF:Optional.empty | [optional] [default to null]
**Reuse** | **int64** | Optional[Reuse threshold — Arbitrary value below which a suppressed route can be used again.] REF:Optional.empty | [optional] [default to null]
**Suppress** | **int64** | Optional[Suppress threshold — Arbitrary value above which a route can no longer be used or included in advertisements. Must be greater than the reuse threshold.] REF:Optional.empty | [optional] [default to null]
**HalfLife** | **int64** | Optional[Decay half life (in minutes). Must be less than the max-suppress value.] REF:Optional.empty | [optional] [default to null]
**Enabled** | **bool** | Optional[Desired state of damping] REF:Optional.empty | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


