package main

import (
	"github.com/HardySimpson/pilotage/pkg/session"
	"github.com/HardySimpson/pilotage/pkg/k8sclient"
)

func main() {
	kubecli, config := k8sclient.NewKubeClient()
	s := session.New(kubecli, config)
	// run shell in session
	s.Run()
}