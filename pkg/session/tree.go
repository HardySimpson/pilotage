package session

import (
	"github.com/HardySimpson/pilotage/pkg/vfs"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

func addDefaultTree(node *vfs.Node, cli kubernetes.Interface) {
	nsd := node.AddChild("namespaces", nil, nil)

	nss, _ := cli.CoreV1().Namespaces().List(metav1.ListOptions{})
	for _, ns := range nss.Items {
		nsd.AddChild(ns.Name, ns, nil)
	}

}
