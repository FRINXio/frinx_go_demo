# FrinxOpenconfigMplsTetunnelstopTunnelsTunnelConfig

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ReoptimizeTimer** | **int32** | Optional[frequency of reoptimization of a traffic engineered LSP] REF:Optional.empty | [optional] [default to null]
**ShortcutEligible** | **bool** | Optional[Whether this LSP is considered to be eligible for us as a shortcut in the IGP. In the case that this leaf is set to true, the IGP SPF calculation uses the metric specified to determine whether traffic should be carried over this LSP] REF:Optional.empty | [optional] [default to null]
**MetricType** | [***FrinxOpenconfigMplsMetricTypeIdentityref**](frinx.openconfig.mpls.MetricTypeIdentityref.md) | Optional[The type of metric specification that should be used to set the LSP(s) metric] REF:Optional.empty | [optional] [default to null]
**Preference** | **int32** | Optional[Specifies a preference for this tunnel. A lower number signifies a better preference] REF:Optional.empty | [optional] [default to null]
**Description** | **string** | Optional[optional text description for the tunnel] REF:Optional.empty | [optional] [default to null]
**Source** | **string** | Optional[RSVP-TE tunnel source address] REF:Optional.empty | [optional] [default to null]
**Type_** | [***FrinxOpenconfigMplsTypeIdentityref**](frinx.openconfig.mpls.TypeIdentityref.md) | Optional[Tunnel type, p2p or p2mp] REF:Optional.empty | [optional] [default to null]
**HoldPriority** | **int32** | Optional[preemption priority once the LSP is established, lower is higher priority; default 0 indicates other LSPs will not preempt the LSPs once established] REF:Optional[RFC 3209 - RSVP-TE: Extensions to RSVP for LSP Tunnels] | [optional] [default to null]
**ProtectionStyleRequested** | [***FrinxOpenconfigMplsProtectionStyleRequestedIdentityref**](frinx.openconfig.mpls.ProtectionStyleRequestedIdentityref.md) | Optional[style of mpls frr protection desired: can be link, link-node or unprotected.] REF:Optional.empty | [optional] [default to null]
**AdminStatus** | [***FrinxOpenconfigMplsAdminStatusIdentityref**](frinx.openconfig.mpls.AdminStatusIdentityref.md) | Optional[TE tunnel administrative state.] REF:Optional.empty | [optional] [default to null]
**SoftPreemption** | **bool** | Optional[Enables RSVP soft-preemption on this LSP] REF:Optional.empty | [optional] [default to null]
**Metric** | **int32** | Optional[The value of the metric that should be specified. The value supplied in this leaf is used in conjunction with the metric type to determine the value of the metric used by the system. Where the metric-type is set to LSP_METRIC_ABSOLUTE - the value of this leaf is used directly; where it is set to LSP_METRIC_RELATIVE, the relevant (positive or negative) offset is used to formulate the metric; where metric-type is LSP_METRIC_INHERITED, the value of this leaf is not utilised] REF:Optional.empty | [optional] [default to null]
**SignalingProtocol** | [***FrinxOpenconfigMplsSignalingProtocolIdentityref**](frinx.openconfig.mpls.SignalingProtocolIdentityref.md) | Optional[Signaling protocol used to set up this tunnel] REF:Optional.empty | [optional] [default to null]
**Name** | **string** | Optional[The tunnel name] REF:Optional.empty | [optional] [default to null]
**SetupPriority** | **int32** | Optional[RSVP-TE preemption priority during LSP setup, lower is higher priority; default 7 indicates that LSP will not preempt established LSPs during setup] REF:Optional[RFC 3209 - RSVP-TE: Extensions to RSVP for LSP Tunnels] | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


