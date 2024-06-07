package util

import (
	"context"

	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

const OVN_INTERCONNECT_CONFIGMAP_NAME = "ovn-interconnect-configuration"
const OVN_NAMESPACE = "openshift-ovn-kubernetes"
const OVN_MASTER = "ovnkube-master"
const OVN_CONTROL_PLANE = "ovnkube-control-plane"
const OVN_NODE = "ovnkube-node"
const OVN_CONTROLLER = "ovnkube-controller"
const OVN_IPSEC = "ovn-ipsec"                             // 4.13 ipsec daemonset
const OVN_IPSEC_HOST = "ovn-ipsec-host"                   // 4.14 ipsec daemonset
const OVN_IPSEC_CONTAINERIZED = "ovn-ipsec-containerized" // 4.14 ipsec daemonset

func GetInterConnectConfigMap(kubeClient kubernetes.Interface) (*corev1.ConfigMap, error) {
	return kubeClient.CoreV1().ConfigMaps(OVN_NAMESPACE).Get(context.TODO(), OVN_INTERCONNECT_CONFIGMAP_NAME, metav1.GetOptions{})
}
