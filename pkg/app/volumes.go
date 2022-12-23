package app

import (
	appsv1 "k8s.io/api/apps/v1"

	composeSpec "github.com/compose-spec/compose-go/types"
)

func MappingVolumes(deploy *appsv1.Deployment) composeSpec.Volumes {
	volumes := composeSpec.Volumes{}
	for _, v := range deploy.Spec.Template.Spec.Volumes {
		volumes[v.Name] = composeSpec.VolumeConfig{
			Name:       v.Name,
			Driver:     "",
			DriverOpts: nil,
			External:   composeSpec.External{},
			Labels:     nil,
			Extensions: nil,
		}
	}
	return volumes
}
