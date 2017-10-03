package main

import (
	"github.com/HardySimpson/pilotage/pkg/session"
	"github.com/HardySimpson/pilotage/pkg/k8sclient"
)

func main() {
	kubecli := k8sclient.NewKubeClient()
	s := session.New(kubecli)
	// run shell in session
	s.Run()
}