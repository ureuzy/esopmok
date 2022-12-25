package app

import (
	"fmt"

	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"

	composeSpec "github.com/compose-spec/compose-go/types"
	"github.com/ureuzy/esopmok/pkg/utils"
)

func MappingServices(deploy *appsv1.Deployment) (services composeSpec.Services) {
	for _, c := range deploy.Spec.Template.Spec.Containers {
		services = append(services, mappingService(deploy, &c))
	}
	for _, c := range deploy.Spec.Template.Spec.InitContainers {
		services = append(services, mappingService(deploy, &c))
	}
	return
}

func mappingService(deploy *appsv1.Deployment, container *corev1.Container) (service composeSpec.ServiceConfig) {
	service.Name = fmt.Sprintf("%s-%s", deploy.Name, container.Name)
	service.Image = container.Image
	service.Deploy = mappingDeploy(deploy)
	service.PullPolicy = mappingPullPolicy(container.ImagePullPolicy)
	service.Volumes = *mappingVolumeMounts(container.VolumeMounts)
	service.Ports = *mappingPorts(container.Ports)
	service.Environment = *mappingEnvironment(container.Env)
	service.Command = container.Command

	if container.SecurityContext != nil && container.SecurityContext.Capabilities != nil {
		service.CapAdd = *mappingCapabilities(container.SecurityContext.Capabilities.Add)
		service.CapDrop = *mappingCapabilities(container.SecurityContext.Capabilities.Drop)
	}

	return
}

func mappingDeploy(deployment *appsv1.Deployment) *composeSpec.DeployConfig {
	deploy := composeSpec.DeployConfig{}
	deploy.Replicas = utils.Pointer(uint64(*deployment.Spec.Replicas))
	return &deploy
}

// MapVolumeMounts TODO needs to be changed depending on the type of volumes. Now fixed "volume" type
func mappingVolumeMounts(volumeMounts []corev1.VolumeMount) *[]composeSpec.ServiceVolumeConfig {
	return utils.MapP(&volumeMounts, func(v corev1.VolumeMount) (result composeSpec.ServiceVolumeConfig) {
		result.Type = composeSpec.VolumeTypeVolume
		result.Source = v.Name
		result.Target = v.MountPath
		return
	})
}

func mappingPorts(containerPorts []corev1.ContainerPort) *[]composeSpec.ServicePortConfig {
	return utils.MapP(&containerPorts, func(v corev1.ContainerPort) (result composeSpec.ServicePortConfig) {
		result.Target = uint32(v.ContainerPort)
		result.Protocol = string(v.Protocol)
		return
	})
}

func mappingEnvironment(envVar []corev1.EnvVar) *composeSpec.MappingWithEquals {
	environments := composeSpec.MappingWithEquals{}
	for _, env := range envVar {
		environments[env.Name] = utils.Pointer(env.Value)
	}
	return &environments
}

func mappingCapabilities(capabilities []corev1.Capability) *[]string {
	return utils.MapP(&capabilities, func(c corev1.Capability) string {
		return string(c)
	})
}

func mappingPullPolicy(policy corev1.PullPolicy) string {
	switch policy {
	case corev1.PullAlways :
		return composeSpec.PullPolicyAlways
	case corev1.PullNever:
		return composeSpec.PullPolicyNever
	case corev1.PullIfNotPresent:
		return composeSpec.PullPolicyIfNotPresent
	default:
		return "auto"
	}
}
