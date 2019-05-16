# \JournalApi

All URIs are relative to *http://172.17.0.1:8181/restconf*

Method | HTTP request | Description
------------- | ------------- | -------------
[**RpcJournalClearJournal**](JournalApi.md#RpcJournalClearJournal) | **Post** /operations/journal:clear-journal/ | 
[**RpcJournalReadJournal**](JournalApi.md#RpcJournalReadJournal) | **Post** /operations/journal:read-journal/ | 


# **RpcJournalClearJournal**
> JournalClearJournal RpcJournalClearJournal(ctx, )


### Required Parameters
This endpoint does not need any parameter.

### Return type

[**JournalClearJournal**](journal.ClearJournal.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RpcJournalReadJournal**
> JournalReadJournal RpcJournalReadJournal(ctx, )


### Required Parameters
This endpoint does not need any parameter.

### Return type

[**JournalReadJournal**](journal.ReadJournal.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

