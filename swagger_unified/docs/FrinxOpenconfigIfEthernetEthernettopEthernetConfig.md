# FrinxOpenconfigIfEthernetEthernettopEthernetConfig

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FrinxIfAggregateExtensionadminKey** | **int32** | Optional[Member port of LACP has key value. All member ports in one aggregator have same key values. To make the aggregator consisted of specified member ports, configure the different key value with the key value of another port.] REF:Optional.empty | [optional] [default to null]
**MacAddress** | **string** | Optional[Assigns a MAC address to the Ethernet interface.  If not specified, the corresponding operational state leaf is expected to show the system-assigned MAC address.] REF:Optional.empty | [optional] [default to null]
**PortSpeed** | [***FrinxOpenconfigIfEthernetPortSpeedIdentityref**](frinx.openconfig.if.ethernet.PortSpeedIdentityref.md) | Optional[When auto-negotiate is TRUE, this optionally sets the port-speed mode that will be advertised to the peer for negotiation.  If unspecified, it is expected that the interface will select the highest speed available based on negotiation.  When auto-negotiate is set to FALSE, sets the link speed to a fixed value -- supported values are defined by ETHERNET_SPEED identities] REF:Optional.empty | [optional] [default to null]
**EnableFlowControl** | **bool** | Optional[Enable or disable flow control for this interface. Ethernet flow control is a mechanism by which a receiver may send PAUSE frames to a sender to stop transmission for a specified time.  This setting should override auto-negotiated flow control settings.  If left unspecified, and auto-negotiate is TRUE, flow control mode is negotiated with the peer interface.] REF:Optional[IEEE 802.3x] | [optional] [default to null]
**FrinxLacpLagMemberinterval** | [***FrinxOpenconfigLacpLacpPeriodType**](frinx.openconfig.lacp.LacpPeriodType.md) | Optional[Set the period between LACP messages -- uses the lacp-period-type enumeration.] REF:Optional.empty | [optional] [default to null]
**FrinxLacpLagMemberlacpMode** | [***FrinxOpenconfigLacpLacpActivityType**](frinx.openconfig.lacp.LacpActivityType.md) | Optional[ACTIVE is to initiate the transmission of LACP packets. PASSIVE is to wait for peer to initiate the transmission of LACP packets.] REF:Optional.empty | [optional] [default to null]
**FrinxOpenconfigIfAggregateaggregateId** | **string** | Optional[Specify the logical aggregate interface to which this interface belongs] REF:Optional.empty | [optional] [default to null]
**AutoNegotiate** | **bool** | Optional[Set to TRUE to request the interface to auto-negotiate transmission parameters with its peer interface.  When set to FALSE, the transmission parameters are specified manually.] REF:Optional[IEEE 802.3-2012 auto-negotiation transmission parameters] | [optional] [default to null]
**DuplexMode** | [***FrinxOpenconfigIfEthernetDuplexModeEnumeration**](frinx.openconfig.if.ethernet.DuplexModeEnumeration.md) | Optional[When auto-negotiate is TRUE, this optionally sets the duplex mode that will be advertised to the peer.  If unspecified, the interface should negotiate the duplex mode directly (typically full-duplex).  When auto-negotiate is FALSE, this sets the duplex mode on the interface directly.] REF:Optional.empty | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


