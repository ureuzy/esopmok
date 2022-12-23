package app

import (
	composeSpec "github.com/compose-spec/compose-go/types"
	v1 "k8s.io/api/apps/v1"
)

type Compose struct {
	*composeSpec.Config `yaml:",inline"`
}

func (c *Compose) Mapping(deploy *v1.Deployment) {
	c.Services = MappingServices(deploy)
	c.Volumes = MappingVolumes(deploy)
}
