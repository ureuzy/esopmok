package main

import (
	"fmt"
	"os"

	_ "k8s.io/client-go/plugin/pkg/client/auth"

	"github.com/ureuzy/esopmok/cmd/cli"
)

func main() {
	if err := cli.RootCmd().Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
