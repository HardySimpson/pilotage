package main

import (
	"fmt"

	"github.com/HardySimpson/pilotage/pkg/version"
	"github.com/HardySimpson/pilotage/pkg/cmd"
	"github.com/HardySimpson/pilotage/pkg/k8sclient"
)

func main() {

	kubecli := k8sclient.NewKubeClient()


	shell := cmd.New(kubecli)

	// run shell
	shell.Run()

	fmt.Println("hello, world, ", version.GitSHA)
}