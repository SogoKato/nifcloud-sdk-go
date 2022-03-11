// Code generated by smithy-go-codegen DO NOT EDIT.


package types

import (
	smithydocument "github.com/aws/smithy-go/document"
)

type AddonsConfig struct {
	
	HttpLoadBalancing *HttpLoadBalancing
	
	noSmithyDocumentSerde
}

type Cluster struct {
	
	AddonsConfig *AddonsConfig
	
	CreateTime *string
	
	Description *string
	
	FirewallGroup *string
	
	InitialKubernetesVersion *string
	
	InitialNodeCount *int32
	
	KubernetesVersion *string
	
	Locations []string
	
	Name *string
	
	NetworkConfig *NetworkConfig
	
	NodeCount *int32
	
	NodePools []NodePools
	
	Status *string
	
	noSmithyDocumentSerde
}

type ClusterOfCreateSnapshot struct {
	
	KubernetesVersion *string
	
	Name *string
	
	NodePools []NodePoolsOfCreateSnapshot
	
	noSmithyDocumentSerde
}

type ClusterOfDeleteSnapshot struct {
	
	KubernetesVersion *string
	
	Name *string
	
	NodePools []NodePoolsOfDeleteSnapshot
	
	noSmithyDocumentSerde
}

type ClusterOfDeleteSnapshots struct {
	
	KubernetesVersion *string
	
	Name *string
	
	NodePools []NodePoolsOfDeleteSnapshots
	
	noSmithyDocumentSerde
}

type ClusterOfGetSnapshot struct {
	
	KubernetesVersion *string
	
	Name *string
	
	NodePools []NodePoolsOfGetSnapshot
	
	noSmithyDocumentSerde
}

type ClusterOfListSnapshots struct {
	
	KubernetesVersion *string
	
	Name *string
	
	NodePools []NodePoolsOfListSnapshots
	
	noSmithyDocumentSerde
}

type ClusterOfUpdateSnapshot struct {
	
	KubernetesVersion *string
	
	Name *string
	
	NodePools []NodePoolsOfUpdateSnapshot
	
	noSmithyDocumentSerde
}

type Clusters struct {
	
	AddonsConfig *AddonsConfig
	
	CreateTime *string
	
	Description *string
	
	FirewallGroup *string
	
	InitialKubernetesVersion *string
	
	InitialNodeCount *int32
	
	KubernetesVersion *string
	
	Locations []string
	
	Name *string
	
	NetworkConfig *NetworkConfig
	
	NodeCount *int32
	
	NodePools []NodePools
	
	Status *string
	
	noSmithyDocumentSerde
}

type ClustersOfGetLoadBalancer struct {
	
	Name *string
	
	NodePools []NodePoolsOfGetLoadBalancer
	
	noSmithyDocumentSerde
}

type ClustersOfListLoadBalancers struct {
	
	Name *string
	
	NodePools []NodePoolsOfListLoadBalancers
	
	noSmithyDocumentSerde
}

type Filter struct {
	
	FilterType *string
	
	IpAddresses *string
	
	noSmithyDocumentSerde
}

type FirewallGroup struct {
	
	Description *string
	
	Name *string
	
	Rules []Rules
	
	noSmithyDocumentSerde
}

type FirewallGroups struct {
	
	Description *string
	
	Name *string
	
	Rules []Rules
	
	noSmithyDocumentSerde
}

type HealthCheck struct {
	
	HealthyThreshold *int32
	
	Interval *int32
	
	Target *string
	
	Timeout *int32
	
	UnhealthyThreshold *int32
	
	noSmithyDocumentSerde
}

type HttpLoadBalancing struct {
	
	Disabled *bool
	
	noSmithyDocumentSerde
}

type Listener struct {
	
	BalancingType *string
	
	InstancePort *string
	
	LoadBalancerPort *string
	
	Protocol *string
	
	SslCertificateId *string
	
	noSmithyDocumentSerde
}

type ListenerDescriptions struct {
	
	Listener *Listener
	
	noSmithyDocumentSerde
}

type LoadBalancers struct {
	
	AccountingType *int32
	
	AvailabilityZones []string
	
	Clusters []ClustersOfGetLoadBalancer
	
	CreatedTime *string
	
	Description *string
	
	DnsName *string
	
	Filter *Filter
	
	HealthCheck *HealthCheck
	
	ListenerDescriptions []ListenerDescriptions
	
	LoadBalancerName *string
	
	NetworkVolume *string
	
	NextMonthAccountingType *int32
	
	Option *Option
	
	PolicyType *string
	
	noSmithyDocumentSerde
}

type LoadBalancersOfListLoadBalancers struct {
	
	AccountingType *int32
	
	AvailabilityZones []string
	
	Clusters []ClustersOfListLoadBalancers
	
	CreatedTime *string
	
	Description *string
	
	DnsName *string
	
	Filter *Filter
	
	HealthCheck *HealthCheck
	
	ListenerDescriptions []ListenerDescriptions
	
	LoadBalancerName *string
	
	NetworkVolume *string
	
	NextMonthAccountingType *int32
	
	Option *Option
	
	PolicyType *string
	
	noSmithyDocumentSerde
}

type NetworkConfig struct {
	
	NetworkId *string
	
	noSmithyDocumentSerde
}

type NodePool struct {
	
	InitialNodeCount *int32
	
	InstanceType *string
	
	Name *string
	
	NodeCount *int32
	
	Nodes []Nodes
	
	Status *string
	
	noSmithyDocumentSerde
}

type NodePools struct {
	
	InitialNodeCount *int32
	
	InstanceType *string
	
	Name *string
	
	NodeCount *int32
	
	Nodes []Nodes
	
	Status *string
	
	noSmithyDocumentSerde
}

type NodePoolsOfCreateSnapshot struct {
	
	InstanceType *string
	
	Name *string
	
	NodeCount *int32
	
	noSmithyDocumentSerde
}

type NodePoolsOfDeleteSnapshot struct {
	
	InstanceType *string
	
	Name *string
	
	NodeCount *int32
	
	noSmithyDocumentSerde
}

type NodePoolsOfDeleteSnapshots struct {
	
	InstanceType *string
	
	Name *string
	
	NodeCount *int32
	
	noSmithyDocumentSerde
}

type NodePoolsOfGetLoadBalancer struct {
	
	Name *string
	
	NodeCount *int32
	
	Nodes []NodesOfGetLoadBalancer
	
	noSmithyDocumentSerde
}

type NodePoolsOfGetSnapshot struct {
	
	InstanceType *string
	
	Name *string
	
	NodeCount *int32
	
	noSmithyDocumentSerde
}

type NodePoolsOfListLoadBalancers struct {
	
	Name *string
	
	NodeCount *int32
	
	Nodes []NodesOfListLoadBalancers
	
	noSmithyDocumentSerde
}

type NodePoolsOfListSnapshots struct {
	
	InstanceType *string
	
	Name *string
	
	NodeCount *int32
	
	noSmithyDocumentSerde
}

type NodePoolsOfUpdateSnapshot struct {
	
	InstanceType *string
	
	Name *string
	
	NodeCount *int32
	
	noSmithyDocumentSerde
}

type Nodes struct {
	
	AvailabilityZone *string
	
	Name *string
	
	PrivateIpAddress *string
	
	PublicIpAddress *string
	
	Status *string
	
	noSmithyDocumentSerde
}

type NodesOfGetLoadBalancer struct {
	
	AvailabilityZone *string
	
	HealthCheckState *string
	
	Name *string
	
	PublicIpAddress *string
	
	noSmithyDocumentSerde
}

type NodesOfListLoadBalancers struct {
	
	AvailabilityZone *string
	
	HealthCheckState *string
	
	Name *string
	
	PublicIpAddress *string
	
	noSmithyDocumentSerde
}

type Option struct {
	
	SessionStickinessPolicy *SessionStickinessPolicy
	
	SorryPage *SorryPage
	
	noSmithyDocumentSerde
}

type RequestAddonsConfig struct {
	
	RequestHttpLoadBalancing *RequestHttpLoadBalancing
	
	noSmithyDocumentSerde
}

type RequestCluster struct {
	
	// This member is required.
	FirewallGroup *string
	
	// This member is required.
	ListOfRequestLocations []string
	
	// This member is required.
	ListOfRequestNodePools []RequestNodePools
	
	// This member is required.
	Name *string
	
	Description *string
	
	KubernetesVersion KubernetesVersionOfclusterForCreateCluster
	
	RequestAddonsConfig *RequestAddonsConfig
	
	RequestNetworkConfig *RequestNetworkConfig
	
	noSmithyDocumentSerde
}

type RequestClusterOfCreateSnapshot struct {
	
	// This member is required.
	Name *string
	
	noSmithyDocumentSerde
}

type RequestClusterOfRestoreClusterFromSnapshot struct {
	
	// This member is required.
	FirewallGroup *string
	
	// This member is required.
	ListOfRequestLocations []string
	
	// This member is required.
	Name *string
	
	Description *string
	
	RequestAddonsConfig *RequestAddonsConfig
	
	RequestNetworkConfig *RequestNetworkConfig
	
	noSmithyDocumentSerde
}

type RequestClusterOfUpdateCluster struct {
	
	Description *string
	
	KubernetesVersion KubernetesVersionOfclusterForUpdateCluster
	
	Name *string
	
	RequestAddonsConfig *RequestAddonsConfig
	
	noSmithyDocumentSerde
}

type RequestFirewallGroup struct {
	
	// This member is required.
	Name *string
	
	Description *string
	
	noSmithyDocumentSerde
}

type RequestFirewallGroupOfUpdateFirewallGroup struct {
	
	Description *string
	
	Name *string
	
	noSmithyDocumentSerde
}

type RequestHttpLoadBalancing struct {
	
	Disabled *bool
	
	noSmithyDocumentSerde
}

type RequestNetworkConfig struct {
	
	NetworkId *string
	
	noSmithyDocumentSerde
}

type RequestNodePool struct {
	
	// This member is required.
	InstanceType InstanceTypeOfnodePoolForCreateNodePool
	
	// This member is required.
	Name *string
	
	NodeCount *int32
	
	noSmithyDocumentSerde
}

type RequestNodePools struct {
	
	// This member is required.
	InstanceType InstanceTypeOfclusterForCreateCluster
	
	// This member is required.
	Name *string
	
	NodeCount *int32
	
	noSmithyDocumentSerde
}

type RequestRules struct {
	
	// This member is required.
	CidrIp *string
	
	Description *string
	
	Direction DirectionOfrulesForAuthorizeFirewallGroup
	
	FromPort *int32
	
	Protocol ProtocolOfrulesForAuthorizeFirewallGroup
	
	ToPort *int32
	
	noSmithyDocumentSerde
}

type RequestSnapshot struct {
	
	// This member is required.
	Name *string
	
	// This member is required.
	RequestCluster *RequestClusterOfCreateSnapshot
	
	Description *string
	
	ExpirationTime *string
	
	noSmithyDocumentSerde
}

type RequestSnapshotOfUpdateSnapshot struct {
	
	Description *string
	
	ExpirationTime *string
	
	Name *string
	
	noSmithyDocumentSerde
}

type Rules struct {
	
	CidrIp *string
	
	Description *string
	
	Direction *string
	
	FromPort *int32
	
	Id *string
	
	Protocol *string
	
	Status *string
	
	ToPort *int32
	
	noSmithyDocumentSerde
}

type ServerConfig struct {
	
	DefaultKubernetesVersion *string
	
	ValidKubernetesVersions []string
	
	noSmithyDocumentSerde
}

type SessionStickinessPolicy struct {
	
	Enabled *bool
	
	ExpirationPeriod *int32
	
	noSmithyDocumentSerde
}

type Snapshot struct {
	
	Cluster *ClusterOfGetSnapshot
	
	CreateTime *string
	
	Description *string
	
	ExpirationTime *string
	
	Name *string
	
	ResourceVersion *string
	
	Status *string
	
	noSmithyDocumentSerde
}

type SnapshotOfCreateSnapshot struct {
	
	Cluster *ClusterOfCreateSnapshot
	
	CreateTime *string
	
	Description *string
	
	ExpirationTime *string
	
	Name *string
	
	ResourceVersion *string
	
	Status *string
	
	noSmithyDocumentSerde
}

type SnapshotOfDeleteSnapshot struct {
	
	Cluster *ClusterOfDeleteSnapshot
	
	CreateTime *string
	
	Description *string
	
	ExpirationTime *string
	
	Name *string
	
	ResourceVersion *string
	
	Status *string
	
	noSmithyDocumentSerde
}

type SnapshotOfUpdateSnapshot struct {
	
	Cluster *ClusterOfUpdateSnapshot
	
	CreateTime *string
	
	Description *string
	
	ExpirationTime *string
	
	Name *string
	
	ResourceVersion *string
	
	Status *string
	
	noSmithyDocumentSerde
}

type Snapshots struct {
	
	Cluster *ClusterOfListSnapshots
	
	CreateTime *string
	
	Description *string
	
	ExpirationTime *string
	
	Name *string
	
	ResourceVersion *string
	
	Status *string
	
	noSmithyDocumentSerde
}

type SnapshotsOfDeleteSnapshots struct {
	
	Cluster *ClusterOfDeleteSnapshots
	
	CreateTime *string
	
	Description *string
	
	ExpirationTime *string
	
	Name *string
	
	ResourceVersion *string
	
	Status *string
	
	noSmithyDocumentSerde
}

type SorryPage struct {
	
	Enabled *bool
	
	StatusCode *int32
	
	noSmithyDocumentSerde
}

type noSmithyDocumentSerde = smithydocument.NoSerde
