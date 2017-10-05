package session

import (
	"github.com/HardySimpson/pilotage/pkg/vfs"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)


func addDefaultTree(node *vfs.Node, cli kubernetes.Interface) {

	nsDirNode := node.AddChild(&vfs.Node{
		Name: "namespaces",
		FreshFunc: func(prevObj interface{}) ([]*vfs.Node, interface{}, error) {
			mp := []*vfs.Node{}
			nsList, err := cli.CoreV1().Namespaces().List(metav1.ListOptions{})
			if err != nil {
				return nil, nil, err
			}
			for _, nsObj := range nsList.Items {
				mp = append(mp, &vfs.Node{Name:nsObj.Name})
			}
			return mp, nsList, nil
		},
	}).FreshChildren()

	for _, nsNode := range nsDirNode.Children {
		nsNode.AddChild(&vfs.Node{
			Name: "pods",
			FreshFunc: func(prevObj interface{}) ([]*vfs.Node, interface{}, error) {
				mp := []*vfs.Node{}
				podList, err := cli.CoreV1().Pods(nsNode.Name).List(metav1.ListOptions{})
				if err != nil {
					return nil, nil, err
				}
				for _, podObj := range podList.Items {
					mp = append(mp, &vfs.Node{Name:podObj.Name})
				}
				return mp, podList, nil
			},

		}).FreshChildren()

		nsNode.AddChild(&vfs.Node{
			Name: "configmaps",
			FreshFunc: func(prevObj interface{}) ([]*vfs.Node, interface{}, error) {
				mp := []*vfs.Node{}
				list, err := cli.CoreV1().ConfigMaps(nsNode.Name).List(metav1.ListOptions{})
				if err != nil {
					return nil, nil, err
				}
				for _, item := range list.Items {
					mp = append(mp, &vfs.Node{Name:item.Name})
				}
				return mp, list, nil
			},

		}).FreshChildren()


		nsNode.AddChild(&vfs.Node{
			Name: "endpoints",
			FreshFunc: func(prevObj interface{}) ([]*vfs.Node, interface{}, error) {
				mp := []*vfs.Node{}
				list, err := cli.CoreV1().Endpoints(nsNode.Name).List(metav1.ListOptions{})
				if err != nil {
					return nil, nil, err
				}
				for _, item := range list.Items {
					mp = append(mp, &vfs.Node{Name:item.Name})
				}
				return mp, list, nil
			},

		}).FreshChildren()

		nsNode.AddChild(&vfs.Node{
			Name: "services",
			FreshFunc: func(prevObj interface{}) ([]*vfs.Node, interface{}, error) {
				mp := []*vfs.Node{}
				list, err := cli.CoreV1().Services(nsNode.Name).List(metav1.ListOptions{})
				if err != nil {
					return nil, nil, err
				}
				for _, item := range list.Items {
					mp = append(mp, &vfs.Node{Name:item.Name})
				}
				return mp, list, nil
			},

		}).FreshChildren()

		nsNode.AddChild(&vfs.Node{
			Name: "secrets",
			FreshFunc: func(prevObj interface{}) ([]*vfs.Node, interface{}, error) {
				mp := []*vfs.Node{}
				list, err := cli.CoreV1().Secrets(nsNode.Name).List(metav1.ListOptions{})
				if err != nil {
					return nil, nil, err
				}
				for _, item := range list.Items {
					mp = append(mp, &vfs.Node{Name:item.Name})
				}
				return mp, list, nil
			},

		}).FreshChildren()


		/*
		for _, nsObj := range nsList.Items {
			nsNode := nsDirNode.AddChild(nsObj.Name, nsObj, nil)

			podDirNode := nsNode.AddChild("pods", nil, nil)
			podList, _ := cli.CoreV1().Pods(nsObj.Name).List(metav1.ListOptions{})
			for _, podObj := range podList.Items {

				podDirNode.AddChild(podObj.Name, podObj, nil)
			}

			deployDirNode := nsNode.AddChild("deployments", nil, nil)
			deployList, _ := cli.Extensions().Deployments(nsObj.Name).List(metav1.ListOptions{})
			for _, deployObj := range deployList.Items {
				deployDirNode.AddChild(deployObj.Name, deployObj, nil)
			}

		}
		*/
	}
}
