# FrinxOpenconfigQosQosscheduler2r3ctopTwoRateThreeColor

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExceedAction** | [***FrinxOpenconfigQosQosscheduler2r3ctopTworatethreecolorExceedAction**](frinx.openconfig.qos.qosscheduler2r3ctop.tworatethreecolor.ExceedAction.md) | Optional[Action to be applied to the packets that are scheduled within the PIR of the policer. Packets that receive a token from within the PIR allocation have all the specified actions applied to them] REF:Optional.empty | [optional] [default to null]
**Config** | [***FrinxOpenconfigQosQosscheduler2r3ctopTworatethreecolorConfig**](frinx.openconfig.qos.qosscheduler2r3ctop.tworatethreecolor.Config.md) | Optional[Configuration data for 2 rate, 3 color policers.] REF:Optional.empty | [optional] [default to null]
**ViolateAction** | [***FrinxOpenconfigQosQosscheduler2r3ctopTworatethreecolorViolateAction**](frinx.openconfig.qos.qosscheduler2r3ctop.tworatethreecolor.ViolateAction.md) | Optional[Action to be applied to the packets that are scheduled above the PIR of the policer. Packets that do not receive a token from either bucket have all specified actions applied to them.] REF:Optional.empty | [optional] [default to null]
**ConformAction** | [***FrinxOpenconfigQosQosscheduler2r3ctopTworatethreecolorConformAction**](frinx.openconfig.qos.qosscheduler2r3ctop.tworatethreecolor.ConformAction.md) | Optional[Action to be applied to the packets that are scheduled within the CIR of the policer. All packets that receive a token from this bucket have all actions specified applied to them] REF:Optional.empty | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


