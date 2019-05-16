# FrinxOpenconfigMplsP2pprimarypathstopP2pprimarypathP2pPrimaryPath

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Optional[Path name] REF:Optional.empty | [optional] [default to null]
**AdminGroups** | [***FrinxOpenconfigMplsTepathplacementconstraintstopAdminGroups**](frinx.openconfig.mpls.tepathplacementconstraintstop.AdminGroups.md) | Optional[Top-level container for include/exclude constraints for link affinities] REF:Optional.empty | [optional] [default to null]
**CandidateSecondaryPaths** | [***FrinxOpenconfigMplsP2pprimarypathstopP2pprimarypathP2pprimarypathCandidateSecondaryPaths**](frinx.openconfig.mpls.p2pprimarypathstop.p2pprimarypath.p2pprimarypath.CandidateSecondaryPaths.md) | Optional[The set of candidate secondary paths which may be used for this primary path. When secondary paths are specified in the list the path of the secondary LSP in use must be restricted to those path options referenced. The priority of the secondary paths is specified within the list. Higher priority values are less preferred - that is to say that a path with priority 0 is the most preferred path. In the case that the list is empty, any secondary path option may be utilised when the current primary path is in use.] REF:Optional.empty | [optional] [default to null]
**Config** | [***FrinxOpenconfigMplsP2pprimarypathstopP2pprimarypathP2pprimarypathConfig**](frinx.openconfig.mpls.p2pprimarypathstop.p2pprimarypath.p2pprimarypath.Config.md) | Optional[Configuration parameters related to paths] REF:Optional.empty | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


