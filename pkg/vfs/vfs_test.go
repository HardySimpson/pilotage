package vfs

import (
	"fmt"
	"testing"
	"strings"

	"k8s.io/api/core/v1"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func TestNewVFS(t *testing.T) {
	vfs := NewVFS()

	nd := vfs.RootNode.AddChild(nil)

	fmt.Println(strings.Join(nd.ListChildrenName(), " "))
}

func TestCatNode(t *testing.T) {
	n := &Node{
		Obj: &v1.Pod {
			TypeMeta: metav1.TypeMeta{
				Kind: "aa",
			},
		},
	}
	fmt.Println(n.CatNode())
}
