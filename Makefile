ROOT_DIR    = $(shell pwd)
NAMESPACE   = "ximubuqi"
DEPLOY_NAME = "iot_device_simulation_arm"
IMAGE_TAG   = "v0.0.1"
IMAGE_NAME = "iot-device-simulation"
DOCKER_NAME = "iotds-docker"

include ./hack/hack.mk
docker:
	@rm $(DEPLOY_NAME) || true
	@go mod tidy
	@set GOARCH=arm
	@go env -w GOARCH=arm
	@set GOOS=linux
	@go env -w GOOS=linux
	@go build -tags=k8s -o $(DEPLOY_NAME)  .
	@set GOARCH=amd64
	@go env -w GOARCH=amd64
	@set GOOS=windows
	@go env -w GOOS=windows
	@docker rmi -f $(NAMESPACE)/$(IMAGE_NAME):$(IMAGE_TAG) || true
	@docker pull ubuntu:20.04
	@docker build -t $(NAMESPACE)/$(IMAGE_NAME):$(IMAGE_TAG) .