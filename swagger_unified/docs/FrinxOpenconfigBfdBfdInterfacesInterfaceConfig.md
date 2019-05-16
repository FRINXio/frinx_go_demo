# FrinxOpenconfigBfdBfdInterfacesInterfaceConfig

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DetectionMultiplier** | **int32** | Optional[The number of packets that must be missed to declare this session as down. The detection interval for the BFD session is calculated by multiplying the value of the negotiated transmission interval by this value.] REF:Optional.empty | [optional] [default to null]
**DesiredMinimumTxInterval** | **int64** | Optional[The minimum interval between transmission of BFD control packets that the operator desires. This value is advertised to the peer, however the actual interval used is specified by taking the maximum of desired-minimum-tx-interval and the value of the remote required-minimum-receive interval value.  This value is specified as an integer number of microseconds.] REF:Optional.empty | [optional] [default to null]
**FrinxBfdExtensionremoteAddress** | **string** | Optional[The IP address used by the remote system for this BFD session.] REF:Optional.empty | [optional] [default to null]
**EnablePerMemberLink** | **bool** | Optional[When this leaf is set to true - BFD will be enabled on each member interface of the aggregated Ethernet bundle.] REF:Optional.empty | [optional] [default to null]
**Id** | **string** | Optional[A unique identifier for the interface.] REF:Optional.empty | [optional] [default to null]
**Enabled** | **bool** | Optional[When this leaf is set to true then the BFD session is enabled on the specified interface - if it is set to false, it is administratively disabled.] REF:Optional.empty | [optional] [default to null]
**RequiredMinimumReceive** | **int64** | Optional[The minimum interval between received BFD control packets that this system should support. This value is advertised to the remote peer to indicate the maximum frequency (i.e., minimum inter-packet interval) between BFD control packets that is acceptable to the local system.] REF:Optional.empty | [optional] [default to null]
**LocalAddress** | **string** | Optional[The source IP address to be used for BFD sessions over this interface.] REF:Optional.empty | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


