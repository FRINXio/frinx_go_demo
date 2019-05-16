# NotificationsCreatesubscriptionInput

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Filter** | **string** | Optional[An optional parameter that indicates which subset of all possible events is of interest. The format of this parameter is the same as that of the filter parameter in the NETCONF protocol operations. If not present, all events not precluded by other parameters will be sent.] REF:Optional.empty | [optional] [default to null]
**StopTime** | **string** | Optional[An optional parameter used with the optional replay feature to indicate the newest notifications of interest. If stop time is not present, the notifications will continue until the subscription is terminated. Must be used with startTime.] REF:Optional.empty | [optional] [default to null]
**StartTime** | **string** | Optional[A parameter used to trigger the replay feature and indicates that the replay should start at the time specified. If start time is not present, this is not a replay subscription.] REF:Optional.empty | [optional] [default to null]
**Stream** | **string** | Optional[An optional parameter that indicates which stream of events is of interest. If not present, then events in the default NETCONF stream will be sent.] REF:Optional.empty | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


