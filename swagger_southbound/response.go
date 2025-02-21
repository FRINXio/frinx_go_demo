/*
 * southbound
 *
 * API generated from yang definitions: [aaa-encrypt-service-config,cli-topology,cli-translate-registry,cluster-singleton-service-impl,cluster-singleton-service-spi,config,fake,general-entity,ietf-inet-types,ietf-netconf,ietf-netconf-monitoring,ietf-netconf-monitoring-extension,ietf-netconf-notifications,ietf-yang-library,ietf-yang-types,journal,nc-notifications,netconf-keystore,netconf-node-inventory,netconf-node-topology,network-topology,network-topology,notifications,opendaylight-config-dom-datastore,opendaylight-entity-ownership-service,opendaylight-inmemory-datastore-provider,opendaylight-inventory,opendaylight-legacy-entity-ownership-service-provider,opendaylight-md-sal-common,opendaylight-md-sal-dom,opendaylight-operational-dom-datastore,rpc-context,yang-ext]
 *
 * API version: 4.2.0.frinx
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger_southbound

import (
	"net/http"
)

type APIResponse struct {
	*http.Response `json:"-"`
	Message        string `json:"message,omitempty"`
	// Operation is the name of the swagger operation.
	Operation string `json:"operation,omitempty"`
	// RequestURL is the request URL. This value is always available, even if the
	// embedded *http.Response is nil.
	RequestURL string `json:"url,omitempty"`
	// Method is the HTTP method used for the request.  This value is always
	// available, even if the embedded *http.Response is nil.
	Method string `json:"method,omitempty"`
	// Payload holds the contents of the response body (which may be nil or empty).
	// This is provided here as the raw response.Body() reader will have already
	// been drained.
	Payload []byte `json:"-"`
}

func NewAPIResponse(r *http.Response) *APIResponse {

	response := &APIResponse{Response: r}
	return response
}

func NewAPIResponseWithError(errorMessage string) *APIResponse {

	response := &APIResponse{Message: errorMessage}
	return response
}
