// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package hatoba

type DirectionOfrulesForAuthorizeFirewallGroup string

// Enum values for DirectionOfrulesForAuthorizeFirewallGroup
const (
	DirectionOfrulesForAuthorizeFirewallGroupIn  DirectionOfrulesForAuthorizeFirewallGroup = "IN"
	DirectionOfrulesForAuthorizeFirewallGroupOut DirectionOfrulesForAuthorizeFirewallGroup = "OUT"
)

func (enum DirectionOfrulesForAuthorizeFirewallGroup) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum DirectionOfrulesForAuthorizeFirewallGroup) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type InstanceTypeOfclusterForCreateCluster string

// Enum values for InstanceTypeOfclusterForCreateCluster
const (
	InstanceTypeOfclusterForCreateClusterCMedium        InstanceTypeOfclusterForCreateCluster = "c-medium"
	InstanceTypeOfclusterForCreateClusterEMedium        InstanceTypeOfclusterForCreateCluster = "e-medium"
	InstanceTypeOfclusterForCreateClusterMedium         InstanceTypeOfclusterForCreateCluster = "medium"
	InstanceTypeOfclusterForCreateClusterCMedium4       InstanceTypeOfclusterForCreateCluster = "c-medium4"
	InstanceTypeOfclusterForCreateClusterEMedium4       InstanceTypeOfclusterForCreateCluster = "e-medium4"
	InstanceTypeOfclusterForCreateClusterMedium4        InstanceTypeOfclusterForCreateCluster = "medium4"
	InstanceTypeOfclusterForCreateClusterCMedium8       InstanceTypeOfclusterForCreateCluster = "c-medium8"
	InstanceTypeOfclusterForCreateClusterEMedium8       InstanceTypeOfclusterForCreateCluster = "e-medium8"
	InstanceTypeOfclusterForCreateClusterMedium8        InstanceTypeOfclusterForCreateCluster = "medium8"
	InstanceTypeOfclusterForCreateClusterEMedium16      InstanceTypeOfclusterForCreateCluster = "e-medium16"
	InstanceTypeOfclusterForCreateClusterMedium16       InstanceTypeOfclusterForCreateCluster = "medium16"
	InstanceTypeOfclusterForCreateClusterEMedium24      InstanceTypeOfclusterForCreateCluster = "e-medium24"
	InstanceTypeOfclusterForCreateClusterMedium24       InstanceTypeOfclusterForCreateCluster = "medium24"
	InstanceTypeOfclusterForCreateClusterCLarge         InstanceTypeOfclusterForCreateCluster = "c-large"
	InstanceTypeOfclusterForCreateClusterELarge         InstanceTypeOfclusterForCreateCluster = "e-large"
	InstanceTypeOfclusterForCreateClusterLarge          InstanceTypeOfclusterForCreateCluster = "large"
	InstanceTypeOfclusterForCreateClusterCLarge8        InstanceTypeOfclusterForCreateCluster = "c-large8"
	InstanceTypeOfclusterForCreateClusterELarge8        InstanceTypeOfclusterForCreateCluster = "e-large8"
	InstanceTypeOfclusterForCreateClusterLarge8         InstanceTypeOfclusterForCreateCluster = "large8"
	InstanceTypeOfclusterForCreateClusterELarge16       InstanceTypeOfclusterForCreateCluster = "e-large16"
	InstanceTypeOfclusterForCreateClusterLarge16        InstanceTypeOfclusterForCreateCluster = "large16"
	InstanceTypeOfclusterForCreateClusterELarge24       InstanceTypeOfclusterForCreateCluster = "e-large24"
	InstanceTypeOfclusterForCreateClusterLarge24        InstanceTypeOfclusterForCreateCluster = "large24"
	InstanceTypeOfclusterForCreateClusterELarge32       InstanceTypeOfclusterForCreateCluster = "e-large32"
	InstanceTypeOfclusterForCreateClusterLarge32        InstanceTypeOfclusterForCreateCluster = "large32"
	InstanceTypeOfclusterForCreateClusterEExtraLarge8   InstanceTypeOfclusterForCreateCluster = "e-extra-large8"
	InstanceTypeOfclusterForCreateClusterExtraLarge8    InstanceTypeOfclusterForCreateCluster = "extra-large8"
	InstanceTypeOfclusterForCreateClusterEExtraLarge16  InstanceTypeOfclusterForCreateCluster = "e-extra-large16"
	InstanceTypeOfclusterForCreateClusterExtraLarge16   InstanceTypeOfclusterForCreateCluster = "extra-large16"
	InstanceTypeOfclusterForCreateClusterEExtraLarge24  InstanceTypeOfclusterForCreateCluster = "e-extra-large24"
	InstanceTypeOfclusterForCreateClusterExtraLarge24   InstanceTypeOfclusterForCreateCluster = "extra-large24"
	InstanceTypeOfclusterForCreateClusterEExtraLarge32  InstanceTypeOfclusterForCreateCluster = "e-extra-large32"
	InstanceTypeOfclusterForCreateClusterExtraLarge32   InstanceTypeOfclusterForCreateCluster = "extra-large32"
	InstanceTypeOfclusterForCreateClusterEExtraLarge48  InstanceTypeOfclusterForCreateCluster = "e-extra-large48"
	InstanceTypeOfclusterForCreateClusterExtraLarge48   InstanceTypeOfclusterForCreateCluster = "extra-large48"
	InstanceTypeOfclusterForCreateClusterEDoubleLarge16 InstanceTypeOfclusterForCreateCluster = "e-double-large16"
	InstanceTypeOfclusterForCreateClusterDoubleLarge16  InstanceTypeOfclusterForCreateCluster = "double-large16"
	InstanceTypeOfclusterForCreateClusterEDoubleLarge24 InstanceTypeOfclusterForCreateCluster = "e-double-large24"
	InstanceTypeOfclusterForCreateClusterDoubleLarge24  InstanceTypeOfclusterForCreateCluster = "double-large24"
	InstanceTypeOfclusterForCreateClusterEDoubleLarge32 InstanceTypeOfclusterForCreateCluster = "e-double-large32"
	InstanceTypeOfclusterForCreateClusterDoubleLarge32  InstanceTypeOfclusterForCreateCluster = "double-large32"
	InstanceTypeOfclusterForCreateClusterEDoubleLarge48 InstanceTypeOfclusterForCreateCluster = "e-double-large48"
	InstanceTypeOfclusterForCreateClusterDoubleLarge48  InstanceTypeOfclusterForCreateCluster = "double-large48"
	InstanceTypeOfclusterForCreateClusterEDoubleLarge64 InstanceTypeOfclusterForCreateCluster = "e-double-large64"
	InstanceTypeOfclusterForCreateClusterDoubleLarge64  InstanceTypeOfclusterForCreateCluster = "double-large64"
	InstanceTypeOfclusterForCreateClusterEDoubleLarge96 InstanceTypeOfclusterForCreateCluster = "e-double-large96"
	InstanceTypeOfclusterForCreateClusterDoubleLarge96  InstanceTypeOfclusterForCreateCluster = "double-large96"
)

func (enum InstanceTypeOfclusterForCreateCluster) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum InstanceTypeOfclusterForCreateCluster) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type InstanceTypeOfnodePoolForCreateNodePool string

// Enum values for InstanceTypeOfnodePoolForCreateNodePool
const (
	InstanceTypeOfnodePoolForCreateNodePoolCMedium        InstanceTypeOfnodePoolForCreateNodePool = "c-medium"
	InstanceTypeOfnodePoolForCreateNodePoolEMedium        InstanceTypeOfnodePoolForCreateNodePool = "e-medium"
	InstanceTypeOfnodePoolForCreateNodePoolMedium         InstanceTypeOfnodePoolForCreateNodePool = "medium"
	InstanceTypeOfnodePoolForCreateNodePoolCMedium4       InstanceTypeOfnodePoolForCreateNodePool = "c-medium4"
	InstanceTypeOfnodePoolForCreateNodePoolEMedium4       InstanceTypeOfnodePoolForCreateNodePool = "e-medium4"
	InstanceTypeOfnodePoolForCreateNodePoolMedium4        InstanceTypeOfnodePoolForCreateNodePool = "medium4"
	InstanceTypeOfnodePoolForCreateNodePoolCMedium8       InstanceTypeOfnodePoolForCreateNodePool = "c-medium8"
	InstanceTypeOfnodePoolForCreateNodePoolEMedium8       InstanceTypeOfnodePoolForCreateNodePool = "e-medium8"
	InstanceTypeOfnodePoolForCreateNodePoolMedium8        InstanceTypeOfnodePoolForCreateNodePool = "medium8"
	InstanceTypeOfnodePoolForCreateNodePoolEMedium16      InstanceTypeOfnodePoolForCreateNodePool = "e-medium16"
	InstanceTypeOfnodePoolForCreateNodePoolMedium16       InstanceTypeOfnodePoolForCreateNodePool = "medium16"
	InstanceTypeOfnodePoolForCreateNodePoolEMedium24      InstanceTypeOfnodePoolForCreateNodePool = "e-medium24"
	InstanceTypeOfnodePoolForCreateNodePoolMedium24       InstanceTypeOfnodePoolForCreateNodePool = "medium24"
	InstanceTypeOfnodePoolForCreateNodePoolCLarge         InstanceTypeOfnodePoolForCreateNodePool = "c-large"
	InstanceTypeOfnodePoolForCreateNodePoolELarge         InstanceTypeOfnodePoolForCreateNodePool = "e-large"
	InstanceTypeOfnodePoolForCreateNodePoolLarge          InstanceTypeOfnodePoolForCreateNodePool = "large"
	InstanceTypeOfnodePoolForCreateNodePoolCLarge8        InstanceTypeOfnodePoolForCreateNodePool = "c-large8"
	InstanceTypeOfnodePoolForCreateNodePoolELarge8        InstanceTypeOfnodePoolForCreateNodePool = "e-large8"
	InstanceTypeOfnodePoolForCreateNodePoolLarge8         InstanceTypeOfnodePoolForCreateNodePool = "large8"
	InstanceTypeOfnodePoolForCreateNodePoolELarge16       InstanceTypeOfnodePoolForCreateNodePool = "e-large16"
	InstanceTypeOfnodePoolForCreateNodePoolLarge16        InstanceTypeOfnodePoolForCreateNodePool = "large16"
	InstanceTypeOfnodePoolForCreateNodePoolELarge24       InstanceTypeOfnodePoolForCreateNodePool = "e-large24"
	InstanceTypeOfnodePoolForCreateNodePoolLarge24        InstanceTypeOfnodePoolForCreateNodePool = "large24"
	InstanceTypeOfnodePoolForCreateNodePoolELarge32       InstanceTypeOfnodePoolForCreateNodePool = "e-large32"
	InstanceTypeOfnodePoolForCreateNodePoolLarge32        InstanceTypeOfnodePoolForCreateNodePool = "large32"
	InstanceTypeOfnodePoolForCreateNodePoolEExtraLarge8   InstanceTypeOfnodePoolForCreateNodePool = "e-extra-large8"
	InstanceTypeOfnodePoolForCreateNodePoolExtraLarge8    InstanceTypeOfnodePoolForCreateNodePool = "extra-large8"
	InstanceTypeOfnodePoolForCreateNodePoolEExtraLarge16  InstanceTypeOfnodePoolForCreateNodePool = "e-extra-large16"
	InstanceTypeOfnodePoolForCreateNodePoolExtraLarge16   InstanceTypeOfnodePoolForCreateNodePool = "extra-large16"
	InstanceTypeOfnodePoolForCreateNodePoolEExtraLarge24  InstanceTypeOfnodePoolForCreateNodePool = "e-extra-large24"
	InstanceTypeOfnodePoolForCreateNodePoolExtraLarge24   InstanceTypeOfnodePoolForCreateNodePool = "extra-large24"
	InstanceTypeOfnodePoolForCreateNodePoolEExtraLarge32  InstanceTypeOfnodePoolForCreateNodePool = "e-extra-large32"
	InstanceTypeOfnodePoolForCreateNodePoolExtraLarge32   InstanceTypeOfnodePoolForCreateNodePool = "extra-large32"
	InstanceTypeOfnodePoolForCreateNodePoolEExtraLarge48  InstanceTypeOfnodePoolForCreateNodePool = "e-extra-large48"
	InstanceTypeOfnodePoolForCreateNodePoolExtraLarge48   InstanceTypeOfnodePoolForCreateNodePool = "extra-large48"
	InstanceTypeOfnodePoolForCreateNodePoolEDoubleLarge16 InstanceTypeOfnodePoolForCreateNodePool = "e-double-large16"
	InstanceTypeOfnodePoolForCreateNodePoolDoubleLarge16  InstanceTypeOfnodePoolForCreateNodePool = "double-large16"
	InstanceTypeOfnodePoolForCreateNodePoolEDoubleLarge24 InstanceTypeOfnodePoolForCreateNodePool = "e-double-large24"
	InstanceTypeOfnodePoolForCreateNodePoolDoubleLarge24  InstanceTypeOfnodePoolForCreateNodePool = "double-large24"
	InstanceTypeOfnodePoolForCreateNodePoolEDoubleLarge32 InstanceTypeOfnodePoolForCreateNodePool = "e-double-large32"
	InstanceTypeOfnodePoolForCreateNodePoolDoubleLarge32  InstanceTypeOfnodePoolForCreateNodePool = "double-large32"
	InstanceTypeOfnodePoolForCreateNodePoolEDoubleLarge48 InstanceTypeOfnodePoolForCreateNodePool = "e-double-large48"
	InstanceTypeOfnodePoolForCreateNodePoolDoubleLarge48  InstanceTypeOfnodePoolForCreateNodePool = "double-large48"
	InstanceTypeOfnodePoolForCreateNodePoolEDoubleLarge64 InstanceTypeOfnodePoolForCreateNodePool = "e-double-large64"
	InstanceTypeOfnodePoolForCreateNodePoolDoubleLarge64  InstanceTypeOfnodePoolForCreateNodePool = "double-large64"
	InstanceTypeOfnodePoolForCreateNodePoolEDoubleLarge96 InstanceTypeOfnodePoolForCreateNodePool = "e-double-large96"
	InstanceTypeOfnodePoolForCreateNodePoolDoubleLarge96  InstanceTypeOfnodePoolForCreateNodePool = "double-large96"
)

func (enum InstanceTypeOfnodePoolForCreateNodePool) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum InstanceTypeOfnodePoolForCreateNodePool) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type KubernetesVersionOfclusterForCreateCluster string

// Enum values for KubernetesVersionOfclusterForCreateCluster
const (
	KubernetesVersionOfclusterForCreateClusterV1196  KubernetesVersionOfclusterForCreateCluster = "v1.19.6"
	KubernetesVersionOfclusterForCreateClusterV11911 KubernetesVersionOfclusterForCreateCluster = "v1.19.11"
	KubernetesVersionOfclusterForCreateClusterV11915 KubernetesVersionOfclusterForCreateCluster = "v1.19.15"
	KubernetesVersionOfclusterForCreateClusterV1201  KubernetesVersionOfclusterForCreateCluster = "v1.20.1"
	KubernetesVersionOfclusterForCreateClusterV1207  KubernetesVersionOfclusterForCreateCluster = "v1.20.7"
	KubernetesVersionOfclusterForCreateClusterV12011 KubernetesVersionOfclusterForCreateCluster = "v1.20.11"
	KubernetesVersionOfclusterForCreateClusterV1211  KubernetesVersionOfclusterForCreateCluster = "v1.21.1"
	KubernetesVersionOfclusterForCreateClusterV1215  KubernetesVersionOfclusterForCreateCluster = "v1.21.5"
	KubernetesVersionOfclusterForCreateClusterV1222  KubernetesVersionOfclusterForCreateCluster = "v1.22.2"
)

func (enum KubernetesVersionOfclusterForCreateCluster) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum KubernetesVersionOfclusterForCreateCluster) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type KubernetesVersionOfclusterForUpdateCluster string

// Enum values for KubernetesVersionOfclusterForUpdateCluster
const (
	KubernetesVersionOfclusterForUpdateClusterV1196  KubernetesVersionOfclusterForUpdateCluster = "v1.19.6"
	KubernetesVersionOfclusterForUpdateClusterV11911 KubernetesVersionOfclusterForUpdateCluster = "v1.19.11"
	KubernetesVersionOfclusterForUpdateClusterV11915 KubernetesVersionOfclusterForUpdateCluster = "v1.19.15"
	KubernetesVersionOfclusterForUpdateClusterV1201  KubernetesVersionOfclusterForUpdateCluster = "v1.20.1"
	KubernetesVersionOfclusterForUpdateClusterV1207  KubernetesVersionOfclusterForUpdateCluster = "v1.20.7"
	KubernetesVersionOfclusterForUpdateClusterV12011 KubernetesVersionOfclusterForUpdateCluster = "v1.20.11"
	KubernetesVersionOfclusterForUpdateClusterV1211  KubernetesVersionOfclusterForUpdateCluster = "v1.21.1"
	KubernetesVersionOfclusterForUpdateClusterV1215  KubernetesVersionOfclusterForUpdateCluster = "v1.21.5"
	KubernetesVersionOfclusterForUpdateClusterV1222  KubernetesVersionOfclusterForUpdateCluster = "v1.22.2"
)

func (enum KubernetesVersionOfclusterForUpdateCluster) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum KubernetesVersionOfclusterForUpdateCluster) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type ProtocolOfrulesForAuthorizeFirewallGroup string

// Enum values for ProtocolOfrulesForAuthorizeFirewallGroup
const (
	ProtocolOfrulesForAuthorizeFirewallGroupAny   ProtocolOfrulesForAuthorizeFirewallGroup = "ANY"
	ProtocolOfrulesForAuthorizeFirewallGroupTcp   ProtocolOfrulesForAuthorizeFirewallGroup = "TCP"
	ProtocolOfrulesForAuthorizeFirewallGroupUdp   ProtocolOfrulesForAuthorizeFirewallGroup = "UDP"
	ProtocolOfrulesForAuthorizeFirewallGroupIcmp  ProtocolOfrulesForAuthorizeFirewallGroup = "ICMP"
	ProtocolOfrulesForAuthorizeFirewallGroupSsh   ProtocolOfrulesForAuthorizeFirewallGroup = "SSH"
	ProtocolOfrulesForAuthorizeFirewallGroupHttp  ProtocolOfrulesForAuthorizeFirewallGroup = "HTTP"
	ProtocolOfrulesForAuthorizeFirewallGroupHttps ProtocolOfrulesForAuthorizeFirewallGroup = "HTTPS"
	ProtocolOfrulesForAuthorizeFirewallGroupRdp   ProtocolOfrulesForAuthorizeFirewallGroup = "RDP"
	ProtocolOfrulesForAuthorizeFirewallGroupGre   ProtocolOfrulesForAuthorizeFirewallGroup = "GRE"
	ProtocolOfrulesForAuthorizeFirewallGroupEsp   ProtocolOfrulesForAuthorizeFirewallGroup = "ESP"
	ProtocolOfrulesForAuthorizeFirewallGroupAh    ProtocolOfrulesForAuthorizeFirewallGroup = "AH"
	ProtocolOfrulesForAuthorizeFirewallGroupVrrp  ProtocolOfrulesForAuthorizeFirewallGroup = "VRRP"
	ProtocolOfrulesForAuthorizeFirewallGroupL2tp  ProtocolOfrulesForAuthorizeFirewallGroup = "L2TP"
)

func (enum ProtocolOfrulesForAuthorizeFirewallGroup) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum ProtocolOfrulesForAuthorizeFirewallGroup) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}
