package cluster

import (
	"k8s.io/client-go/rest"
)

const (
	//DefaultClusterName is the cluster for the unspecified cluster
	DefaultClusterName = "_default"
)

//Info represents a Cluster,
type Info struct {
	// Name is the cluster name, usually the Cluster resource's name
	Name string
	// Endpoint the apiserver's endpoint of the cluster
	Endpoint string
	// Token is a admin token , it should have all the access to the cluster
	Token string

	// Namespace the namespace which the chart will be installed to
	Namespace string
}

//GetContext is the context name for this cluster, this name format is generated from k8s code
func (i *Info) GetContext() string {
	return i.Name + "@" + i.Name
}

//ToRestConfig generate rest.Config from cluster info.
func (i *Info) ToRestConfig() *rest.Config {
	return &rest.Config{
		Host:        i.Endpoint,
		BearerToken: i.Token,
		TLSClientConfig: rest.TLSClientConfig{
			Insecure: true,
		},
	}
}

//RestConfigToCluster generate a cluster Info from a rest config
// This method and the Info.ToRestConfig both only support bearer token for now
// luckily, the in-cluster rest config also use bearer token
func RestConfigToCluster(config *rest.Config, generatedName string) *Info {
	var i Info
	i.Token = config.BearerToken
	i.Endpoint = config.Host
	i.Name = generatedName
	return &i
}
