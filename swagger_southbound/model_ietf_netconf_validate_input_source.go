/*
 * southbound
 *
 * API generated from yang definitions: [aaa-encrypt-service-config,cli-topology,cli-translate-registry,cluster-singleton-service-impl,cluster-singleton-service-spi,config,fake,general-entity,ietf-inet-types,ietf-netconf,ietf-netconf-monitoring,ietf-netconf-monitoring-extension,ietf-netconf-notifications,ietf-yang-library,ietf-yang-types,journal,nc-notifications,netconf-keystore,netconf-node-inventory,netconf-node-topology,network-topology,network-topology,notifications,opendaylight-config-dom-datastore,opendaylight-entity-ownership-service,opendaylight-inmemory-datastore-provider,opendaylight-inventory,opendaylight-legacy-entity-ownership-service-provider,opendaylight-md-sal-common,opendaylight-md-sal-dom,opendaylight-operational-dom-datastore,rpc-context,yang-ext]
 *
 * API version: 4.2.0.frinx
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger_southbound

// Optional[Particular configuration to validate.] REF:Optional.empty
type IetfNetconfValidateInputSource struct {
	// Optional[The running configuration is the config source.] REF:Optional.empty
	Running string `json:"running,omitempty"`
	// Optional[The candidate configuration is the config source.] REF:Optional.empty
	Candidate string `json:"candidate,omitempty"`
	// Optional[Inline Config content: <config> element.  Represents an entire configuration datastore, not a subset of the running datastore.] REF:Optional.empty
	Config string `json:"config,omitempty"`
	// Optional[The startup configuration is the config source.] REF:Optional.empty
	Startup string `json:"startup,omitempty"`
	// Optional[The URL-based configuration is the config source.] REF:Optional.empty
	Url string `json:"url,omitempty"`
}
