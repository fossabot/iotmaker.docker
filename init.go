package iotmaker_docker

// Must be first function call
func (el *DockerSystem) Init() error {
	el.contextCreate()
	return el.clientCreate()
}