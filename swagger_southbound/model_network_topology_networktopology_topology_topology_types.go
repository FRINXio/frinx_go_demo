/*
 * southbound
 *
 * API generated from yang definitions: [aaa-encrypt-service-config,cli-topology,cli-translate-registry,cluster-singleton-service-impl,cluster-singleton-service-spi,config,fake,general-entity,ietf-inet-types,ietf-netconf,ietf-netconf-monitoring,ietf-netconf-monitoring-extension,ietf-netconf-notifications,ietf-yang-library,ietf-yang-types,journal,nc-notifications,netconf-keystore,netconf-node-inventory,netconf-node-topology,network-topology,network-topology,notifications,opendaylight-config-dom-datastore,opendaylight-entity-ownership-service,opendaylight-inmemory-datastore-provider,opendaylight-inventory,opendaylight-legacy-entity-ownership-service-provider,opendaylight-md-sal-common,opendaylight-md-sal-dom,opendaylight-operational-dom-datastore,rpc-context,yang-ext]
 *
 * API version: 4.2.0.frinx
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger_southbound

// Optional[This container is used to identify the type, or types (as a topology can support several types simultaneously), of the topology. Topology types are the subject of several integrity constraints that an implementing server can validate in order to maintain integrity of the datastore. Topology types are indicated through separate data nodes; the set of topology types is expected to increase over time. To add support for a new topology, an augmenting module needs to augment this container with a new empty optional container to indicate the new topology type. The use of a container allows to indicate a subcategorization of topology types. The container SHALL NOT be augmented with any data nodes that serve a purpose other than identifying a particular topology type. ] REF:Optional.empty
type NetworkTopologyNetworktopologyTopologyTopologyTypes struct {
	// Optional.empty REF:Optional.empty
	NetconfNodeTopologytopologyNetconf *interface{} `json:"netconf-node-topology:topology-netconf,omitempty"`
}
