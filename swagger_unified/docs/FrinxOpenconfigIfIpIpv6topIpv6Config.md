# FrinxOpenconfigIfIpIpv6topIpv6Config

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DupAddrDetectTransmits** | **int64** | Optional[[adapted from IETF IP model RFC 7277]  The number of consecutive Neighbor Solicitation messages sent while performing Duplicate Address Detection on a tentative address.  A value of zero indicates that Duplicate Address Detection is not performed on tentative addresses.  A value of one indicates a single transmission with no follow-up retransmissions.] REF:Optional[RFC 4862: IPv6 Stateless Address Autoconfiguration] | [optional] [default to null]
**Enabled** | **bool** | Optional[[adapted from IETF IP model RFC 7277]  Controls whether IPv6 is enabled or disabled on this interface.  When IPv6 is enabled, this interface is connected to an IPv6 stack, and the interface can send and receive IPv6 packets.] REF:Optional.empty | [optional] [default to null]
**Mtu** | **int64** | Optional[[adapted from IETF IP model RFC 7277]  The size, in octets, of the largest IPv6 packet that the interface will send and receive.  The server may restrict the allowed values for this leaf, depending on the interface&#39;s type.  If this leaf is not configured, the operationally used MTU depends on the interface&#39;s type.] REF:Optional[RFC 2460: Internet Protocol, Version 6 (IPv6) Specification          Section 5] | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


