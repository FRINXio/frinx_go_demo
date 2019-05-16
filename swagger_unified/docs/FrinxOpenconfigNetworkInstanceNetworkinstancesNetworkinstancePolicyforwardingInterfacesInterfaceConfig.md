# FrinxOpenconfigNetworkInstanceNetworkinstancesNetworkinstancePolicyforwardingInterfacesInterfaceConfig

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApplyForwardingPolicy** | **string** | Optional[The policy to be applied on the interface. Packets ingress on the referenced interface should be compared to the match criteria within the specified policy, and in the case that these criteria are met, the forwarding actions specified applied. These policies should be applied following quality of service classification, and ACL actions if such entities are referenced by the corresponding interface.] REF:Optional.empty | [optional] [default to null]
**FrinxJuniperPfInterfacesExtensionclassifiers** | [***FrinxJuniperPfInterfacesExtensionJuniperpfinterfaceextensionconfigClassifiers**](frinx.juniper.pf.interfaces.extension.juniperpfinterfaceextensionconfig.Classifiers.md) | Optional[Classify incoming packets based on code point value] REF:Optional.empty | [optional] [default to null]
**FrinxJuniperPfInterfacesExtensionschedulerMap** | **string** | Optional[Output scheduler map] REF:Optional.empty | [optional] [default to null]
**FrinxCiscoPfInterfacesExtensioninputServicePolicy** | **string** | Optional[Service policy which is applied on packets in ingress direction] REF:Optional.empty | [optional] [default to null]
**FrinxCiscoPfInterfacesExtensionoutputServicePolicy** | **string** | Optional[Service policy which is applied on packets in egress direction] REF:Optional.empty | [optional] [default to null]
**InterfaceId** | **string** | Optional[A unique identifier for the interface.] REF:Optional.empty | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


