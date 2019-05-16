# IetfNetconfCommitInput

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Persist** | **string** | Optional[This parameter is used to make a confirmed commit persistent.  A persistent confirmed commit is not aborted if the NETCONF session terminates.  The only way to abort a persistent confirmed commit is to let the timer expire, or to use the &lt;cancel-commit&gt; operation.  The value of this parameter is a token that must be given in the &#39;persist-id&#39; parameter of &lt;commit&gt; or &lt;cancel-commit&gt; operations in order to confirm or cancel the persistent confirmed commit.  The token should be a random string.] REF:Optional[RFC 6241, Section 8.3.4.1] | [optional] [default to null]
**Confirmed** | **string** | Optional[Requests a confirmed commit.] REF:Optional[RFC 6241, Section 8.3.4.1] | [optional] [default to null]
**ConfirmTimeout** | **int64** | Optional[The timeout interval for a confirmed commit.] REF:Optional[RFC 6241, Section 8.3.4.1] | [optional] [default to null]
**PersistId** | **string** | Optional[This parameter is given in order to commit a persistent confirmed commit.  The value must be equal to the value given in the &#39;persist&#39; parameter to the &lt;commit&gt; operation. If it does not match, the operation fails with an &#39;invalid-value&#39; error.] REF:Optional[RFC 6241, Section 8.3.4.1] | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


