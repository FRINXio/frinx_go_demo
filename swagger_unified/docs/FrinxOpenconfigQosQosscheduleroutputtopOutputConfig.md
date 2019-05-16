# FrinxOpenconfigQosQosscheduleroutputtopOutputConfig

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OutputType** | [***FrinxOpenconfigQosOutputTypeEnumeration**](frinx.openconfig.qos.OutputTypeEnumeration.md) | Optional[Describes the type of output sink for the scheduler.] REF:Optional.empty | [optional] [default to null]
**OutputFwdGroup** | **string** | Optional[When the scheduler output type is a forwarding group, this leaf provides a reference to the forwarding group.] REF:Optional.empty | [optional] [default to null]
**ChildScheduler** | **string** | Optional[When the scheduler output type is a child scheduler, this leaf provides a reference to the downstream scheduler.] REF:Optional.empty | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


