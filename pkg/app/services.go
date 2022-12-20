package app

import (
	"fmt"
	composeSpec "github.com/compose-spec/compose-go/types"
	"github.com/ureuzy/esopmok/pkg/utils"
	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
)

func MappingServices(deploy *appsv1.Deployment) composeSpec.Services {
	services := composeSpec.Services{}
	for _, c := range deploy.Spec.Template.Spec.Containers {
		services = append(services, mapC(deploy, c))
	}
	for _, c := range deploy.Spec.Template.Spec.InitContainers {
		services = append(services, mapC(deploy, c))
	}
	return services
}

func mapC(deploy *appsv1.Deployment, container corev1.Container) composeSpec.ServiceConfig {
	serviceConfig := composeSpec.ServiceConfig{
		Name:  fmt.Sprintf("%s-%s", deploy.Name, container.Name),
		Image: container.Image,
		Deploy: &composeSpec.DeployConfig{
			Mode:           "",
			Replicas:       utils.Pointer(uint64(*deploy.Spec.Replicas)),
			Labels:         nil,
			UpdateConfig:   nil,
			RollbackConfig: nil,
			Resources:      composeSpec.Resources{}, // TODO Bind resource quota
			RestartPolicy:  nil,
			Placement:      composeSpec.Placement{},
			EndpointMode:   "",
			Extensions:     nil,
		},
		PullPolicy: convPullPolicy(container.ImagePullPolicy),
		Volumes:    volumeMounts(container.VolumeMounts),
	}
	return serviceConfig
}

// TODO needs to be changed depending on the type of volumes. Now fixed "volume" type
func volumeMounts(mounts []corev1.VolumeMount) []composeSpec.ServiceVolumeConfig {
	var serviceVolumeConfigs []composeSpec.ServiceVolumeConfig
	for _, mount := range mounts {
		config := composeSpec.ServiceVolumeConfig{
			Type:        composeSpec.VolumeTypeVolume,
			Source:      mount.Name,
			Target:      mount.MountPath,
			ReadOnly:    false,
			Consistency: "",
			Bind:        nil,
			Volume:      nil,
			Tmpfs:       nil,
			Extensions:  nil,
		}
		serviceVolumeConfigs = append(serviceVolumeConfigs, config)
	}
	return serviceVolumeConfigs
}

func convPullPolicy(policy corev1.PullPolicy) string {
	switch policy {
	case "Always":
		return "always"
	case "IfNotPresent":
		return "if_not_present"
	case "Never":
		return "never"
	default:
		return "auto"
	}
}
