/*
 * southbound
 *
 * API generated from yang definitions: [aaa-encrypt-service-config,cli-topology,cli-translate-registry,cluster-singleton-service-impl,cluster-singleton-service-spi,config,fake,general-entity,ietf-inet-types,ietf-netconf,ietf-netconf-monitoring,ietf-netconf-monitoring-extension,ietf-netconf-notifications,ietf-yang-library,ietf-yang-types,journal,nc-notifications,netconf-keystore,netconf-node-inventory,netconf-node-topology,network-topology,network-topology,notifications,opendaylight-config-dom-datastore,opendaylight-entity-ownership-service,opendaylight-inmemory-datastore-provider,opendaylight-inventory,opendaylight-legacy-entity-ownership-service-provider,opendaylight-md-sal-common,opendaylight-md-sal-dom,opendaylight-operational-dom-datastore,rpc-context,yang-ext]
 *
 * API version: 4.2.0.frinx
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger_southbound

// Optional[Device specific list of error patterns. This list is the primary source of error checking on the device. This list can be overridden from the code.] REF:Optional.empty
type CliTranslateRegistryErrorpatternsDefaultErrorPatterns struct {
	// Optional[Device specific error patterns.] REF:Optional.empty
	ErrorPattern []string `json:"error-pattern,omitempty"`
}
