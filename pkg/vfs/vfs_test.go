package vfs

import (
	"fmt"
	"testing"
	"strings"
)

func TestNewVFS(t *testing.T) {
	vfs := NewVFS()

	nd := vfs.RootNode.AddChild("namespaces", nil, nil)

	nd.AddChild("aa", nil, nil)
	nd.AddChild("bb", nil, nil)

	fmt.Println(strings.Join(nd.ListChildrenName(), " "))
}
