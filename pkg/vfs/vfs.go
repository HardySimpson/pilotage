package vfs

import "path"

type Node struct {
	Name    string
	Children map[string]*Node
	Path 	string
	Parent  *Node
	Obj 	interface{}
	FreshFunc   FreshChildFunc
}

type FreshChildFunc func(PrevObj interface{}) ([]*Node, interface{}, error)


func (n *Node) AddChild(child *Node) *Node {
	child.Parent = n
	child.Path = path.Join(n.Path, child.Name)
	if n.Children == nil {
		n.Children = make(map[string]*Node)
	}
	n.Children[child.Name] = child
	return child
}

func (n *Node) AddChildren(children []*Node) {
	for _, child := range children {
		n.AddChild(child)
	}
}

func (n *Node) FreshChildren() *Node{
	if n.FreshFunc == nil {
		return n
	}
	children, obj, err := n.FreshFunc(n.Obj)
	if err != nil {
		// log error
	}
	n.Children = nil
	n.AddChildren(children)
	n.Obj = obj
	return n
}

func (n *Node) GetChild(name string) *Node {
	return n.Children[name]
}

func (n *Node) ListChildrenName() []string {
	l := []string{}

	for k, _ := range n.Children {
		l = append(l, k)
	}
	return l
}

type VFS struct {
	RootNode *Node
}

func NewVFS() *VFS {
	vfs := &VFS{
		RootNode: &Node{
			Name: "",
			Children: make(map[string]*Node),
			Parent: nil,
		},
	}
	return vfs
}




