# FrinxOpenconfigQosQosschedulerinputstopInputsInputConfig

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Weight** | **int32** | Optional[For priority schedulers, this indicates the priority of the corresponding input.  Higher values indicates lower priority.  For weighted round-robin schedulers, this leaf indicates the weight of the corresponding input.] REF:Optional.empty | [optional] [default to null]
**InputType** | [***FrinxOpenconfigQosInputTypeEnumeration**](frinx.openconfig.qos.InputTypeEnumeration.md) | Optional[Describes the type of input source for the scheduler] REF:Optional.empty | [optional] [default to null]
**Id** | **string** | Optional[User-defined identifier for the scheduler input] REF:Optional.empty | [optional] [default to null]
**Queue** | **string** | Optional[Reference to a queue that is an input source for the scheduler] REF:Optional.empty | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


