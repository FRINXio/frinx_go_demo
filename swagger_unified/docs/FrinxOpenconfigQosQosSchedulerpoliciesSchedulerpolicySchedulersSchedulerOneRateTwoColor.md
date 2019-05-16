# FrinxOpenconfigQosQosSchedulerpoliciesSchedulerpolicySchedulersSchedulerOneRateTwoColor

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExceedAction** | [***FrinxOpenconfigQosQosscheduler1r2ctopOneratetwocolorExceedAction**](frinx.openconfig.qos.qosscheduler1r2ctop.oneratetwocolor.ExceedAction.md) | Optional[Action to be applied to packets that are scheduled above the CIR of the one-rate, two-colour shaper. Packets that do not receive a token from the in-CIR bucket are said to be exceeding, and have all of the specified actions applied.] REF:Optional.empty | [optional] [default to null]
**Config** | [***FrinxOpenconfigQosQosSchedulerpoliciesSchedulerpolicySchedulersSchedulerOneratetwocolorConfig**](frinx.openconfig.qos.qos.schedulerpolicies.schedulerpolicy.schedulers.scheduler.oneratetwocolor.Config.md) | Optional[Configuration data for 1 rate, 2 color shapers] REF:Optional.empty | [optional] [default to null]
**ConformAction** | [***FrinxOpenconfigQosQosscheduler1r2ctopOneratetwocolorConformAction**](frinx.openconfig.qos.qosscheduler1r2ctop.oneratetwocolor.ConformAction.md) | Optional[Action to be applied to packets that are scheduled within the CIR of the one-rate, two-colour scheduler. Packets that receive a token from the in-CIR bucket are said to be conforming and have all of the specified actions applied.] REF:Optional.empty | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


