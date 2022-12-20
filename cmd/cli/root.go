package cli

import (
	"fmt"
	"github.com/ureuzy/esopmok/pkg/app"
	"k8s.io/cli-runtime/pkg/genericclioptions"
	"k8s.io/kubectl/pkg/cmd/util"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func RootCmd() *cobra.Command {

	configFlags := genericclioptions.NewConfigFlags(false).
		WithDeprecatedPasswordFlag().
		WithDiscoveryBurst(300).
		WithDiscoveryQPS(50.0)

	resourceBuilderFlags := genericclioptions.NewResourceBuilderFlags().
		WithAll(false).
		WithAllNamespaces(false).
		WithFile(false).
		WithLabelSelector("").
		WithFieldSelector("").
		WithLatest()

	f := util.NewFactory(configFlags)

	cmd := &cobra.Command{
		Use:     "kubectl-esopmok",
		Short:   "sample short",
		Long:    "sample long",
		Example: "kubectl esopmok deploy [deployment name]",
		PreRun: func(cmd *cobra.Command, args []string) {
			if err := viper.BindPFlags(cmd.Flags()); err != nil {
				fmt.Println(err)
			}
		},
		Run: func(cmd *cobra.Command, args []string) {
			esopmok, err := app.NewEsopmok(f, args)
			if err != nil {
				fmt.Println(err)
			}
			if err = esopmok.Run(); err != nil {
				fmt.Println(err)
			}
		},
		Version: "0.0.1",
	}

	configFlags.AddFlags(cmd.Flags())
	resourceBuilderFlags.AddFlags(cmd.Flags())

	return cmd
}
