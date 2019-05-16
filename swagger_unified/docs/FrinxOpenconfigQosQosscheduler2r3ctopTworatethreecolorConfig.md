# FrinxOpenconfigQosQosscheduler2r3ctopTworatethreecolorConfig

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Bc** | **int64** | Optional[Committed burst size for the dual-rate token bucket policer.  This value represents the depth of the token bucket.] REF:Optional.empty | [optional] [default to null]
**CirPct** | **int32** | Optional[Committed information rate for the dual-rate token bucket policer. This value represents the rate at which tokens are added to the primary bucket. It is expressed as a percentage of the total bandwidth available within the context the scheduler is instantiated.] REF:Optional.empty | [optional] [default to null]
**Be** | **int64** | Optional[Excess burst size for the dual-rate token bucket policer. This value represents the depth of the secondary bucket.] REF:Optional.empty | [optional] [default to null]
**CirPctRemaining** | **int32** | Optional[Committed information rate for the dual-rate token bucket policer. This value represents the rate at which tokens are added to the primary bucket. It is expressed as a percentage of the remaining bandwidth within the context the scheduler is instantiated.] REF:Optional.empty | [optional] [default to null]
**Pir** | **int32** | Optional[Peak information rate for the dual-rate token bucket policer.  This value represents the rate at which tokens are added to the secondary bucket.] REF:Optional.empty | [optional] [default to null]
**PirPctRemaining** | **int32** | Optional[Peak information rate for the dual-rate token bucket policer. This value represents the rate at which tokens are added to the secondary bucket. It is expressed as a percentage of the remaining bandwidth within the context the scheduler is instantiated.] REF:Optional.empty | [optional] [default to null]
**Cir** | **int32** | Optional[Committed information rate for the dual-rate token bucket policer.  This value represents the rate at which tokens are added to the primary bucket.] REF:Optional.empty | [optional] [default to null]
**PirPct** | **int32** | Optional[Peak information rate for the dual-rate token bucket policer. This value represents the rate at which tokens are added to the secondary bucket. The value is expressed as a percentage of the total bandwidth available in the context in which the scheduler is instantiated.] REF:Optional.empty | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


