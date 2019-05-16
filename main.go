package main

import (
	"encoding/json"
	"time"
	"context"
	"fmt"
	"os"
	"frinx_go_demo/swagger_unified"
	"frinx_go_demo/swagger_southbound"
)

func putNetconfNode(ctx context.Context, c *swagger_southbound.APIClient) {
	var node = swagger_southbound.NetworkTopologyNetworktopologyTopologyNode{
		NetconfNodeTopologyhost: "<IP>",
		NetconfNodeTopologyport: 2424,
		Username: "USERNAME",
		Password: "PASSWORD",
		NodeId: "PE1",
		NetconfNodeTopologytcpOnly: false,
		NetconfNodeTopologykeepaliveDelay: 100,
	}

	var nodeBodyParam = swagger_southbound.NetworkTopologyNetworktopologyTopologyNodeRequest{NetworkTopologynode: []swagger_southbound.NetworkTopologyNetworktopologyTopologyNode{node}}
	
	// var b, _ = json.Marshal(nodeBodyParam)
	// println(string(b))

	var http, e = c.NetworkTopologyApi.PutNetworkTopologyNetworkTopologyTopologyNode(ctx, "topology-netconf", "PE1", nodeBodyParam)
	if e != nil {
		println("Error response from ODL when adding netconf node: " + e.Error())
		os.Exit(1)
	}
	println("Status code from ODL when adding netconf node: " + http.Status)
}

func waitUntilNetconfNodeConnects(ctx context.Context, c *swagger_southbound.APIClient) {

	var connectionStatus = swagger_southbound.NetconfNodeTopologyConnectionStatusEnumeration_CONNECTING

	for (connectionStatus != swagger_southbound.NetconfNodeTopologyConnectionStatusEnumeration_CONNECTED) {
		time.Sleep(1000 * time.Millisecond)

		var r, http, e = c.NetworkTopologyApi.OperGetNetworkTopologyNetworkTopologyTopologyNode(ctx, "topology-netconf", "PE1")
		println("Status code from ODL when adding netconf node: " + http.Status)
		if e != nil {
			println("Error response from ODL when checking netconf node connection status: " + e.Error())
			continue
		}

		// var b, _ = json.Marshal(r)
		// println(string(b))

		println(http.Request.URL.String())

		if len(r.NetworkTopologynode) == 0 {
			println("No response from ODL when checking netconf node connection status")
			continue
		}	

		connectionStatus = *r.NetworkTopologynode[0].NetconfNodeTopologyconnectionStatus
		println("Status code from ODL when adding netconf node: " + http.Status)
		println("Connection status from ODL: " + connectionStatus)
	}
}

func printInterfaces(ctx context.Context, c *swagger_unified.APIClient) {
	var r, http, e = c.FrinxOpenconfigInterfacesApi.GetFrinxOpenconfigInterfacesInterfaces(ctx, "PE1")
	if e != nil {
		println("Error response from ODL when reading interfaces: " + e.Error())
		os.Exit(1)
	}

	println("Status code from ODL when reading interfaces: " + http.Status)
	var b, _ = json.Marshal(r)
	println(string(b))
}

func main() {
	var dockerContainerIp = "172.17.0.2"

	fmt.Println("Frinx Go Client.")
	fmt.Println("=============")

	var configS = swagger_southbound.NewConfiguration()
	configS.BasePath = "http://" + dockerContainerIp + ":8181/restconf"
	configS.AddDefaultHeader("Authorization", "Basic YWRtaW46YWRtaW4=")
	var cS = swagger_southbound.NewAPIClient(configS)
	var ctxS = context.Background()
	
	putNetconfNode(ctxS, cS)
	waitUntilNetconfNodeConnects(ctxS, cS)
	
	var config = swagger_unified.NewConfiguration()
	config.BasePath = "http://" + dockerContainerIp + ":8181/restconf"
	config.AddDefaultHeader("Authorization", "Basic YWRtaW46YWRtaW4=")
	var c = swagger_unified.NewAPIClient(config)
	var ctx = context.Background()

	printInterfaces(ctx, c)
}
