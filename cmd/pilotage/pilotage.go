package main

import (
	"fmt"

	"github.com/HardySimpson/pilotage/pkg/version"
)

func main() {
	fmt.Println("hello, world, ", version.GitSHA)
}