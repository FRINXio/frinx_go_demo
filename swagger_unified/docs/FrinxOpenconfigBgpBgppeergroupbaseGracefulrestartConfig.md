# FrinxOpenconfigBgpBgppeergroupbaseGracefulrestartConfig

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RestartTime** | **int32** | Optional[Estimated time (in seconds) for the local BGP speaker to restart a session. This value is advertise in the graceful restart BGP capability.  This is a 12-bit value, referred to as Restart Time in RFC4724.  Per RFC4724, the suggested default value is &lt;&#x3D; the hold-time value.] REF:Optional.empty | [optional] [default to null]
**HelperOnly** | **bool** | Optional[Enable graceful-restart in helper mode only. When this leaf is set, the local system does not retain forwarding its own state during a restart, but supports procedures for the receiving speaker, as defined in RFC4724.] REF:Optional.empty | [optional] [default to null]
**StaleRoutesTime** | **int32** | Optional[An upper-bound on the time thate stale routes will be retained by a router after a session is restarted. If an End-of-RIB (EOR) marker is received prior to this timer expiring stale-routes will be flushed upon its receipt - if no EOR is received, then when this timer expires stale paths will be purged. This timer is referred to as the Selection_Deferral_Timer in RFC4724] REF:Optional.empty | [optional] [default to null]
**Enabled** | **bool** | Optional[Enable or disable the graceful-restart capability.] REF:Optional.empty | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


