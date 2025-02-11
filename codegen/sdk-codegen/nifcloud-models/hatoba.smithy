namespace com.nifcloud.api.hatoba

use aws.auth#sigv4
use aws.api#service
use aws.protocols#restJson1
use aws.protocols#restXml
use aws.protocols#awsQuery
use aws.protocols#ec2Query
use smithy.waiters#waitable
use aws.protocols#ec2QueryName


structure GetClusterRequest {
    @required
    @httpLabel
    @jsonName("ClusterName")
    ClusterName: String,
}

structure Cluster {
    @jsonName("addonsConfig")
    AddonsConfig: AddonsConfig,
    @jsonName("createTime")
    CreateTime: String,
    @jsonName("description")
    Description: String,
    @jsonName("firewallGroup")
    FirewallGroup: String,
    @jsonName("initialKubernetesVersion")
    InitialKubernetesVersion: String,
    @jsonName("initialNodeCount")
    InitialNodeCount: Integer,
    @jsonName("kubernetesVersion")
    KubernetesVersion: String,
    @jsonName("locations")
    Locations: ListOfLocations,
    @jsonName("name")
    Name: String,
    @jsonName("networkConfig")
    NetworkConfig: NetworkConfig,
    @jsonName("nodeCount")
    NodeCount: Integer,
    @jsonName("nodePools")
    NodePools: ListOfNodePools,
    @jsonName("nrn")
    Nrn: String,
    @jsonName("status")
    Status: String,
    @jsonName("tags")
    Tags: ListOfTags,
}

structure AddonsConfig {
    @jsonName("httpLoadBalancing")
    HttpLoadBalancing: HttpLoadBalancing,
}

structure HttpLoadBalancing {
    @jsonName("disabled")
    Disabled: Boolean,
}

list ListOfNodePools {
    member: NodePools,
}

structure NodePools {
    @jsonName("initialNodeCount")
    InitialNodeCount: Integer,
    @jsonName("instanceType")
    InstanceType: String,
    @jsonName("name")
    Name: String,
    @jsonName("nodeCount")
    NodeCount: Integer,
    @jsonName("nodes")
    Nodes: ListOfNodes,
    @jsonName("nrn")
    Nrn: String,
    @jsonName("status")
    Status: String,
    @jsonName("tags")
    Tags: ListOfTags,
}

list ListOfNodes {
    member: Nodes,
}

structure Nodes {
    @jsonName("availabilityZone")
    AvailabilityZone: String,
    @jsonName("name")
    Name: String,
    @jsonName("privateIpAddress")
    PrivateIpAddress: String,
    @jsonName("publicIpAddress")
    PublicIpAddress: String,
    @jsonName("status")
    Status: String,
}

list ListOfTags {
    member: Tags,
}

structure Tags {
    @jsonName("id")
    Id: String,
    @jsonName("key")
    Key: String,
    @jsonName("value")
    Value: String,
}

list ListOfLocations {
    member: String,
}

structure NetworkConfig {
    @jsonName("networkId")
    NetworkId: String,
}

structure GetClusterResult {
    @jsonName("cluster")
    Cluster: Cluster,
}

structure ListClustersRequest {
    @httpQuery("filters")
    @jsonName("filters")
    Filters: String,
}

list ListOfClusters {
    member: Clusters,
}

structure Clusters {
    @jsonName("addonsConfig")
    AddonsConfig: AddonsConfig,
    @jsonName("createTime")
    CreateTime: String,
    @jsonName("description")
    Description: String,
    @jsonName("firewallGroup")
    FirewallGroup: String,
    @jsonName("initialKubernetesVersion")
    InitialKubernetesVersion: String,
    @jsonName("initialNodeCount")
    InitialNodeCount: Integer,
    @jsonName("kubernetesVersion")
    KubernetesVersion: String,
    @jsonName("locations")
    Locations: ListOfLocations,
    @jsonName("name")
    Name: String,
    @jsonName("networkConfig")
    NetworkConfig: NetworkConfig,
    @jsonName("nodeCount")
    NodeCount: Integer,
    @jsonName("nodePools")
    NodePools: ListOfNodePools,
    @jsonName("nrn")
    Nrn: String,
    @jsonName("status")
    Status: String,
    @jsonName("tags")
    Tags: ListOfTags,
}

structure ListClustersResult {
    @jsonName("clusters")
    Clusters: ListOfClusters,
}

structure RequestCluster {
    @jsonName("description")
    Description: String,
    @required
    @jsonName("firewallGroup")
    FirewallGroup: String,
    @jsonName("kubernetesVersion")
    KubernetesVersion: KubernetesVersionOfclusterForCreateCluster,
    @required
    @jsonName("locations")
    ListOfRequestLocations: ListOfRequestLocations,
    @required
    @jsonName("nodePools")
    ListOfRequestNodePools: ListOfRequestNodePools,
    @jsonName("tags")
    ListOfRequestTags: ListOfRequestTags,
    @required
    @jsonName("name")
    Name: String,
    @jsonName("addonsConfig")
    RequestAddonsConfig: RequestAddonsConfig,
    @jsonName("networkConfig")
    RequestNetworkConfig: RequestNetworkConfig,
}

@enum([
  {
      name: "V1_23_3",
      value: "v1.23.3",
  },
  {
      name: "V1_23_9",
      value: "v1.23.9",
  },
  {
      name: "V1_24_3",
      value: "v1.24.3",
  },
])
string KubernetesVersionOfclusterForCreateCluster

list ListOfRequestLocations {
    member: String,
}

structure RequestHttpLoadBalancing {
    @jsonName("disabled")
    Disabled: Boolean,
}

structure RequestAddonsConfig {
    @jsonName("httpLoadBalancing")
    RequestHttpLoadBalancing: RequestHttpLoadBalancing,
}

structure RequestNodePools {
    @required
    @jsonName("instanceType")
    InstanceType: InstanceTypeOfclusterForCreateCluster,
    @jsonName("tags")
    ListOfRequestTags: ListOfRequestTags,
    @required
    @jsonName("name")
    Name: String,
    @jsonName("nodeCount")
    NodeCount: Integer,
}

list ListOfRequestNodePools {
    member: RequestNodePools,
}

@enum([
  {
      name: "C_MEDIUM",
      value: "c-medium",
  },
  {
      name: "E_MEDIUM",
      value: "e-medium",
  },
  {
      name: "MEDIUM",
      value: "medium",
  },
  {
      name: "C_MEDIUM4",
      value: "c-medium4",
  },
  {
      name: "E_MEDIUM4",
      value: "e-medium4",
  },
  {
      name: "MEDIUM4",
      value: "medium4",
  },
  {
      name: "C_MEDIUM8",
      value: "c-medium8",
  },
  {
      name: "E_MEDIUM8",
      value: "e-medium8",
  },
  {
      name: "MEDIUM8",
      value: "medium8",
  },
  {
      name: "E_MEDIUM16",
      value: "e-medium16",
  },
  {
      name: "MEDIUM16",
      value: "medium16",
  },
  {
      name: "E_MEDIUM24",
      value: "e-medium24",
  },
  {
      name: "MEDIUM24",
      value: "medium24",
  },
  {
      name: "C_LARGE",
      value: "c-large",
  },
  {
      name: "E_LARGE",
      value: "e-large",
  },
  {
      name: "LARGE",
      value: "large",
  },
  {
      name: "C_LARGE8",
      value: "c-large8",
  },
  {
      name: "E_LARGE8",
      value: "e-large8",
  },
  {
      name: "LARGE8",
      value: "large8",
  },
  {
      name: "E_LARGE16",
      value: "e-large16",
  },
  {
      name: "LARGE16",
      value: "large16",
  },
  {
      name: "E_LARGE24",
      value: "e-large24",
  },
  {
      name: "LARGE24",
      value: "large24",
  },
  {
      name: "E_LARGE32",
      value: "e-large32",
  },
  {
      name: "LARGE32",
      value: "large32",
  },
  {
      name: "E_EXTRA_LARGE8",
      value: "e-extra-large8",
  },
  {
      name: "EXTRA_LARGE8",
      value: "extra-large8",
  },
  {
      name: "E_EXTRA_LARGE16",
      value: "e-extra-large16",
  },
  {
      name: "EXTRA_LARGE16",
      value: "extra-large16",
  },
  {
      name: "E_EXTRA_LARGE24",
      value: "e-extra-large24",
  },
  {
      name: "EXTRA_LARGE24",
      value: "extra-large24",
  },
  {
      name: "E_EXTRA_LARGE32",
      value: "e-extra-large32",
  },
  {
      name: "EXTRA_LARGE32",
      value: "extra-large32",
  },
  {
      name: "E_EXTRA_LARGE48",
      value: "e-extra-large48",
  },
  {
      name: "EXTRA_LARGE48",
      value: "extra-large48",
  },
  {
      name: "E_DOUBLE_LARGE16",
      value: "e-double-large16",
  },
  {
      name: "DOUBLE_LARGE16",
      value: "double-large16",
  },
  {
      name: "E_DOUBLE_LARGE24",
      value: "e-double-large24",
  },
  {
      name: "DOUBLE_LARGE24",
      value: "double-large24",
  },
  {
      name: "E_DOUBLE_LARGE32",
      value: "e-double-large32",
  },
  {
      name: "DOUBLE_LARGE32",
      value: "double-large32",
  },
  {
      name: "E_DOUBLE_LARGE48",
      value: "e-double-large48",
  },
  {
      name: "DOUBLE_LARGE48",
      value: "double-large48",
  },
  {
      name: "E_DOUBLE_LARGE64",
      value: "e-double-large64",
  },
  {
      name: "DOUBLE_LARGE64",
      value: "double-large64",
  },
  {
      name: "E_DOUBLE_LARGE96",
      value: "e-double-large96",
  },
  {
      name: "DOUBLE_LARGE96",
      value: "double-large96",
  },
])
string InstanceTypeOfclusterForCreateCluster

structure RequestTags {
    @jsonName("key")
    Key: String,
    @jsonName("value")
    Value: String,
}

list ListOfRequestTags {
    member: RequestTags,
}

structure RequestNetworkConfig {
    @jsonName("networkId")
    NetworkId: String,
}

structure CreateClusterRequest {
    @required
    @jsonName("cluster")
    Cluster: RequestCluster,
}

structure CreateClusterResult {
    @jsonName("cluster")
    Cluster: Cluster,
}

structure RequestClusterOfUpdateCluster {
    @jsonName("description")
    Description: String,
    @jsonName("kubernetesVersion")
    KubernetesVersion: KubernetesVersionOfclusterForUpdateCluster,
    @jsonName("tags")
    ListOfRequestTags: ListOfRequestTags,
    @jsonName("name")
    Name: String,
    @jsonName("addonsConfig")
    RequestAddonsConfig: RequestAddonsConfig,
}

@enum([
  {
      name: "V1_23_3",
      value: "v1.23.3",
  },
  {
      name: "V1_23_9",
      value: "v1.23.9",
  },
  {
      name: "V1_24_3",
      value: "v1.24.3",
  },
])
string KubernetesVersionOfclusterForUpdateCluster

structure UpdateClusterRequest {
    @jsonName("cluster")
    Cluster: RequestClusterOfUpdateCluster,
    @required
    @httpLabel
    @jsonName("ClusterName")
    ClusterName: String,
}

structure UpdateClusterResult {
    @jsonName("cluster")
    Cluster: Cluster,
}

structure DeleteClusterRequest {
    @required
    @httpLabel
    @jsonName("ClusterName")
    ClusterName: String,
}

structure DeleteClusterResult {
    @jsonName("cluster")
    Cluster: Cluster,
}

structure DeleteClustersRequest {
    @required
    @httpQuery("names")
    @jsonName("names")
    Names: String,
}

structure DeleteClustersResult {
    @jsonName("clusters")
    Clusters: ListOfClusters,
}

structure GetClusterCredentialsRequest {
    @required
    @httpLabel
    @jsonName("ClusterName")
    ClusterName: String,
}

structure GetClusterCredentialsResult {
    @jsonName("credentials")
    Credentials: String,
}

structure GetNodePoolRequest {
    @required
    @httpLabel
    @jsonName("ClusterName")
    ClusterName: String,
    @required
    @httpLabel
    @jsonName("NodePoolName")
    NodePoolName: String,
}

structure NodePool {
    @jsonName("initialNodeCount")
    InitialNodeCount: Integer,
    @jsonName("instanceType")
    InstanceType: String,
    @jsonName("name")
    Name: String,
    @jsonName("nodeCount")
    NodeCount: Integer,
    @jsonName("nodes")
    Nodes: ListOfNodes,
    @jsonName("nrn")
    Nrn: String,
    @jsonName("status")
    Status: String,
    @jsonName("tags")
    Tags: ListOfTags,
}

structure GetNodePoolResult {
    @jsonName("nodePool")
    NodePool: NodePool,
}

structure ListNodePoolsRequest {
    @required
    @httpLabel
    @jsonName("ClusterName")
    ClusterName: String,
    @httpQuery("filters")
    @jsonName("filters")
    Filters: String,
}

structure ListNodePoolsResult {
    @jsonName("nodePools")
    NodePools: ListOfNodePools,
}

structure RequestNodePool {
    @required
    @jsonName("instanceType")
    InstanceType: InstanceTypeOfnodePoolForCreateNodePool,
    @jsonName("tags")
    ListOfRequestTags: ListOfRequestTags,
    @required
    @jsonName("name")
    Name: String,
    @jsonName("nodeCount")
    NodeCount: Integer,
}

@enum([
  {
      name: "C_MEDIUM",
      value: "c-medium",
  },
  {
      name: "E_MEDIUM",
      value: "e-medium",
  },
  {
      name: "MEDIUM",
      value: "medium",
  },
  {
      name: "C_MEDIUM4",
      value: "c-medium4",
  },
  {
      name: "E_MEDIUM4",
      value: "e-medium4",
  },
  {
      name: "MEDIUM4",
      value: "medium4",
  },
  {
      name: "C_MEDIUM8",
      value: "c-medium8",
  },
  {
      name: "E_MEDIUM8",
      value: "e-medium8",
  },
  {
      name: "MEDIUM8",
      value: "medium8",
  },
  {
      name: "E_MEDIUM16",
      value: "e-medium16",
  },
  {
      name: "MEDIUM16",
      value: "medium16",
  },
  {
      name: "E_MEDIUM24",
      value: "e-medium24",
  },
  {
      name: "MEDIUM24",
      value: "medium24",
  },
  {
      name: "C_LARGE",
      value: "c-large",
  },
  {
      name: "E_LARGE",
      value: "e-large",
  },
  {
      name: "LARGE",
      value: "large",
  },
  {
      name: "C_LARGE8",
      value: "c-large8",
  },
  {
      name: "E_LARGE8",
      value: "e-large8",
  },
  {
      name: "LARGE8",
      value: "large8",
  },
  {
      name: "E_LARGE16",
      value: "e-large16",
  },
  {
      name: "LARGE16",
      value: "large16",
  },
  {
      name: "E_LARGE24",
      value: "e-large24",
  },
  {
      name: "LARGE24",
      value: "large24",
  },
  {
      name: "E_LARGE32",
      value: "e-large32",
  },
  {
      name: "LARGE32",
      value: "large32",
  },
  {
      name: "E_EXTRA_LARGE8",
      value: "e-extra-large8",
  },
  {
      name: "EXTRA_LARGE8",
      value: "extra-large8",
  },
  {
      name: "E_EXTRA_LARGE16",
      value: "e-extra-large16",
  },
  {
      name: "EXTRA_LARGE16",
      value: "extra-large16",
  },
  {
      name: "E_EXTRA_LARGE24",
      value: "e-extra-large24",
  },
  {
      name: "EXTRA_LARGE24",
      value: "extra-large24",
  },
  {
      name: "E_EXTRA_LARGE32",
      value: "e-extra-large32",
  },
  {
      name: "EXTRA_LARGE32",
      value: "extra-large32",
  },
  {
      name: "E_EXTRA_LARGE48",
      value: "e-extra-large48",
  },
  {
      name: "EXTRA_LARGE48",
      value: "extra-large48",
  },
  {
      name: "E_DOUBLE_LARGE16",
      value: "e-double-large16",
  },
  {
      name: "DOUBLE_LARGE16",
      value: "double-large16",
  },
  {
      name: "E_DOUBLE_LARGE24",
      value: "e-double-large24",
  },
  {
      name: "DOUBLE_LARGE24",
      value: "double-large24",
  },
  {
      name: "E_DOUBLE_LARGE32",
      value: "e-double-large32",
  },
  {
      name: "DOUBLE_LARGE32",
      value: "double-large32",
  },
  {
      name: "E_DOUBLE_LARGE48",
      value: "e-double-large48",
  },
  {
      name: "DOUBLE_LARGE48",
      value: "double-large48",
  },
  {
      name: "E_DOUBLE_LARGE64",
      value: "e-double-large64",
  },
  {
      name: "DOUBLE_LARGE64",
      value: "double-large64",
  },
  {
      name: "E_DOUBLE_LARGE96",
      value: "e-double-large96",
  },
  {
      name: "DOUBLE_LARGE96",
      value: "double-large96",
  },
])
string InstanceTypeOfnodePoolForCreateNodePool

structure CreateNodePoolRequest {
    @required
    @httpLabel
    @jsonName("ClusterName")
    ClusterName: String,
    @required
    @jsonName("nodePool")
    NodePool: RequestNodePool,
}

structure CreateNodePoolResult {
    @jsonName("nodePool")
    NodePool: NodePool,
}

structure RequestNodePoolOfUpdateNodePool {
    @jsonName("tags")
    ListOfRequestTags: ListOfRequestTags,
}

structure UpdateNodePoolRequest {
    @required
    @httpLabel
    @jsonName("ClusterName")
    ClusterName: String,
    @jsonName("nodePool")
    NodePool: RequestNodePoolOfUpdateNodePool,
    @required
    @httpLabel
    @jsonName("NodePoolName")
    NodePoolName: String,
}

structure UpdateNodePoolResult {
    @jsonName("nodePool")
    NodePool: NodePool,
}

structure SetNodePoolSizeRequest {
    @required
    @httpLabel
    @jsonName("ClusterName")
    ClusterName: String,
    @required
    @jsonName("nodeCount")
    NodeCount: Integer,
    @required
    @httpLabel
    @jsonName("NodePoolName")
    NodePoolName: String,
}

structure SetNodePoolSizeResult {
    @jsonName("nodePool")
    NodePool: NodePool,
}

structure DeleteNodePoolRequest {
    @required
    @httpLabel
    @jsonName("ClusterName")
    ClusterName: String,
    @required
    @httpLabel
    @jsonName("NodePoolName")
    NodePoolName: String,
}

structure DeleteNodePoolResult {
    @jsonName("nodePool")
    NodePool: NodePool,
}

structure DeleteNodePoolsRequest {
    @required
    @httpLabel
    @jsonName("ClusterName")
    ClusterName: String,
    @required
    @httpQuery("names")
    @jsonName("names")
    Names: String,
}

structure DeleteNodePoolsResult {
    @jsonName("nodePools")
    NodePools: ListOfNodePools,
}

structure RebootNodeRequest {
    @required
    @httpLabel
    @jsonName("ClusterName")
    ClusterName: String,
    @jsonName("force")
    Force: Boolean,
    @required
    @httpLabel
    @jsonName("NodeName")
    NodeName: String,
    @required
    @httpLabel
    @jsonName("NodePoolName")
    NodePoolName: String,
}

structure Node {
    @jsonName("availabilityZone")
    AvailabilityZone: String,
    @jsonName("name")
    Name: String,
    @jsonName("privateIpAddress")
    PrivateIpAddress: String,
    @jsonName("publicIpAddress")
    PublicIpAddress: String,
    @jsonName("status")
    Status: String,
}

structure RebootNodeResult {
    @jsonName("node")
    Node: Node,
}

structure GetServerConfigRequest {}

structure ServerConfig {
    @jsonName("defaultKubernetesVersion")
    DefaultKubernetesVersion: String,
    @jsonName("validKubernetesVersions")
    ValidKubernetesVersions: ListOfValidKubernetesVersions,
}

list ListOfValidKubernetesVersions {
    member: String,
}

structure GetServerConfigResult {
    @jsonName("serverConfig")
    ServerConfig: ServerConfig,
}

structure GetFirewallGroupRequest {
    @required
    @httpLabel
    @jsonName("FirewallGroupName")
    FirewallGroupName: String,
}

structure FirewallGroup {
    @jsonName("description")
    Description: String,
    @jsonName("name")
    Name: String,
    @jsonName("nrn")
    Nrn: String,
    @jsonName("rules")
    Rules: ListOfRules,
    @jsonName("tags")
    Tags: ListOfTags,
}

list ListOfRules {
    member: Rules,
}

structure Rules {
    @jsonName("cidrIp")
    CidrIp: String,
    @jsonName("description")
    Description: String,
    @jsonName("direction")
    Direction: String,
    @jsonName("fromPort")
    FromPort: Integer,
    @jsonName("id")
    Id: String,
    @jsonName("protocol")
    Protocol: String,
    @jsonName("status")
    Status: String,
    @jsonName("toPort")
    ToPort: Integer,
}

structure GetFirewallGroupResult {
    @jsonName("firewallGroup")
    FirewallGroup: FirewallGroup,
}

structure ListFirewallGroupsRequest {
    @httpQuery("filters")
    @jsonName("filters")
    Filters: String,
}

list ListOfFirewallGroups {
    member: FirewallGroups,
}

structure FirewallGroups {
    @jsonName("description")
    Description: String,
    @jsonName("name")
    Name: String,
    @jsonName("nrn")
    Nrn: String,
    @jsonName("rules")
    Rules: ListOfRules,
    @jsonName("tags")
    Tags: ListOfTags,
}

structure ListFirewallGroupsResult {
    @jsonName("firewallGroups")
    FirewallGroups: ListOfFirewallGroups,
}

structure RequestFirewallGroup {
    @jsonName("description")
    Description: String,
    @jsonName("tags")
    ListOfRequestTags: ListOfRequestTags,
    @required
    @jsonName("name")
    Name: String,
}

structure CreateFirewallGroupRequest {
    @required
    @jsonName("firewallGroup")
    FirewallGroup: RequestFirewallGroup,
}

structure CreateFirewallGroupResult {
    @jsonName("firewallGroup")
    FirewallGroup: FirewallGroup,
}

structure RequestFirewallGroupOfUpdateFirewallGroup {
    @jsonName("description")
    Description: String,
    @jsonName("tags")
    ListOfRequestTags: ListOfRequestTags,
    @jsonName("name")
    Name: String,
}

structure UpdateFirewallGroupRequest {
    @jsonName("firewallGroup")
    FirewallGroup: RequestFirewallGroupOfUpdateFirewallGroup,
    @required
    @httpLabel
    @jsonName("FirewallGroupName")
    FirewallGroupName: String,
}

structure UpdateFirewallGroupResult {
    @jsonName("firewallGroup")
    FirewallGroup: FirewallGroup,
}

structure DeleteFirewallGroupRequest {
    @required
    @httpLabel
    @jsonName("FirewallGroupName")
    FirewallGroupName: String,
}

structure DeleteFirewallGroupResult {
    @jsonName("firewallGroup")
    FirewallGroup: FirewallGroup,
}

structure DeleteFirewallGroupsRequest {
    @httpQuery("names")
    @jsonName("names")
    Names: String,
}

structure DeleteFirewallGroupsResult {
    @jsonName("firewallGroups")
    FirewallGroups: ListOfFirewallGroups,
}

@enum([
  {
      name: "ANY",
      value: "ANY",
  },
  {
      name: "TCP",
      value: "TCP",
  },
  {
      name: "UDP",
      value: "UDP",
  },
  {
      name: "ICMP",
      value: "ICMP",
  },
  {
      name: "SSH",
      value: "SSH",
  },
  {
      name: "HTTP",
      value: "HTTP",
  },
  {
      name: "HTTPS",
      value: "HTTPS",
  },
  {
      name: "RDP",
      value: "RDP",
  },
  {
      name: "GRE",
      value: "GRE",
  },
  {
      name: "ESP",
      value: "ESP",
  },
  {
      name: "AH",
      value: "AH",
  },
  {
      name: "VRRP",
      value: "VRRP",
  },
  {
      name: "L2TP",
      value: "L2TP",
  },
])
string ProtocolOfrulesForAuthorizeFirewallGroup

structure RequestRules {
    @required
    @jsonName("cidrIp")
    CidrIp: String,
    @jsonName("description")
    Description: String,
    @jsonName("direction")
    Direction: DirectionOfrulesForAuthorizeFirewallGroup,
    @jsonName("fromPort")
    FromPort: Integer,
    @jsonName("protocol")
    Protocol: ProtocolOfrulesForAuthorizeFirewallGroup,
    @jsonName("toPort")
    ToPort: Integer,
}

list ListOfRequestRules {
    member: RequestRules,
}

@enum([
  {
      name: "IN",
      value: "IN",
  },
  {
      name: "OUT",
      value: "OUT",
  },
])
string DirectionOfrulesForAuthorizeFirewallGroup

structure AuthorizeFirewallGroupRequest {
    @required
    @httpLabel
    @jsonName("FirewallGroupName")
    FirewallGroupName: String,
    @required
    @jsonName("rules")
    Rules: ListOfRequestRules,
}

structure AuthorizeFirewallGroupResult {
    @jsonName("firewallGroup")
    FirewallGroup: FirewallGroup,
}

structure RevokeFirewallGroupRequest {
    @required
    @httpLabel
    @jsonName("FirewallGroupName")
    FirewallGroupName: String,
    @httpQuery("ids")
    @jsonName("ids")
    Ids: String,
}

structure RevokeFirewallGroupResult {
    @jsonName("firewallGroup")
    FirewallGroup: FirewallGroup,
}

structure GetDiskRequest {
    @required
    @httpLabel
    @jsonName("DiskName")
    DiskName: String,
}

structure Disk {
    @jsonName("attachments")
    Attachments: ListOfAttachments,
    @jsonName("availability_zone")
    AvailabilityZone: String,
    @jsonName("cluster")
    Cluster: ClusterOfGetDisk,
    @jsonName("createTime")
    CreateTime: String,
    @jsonName("description")
    Description: String,
    @jsonName("name")
    Name: String,
    @jsonName("nrn")
    Nrn: String,
    @jsonName("size")
    Size: Integer,
    @jsonName("status")
    Status: String,
    @jsonName("tags")
    Tags: ListOfTags,
    @jsonName("type")
    Type: String,
}

structure ClusterOfGetDisk {
    @jsonName("name")
    Name: String,
}

list ListOfAttachments {
    member: Attachments,
}

structure Attachments {
    @jsonName("attachTime")
    AttachTime: String,
    @jsonName("devicePath")
    DevicePath: String,
    @jsonName("nodeName")
    NodeName: String,
    @jsonName("status")
    Status: String,
}

structure GetDiskResult {
    @jsonName("disk")
    Disk: Disk,
}

structure ListDisksRequest {
    @httpQuery("filters")
    @jsonName("filters")
    Filters: String,
}

list ListOfDisks {
    member: Disks,
}

structure Disks {
    @jsonName("attachments")
    Attachments: ListOfAttachments,
    @jsonName("availability_zone")
    AvailabilityZone: String,
    @jsonName("cluster")
    Cluster: ClusterOfListDisks,
    @jsonName("createTime")
    CreateTime: String,
    @jsonName("description")
    Description: String,
    @jsonName("name")
    Name: String,
    @jsonName("nrn")
    Nrn: String,
    @jsonName("size")
    Size: Integer,
    @jsonName("status")
    Status: String,
    @jsonName("tags")
    Tags: ListOfTags,
    @jsonName("type")
    Type: String,
}

structure ClusterOfListDisks {
    @jsonName("name")
    Name: String,
}

structure ListDisksResult {
    @jsonName("disks")
    Disks: ListOfDisks,
}

structure RequestDisk {
    @jsonName("availabilityZone")
    AvailabilityZone: String,
    @jsonName("description")
    Description: String,
    @jsonName("tags")
    ListOfRequestTags: ListOfRequestTags,
    @required
    @jsonName("name")
    Name: String,
    @required
    @jsonName("size")
    Size: Integer,
    @required
    @jsonName("type")
    Type: TypeOfdiskForCreateDisk,
}

@enum([
  {
      name: "STANDARD_FLASH_A",
      value: "standard-flash-a",
  },
  {
      name: "STANDARD_FLASH_B",
      value: "standard-flash-b",
  },
  {
      name: "HIGH_SPEED_FLASH_A",
      value: "high-speed-flash-a",
  },
  {
      name: "HIGH_SPEED_FLASH_B",
      value: "high-speed-flash-b",
  },
])
string TypeOfdiskForCreateDisk

structure CreateDiskRequest {
    @required
    @jsonName("disk")
    Disk: RequestDisk,
}

structure DiskOfCreateDisk {
    @jsonName("attachments")
    Attachments: ListOfAttachments,
    @jsonName("availability_zone")
    AvailabilityZone: String,
    @jsonName("cluster")
    Cluster: ClusterOfCreateDisk,
    @jsonName("createTime")
    CreateTime: String,
    @jsonName("description")
    Description: String,
    @jsonName("name")
    Name: String,
    @jsonName("nrn")
    Nrn: String,
    @jsonName("size")
    Size: Integer,
    @jsonName("status")
    Status: String,
    @jsonName("tags")
    Tags: ListOfTags,
    @jsonName("type")
    Type: String,
}

structure ClusterOfCreateDisk {
    @jsonName("name")
    Name: String,
}

structure CreateDiskResult {
    @jsonName("disk")
    Disk: DiskOfCreateDisk,
}

structure RequestDiskOfUpdateDisk {
    @jsonName("description")
    Description: String,
    @jsonName("tags")
    ListOfRequestTags: ListOfRequestTags,
    @jsonName("name")
    Name: String,
    @jsonName("size")
    Size: Integer,
}

structure UpdateDiskRequest {
    @jsonName("disk")
    Disk: RequestDiskOfUpdateDisk,
    @required
    @httpLabel
    @jsonName("DiskName")
    DiskName: String,
}

structure DiskOfUpdateDisk {
    @jsonName("attachments")
    Attachments: ListOfAttachments,
    @jsonName("availability_zone")
    AvailabilityZone: String,
    @jsonName("cluster")
    Cluster: ClusterOfUpdateDisk,
    @jsonName("createTime")
    CreateTime: String,
    @jsonName("description")
    Description: String,
    @jsonName("name")
    Name: String,
    @jsonName("nrn")
    Nrn: String,
    @jsonName("size")
    Size: Integer,
    @jsonName("status")
    Status: String,
    @jsonName("tags")
    Tags: ListOfTags,
    @jsonName("type")
    Type: String,
}

structure ClusterOfUpdateDisk {
    @jsonName("name")
    Name: String,
}

structure UpdateDiskResult {
    @jsonName("disk")
    Disk: DiskOfUpdateDisk,
}

structure DeleteDiskRequest {
    @required
    @httpLabel
    @jsonName("DiskName")
    DiskName: String,
}

structure DiskOfDeleteDisk {
    @jsonName("attachments")
    Attachments: ListOfAttachments,
    @jsonName("availability_zone")
    AvailabilityZone: String,
    @jsonName("cluster")
    Cluster: ClusterOfDeleteDisk,
    @jsonName("createTime")
    CreateTime: String,
    @jsonName("description")
    Description: String,
    @jsonName("name")
    Name: String,
    @jsonName("nrn")
    Nrn: String,
    @jsonName("size")
    Size: Integer,
    @jsonName("status")
    Status: String,
    @jsonName("tags")
    Tags: ListOfTags,
    @jsonName("type")
    Type: String,
}

structure ClusterOfDeleteDisk {
    @jsonName("name")
    Name: String,
}

structure DeleteDiskResult {
    @jsonName("disk")
    Disk: DiskOfDeleteDisk,
}

structure DeleteDisksRequest {
    @required
    @httpQuery("names")
    @jsonName("names")
    Names: String,
}

list ListOfDisksOfDeleteDisks {
    member: DisksOfDeleteDisks,
}

structure DisksOfDeleteDisks {
    @jsonName("attachments")
    Attachments: ListOfAttachments,
    @jsonName("availability_zone")
    AvailabilityZone: String,
    @jsonName("cluster")
    Cluster: ClusterOfDeleteDisks,
    @jsonName("createTime")
    CreateTime: String,
    @jsonName("description")
    Description: String,
    @jsonName("name")
    Name: String,
    @jsonName("nrn")
    Nrn: String,
    @jsonName("size")
    Size: Integer,
    @jsonName("status")
    Status: String,
    @jsonName("tags")
    Tags: ListOfTags,
    @jsonName("type")
    Type: String,
}

structure ClusterOfDeleteDisks {
    @jsonName("name")
    Name: String,
}

structure DeleteDisksResult {
    @jsonName("disks")
    Disks: ListOfDisksOfDeleteDisks,
}

structure AttachDiskRequest {
    @required
    @httpLabel
    @jsonName("DiskName")
    DiskName: String,
    @required
    @jsonName("nodeName")
    NodeName: String,
}

structure DiskOfAttachDisk {
    @jsonName("attachments")
    Attachments: ListOfAttachments,
    @jsonName("availability_zone")
    AvailabilityZone: String,
    @jsonName("cluster")
    Cluster: ClusterOfAttachDisk,
    @jsonName("createTime")
    CreateTime: String,
    @jsonName("description")
    Description: String,
    @jsonName("name")
    Name: String,
    @jsonName("nrn")
    Nrn: String,
    @jsonName("size")
    Size: Integer,
    @jsonName("status")
    Status: String,
    @jsonName("tags")
    Tags: ListOfTags,
    @jsonName("type")
    Type: String,
}

structure ClusterOfAttachDisk {
    @jsonName("name")
    Name: String,
}

structure AttachDiskResult {
    @jsonName("disk")
    Disk: DiskOfAttachDisk,
}

structure DetachDiskRequest {
    @required
    @httpLabel
    @jsonName("DiskName")
    DiskName: String,
    @required
    @jsonName("nodeName")
    NodeName: String,
}

structure DiskOfDetachDisk {
    @jsonName("attachments")
    Attachments: ListOfAttachments,
    @jsonName("availability_zone")
    AvailabilityZone: String,
    @jsonName("cluster")
    Cluster: ClusterOfDetachDisk,
    @jsonName("createTime")
    CreateTime: String,
    @jsonName("description")
    Description: String,
    @jsonName("name")
    Name: String,
    @jsonName("nrn")
    Nrn: String,
    @jsonName("size")
    Size: Integer,
    @jsonName("status")
    Status: String,
    @jsonName("tags")
    Tags: ListOfTags,
    @jsonName("type")
    Type: String,
}

structure ClusterOfDetachDisk {
    @jsonName("name")
    Name: String,
}

structure DetachDiskResult {
    @jsonName("disk")
    Disk: DiskOfDetachDisk,
}

structure GetSnapshotRequest {
    @required
    @httpLabel
    @jsonName("SnapshotName")
    SnapshotName: String,
}

structure Snapshot {
    @jsonName("cluster")
    Cluster: ClusterOfGetSnapshot,
    @jsonName("createTime")
    CreateTime: String,
    @jsonName("description")
    Description: String,
    @jsonName("expirationTime")
    ExpirationTime: String,
    @jsonName("name")
    Name: String,
    @jsonName("nrn")
    Nrn: String,
    @jsonName("resourceVersion")
    ResourceVersion: String,
    @jsonName("status")
    Status: String,
    @jsonName("tags")
    Tags: ListOfTags,
}

structure ClusterOfGetSnapshot {
    @jsonName("kubernetesVersion")
    KubernetesVersion: String,
    @jsonName("name")
    Name: String,
    @jsonName("nodePools")
    NodePools: ListOfNodePoolsOfGetSnapshot,
}

list ListOfNodePoolsOfGetSnapshot {
    member: NodePoolsOfGetSnapshot,
}

structure NodePoolsOfGetSnapshot {
    @jsonName("instanceType")
    InstanceType: String,
    @jsonName("name")
    Name: String,
    @jsonName("nodeCount")
    NodeCount: Integer,
}

structure GetSnapshotResult {
    @jsonName("snapshot")
    Snapshot: Snapshot,
}

structure ListSnapshotsRequest {
    @httpQuery("filters")
    @jsonName("filters")
    Filters: String,
}

list ListOfSnapshots {
    member: Snapshots,
}

structure Snapshots {
    @jsonName("cluster")
    Cluster: ClusterOfListSnapshots,
    @jsonName("createTime")
    CreateTime: String,
    @jsonName("description")
    Description: String,
    @jsonName("expirationTime")
    ExpirationTime: String,
    @jsonName("name")
    Name: String,
    @jsonName("nrn")
    Nrn: String,
    @jsonName("resourceVersion")
    ResourceVersion: String,
    @jsonName("status")
    Status: String,
    @jsonName("tags")
    Tags: ListOfTags,
}

structure ClusterOfListSnapshots {
    @jsonName("kubernetesVersion")
    KubernetesVersion: String,
    @jsonName("name")
    Name: String,
    @jsonName("nodePools")
    NodePools: ListOfNodePoolsOfListSnapshots,
}

list ListOfNodePoolsOfListSnapshots {
    member: NodePoolsOfListSnapshots,
}

structure NodePoolsOfListSnapshots {
    @jsonName("instanceType")
    InstanceType: String,
    @jsonName("name")
    Name: String,
    @jsonName("nodeCount")
    NodeCount: Integer,
}

structure ListSnapshotsResult {
    @jsonName("snapshots")
    Snapshots: ListOfSnapshots,
}

structure RequestSnapshot {
    @jsonName("description")
    Description: String,
    @jsonName("expirationTime")
    ExpirationTime: String,
    @jsonName("tags")
    ListOfRequestTags: ListOfRequestTags,
    @required
    @jsonName("name")
    Name: String,
    @required
    @jsonName("cluster")
    RequestCluster: RequestClusterOfCreateSnapshot,
}

structure RequestClusterOfCreateSnapshot {
    @required
    @jsonName("name")
    Name: String,
}

structure CreateSnapshotRequest {
    @required
    @jsonName("snapshot")
    Snapshot: RequestSnapshot,
}

structure SnapshotOfCreateSnapshot {
    @jsonName("cluster")
    Cluster: ClusterOfCreateSnapshot,
    @jsonName("createTime")
    CreateTime: String,
    @jsonName("description")
    Description: String,
    @jsonName("expirationTime")
    ExpirationTime: String,
    @jsonName("name")
    Name: String,
    @jsonName("nrn")
    Nrn: String,
    @jsonName("resourceVersion")
    ResourceVersion: String,
    @jsonName("status")
    Status: String,
    @jsonName("tags")
    Tags: ListOfTags,
}

structure ClusterOfCreateSnapshot {
    @jsonName("kubernetesVersion")
    KubernetesVersion: String,
    @jsonName("name")
    Name: String,
    @jsonName("nodePools")
    NodePools: ListOfNodePoolsOfCreateSnapshot,
}

list ListOfNodePoolsOfCreateSnapshot {
    member: NodePoolsOfCreateSnapshot,
}

structure NodePoolsOfCreateSnapshot {
    @jsonName("instanceType")
    InstanceType: String,
    @jsonName("name")
    Name: String,
    @jsonName("nodeCount")
    NodeCount: Integer,
}

structure CreateSnapshotResult {
    @jsonName("snapshot")
    Snapshot: SnapshotOfCreateSnapshot,
}

structure RequestSnapshotOfUpdateSnapshot {
    @jsonName("description")
    Description: String,
    @jsonName("expirationTime")
    ExpirationTime: String,
    @jsonName("tags")
    ListOfRequestTags: ListOfRequestTags,
    @jsonName("name")
    Name: String,
}

structure UpdateSnapshotRequest {
    @jsonName("snapshot")
    Snapshot: RequestSnapshotOfUpdateSnapshot,
    @required
    @httpLabel
    @jsonName("SnapshotName")
    SnapshotName: String,
}

structure SnapshotOfUpdateSnapshot {
    @jsonName("cluster")
    Cluster: ClusterOfUpdateSnapshot,
    @jsonName("createTime")
    CreateTime: String,
    @jsonName("description")
    Description: String,
    @jsonName("expirationTime")
    ExpirationTime: String,
    @jsonName("name")
    Name: String,
    @jsonName("nrn")
    Nrn: String,
    @jsonName("resourceVersion")
    ResourceVersion: String,
    @jsonName("status")
    Status: String,
    @jsonName("tags")
    Tags: ListOfTags,
}

structure ClusterOfUpdateSnapshot {
    @jsonName("kubernetesVersion")
    KubernetesVersion: String,
    @jsonName("name")
    Name: String,
    @jsonName("nodePools")
    NodePools: ListOfNodePoolsOfUpdateSnapshot,
}

list ListOfNodePoolsOfUpdateSnapshot {
    member: NodePoolsOfUpdateSnapshot,
}

structure NodePoolsOfUpdateSnapshot {
    @jsonName("instanceType")
    InstanceType: String,
    @jsonName("name")
    Name: String,
    @jsonName("nodeCount")
    NodeCount: Integer,
}

structure UpdateSnapshotResult {
    @jsonName("snapshot")
    Snapshot: SnapshotOfUpdateSnapshot,
}

structure DeleteSnapshotRequest {
    @required
    @httpLabel
    @jsonName("SnapshotName")
    SnapshotName: String,
}

structure SnapshotOfDeleteSnapshot {
    @jsonName("cluster")
    Cluster: ClusterOfDeleteSnapshot,
    @jsonName("createTime")
    CreateTime: String,
    @jsonName("description")
    Description: String,
    @jsonName("expirationTime")
    ExpirationTime: String,
    @jsonName("name")
    Name: String,
    @jsonName("nrn")
    Nrn: String,
    @jsonName("resourceVersion")
    ResourceVersion: String,
    @jsonName("status")
    Status: String,
    @jsonName("tags")
    Tags: ListOfTags,
}

structure ClusterOfDeleteSnapshot {
    @jsonName("kubernetesVersion")
    KubernetesVersion: String,
    @jsonName("name")
    Name: String,
    @jsonName("nodePools")
    NodePools: ListOfNodePoolsOfDeleteSnapshot,
}

list ListOfNodePoolsOfDeleteSnapshot {
    member: NodePoolsOfDeleteSnapshot,
}

structure NodePoolsOfDeleteSnapshot {
    @jsonName("instanceType")
    InstanceType: String,
    @jsonName("name")
    Name: String,
    @jsonName("nodeCount")
    NodeCount: Integer,
}

structure DeleteSnapshotResult {
    @jsonName("snapshot")
    Snapshot: SnapshotOfDeleteSnapshot,
}

structure DeleteSnapshotsRequest {
    @httpQuery("names")
    @jsonName("names")
    Names: String,
}

list ListOfSnapshotsOfDeleteSnapshots {
    member: SnapshotsOfDeleteSnapshots,
}

structure SnapshotsOfDeleteSnapshots {
    @jsonName("cluster")
    Cluster: ClusterOfDeleteSnapshots,
    @jsonName("createTime")
    CreateTime: String,
    @jsonName("description")
    Description: String,
    @jsonName("expirationTime")
    ExpirationTime: String,
    @jsonName("name")
    Name: String,
    @jsonName("nrn")
    Nrn: String,
    @jsonName("resourceVersion")
    ResourceVersion: String,
    @jsonName("status")
    Status: String,
    @jsonName("tags")
    Tags: ListOfTags,
}

structure ClusterOfDeleteSnapshots {
    @jsonName("kubernetesVersion")
    KubernetesVersion: String,
    @jsonName("name")
    Name: String,
    @jsonName("nodePools")
    NodePools: ListOfNodePoolsOfDeleteSnapshots,
}

list ListOfNodePoolsOfDeleteSnapshots {
    member: NodePoolsOfDeleteSnapshots,
}

structure NodePoolsOfDeleteSnapshots {
    @jsonName("instanceType")
    InstanceType: String,
    @jsonName("name")
    Name: String,
    @jsonName("nodeCount")
    NodeCount: Integer,
}

structure DeleteSnapshotsResult {
    @jsonName("snapshots")
    Snapshots: ListOfSnapshotsOfDeleteSnapshots,
}

structure RequestClusterOfRestoreClusterFromSnapshot {
    @jsonName("description")
    Description: String,
    @required
    @jsonName("firewallGroup")
    FirewallGroup: String,
    @required
    @jsonName("locations")
    ListOfRequestLocations: ListOfRequestLocations,
    @jsonName("tags")
    ListOfRequestTags: ListOfRequestTags,
    @required
    @jsonName("name")
    Name: String,
    @jsonName("addonsConfig")
    RequestAddonsConfig: RequestAddonsConfig,
    @jsonName("networkConfig")
    RequestNetworkConfig: RequestNetworkConfig,
}

structure RestoreClusterFromSnapshotRequest {
    @required
    @jsonName("cluster")
    Cluster: RequestClusterOfRestoreClusterFromSnapshot,
    @required
    @httpLabel
    @jsonName("SnapshotName")
    SnapshotName: String,
}

structure RestoreClusterFromSnapshotResult {
    @jsonName("cluster")
    Cluster: Cluster,
}

structure GetLoadBalancerRequest {
    @required
    @httpQuery("instancePort")
    @jsonName("instancePort")
    InstancePort: Integer,
    @required
    @httpLabel
    @jsonName("LoadBalancerName")
    LoadBalancerName: String,
    @required
    @httpQuery("loadBalancerPort")
    @jsonName("loadBalancerPort")
    LoadBalancerPort: Integer,
}

structure LoadBalancers {
    @jsonName("accountingType")
    AccountingType: Integer,
    @jsonName("availabilityZones")
    AvailabilityZones: ListOfAvailabilityZones,
    @jsonName("clusters")
    Clusters: ListOfClustersOfGetLoadBalancer,
    @jsonName("createdTime")
    CreatedTime: String,
    @jsonName("description")
    Description: String,
    @jsonName("dnsName")
    DnsName: String,
    @jsonName("filter")
    Filter: Filter,
    @jsonName("healthCheck")
    HealthCheck: HealthCheck,
    @jsonName("listenerDescriptions")
    ListenerDescriptions: ListOfListenerDescriptions,
    @jsonName("loadBalancerName")
    LoadBalancerName: String,
    @jsonName("networkVolume")
    NetworkVolume: String,
    @jsonName("nextMonthAccountingType")
    NextMonthAccountingType: Integer,
    @jsonName("option")
    Option: Option,
    @jsonName("policyType")
    PolicyType: String,
}

list ListOfAvailabilityZones {
    member: String,
}

list ListOfClustersOfGetLoadBalancer {
    member: ClustersOfGetLoadBalancer,
}

structure ClustersOfGetLoadBalancer {
    @jsonName("name")
    Name: String,
    @jsonName("nodePools")
    NodePools: ListOfNodePoolsOfGetLoadBalancer,
}

list ListOfNodePoolsOfGetLoadBalancer {
    member: NodePoolsOfGetLoadBalancer,
}

structure NodePoolsOfGetLoadBalancer {
    @jsonName("name")
    Name: String,
    @jsonName("nodeCount")
    NodeCount: Integer,
    @jsonName("nodes")
    Nodes: ListOfNodesOfGetLoadBalancer,
}

list ListOfNodesOfGetLoadBalancer {
    member: NodesOfGetLoadBalancer,
}

structure NodesOfGetLoadBalancer {
    @jsonName("availabilityZone")
    AvailabilityZone: String,
    @jsonName("healthCheckState")
    HealthCheckState: String,
    @jsonName("name")
    Name: String,
    @jsonName("publicIpAddress")
    PublicIpAddress: String,
}

structure Filter {
    @jsonName("filterType")
    FilterType: String,
    @jsonName("ipAddresses")
    IpAddresses: String,
}

structure HealthCheck {
    @jsonName("healthyThreshold")
    HealthyThreshold: Integer,
    @jsonName("interval")
    Interval: Integer,
    @jsonName("target")
    Target: String,
    @jsonName("timeout")
    Timeout: Integer,
    @jsonName("unhealthyThreshold")
    UnhealthyThreshold: Integer,
}

list ListOfListenerDescriptions {
    member: ListenerDescriptions,
}

structure ListenerDescriptions {
    @jsonName("listener")
    Listener: Listener,
}

structure Listener {
    @jsonName("balancingType")
    BalancingType: String,
    @jsonName("instancePort")
    InstancePort: String,
    @jsonName("loadBalancerPort")
    LoadBalancerPort: String,
    @jsonName("protocol")
    Protocol: String,
    @jsonName("sslCertificateId")
    SslCertificateId: String,
}

structure Option {
    @jsonName("sessionStickinessPolicy")
    SessionStickinessPolicy: SessionStickinessPolicy,
    @jsonName("sorryPage")
    SorryPage: SorryPage,
}

structure SessionStickinessPolicy {
    @jsonName("enabled")
    Enabled: Boolean,
    @jsonName("expirationPeriod")
    ExpirationPeriod: Integer,
}

structure SorryPage {
    @jsonName("enabled")
    Enabled: Boolean,
    @jsonName("statusCode")
    StatusCode: Integer,
}

structure GetLoadBalancerResult {
    @jsonName("loadBalancers")
    LoadBalancers: LoadBalancers,
}

structure ListLoadBalancersRequest {}

list ListOfLoadBalancersOfListLoadBalancers {
    member: LoadBalancersOfListLoadBalancers,
}

structure LoadBalancersOfListLoadBalancers {
    @jsonName("accountingType")
    AccountingType: Integer,
    @jsonName("availabilityZones")
    AvailabilityZones: ListOfAvailabilityZones,
    @jsonName("clusters")
    Clusters: ListOfClustersOfListLoadBalancers,
    @jsonName("createdTime")
    CreatedTime: String,
    @jsonName("description")
    Description: String,
    @jsonName("dnsName")
    DnsName: String,
    @jsonName("filter")
    Filter: Filter,
    @jsonName("healthCheck")
    HealthCheck: HealthCheck,
    @jsonName("listenerDescriptions")
    ListenerDescriptions: ListOfListenerDescriptions,
    @jsonName("loadBalancerName")
    LoadBalancerName: String,
    @jsonName("networkVolume")
    NetworkVolume: String,
    @jsonName("nextMonthAccountingType")
    NextMonthAccountingType: Integer,
    @jsonName("option")
    Option: Option,
    @jsonName("policyType")
    PolicyType: String,
}

list ListOfClustersOfListLoadBalancers {
    member: ClustersOfListLoadBalancers,
}

structure ClustersOfListLoadBalancers {
    @jsonName("name")
    Name: String,
    @jsonName("nodePools")
    NodePools: ListOfNodePoolsOfListLoadBalancers,
}

list ListOfNodePoolsOfListLoadBalancers {
    member: NodePoolsOfListLoadBalancers,
}

structure NodePoolsOfListLoadBalancers {
    @jsonName("name")
    Name: String,
    @jsonName("nodeCount")
    NodeCount: Integer,
    @jsonName("nodes")
    Nodes: ListOfNodesOfListLoadBalancers,
}

list ListOfNodesOfListLoadBalancers {
    member: NodesOfListLoadBalancers,
}

structure NodesOfListLoadBalancers {
    @jsonName("availabilityZone")
    AvailabilityZone: String,
    @jsonName("healthCheckState")
    HealthCheckState: String,
    @jsonName("name")
    Name: String,
    @jsonName("publicIpAddress")
    PublicIpAddress: String,
}

structure ListLoadBalancersResult {
    @jsonName("loadBalancers")
    LoadBalancers: ListOfLoadBalancersOfListLoadBalancers,
}

structure ListTagsRequest {
    @httpQuery("nrn")
    @jsonName("nrn")
    Nrn: String,
}

list ListOfTagsOfListTags {
    member: TagsOfListTags,
}

structure TagsOfListTags {
    @jsonName("id")
    Id: String,
    @jsonName("key")
    Key: String,
    @jsonName("nrn")
    Nrn: String,
    @jsonName("value")
    Value: String,
}

structure ListTagsResult {
    @jsonName("tags")
    Tags: ListOfTagsOfListTags,
}

structure RequestTagsOfCreateTags {
    @required
    @jsonName("key")
    Key: String,
    @required
    @jsonName("nrn")
    Nrn: String,
    @required
    @jsonName("value")
    Value: String,
}

list ListOfRequestTagsOfCreateTags {
    member: RequestTagsOfCreateTags,
}

structure CreateTagsRequest {
    @required
    @jsonName("tags")
    Tags: ListOfRequestTagsOfCreateTags,
}

list ListOfTagsOfCreateTags {
    member: TagsOfCreateTags,
}

structure TagsOfCreateTags {
    @jsonName("id")
    Id: String,
    @jsonName("key")
    Key: String,
    @jsonName("nrn")
    Nrn: String,
    @jsonName("value")
    Value: String,
}

structure CreateTagsResult {
    @jsonName("tags")
    Tags: ListOfTagsOfCreateTags,
}

structure RequestTagsOfUpdateTags {
    @required
    @jsonName("key")
    Key: String,
    @required
    @jsonName("nrn")
    Nrn: String,
    @required
    @jsonName("value")
    Value: String,
}

list ListOfRequestTagsOfUpdateTags {
    member: RequestTagsOfUpdateTags,
}

structure UpdateTagsRequest {
    @required
    @jsonName("tags")
    Tags: ListOfRequestTagsOfUpdateTags,
}

list ListOfTagsOfUpdateTags {
    member: TagsOfUpdateTags,
}

structure TagsOfUpdateTags {
    @jsonName("id")
    Id: String,
    @jsonName("key")
    Key: String,
    @jsonName("nrn")
    Nrn: String,
    @jsonName("value")
    Value: String,
}

structure UpdateTagsResult {
    @jsonName("tags")
    Tags: ListOfTagsOfUpdateTags,
}

structure DeleteTagsRequest {
    @required
    @httpQuery("ids")
    @jsonName("ids")
    Ids: String,
}

list ListOfTagsOfDeleteTags {
    member: TagsOfDeleteTags,
}

structure TagsOfDeleteTags {
    @jsonName("id")
    Id: String,
    @jsonName("key")
    Key: String,
    @jsonName("nrn")
    Nrn: String,
    @jsonName("value")
    Value: String,
}

structure DeleteTagsResult {
    @jsonName("tags")
    Tags: ListOfTagsOfDeleteTags,
}

@restJson1
@sigv4(name: "hatoba")
@service(
    sdkId: "hatoba",
    arnNamespace: "hatoba",
    endpointPrefix: "hatoba"
)
service KubernetesServiceHatoba {
    version: "v1",
    operations: [
        GetCluster,
        ListClusters,
        CreateCluster,
        UpdateCluster,
        DeleteCluster,
        DeleteClusters,
        GetClusterCredentials,
        GetNodePool,
        ListNodePools,
        CreateNodePool,
        UpdateNodePool,
        SetNodePoolSize,
        DeleteNodePool,
        DeleteNodePools,
        RebootNode,
        GetServerConfig,
        GetFirewallGroup,
        ListFirewallGroups,
        CreateFirewallGroup,
        UpdateFirewallGroup,
        DeleteFirewallGroup,
        DeleteFirewallGroups,
        AuthorizeFirewallGroup,
        RevokeFirewallGroup,
        GetDisk,
        ListDisks,
        CreateDisk,
        UpdateDisk,
        DeleteDisk,
        DeleteDisks,
        AttachDisk,
        DetachDisk,
        GetSnapshot,
        ListSnapshots,
        CreateSnapshot,
        UpdateSnapshot,
        DeleteSnapshot,
        DeleteSnapshots,
        RestoreClusterFromSnapshot,
        GetLoadBalancer,
        ListLoadBalancers,
        ListTags,
        CreateTags,
        UpdateTags,
        DeleteTags,
    ],
}


@http(method: "GET", uri: "/v1/clusters/{ClusterName}" )
@readonly
@waitable(
    ClusterRunning: {
        acceptors: [
            {
                state: "success",
                matcher: {
                    output: {
                         path: "Cluster.Status",
                         comparator: "stringEquals",
                         expected: "RUNNING",
                    },
                },
            },
            {
                state: "failure",
                matcher: {
                    output: {
                         path: "Cluster.Status",
                         comparator: "stringEquals",
                         expected: "ERROR",
                    },
                },
            },
            {
                state: "failure",
                matcher: {
                    output: {
                         path: "Cluster.Status",
                         comparator: "stringEquals",
                         expected: "DEGRADED",
                    },
                },
            },
            {
                state: "failure",
                matcher: {
                    output: {
                         path: "Cluster.Status",
                         comparator: "stringEquals",
                         expected: "STOPPING",
                    },
                },
            },
        ],
        minDelay: 60,
    },
    ClusterDeleted: {
        acceptors: [
            {
                state: "success",
                matcher: {
                    errorType: "Client.InvalidParameterNotFound.Cluster",
                },
            },
            {
                state: "failure",
                matcher: {
                    output: {
                         path: "Cluster.Status",
                         comparator: "stringEquals",
                         expected: "RUNNING",
                    },
                },
            },
            {
                state: "failure",
                matcher: {
                    output: {
                         path: "Cluster.Status",
                         comparator: "stringEquals",
                         expected: "ERROR",
                    },
                },
            },
            {
                state: "failure",
                matcher: {
                    output: {
                         path: "Cluster.Status",
                         comparator: "stringEquals",
                         expected: "DEGRADED",
                    },
                },
            },
        ],
        minDelay: 30,
    },
)
operation GetCluster {
    input: GetClusterRequest,
    output: GetClusterResult,
}

@http(method: "GET", uri: "/v1/clusters" )
@readonly
operation ListClusters {
    input: ListClustersRequest,
    output: ListClustersResult,
}

@http(method: "POST", uri: "/v1/clusters" , code: 201)
operation CreateCluster {
    input: CreateClusterRequest,
    output: CreateClusterResult,
}

@http(method: "PUT", uri: "/v1/clusters/{ClusterName}" )
@idempotent
operation UpdateCluster {
    input: UpdateClusterRequest,
    output: UpdateClusterResult,
}

@http(method: "DELETE", uri: "/v1/clusters/{ClusterName}" )
@idempotent
operation DeleteCluster {
    input: DeleteClusterRequest,
    output: DeleteClusterResult,
}

@http(method: "DELETE", uri: "/v1/clusters" )
@idempotent
operation DeleteClusters {
    input: DeleteClustersRequest,
    output: DeleteClustersResult,
}

@http(method: "GET", uri: "/v1/clusters/{ClusterName}/credentials" )
@readonly
operation GetClusterCredentials {
    input: GetClusterCredentialsRequest,
    output: GetClusterCredentialsResult,
}

@http(method: "GET", uri: "/v1/clusters/{ClusterName}/nodePools/{NodePoolName}" )
@readonly
operation GetNodePool {
    input: GetNodePoolRequest,
    output: GetNodePoolResult,
}

@http(method: "GET", uri: "/v1/clusters/{ClusterName}/nodePools" )
@readonly
operation ListNodePools {
    input: ListNodePoolsRequest,
    output: ListNodePoolsResult,
}

@http(method: "POST", uri: "/v1/clusters/{ClusterName}/nodePools" , code: 201)
operation CreateNodePool {
    input: CreateNodePoolRequest,
    output: CreateNodePoolResult,
}

@http(method: "PUT", uri: "/v1/clusters/{ClusterName}/nodePools/{NodePoolName}" )
@idempotent
operation UpdateNodePool {
    input: UpdateNodePoolRequest,
    output: UpdateNodePoolResult,
}

@http(method: "POST", uri: "/v1/clusters/{ClusterName}/nodePools/{NodePoolName}/:setSize" , code: 201)
operation SetNodePoolSize {
    input: SetNodePoolSizeRequest,
    output: SetNodePoolSizeResult,
}

@http(method: "DELETE", uri: "/v1/clusters/{ClusterName}/nodePools/{NodePoolName}" )
@idempotent
operation DeleteNodePool {
    input: DeleteNodePoolRequest,
    output: DeleteNodePoolResult,
}

@http(method: "DELETE", uri: "/v1/clusters/{ClusterName}/nodePools" )
@idempotent
operation DeleteNodePools {
    input: DeleteNodePoolsRequest,
    output: DeleteNodePoolsResult,
}

@http(method: "POST", uri: "/v1/clusters/{ClusterName}/nodePools/{NodePoolName}/nodes/{NodeName}/:reboot" , code: 201)
operation RebootNode {
    input: RebootNodeRequest,
    output: RebootNodeResult,
}

@http(method: "GET", uri: "/v1/serverConfig" )
@readonly
operation GetServerConfig {
    input: GetServerConfigRequest,
    output: GetServerConfigResult,
}

@http(method: "GET", uri: "/v1/firewallGroups/{FirewallGroupName}" )
@readonly
@waitable(
    FirewallRuleAuthorized: {
        acceptors: [
            {
                state: "success",
                matcher: {
                    output: {
                         path: "FirewallGroup.Rules[].Status",
                         comparator: "allStringEquals",
                         expected: "AUTHORIZED",
                    },
                },
            },
            {
                state: "success",
                matcher: {
                    output: {
                         path: "length(FirewallGroup.Rules[]) == `0`",
                         comparator: "booleanEquals",
                         expected: "true",
                    },
                },
            },
        ],
        minDelay: 20,
    },
)
operation GetFirewallGroup {
    input: GetFirewallGroupRequest,
    output: GetFirewallGroupResult,
}

@http(method: "GET", uri: "/v1/firewallGroups" )
@readonly
operation ListFirewallGroups {
    input: ListFirewallGroupsRequest,
    output: ListFirewallGroupsResult,
}

@http(method: "POST", uri: "/v1/firewallGroups" , code: 201)
operation CreateFirewallGroup {
    input: CreateFirewallGroupRequest,
    output: CreateFirewallGroupResult,
}

@http(method: "PUT", uri: "/v1/firewallGroups/{FirewallGroupName}" )
@idempotent
operation UpdateFirewallGroup {
    input: UpdateFirewallGroupRequest,
    output: UpdateFirewallGroupResult,
}

@http(method: "DELETE", uri: "/v1/firewallGroups/{FirewallGroupName}" )
@idempotent
operation DeleteFirewallGroup {
    input: DeleteFirewallGroupRequest,
    output: DeleteFirewallGroupResult,
}

@http(method: "DELETE", uri: "/v1/firewallGroups" )
@idempotent
operation DeleteFirewallGroups {
    input: DeleteFirewallGroupsRequest,
    output: DeleteFirewallGroupsResult,
}

@http(method: "POST", uri: "/v1/firewallGroups/{FirewallGroupName}/rules" , code: 201)
operation AuthorizeFirewallGroup {
    input: AuthorizeFirewallGroupRequest,
    output: AuthorizeFirewallGroupResult,
}

@http(method: "DELETE", uri: "/v1/firewallGroups/{FirewallGroupName}/rules" )
@idempotent
operation RevokeFirewallGroup {
    input: RevokeFirewallGroupRequest,
    output: RevokeFirewallGroupResult,
}

@http(method: "GET", uri: "/v1/disks/{DiskName}" )
@readonly
operation GetDisk {
    input: GetDiskRequest,
    output: GetDiskResult,
}

@http(method: "GET", uri: "/v1/disks" )
@readonly
operation ListDisks {
    input: ListDisksRequest,
    output: ListDisksResult,
}

@http(method: "POST", uri: "/v1/disks" , code: 201)
operation CreateDisk {
    input: CreateDiskRequest,
    output: CreateDiskResult,
}

@http(method: "PUT", uri: "/v1/disks/{DiskName}" )
@idempotent
operation UpdateDisk {
    input: UpdateDiskRequest,
    output: UpdateDiskResult,
}

@http(method: "DELETE", uri: "/v1/disks/{DiskName}" )
@idempotent
operation DeleteDisk {
    input: DeleteDiskRequest,
    output: DeleteDiskResult,
}

@http(method: "DELETE", uri: "/v1/disks" )
@idempotent
operation DeleteDisks {
    input: DeleteDisksRequest,
    output: DeleteDisksResult,
}

@http(method: "POST", uri: "/v1/disks/{DiskName}/:attach" , code: 201)
operation AttachDisk {
    input: AttachDiskRequest,
    output: AttachDiskResult,
}

@http(method: "POST", uri: "/v1/disks/{DiskName}/:detach" , code: 201)
operation DetachDisk {
    input: DetachDiskRequest,
    output: DetachDiskResult,
}

@http(method: "GET", uri: "/v1/snapshots/{SnapshotName}" )
@readonly
@waitable(
    SnapshotAvailable: {
        acceptors: [
            {
                state: "success",
                matcher: {
                    output: {
                         path: "Snapshot.Status",
                         comparator: "stringEquals",
                         expected: "AVAILABLE",
                    },
                },
            },
            {
                state: "failure",
                matcher: {
                    output: {
                         path: "Snapshot.Status",
                         comparator: "stringEquals",
                         expected: "FAILED",
                    },
                },
            },
        ],
        minDelay: 20,
    },
)
operation GetSnapshot {
    input: GetSnapshotRequest,
    output: GetSnapshotResult,
}

@http(method: "GET", uri: "/v1/snapshots" )
@readonly
operation ListSnapshots {
    input: ListSnapshotsRequest,
    output: ListSnapshotsResult,
}

@http(method: "POST", uri: "/v1/snapshots" , code: 201)
operation CreateSnapshot {
    input: CreateSnapshotRequest,
    output: CreateSnapshotResult,
}

@http(method: "PUT", uri: "/v1/snapshots/{SnapshotName}" )
@idempotent
operation UpdateSnapshot {
    input: UpdateSnapshotRequest,
    output: UpdateSnapshotResult,
}

@http(method: "DELETE", uri: "/v1/snapshots/{SnapshotName}" )
@idempotent
operation DeleteSnapshot {
    input: DeleteSnapshotRequest,
    output: DeleteSnapshotResult,
}

@http(method: "DELETE", uri: "/v1/snapshots" )
@idempotent
operation DeleteSnapshots {
    input: DeleteSnapshotsRequest,
    output: DeleteSnapshotsResult,
}

@http(method: "POST", uri: "/v1/snapshots/{SnapshotName}/:restore" , code: 201)
operation RestoreClusterFromSnapshot {
    input: RestoreClusterFromSnapshotRequest,
    output: RestoreClusterFromSnapshotResult,
}

@http(method: "GET", uri: "/v1/loadBalancers/{LoadBalancerName}" )
@readonly
operation GetLoadBalancer {
    input: GetLoadBalancerRequest,
    output: GetLoadBalancerResult,
}

@http(method: "GET", uri: "/v1/loadBalancers" )
@readonly
operation ListLoadBalancers {
    input: ListLoadBalancersRequest,
    output: ListLoadBalancersResult,
}

@http(method: "GET", uri: "/v1/tags" )
@readonly
operation ListTags {
    input: ListTagsRequest,
    output: ListTagsResult,
}

@http(method: "POST", uri: "/v1/tags" , code: 201)
operation CreateTags {
    input: CreateTagsRequest,
    output: CreateTagsResult,
}

@http(method: "PUT", uri: "/v1/tags" )
@idempotent
operation UpdateTags {
    input: UpdateTagsRequest,
    output: UpdateTagsResult,
}

@http(method: "DELETE", uri: "/v1/tags" )
@idempotent
operation DeleteTags {
    input: DeleteTagsRequest,
    output: DeleteTagsResult,
}
