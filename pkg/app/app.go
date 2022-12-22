package app

import (
	"context"
	"fmt"
	"github.com/ureuzy/esopmok/pkg/consts"
	"gopkg.in/yaml.v3"

	composeSpec "github.com/compose-spec/compose-go/types"
	"github.com/spf13/viper"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/kubectl/pkg/cmd/util"
)

type Esopmok struct {
	kind    string
	name    string
	client  *kubernetes.Clientset
	compose *Compose
}

func Run(f util.Factory, args []string) error {
	kind, name, err := figureOutArgs(args)
	if err != nil {
		return err
	}

	clientSet, err := f.KubernetesClientSet()
	if err != nil {
		return err
	}

	esopmpk := Esopmok{
		kind,
		name,
		clientSet,
		&Compose{&composeSpec.Config{}},
	}

	return esopmpk.Run()
}

func (e *Esopmok) Run() error {
	deploy, err := e.client.AppsV1().Deployments(viper.GetString("namespace")).Get(context.Background(), e.name, metav1.GetOptions{})
	if err != nil {
		return err
	}

	e.compose.Mapping(deploy)
	e.compose.Name = consts.ProjectName

	out, err := yaml.Marshal(e.compose)
	if err != nil {
		return err
	}

	fmt.Println(string(out))

	return nil
}

func figureOutArgs(args []string) (string, string, error) {
	if l := len(args); l != 2 {
		return "", "", fmt.Errorf("accepts 2 args, received %d", l)
	}
	return args[0], args[1], nil
}
