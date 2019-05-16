# NetconfKeystoreKeystoreentryKeyCredential

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Passphrase** | **string** | Optional[If the provided key is encrypted by a passphrase this needs to be included. Leave empty if the key does not have a passphrase. DO NOT write write this directly into the datastore, use the provided rpc&#39;s as these will encrypt the passhprase before the entry is written into the datastore.] REF:Optional.empty | [optional] [default to null]
**KeyId** | **string** | Optional.empty REF:Optional.empty | [optional] [default to null]
**PrivateKey** | **string** | Optional[Base64 encoded private key that should be used for authentication with a netconf device. Do not include a public key as that is calculated from the private key. DO NOT write this directly into the datastore, use the provided rpc&#39;s as these will encrypt the key before the entry is written into the datastore.] REF:Optional.empty | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


