/*
 * southbound
 *
 * API generated from yang definitions: [aaa-encrypt-service-config,cli-topology,cli-translate-registry,cluster-singleton-service-impl,cluster-singleton-service-spi,config,fake,general-entity,ietf-inet-types,ietf-netconf,ietf-netconf-monitoring,ietf-netconf-monitoring-extension,ietf-netconf-notifications,ietf-yang-library,ietf-yang-types,journal,nc-notifications,netconf-keystore,netconf-node-inventory,netconf-node-topology,network-topology,network-topology,notifications,opendaylight-config-dom-datastore,opendaylight-entity-ownership-service,opendaylight-inmemory-datastore-provider,opendaylight-inventory,opendaylight-legacy-entity-ownership-service-provider,opendaylight-md-sal-common,opendaylight-md-sal-dom,opendaylight-operational-dom-datastore,rpc-context,yang-ext]
 *
 * API version: 4.2.0.frinx
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger_southbound

// Optional[The list of network nodes defined for the topology.] REF:Optional.empty
type NetworkTopologyNetworktopologyTopologyNode struct {
	// Optional.empty REF:Optional.empty
	NetconfNodeTopologysleepFactor int32 `json:"netconf-node-topology:sleep-factor,omitempty"`
	// Optional[Timeout for blocking operations within transactions.] REF:Optional.empty
	NetconfNodeTopologydefaultRequestTimeoutMillis int64 `json:"netconf-node-topology:default-request-timeout-millis,omitempty"`
	// Optional.empty REF:Optional.empty
	CliTopologydeviceType string `json:"cli-topology:device-type,omitempty"`
	// Optional.empty REF:Optional.empty
	NetconfNodeTopologytcpOnly bool `json:"netconf-node-topology:tcp-only"`
	// Optional.empty REF:Optional.empty
	CliTopologyconnectedMessage string `json:"cli-topology:connected-message,omitempty"`
	// Optional.empty REF:Optional.empty
	NetconfNodeTopologyport int32 `json:"netconf-node-topology:port,omitempty"`
	// Optional.empty REF:Optional.empty
	Password string `json:"netconf-node-topology:password,omitempty"`
	NetconfNodeTopologyunavailableCapabilities *NetconfNodeTopologyNetconfnodeconnectionstatusUnavailableCapabilities `json:"netconf-node-topology:unavailable-capabilities,omitempty"`
	// Optional[Initial timeout in milliseconds to wait between connection attempts. Will be multiplied by sleep-factor with every additional attempt] REF:Optional.empty
	NetconfNodeTopologybetweenAttemptsTimeoutMillis int32 `json:"netconf-node-topology:between-attempts-timeout-millis,omitempty"`
	NetconfNodeTopologyavailableCapabilities *NetconfNodeTopologyNetconfnodeconnectionstatusAvailableCapabilities `json:"netconf-node-topology:available-capabilities,omitempty"`
	// Optional[Maximum number of connection retries. Non positive value or null is interpreted as infinity.] REF:Optional.empty
	NetconfNodeTopologymaxConnectionAttempts int64 `json:"netconf-node-topology:max-connection-attempts,omitempty"`
	// Optional[Maximal time (in seconds) for connection to keep alive, if no activity was detected in the session and the timeout has been reached, connection will be stopped] REF:Optional.empty
	ConnectionLazyTimeout int32 `json:"connection-lazy-timeout,omitempty"`
	// Optional.empty REF:Optional.empty
	CliTopologyconnectionStatus *CliTopologyConnectionStatusEnumeration `json:"cli-topology:connection-status,omitempty"`
	// Optional[Limit of concurrent messages that can be send before reply messages are received. If value <1 is provided, no limit will be enforced] REF:Optional.empty
	NetconfNodeTopologyconcurrentRpcLimit int32 `json:"netconf-node-topology:concurrent-rpc-limit,omitempty"`
	// Optional[Maximal time (in seconds) for connection establishment, if a connection attempt does not succeed in this time, the attempt is considered a failure] REF:Optional.empty
	ConnectionEstablishTimeout int32 `json:"connection-establish-timeout,omitempty"`
	CliTopologydefaultCommitErrorPatterns *CliTranslateRegistryErrorcommitpatternsDefaultCommitErrorPatterns `json:"cli-topology:default-commit-error-patterns,omitempty"`
	// Optional.empty REF:Optional.empty
	KeepaliveDelay int32 `json:"keepalive-delay,omitempty"`
	LoginPassword *NetconfNodeTopologyNetconfnodecredentialsCredentialsLoginpwunencryptedLoginPasswordUnencrypted `json:"login-password,omitempty"`
	// Optional.empty REF:Optional.empty
	NetconfNodeTopologyconnectedMessage string `json:"netconf-node-topology:connected-message,omitempty"`
	// Optional.empty REF:Optional.empty
	CliTopologyport int32 `json:"cli-topology:port,omitempty"`
	// Optional[Netconf connector sends keepalive RPCs while the session is idle, this delay specifies the delay between keepalive RPC in seconds If a value <1 is provided, no keepalives will be sent] REF:Optional.empty
	NetconfNodeTopologykeepaliveDelay int64 `json:"netconf-node-topology:keepalive-delay,omitempty"`
	// Optional[The identifier of a node in the topology. A node is specific to a topology to which it belongs.] REF:Optional.empty
	NodeId string `json:"node-id,omitempty"`
	NetconfNodeTopologyyangModuleCapabilities *NetconfNodeTopologyNetconfnodeconnectionparametersYangModuleCapabilities `json:"netconf-node-topology:yang-module-capabilities,omitempty"`
	NetconfNodeTopologyyangLibrary *NetconfNodeTopologyNetconfschemastorageYangLibrary `json:"netconf-node-topology:yang-library,omitempty"`
	// Optional[Size of the DRY RUN netconf mountpoint jounral. DRY RUN journal captures netconf RPCs that would be executed when reading/writing some configuration. However the RPCs are not actually sent to the device] REF:Optional.empty
	NetconfNodeTopologydryRunJournalSize int32 `json:"netconf-node-topology:dry-run-journal-size,omitempty"`
	NetconfNodeTopologyodlHelloMessageCapabilities *NetconfNodeTopologyNetconfnodeconnectionparametersOdlHelloMessageCapabilities `json:"netconf-node-topology:odl-hello-message-capabilities,omitempty"`
	// Optional[Privileged EXEC mode password for Cisco IOS devices. If not set credentials password will be used] REF:Optional.empty
	Secret string `json:"secret,omitempty"`
	// Optional[Size of the DRY RUN cli mountpoint jounral. DRY RUN journal captures commands that would be executed when reading/writing some configuration. However the commands are not actually sent to the device] REF:Optional.empty
	CliTopologydryRunJournalSize int32 `json:"cli-topology:dry-run-journal-size,omitempty"`
	// Optional.empty REF:Optional.empty
	KeepaliveInitialDelay int32 `json:"keepalive-initial-delay,omitempty"`
	// Optional[This list defines vertical layering information for nodes. It allows to capture for any given node, which node (or nodes) in the corresponding underlay topology it maps onto. A node can map to zero, one, or more nodes below it; accordingly there can be zero, one, or more elements in the list. If there are specific layering requirements, for example specific to a particular type of topology that only allows for certain layering relationships, the choice below can be augmented with additional cases. A list has been chosen rather than a leaf-list in order to provide room for augmentations, e.g. for statistics or priorization information associated with supporting nodes.] REF:Optional.empty
	SupportingNode []NetworkTopologyNodeattributesSupportingNode `json:"supporting-node,omitempty"`
	// Optional[Specifies timeout in milliseconds after which connection must be established.] REF:Optional.empty
	NetconfNodeTopologyconnectionTimeoutMillis int64 `json:"netconf-node-topology:connection-timeout-millis,omitempty"`
	// Optional.empty REF:Optional.empty
	KeepaliveTimeout int32 `json:"keepalive-timeout,omitempty"`
	// Optional.empty REF:Optional.empty
	NetconfNodeTopologyconnectionStatus *NetconfNodeTopologyConnectionStatusEnumeration `json:"netconf-node-topology:connection-status,omitempty"`
	// Optional[The destination schema repository for yang files relative to the cache directory.  This may be specified per netconf mount so that the loaded yang files are stored to a distinct directory to avoid potential conflict.] REF:Optional.empty
	NetconfNodeTopologyschemaCacheDirectory string `json:"netconf-node-topology:schema-cache-directory,omitempty"`
	// Optional.empty REF:Optional.empty
	CliTopologydeviceVersion string `json:"cli-topology:device-version,omitempty"`
	// Optional[When the underlying node is connected, its NETCONF context is available verbatim under this container through the mount extension.] REF:Optional.empty
	NetconfNodeTopologypassThrough *interface{} `json:"netconf-node-topology:pass-through,omitempty"`
	NetconfNodeTopologynonModuleCapabilities *NetconfNodeTopologyNetconfnodeconnectionparametersNonModuleCapabilities `json:"netconf-node-topology:non-module-capabilities,omitempty"`
	// Optional.empty REF:Optional.empty
	NetconfNodeTopologyschemaless bool `json:"netconf-node-topology:schemaless,omitempty"`
	KeyBased *NetconfNodeTopologyNetconfnodecredentialsCredentialsKeyauthKeyBased `json:"key-based,omitempty"`
	// Optional[When the underlying node is connected, its cli context is available verbatim under this container through the mount extension.] REF:Optional.empty
	CliTopologypassThrough *interface{} `json:"cli-topology:pass-through,omitempty"`
	// Optional.empty REF:Optional.empty
	CliTopologyhost string `json:"cli-topology:host,omitempty"`
	CliTopologydefaultErrorPatterns *CliTranslateRegistryErrorpatternsDefaultErrorPatterns `json:"cli-topology:default-error-patterns,omitempty"`
	// Optional.empty REF:Optional.empty
	CliTopologytransportType *CliTopologyTransportTypeEnumeration `json:"cli-topology:transport-type,omitempty"`
	// Optional[Sets how much information should be stored in the journal. Command-only stores only the actual commands executed on device. Extended records additional information such as: transaction life-cycle, which handlers were invoked etc.] REF:Optional.empty
	CliTopologyjournalLevel *CliTopologyJournalLevel `json:"cli-topology:journal-level,omitempty"`
	// Optional[Size of the cli mountpoint jounral. Journal keeps track of executed commands and makes them available for users/apps for debugging purposes. Value 0 disables journaling] REF:Optional.empty
	CliTopologyjournalSize int32 `json:"cli-topology:journal-size,omitempty"`
	// Optional[A termination point can terminate a link. Depending on the type of topology, a termination point could, for example, refer to a port or an interface.] REF:Optional.empty
	TerminationPoint []NetworkTopologyNetworktopologyTopologyNodeTerminationPoint `json:"termination-point,omitempty"`
	// Optional.empty REF:Optional.empty
	NetconfNodeTopologyhost string `json:"netconf-node-topology:host,omitempty"`
	LoginPasswordUnencrypted *NetconfNodeTopologyNetconfnodecredentialsCredentialsLoginpwunencryptedLoginPasswordUnencrypted `json:"login-password-unencrypted,omitempty"`
	NetconfNodeTopologyclusteredConnectionStatus *NetconfNodeTopologyNetconfnodeconnectionstatusClusteredConnectionStatus `json:"netconf-node-topology:clustered-connection-status,omitempty"`
	// Optional.empty REF:Optional.empty
	NetconfNodeTopologycustomizationFactory string `json:"netconf-node-topology:customization-factory,omitempty"`
	// Optional[Maximal time (in seconds) for command execution, if a command cannot be executed on a device in this time, the execution is considered a failure] REF:Optional.empty
	CommandTimeout int32 `json:"command-timeout,omitempty"`
	// Optional[Time that slave actor will wait for response from master.] REF:Optional.empty
	NetconfNodeTopologyactorResponseWaitTime int32 `json:"netconf-node-topology:actor-response-wait-time,omitempty"`
	// Optional[If true, the connector would auto disconnect/reconnect when schemas are changed in the remote device. The connector subscribes (right after connect) to base netconf notifications and listens for netconf-capability-change notification] REF:Optional.empty
	NetconfNodeTopologyreconnectOnChangedSchema bool `json:"netconf-node-topology:reconnect-on-changed-schema,omitempty"`
	CliTopologyavailableCapabilities *CliTopologyClinodeconnectionstatusAvailableCapabilities `json:"cli-topology:available-capabilities,omitempty"`
	// Optional.empty REF:Optional.empty
	Username string `json:"netconf-node-topology:username,omitempty"`
}
