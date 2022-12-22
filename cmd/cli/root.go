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
			util.CheckErr(viper.BindPFlags(cmd.Flags()))
		},
		Run: func(cmd *cobra.Command, args []string) {
			util.CheckErr(setNamespace(f))
			util.CheckErr(app.Run(f, args))
		},
		Version: "0.0.1",
	}

	configFlags.AddFlags(cmd.Flags())
	resourceBuilderFlags.AddFlags(cmd.Flags())

	return cmd
}

func setNamespace(f util.Factory) error {
	if viper.GetBool("all-namespaces") {
		viper.Set("namespace", "")
		return nil
	}
	namespace, _, err := f.ToRawKubeConfigLoader().Namespace()
	if err != nil {
		return fmt.Errorf("can't determine namespace")
	}
	viper.Set("namespace", namespace)
	return nil
}
