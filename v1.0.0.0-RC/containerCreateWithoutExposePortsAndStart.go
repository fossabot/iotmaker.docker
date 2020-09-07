package iotmakerDocker

import (
	"github.com/docker/docker/api/types/mount"
	"github.com/docker/docker/api/types/network"
)

func (el *DockerSystem) ContainerCreateWithoutExposePortsAndStart(
	imageName,
	containerName string,
	restartPolicy RestartPolicy,
	mountVolumes []mount.Mount,
	containerNetwork *network.NetworkingConfig,
) (
	containerID string,
	err error,
) {

	containerID, err = el.ContainerCreateWithoutExposePorts(
		imageName,
		containerName,
		restartPolicy,
		mountVolumes,
		containerNetwork,
	)
	if err != nil {
		return
	}

	err = el.ContainerStart(containerID)
	return
}