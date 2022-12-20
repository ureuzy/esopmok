package main

import (
	"fmt"
	"os"

	"github.com/ureuzy/esopmok/cmd/cli"
	_ "k8s.io/client-go/plugin/pkg/client/auth"
)

func main() {
	if err := cli.RootCmd().Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
