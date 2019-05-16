# FrinxOpenconfigIsisIsisglobalbaseConfig

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MaxEcmpPaths** | **int32** | Optional[ISIS max-paths count.] REF:Optional.empty | [optional] [default to null]
**Instance** | **string** | Optional[ISIS Instance.] REF:Optional.empty | [optional] [default to null]
**AuthenticationCheck** | **bool** | Optional[When set to true, reject all ISIS protocol PDUs that either have a mismatch in authentication-type or authentication-key.] REF:Optional.empty | [optional] [default to null]
**MaximumAreaAddresses** | **int32** | Optional[Maximum areas supported.] REF:Optional.empty | [optional] [default to null]
**IidTlv** | **bool** | Optional[ISIS Instance Identifier TLV. When set to trues, the IID-TLV identifies the unique instance as well as the topology/topologies to which the PDU applies.] REF:Optional[RFC6822: IS-IS Multi-Instance. TLV 7] | [optional] [default to null]
**LevelCapability** | [***FrinxOpenconfigIsisTypesLevelType**](frinx.openconfig.isis.types.LevelType.md) | Optional[ISIS level capability(level-1, level-2,vlevel-1-2).] REF:Optional.empty | [optional] [default to null]
**FastFlooding** | **bool** | Optional[When set to true, IS will always flood the LSP that triggered an SPF before the router actually runs the SPF computation.] REF:Optional.empty | [optional] [default to null]
**PoiTlv** | **bool** | Optional[ISIS purge TLV. When set to true, a TLV is added to purges to record the system ID  of the IS generating the purge.] REF:Optional[RFC6232: Purge Originator Identification TLV for IS-IS. TLV 13.] | [optional] [default to null]
**Net** | **[]string** | Optional[ISIS network entity title (NET). The first 8 bits are usually 49 (private AFI), next 16 bits represent area, next 48 bits represent system id and final 8 bits are set to 0.] REF:Optional[International Organization for Standardization, Information technology - Open Systems Interconnection-Network service Definition - ISO/ IEC 8348:2002.] | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


