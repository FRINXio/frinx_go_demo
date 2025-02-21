/*
 * southbound
 *
 * API generated from yang definitions: [aaa-encrypt-service-config,cli-topology,cli-translate-registry,cluster-singleton-service-impl,cluster-singleton-service-spi,config,fake,general-entity,ietf-inet-types,ietf-netconf,ietf-netconf-monitoring,ietf-netconf-monitoring-extension,ietf-netconf-notifications,ietf-yang-library,ietf-yang-types,journal,nc-notifications,netconf-keystore,netconf-node-inventory,netconf-node-topology,network-topology,network-topology,notifications,opendaylight-config-dom-datastore,opendaylight-entity-ownership-service,opendaylight-inmemory-datastore-provider,opendaylight-inventory,opendaylight-legacy-entity-ownership-service-provider,opendaylight-md-sal-common,opendaylight-md-sal-dom,opendaylight-operational-dom-datastore,rpc-context,yang-ext]
 *
 * API version: 4.2.0.frinx
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger_southbound

type NetconfKeystorePrivatekeysPrivateKey struct {
	// Optional.empty REF:Optional.empty
	Name string `json:"name,omitempty"`
	// Optional[Base64 encoded private key.] REF:Optional.empty
	Data string `json:"data,omitempty"`
	// Optional[A certificate chain for this public key. Each certificate is an X.509 v3 certificate structure as specified by RFC5280, encoded using the Base64 format.] REF:Optional.empty
	CertificateChain []string `json:"certificate-chain,omitempty"`
}
