package schedulerplugin

import (
	crd_clientset "git.code.oa.com/tkestack/galaxy/pkg/ipam/client/clientset/versioned"
	list "git.code.oa.com/tkestack/galaxy/pkg/ipam/client/listers/galaxy/v1alpha1"
	extensionClient "k8s.io/apiextensions-apiserver/pkg/client/clientset/clientset"
	"k8s.io/client-go/kubernetes"
	appv1 "k8s.io/client-go/listers/apps/v1"
	corev1lister "k8s.io/client-go/listers/core/v1"
	"tkestack.io/tapp-controller/pkg/client/clientset/versioned"
	"tkestack.io/tapp-controller/pkg/client/listers/tappcontroller/v1"
)

type PluginFactoryArgs struct {
	Client            kubernetes.Interface
	TAppClient        versioned.Interface
	PodLister         corev1lister.PodLister
	StatefulSetLister appv1.StatefulSetLister
	DeploymentLister  appv1.DeploymentLister
	TAppLister        v1.TAppLister
	PoolLister        list.PoolLister
	PodHasSynced      func() bool
	StatefulSetSynced func() bool
	DeploymentSynced  func() bool
	TAppHasSynced     func() bool
	PoolSynced        func() bool
	CrdClient         crd_clientset.Interface
	ExtClient         extensionClient.Interface
}
