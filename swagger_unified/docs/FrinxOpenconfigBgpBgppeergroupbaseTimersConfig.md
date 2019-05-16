# FrinxOpenconfigBgpBgppeergroupbaseTimersConfig

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MinimumAdvertisementInterval** | **int32** | Optional[Minimum time which must elapse between subsequent UPDATE messages relating to a common set of NLRI being transmitted to a peer. This timer is referred to as MinRouteAdvertisementIntervalTimer by RFC 4721 and serves to reduce the number of UPDATE messages transmitted when a particular set of NLRI exhibit instability.] REF:Optional[RFC 4271 - A Border Gateway Protocol 4, Sec 9.2.1.1] | [optional] [default to null]
**ConnectRetry** | **int32** | Optional[Time interval in seconds between attempts to establish a session with the peer.] REF:Optional.empty | [optional] [default to null]
**KeepaliveInterval** | **int32** | Optional[Time interval in seconds between transmission of keepalive messages to the neighbor.  Typically set to 1/3 the hold-time.] REF:Optional.empty | [optional] [default to null]
**HoldTime** | **int32** | Optional[Time interval in seconds that a BGP session will be considered active in the absence of keepalive or other messages from the peer.  The hold-time is typically set to 3x the keepalive-interval.] REF:Optional[RFC 4271 - A Border Gateway Protocol 4, Sec. 10] | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


