package cmd

import (
	"sync"
	"path"

	"github.com/abiosoft/ishell"
	"github.com/HardySimpson/pilotage/pkg/vfs"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"strings"
)


type Session struct {
	*ishell.Shell
	kubecli kubernetes.Interface
	path string
	vfs *vfs.VFS
	currentNode *vfs.Node
	lock sync.Mutex
}

func (s *Session) cd(c *ishell.Context) {
	args := c.Args

	if len(args) > 1 {
		c.Println("pilotage: too many arguments")
	} else if len(args) == 0 {
		s.path = "/"
		s.currentNode = s.vfs.RootNode
	} else if args[0] == ".." {
		if s.currentNode.Parent != nil {
			s.currentNode = s.currentNode.Parent
			s.path = path.Dir(s.path)
		}
	}

	n := s.currentNode.GetChild(args[0])
	if n != nil {
		s.currentNode = n
		s.path = path.Join(s.path, args[0])
	}

}

func (s *Session) pwd(c *ishell.Context) {
	s.Shell.Println(s.path)
}

func (s *Session) ls(c *ishell.Context) {
	/*
	if s.currentNode.Fresh != nil {
		s.currentNode.Fresh()
	}
	*/
	s.Shell.Println(strings.Join(s.currentNode.ListChilds(), ""))
}


func New(k kubernetes.Interface) *Session {
	s := &Session{
		Shell: ishell.New(),
		kubecli: k,
		path: "/",
		vfs: vfs.NewVFS(),
	}
	s.currentNode = s.vfs.RootNode

	s.AddCmd(&ishell.Cmd{
		Name: "cd",
		Help: "change to a path",
		Func: s.cd,
	})
	s.AddCmd(&ishell.Cmd{
		Name: "ls",
		Help: "change to a path",
		Func: s.ls,
	})
	s.AddCmd(&ishell.Cmd{
		Name: "pwd",
		Help: "change to a path",
		Func: s.pwd,
	})

	nsd := s.vfs.RootNode.AddChild("namespaces", nil, nil)

	nss, _ := s.kubecli.CoreV1().Namespaces().List(metav1.ListOptions{})
	for _, ns := range nss.Items {
		nsd.AddChild(ns.Name, ns, nil)
	}

	return s
}
