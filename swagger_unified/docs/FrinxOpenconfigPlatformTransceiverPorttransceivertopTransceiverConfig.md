# FrinxOpenconfigPlatformTransceiverPorttransceivertopTransceiverConfig

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EthernetPmdPreconf** | [***FrinxOpenconfigPlatformTransceiverEthernetPmdPreconfIdentityref**](frinx.openconfig.platform.transceiver.EthernetPmdPreconfIdentityref.md) | Optional[The Ethernet PMD is a property of the optical transceiver used on the port, indicating the type of physical connection. It is included in configuration data to allow pre-configuring a port/transceiver with the expected PMD.  The actual PMD is indicated by the ethernet-pmd state leaf.] REF:Optional.empty | [optional] [default to null]
**FormFactorPreconf** | [***FrinxOpenconfigPlatformTransceiverFormFactorPreconfIdentityref**](frinx.openconfig.platform.transceiver.FormFactorPreconfIdentityref.md) | Optional[Indicates the type of optical transceiver used on this port.  If the client port is built into the device and not pluggable, then non-pluggable is the corresponding state. If a device port supports multiple form factors (e.g. QSFP28 and QSFP+, then the value of the transceiver installed shall be reported. If no transceiver is present, then the value of the highest rate form factor shall be reported (QSFP28, for example).  The form factor is included in configuration data to allow pre-configuring a device with the expected type of transceiver ahead of deployment.  The corresponding state leaf should reflect the actual transceiver type plugged into the system.] REF:Optional.empty | [optional] [default to null]
**Enabled** | **bool** | Optional[Turns power on / off to the transceiver -- provides a means to power on/off the transceiver (in the case of SFP, SFP+, QSFP,...) or enable high-power mode (in the case of CFP, CFP2, CFP4) and is optionally supported (device can choose to always enable).  True &#x3D; power on / high power, False &#x3D; powered off] REF:Optional.empty | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


