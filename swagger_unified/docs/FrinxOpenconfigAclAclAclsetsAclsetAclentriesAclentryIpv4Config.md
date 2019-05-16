# FrinxOpenconfigAclAclAclsetsAclsetAclentriesAclentryIpv4Config

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**HopLimit** | **int32** | Optional[The IP packet&#39;s hop limit -- known as TTL (in hops) in IPv4 packets, and hop limit in IPv6] REF:Optional.empty | [optional] [default to null]
**Protocol** | **string** | Optional[The protocol carried in the IP packet, expressed either as its IP protocol number, or by a defined identity.] REF:Optional.empty | [optional] [default to null]
**Dscp** | **int32** | Optional[Value of diffserv codepoint.] REF:Optional.empty | [optional] [default to null]
**FrinxAclExtensionsourceAddressWildcarded** | [***FrinxAclExtensionSrcdstipv4addresswildcardedSourceAddressWildcarded**](frinx.acl.extension.srcdstipv4addresswildcarded.SourceAddressWildcarded.md) | Optional.empty REF:Optional.empty | [optional] [default to null]
**FrinxAclExtensionhopRange** | **string** | Optional[The IP packet&#39;s hop range (TTL)] REF:Optional.empty | [optional] [default to null]
**FrinxAclExtensiondestinationAddressWildcarded** | [***FrinxAclExtensionSrcdstipv4addresswildcardedDestinationAddressWildcarded**](frinx.acl.extension.srcdstipv4addresswildcarded.DestinationAddressWildcarded.md) | Optional.empty REF:Optional.empty | [optional] [default to null]
**DestinationAddress** | **string** | Optional[Destination IPv4 address prefix.] REF:Optional.empty | [optional] [default to null]
**SourceAddress** | **string** | Optional[Source IPv4 address prefix.] REF:Optional.empty | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


