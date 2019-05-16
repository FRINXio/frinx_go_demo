# FrinxOpenconfigQosQosSchedulerpoliciesSchedulerpolicySchedulersSchedulerOneratetwocolorConfig

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MaxQueueDepthBytes** | **int64** | Optional[When the scheduler is specified to be a shaper - the maximum depth of the queue in bytes is the value specified by this leaf.] REF:Optional.empty | [optional] [default to null]
**QueuingBehavior** | [***FrinxOpenconfigQosTypesQueueBehavior**](frinx.openconfig.qos.types.QueueBehavior.md) | Optional[The type of scheduler that is being configured.] REF:Optional.empty | [optional] [default to null]
**FrinxQosExtensionmaxQueueDepthMs** | **int64** | Optional[When the scheduler is specified to be a shaper - the maximum depth of the queue in miliseconds is the value specified by this leaf.] REF:Optional.empty | [optional] [default to null]
**Bc** | **int64** | Optional[Committed burst size for the single-rate token bucket scheduler.  This value represents the depth of the token bucket.] REF:Optional.empty | [optional] [default to null]
**CirPct** | **int32** | Optional[Committed information rate for the single-rate token bucket scheduler. This value represents the rate at which tokens are added to the bucket. It is expressed as a percentage of the total bandwidth allocated to the context in which the scheduler is referenced.] REF:Optional.empty | [optional] [default to null]
**MaxQueueDepthPackets** | **int64** | Optional[When the scheduler is specified to be a shaper - the maximum depth of the queue in packets is the value specified by this leaf.] REF:Optional.empty | [optional] [default to null]
**MaxQueueDepthPercent** | **int32** | Optional[The queue depth specified as a percentage of the total available buffer that is avaialble.] REF:Optional.empty | [optional] [default to null]
**CirPctRemaining** | **int32** | Optional[Committed information rate for the single-rate token bucket scheduler. This value represents the rate at which tokens are added to the bucket. It is expressed as a percentage of the unallocated bandwidth available in the context in which the scheduled is referenced.] REF:Optional.empty | [optional] [default to null]
**Cir** | **int32** | Optional[Committed information rate for the single-rate token bucket scheduler.  This value represents the rate at which tokens are added to the bucket.] REF:Optional.empty | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


