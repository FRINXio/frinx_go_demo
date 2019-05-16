# NetconfNodeTopologyCreatedeviceInput

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConcurrentRpcLimit** | **int32** | Optional[Limit of concurrent messages that can be send before reply messages are received. If value &lt;1 is provided, no limit will be enforced] REF:Optional.empty | [optional] [default to null]
**NodeId** | **string** | Optional.empty REF:Optional.empty | [optional] [default to null]
**MaxConnectionAttempts** | **int64** | Optional[Maximum number of connection retries. Non positive value or null is interpreted as infinity.] REF:Optional.empty | [optional] [default to null]
**TcpOnly** | **bool** | Optional.empty REF:Optional.empty | [optional] [default to null]
**UnavailableCapabilities** | [***NetconfNodeTopologyNetconfnodeconnectionstatusUnavailableCapabilities**](netconf.node.topology.netconfnodeconnectionstatus.UnavailableCapabilities.md) |  | [optional] [default to null]
**ReconnectOnChangedSchema** | **bool** | Optional[If true, the connector would auto disconnect/reconnect when schemas are changed in the remote device. The connector subscribes (right after connect) to base netconf notifications and listens for netconf-capability-change notification] REF:Optional.empty | [optional] [default to null]
**DryRunJournalSize** | **int32** | Optional[Size of the DRY RUN netconf mountpoint jounral. DRY RUN journal captures netconf RPCs that would be executed when reading/writing some configuration. However the RPCs are not actually sent to the device] REF:Optional.empty | [optional] [default to null]
**NonModuleCapabilities** | [***NetconfNodeTopologyNetconfnodeconnectionparametersNonModuleCapabilities**](netconf.node.topology.netconfnodeconnectionparameters.NonModuleCapabilities.md) |  | [optional] [default to null]
**YangModuleCapabilities** | [***NetconfNodeTopologyNetconfnodeconnectionparametersYangModuleCapabilities**](netconf.node.topology.netconfnodeconnectionparameters.YangModuleCapabilities.md) |  | [optional] [default to null]
**Password** | **string** | Optional.empty REF:Optional.empty | [optional] [default to null]
**SleepFactor** | **int32** | Optional.empty REF:Optional.empty | [optional] [default to null]
**Host** | **string** | Optional.empty REF:Optional.empty | [optional] [default to null]
**OdlHelloMessageCapabilities** | [***NetconfNodeTopologyNetconfnodeconnectionparametersOdlHelloMessageCapabilities**](netconf.node.topology.netconfnodeconnectionparameters.OdlHelloMessageCapabilities.md) |  | [optional] [default to null]
**ClusteredConnectionStatus** | [***NetconfNodeTopologyNetconfnodeconnectionstatusClusteredConnectionStatus**](netconf.node.topology.netconfnodeconnectionstatus.ClusteredConnectionStatus.md) |  | [optional] [default to null]
**YangLibrary** | [***NetconfNodeTopologyNetconfschemastorageYangLibrary**](netconf.node.topology.netconfschemastorage.YangLibrary.md) |  | [optional] [default to null]
**ConnectionStatus** | [***NetconfNodeTopologyConnectionStatusEnumeration**](netconf.node.topology.ConnectionStatusEnumeration.md) | Optional.empty REF:Optional.empty | [optional] [default to null]
**BetweenAttemptsTimeoutMillis** | **int32** | Optional[Initial timeout in milliseconds to wait between connection attempts. Will be multiplied by sleep-factor with every additional attempt] REF:Optional.empty | [optional] [default to null]
**KeyBased** | [***NetconfNodeTopologyNetconfnodecredentialsCredentialsKeyauthKeyBased**](netconf.node.topology.netconfnodecredentials.credentials.keyauth.KeyBased.md) |  | [optional] [default to null]
**KeepaliveDelay** | **int64** | Optional[Netconf connector sends keepalive RPCs while the session is idle, this delay specifies the delay between keepalive RPC in seconds If a value &lt;1 is provided, no keepalives will be sent] REF:Optional.empty | [optional] [default to null]
**Schemaless** | **bool** | Optional.empty REF:Optional.empty | [optional] [default to null]
**DefaultRequestTimeoutMillis** | **int64** | Optional[Timeout for blocking operations within transactions.] REF:Optional.empty | [optional] [default to null]
**ConnectionTimeoutMillis** | **int64** | Optional[Specifies timeout in milliseconds after which connection must be established.] REF:Optional.empty | [optional] [default to null]
**Port** | **int32** | Optional.empty REF:Optional.empty | [optional] [default to null]
**AvailableCapabilities** | [***NetconfNodeTopologyNetconfnodeconnectionstatusAvailableCapabilities**](netconf.node.topology.netconfnodeconnectionstatus.AvailableCapabilities.md) |  | [optional] [default to null]
**LoginPasswordUnencrypted** | [***NetconfNodeTopologyNetconfnodecredentialsCredentialsLoginpwunencryptedLoginPasswordUnencrypted**](netconf.node.topology.netconfnodecredentials.credentials.loginpwunencrypted.LoginPasswordUnencrypted.md) |  | [optional] [default to null]
**CustomizationFactory** | **string** | Optional.empty REF:Optional.empty | [optional] [default to null]
**LoginPassword** | [***NetconfNodeTopologyNetconfnodecredentialsCredentialsLoginpwunencryptedLoginPasswordUnencrypted**](netconf.node.topology.netconfnodecredentials.credentials.loginpwunencrypted.LoginPasswordUnencrypted.md) |  | [optional] [default to null]
**SchemaCacheDirectory** | **string** | Optional[The destination schema repository for yang files relative to the cache directory.  This may be specified per netconf mount so that the loaded yang files are stored to a distinct directory to avoid potential conflict.] REF:Optional.empty | [optional] [default to null]
**ConnectedMessage** | **string** | Optional.empty REF:Optional.empty | [optional] [default to null]
**ActorResponseWaitTime** | **int32** | Optional[Time that slave actor will wait for response from master.] REF:Optional.empty | [optional] [default to null]
**PassThrough** | [***interface{}**](interface{}.md) | Optional[When the underlying node is connected, its NETCONF context is available verbatim under this container through the mount extension.] REF:Optional.empty | [optional] [default to null]
**Username** | **string** | Optional.empty REF:Optional.empty | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


