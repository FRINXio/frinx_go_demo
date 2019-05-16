# FrinxOpenconfigBgpBgppeergroupbaseConfig

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PeerType** | [***FrinxOpenconfigBgpTypesPeerType**](frinx.openconfig.bgp.types.PeerType.md) | Optional[Explicitly designate the peer or peer group as internal (iBGP) or external (eBGP).] REF:Optional.empty | [optional] [default to null]
**RemovePrivateAs** | [***FrinxOpenconfigBgpTypesRemovePrivateAsOption**](frinx.openconfig.bgp.types.RemovePrivateAsOption.md) | Optional[Remove private AS numbers from updates sent to peers - when this leaf is not specified, the AS_PATH attribute should be sent to the peer unchanged] REF:Optional.empty | [optional] [default to null]
**SendCommunity** | [***FrinxOpenconfigBgpTypesCommunityType**](frinx.openconfig.bgp.types.CommunityType.md) | Optional[Specify which types of community should be sent to the neighbor or group. The default is to not send the community attribute] REF:Optional.empty | [optional] [default to null]
**Description** | **string** | Optional[An optional textual description (intended primarily for use with a peer or group] REF:Optional.empty | [optional] [default to null]
**PeerGroupName** | **string** | Optional[Name of the BGP peer-group] REF:Optional.empty | [optional] [default to null]
**PeerAs** | **int64** | Optional[AS number of the peer.] REF:Optional.empty | [optional] [default to null]
**RouteFlapDamping** | **bool** | Optional[Enable route flap damping.] REF:Optional.empty | [optional] [default to null]
**AuthPassword** | **string** | Optional[Configures an MD5 authentication password for use with neighboring devices.] REF:Optional.empty | [optional] [default to null]
**LocalAs** | **int64** | Optional[The local autonomous system number that is to be used when establishing sessions with the remote peer or peer group, if this differs from the global BGP router autonomous system number.] REF:Optional.empty | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


