# FrinxOpenconfigMplsLdpMplsldpgracefulrestartGracefulrestartConfig

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ReconnectTime** | **int32** | Optional[Interval for which the remote LDP peers will wait for the local node to reconnect after a failure] REF:Optional[RFC3478 Graceful Restart Mechanism for Label Distribution Protocol] | [optional] [default to null]
**HelperEnable** | **bool** | Optional[Enables the graceful restart helper for LDP.] REF:Optional.empty | [optional] [default to null]
**Enabled** | **bool** | Optional[When set to true, the functionality within which this leaf is defined is enabled, when set to false it is explicitly disabled.] REF:Optional.empty | [optional] [default to null]
**RecoveryTime** | **int32** | Optional[Interval used to specify the time for the remote peer to maintain the MPLS forwarding state after the local node has succesfully reconnected] REF:Optional[RFC3478 Graceful Restart Mechanism for Label Distribution Protocol] | [optional] [default to null]
**ForwardingHoldtime** | **int32** | Optional[Time that defines the interval for keeping the node in recovery mode.] REF:Optional[RFC3478 Graceful Restart Mechanism for Label Distribution Protocol] | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


