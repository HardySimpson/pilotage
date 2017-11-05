package session

import (
	"github.com/HardySimpson/pilotage/pkg/vfs"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

// TODO: flat the code here

func addDefaultTree(node *vfs.Node, cli kubernetes.Interface) {
	node.AddChild(&vfs.Node{
		Name: "namespaces",
		FreshFunc: func(prevObj interface{}) ([]*vfs.Node, interface{}, error) {
			mp := []*vfs.Node{}
			nsList, err := cli.CoreV1().Namespaces().List(metav1.ListOptions{})
			if err != nil {
				return nil, nil, err
			}
			for _, nsObj := range nsList.Items {
				nd := &vfs.Node{Name:nsObj.Name}
				nd.FreshFunc = func(prevObj interface{}) ([]*vfs.Node, interface{}, error) {
					return []*vfs.Node{
						&vfs.Node{
							Name: "pods",
							FreshFunc: func(prevObj interface{}) ([]*vfs.Node, interface{}, error) {
								mp := []*vfs.Node{}
								podList, err := cli.CoreV1().Pods(nd.Name).List(metav1.ListOptions{})
								if err != nil {
									return nil, nil, err
								}
								for _, podObj := range podList.Items {
									mp = append(mp, &vfs.Node{
										Name: podObj.Name,
										Obj: podObj,
									})
								}
								return mp, podList, nil
							},

						},
						&vfs.Node{
							Name: "configmaps",
							FreshFunc: func(prevObj interface{}) ([]*vfs.Node, interface{}, error) {
								mp := []*vfs.Node{}
								list, err := cli.CoreV1().ConfigMaps(nd.Name).List(metav1.ListOptions{})
								if err != nil {
									return nil, nil, err
								}
								for _, item := range list.Items {
									mp = append(mp, &vfs.Node{
										Name: item.Name,
										Obj: item,
									})
								}
								return mp, list, nil
							},
						},
						&vfs.Node{
							Name: "endpoints",
							FreshFunc: func(prevObj interface{}) ([]*vfs.Node, interface{}, error) {
								mp := []*vfs.Node{}
								list, err := cli.CoreV1().Endpoints(nd.Name).List(metav1.ListOptions{})
								if err != nil {
									return nil, nil, err
								}
								for _, item := range list.Items {
									mp = append(mp, &vfs.Node{
										Name:item.Name,
										Obj: item,
									})
								}
								return mp, list, nil
							},
						},
						&vfs.Node{
							Name: "services",
							FreshFunc: func(prevObj interface{}) ([]*vfs.Node, interface{}, error) {
								mp := []*vfs.Node{}
								list, err := cli.CoreV1().Services(nd.Name).List(metav1.ListOptions{})
								if err != nil {
									return nil, nil, err
								}
								for _, item := range list.Items {
									mp = append(mp, &vfs.Node{
										Name:item.Name,
										Obj: item,
									})
								}
								return mp, list, nil
							},
						},
						&vfs.Node{
							Name: "secrets",
							FreshFunc: func(prevObj interface{}) ([]*vfs.Node, interface{}, error) {
								mp := []*vfs.Node{}
								list, err := cli.CoreV1().Secrets(nd.Name).List(metav1.ListOptions{})
								if err != nil {
									return nil, nil, err
								}
								for _, item := range list.Items {
									mp = append(mp, &vfs.Node{
										Name:item.Name,
										Obj: item,
									})
								}
								return mp, list, nil
							},
						},
					}, nd, nil
				}
				mp = append(mp, nd)
			}
			return mp, nsList, nil
		},
	}).FreshChildren()

}
