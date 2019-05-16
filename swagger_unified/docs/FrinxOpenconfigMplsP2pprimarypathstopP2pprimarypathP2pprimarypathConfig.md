# FrinxOpenconfigMplsP2pprimarypathstopP2pprimarypathP2pprimarypathConfig

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PathComputationServer** | **string** | Optional[Address of the external path computation server] REF:Optional.empty | [optional] [default to null]
**CspfTiebreaker** | [***FrinxOpenconfigMplsCspfTieBreaking**](frinx.openconfig.mpls.CspfTieBreaking.md) | Optional[Determine the tie-breaking method to choose between equally desirable paths during CSFP computation] REF:Optional.empty | [optional] [default to null]
**UseCspf** | **bool** | Optional[Flag to enable CSPF for locally computed LSPs] REF:Optional.empty | [optional] [default to null]
**Preference** | **int32** | Optional[Specifies a preference for this path. The lower the number higher the preference] REF:Optional.empty | [optional] [default to null]
**Name** | **string** | Optional[Path name] REF:Optional.empty | [optional] [default to null]
**RetryTimer** | **int32** | Optional[sets the time between attempts to establish the LSP] REF:Optional.empty | [optional] [default to null]
**PathComputationMethod** | [***FrinxOpenconfigMplsPathComputationMethodIdentityref**](frinx.openconfig.mpls.PathComputationMethodIdentityref.md) | Optional[The method used for computing the path, either locally computed, queried from a server or not computed at all (explicitly configured).] REF:Optional.empty | [optional] [default to null]
**SetupPriority** | **int32** | Optional[RSVP-TE preemption priority during LSP setup, lower is higher priority; default 7 indicates that LSP will not preempt established LSPs during setup] REF:Optional[RFC 3209 - RSVP-TE: Extensions to RSVP for LSP Tunnels] | [optional] [default to null]
**HoldPriority** | **int32** | Optional[preemption priority once the LSP is established, lower is higher priority; default 0 indicates other LSPs will not preempt the LSPs once established] REF:Optional[RFC 3209 - RSVP-TE: Extensions to RSVP for LSP Tunnels] | [optional] [default to null]
**ExplicitPathName** | **string** | Optional[reference to a defined path] REF:Optional.empty | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


