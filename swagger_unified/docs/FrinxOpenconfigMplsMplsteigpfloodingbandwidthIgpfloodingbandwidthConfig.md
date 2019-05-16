# FrinxOpenconfigMplsMplsteigpfloodingbandwidthIgpfloodingbandwidthConfig

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UpThresholds** | **[]int32** | Optional[The thresholds (expressed as a percentage of the maximum reservable bandwidth) at which bandwidth updates are to be triggered when the bandwidth is increasing.] REF:Optional.empty | [optional] [default to null]
**DeltaPercentage** | **int32** | Optional[The percentage of the maximum-reservable-bandwidth considered as the delta that results in an IGP update being flooded] REF:Optional.empty | [optional] [default to null]
**DownThresholds** | **[]int32** | Optional[The thresholds (expressed as a percentage of the maximum reservable bandwidth) at which bandwidth updates are to be triggered when the bandwidth is decreasing.] REF:Optional.empty | [optional] [default to null]
**UpDownThresholds** | **[]int32** | Optional[The thresholds (expressed as a percentage of the maximum reservable bandwidth of the interface) at which bandwidth updates are flooded - used both when the bandwidth is increasing and decreasing] REF:Optional.empty | [optional] [default to null]
**ThresholdSpecification** | [***FrinxOpenconfigMplsThresholdSpecificationEnumeration**](frinx.openconfig.mpls.ThresholdSpecificationEnumeration.md) | Optional[This value specifies whether a single set of threshold values should be used for both increasing and decreasing bandwidth when determining whether to trigger updated bandwidth values to be flooded in the IGP TE extensions. MIRRORED-UP-DOWN indicates that a single value (or set of values) should be used for both increasing and decreasing values, where SEPARATE-UP-DOWN specifies that the increasing and decreasing values will be separately specified] REF:Optional.empty | [optional] [default to null]
**ThresholdType** | [***FrinxOpenconfigMplsThresholdTypeEnumeration**](frinx.openconfig.mpls.ThresholdTypeEnumeration.md) | Optional[The type of threshold that should be used to specify the values at which bandwidth is flooded. DELTA indicates that the local system should flood IGP updates when a change in reserved bandwidth &gt;&#x3D; the specified delta occurs on the interface. Where THRESHOLD_CROSSED is specified, the local system should trigger an update (and hence flood) the reserved bandwidth when the reserved bandwidth changes such that it crosses, or becomes equal to one of the threshold values] REF:Optional.empty | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


