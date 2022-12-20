package app

import (
	"context"
	"fmt"
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

func NewEsopmok(f util.Factory, args []string) (*Esopmok, error) {
	kind, name, err := figureOutArgs(args)
	if err != nil {
		return nil, err
	}
	clientSet, err := f.KubernetesClientSet()
	if err != nil {
		return nil, err
	}
	return &Esopmok{
		kind,
		name,
		clientSet,
		&Compose{&composeSpec.Config{}},
	}, nil
}

func (e *Esopmok) Run() error {
	deploy, err := e.client.AppsV1().Deployments(viper.GetString("namespace")).Get(context.Background(), e.name, metav1.GetOptions{})
	if err != nil {
		return err
	}
	e.compose.Mapping(deploy)
	out, err := yaml.Marshal(e.compose)
	if err != nil {
		return err
	}
	fmt.Println(string(out))
	return nil
}

func figureOutArgs(args []string) (string, string, error) {
	if l := len(args); l == 0 || l > 2 {
		return "", "", fmt.Errorf("accepts between 1 and 2 arg(s), received %d", l)
	}
	return args[0], args[1], nil
}
