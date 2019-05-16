# FrinxOpenconfigLldpLldptopLldpConfig

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SystemDescription** | **string** | Optional[The system description field shall contain an alpha-numeric string that is the textual description of the network entity. The system description should include the full name and version identification of the system&#39;s hardware type, software operating system, and networking software. If implementations support IETF RFC 3418, the sysDescr object should be used for this field.] REF:Optional.empty | [optional] [default to null]
**SuppressTlvAdvertisement** | [**[]FrinxOpenconfigLldpSuppressTlvAdvertisementIdentityref**](frinx.openconfig.lldp.SuppressTlvAdvertisementIdentityref.md) | Optional[Indicates whether the local system should suppress the advertisement of particular TLVs with the LLDP PDUs that it transmits. Where a TLV type is specified within this list, it should not be included in any LLDP PDU transmitted by the local agent.] REF:Optional.empty | [optional] [default to null]
**SystemName** | **string** | Optional[The system name field shall contain an alpha-numeric string that indicates the system&#39;s administratively assigned name. The system name should be the system&#39;s fully qualified domain name. If implementations support IETF RFC 3418, the sysName object should be used for this field.] REF:Optional.empty | [optional] [default to null]
**HelloTimer** | **int32** | Optional[System level hello timer for the LLDP protocol.] REF:Optional.empty | [optional] [default to null]
**ChassisIdType** | [***FrinxOpenconfigLldpTypesChassisIdType**](frinx.openconfig.lldp.types.ChassisIdType.md) | Optional[This field identifies the format and source of the chassis identifier string. It is an enumerator defined by the LldpChassisIdSubtype object from IEEE 802.1AB MIB.] REF:Optional.empty | [optional] [default to null]
**ChassisId** | **string** | Optional[The Chassis ID is a mandatory TLV which identifies the chassis component of the endpoint identifier associated with the transmitting LLDP agent] REF:Optional.empty | [optional] [default to null]
**Enabled** | **bool** | Optional[System level state of the LLDP protocol.] REF:Optional.empty | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


