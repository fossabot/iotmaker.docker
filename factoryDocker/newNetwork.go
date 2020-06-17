package factoryDocker

import (
	iotmakerDocker "github.com/helmutkemper/iotmaker.docker"
)

func NewNetwork(networkName string) (err error, networkId string, networkAutoConfiguration *iotmakerDocker.NextNetworkAutoConfiguration) {
	var dockerSys = iotmakerDocker.DockerSystem{}
	dockerSys.Init()

	err, networkId, networkAutoConfiguration = dockerSys.NetworkCreate(
		networkName,
		iotmakerDocker.KNetworkDriveBridge,
		"local",
		"10.0.0.0/16",
		"10.0.0.1",
	)

	return
}