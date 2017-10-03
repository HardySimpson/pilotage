package vfs

type Node struct {
	Name    string
	Children map[string]*Node
	Parent  *Node
	Obj 	interface{}
	Fresh  FreshFunc
}

type FreshFunc func()

func (n *Node) AddChild(name string, obj interface{}, fresh FreshFunc) *Node {
	nd := &Node {
		Name: name,
		Children: make(map[string]*Node),
		Parent: n,
		Obj: obj,
	}
	n.Children[name] = nd
	return nd
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




