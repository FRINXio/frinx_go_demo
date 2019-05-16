# FrinxOpenconfigBgpBgpcommonrouteselectionoptionsRouteselectionoptionsConfig

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AlwaysCompareMed** | **bool** | Optional[Compare multi-exit discriminator (MED) value from different ASes when selecting the best route.  The default behavior is to only compare MEDs for paths received from the same AS.] REF:Optional.empty | [optional] [default to null]
**AdvertiseInactiveRoutes** | **bool** | Optional[Advertise inactive routes to external peers.  The default is to only advertise active routes.] REF:Optional.empty | [optional] [default to null]
**IgnoreAsPathLength** | **bool** | Optional[Ignore the AS path length when selecting the best path. The default is to use the AS path length and prefer paths with shorter length.] REF:Optional.empty | [optional] [default to null]
**EnableAigp** | **bool** | Optional[Flag to enable sending / receiving accumulated IGP attribute in routing updates] REF:Optional.empty | [optional] [default to null]
**IgnoreNextHopIgpMetric** | **bool** | Optional[Ignore the IGP metric to the next-hop when calculating BGP best-path. The default is to select the route for which the metric to the next-hop is lowest] REF:Optional.empty | [optional] [default to null]
**ExternalCompareRouterId** | **bool** | Optional[When comparing similar routes received from external BGP peers, use the router-id as a criterion to select the active path.] REF:Optional.empty | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


