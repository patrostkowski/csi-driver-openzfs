CMDS=openzfsplugin
DOCKER_IMAGE=patrostkowski/openzfsplugin
ARCH=aarch64
OUTPUT_DIR=_output/${ARCH}
BINARY=${OUTPUT_DIR}/${CMDS}
DOCKERFILE_PATH=cmd/openzfsplugin/Dockerfile

all: build docker

build:
	go build -o ${BINARY} ./cmd/openzfsplugin/main.go

docker:
	docker buildx build -f ${DOCKERFILE_PATH} --platform linux/arm64,linux/amd64 -t patrostkowski/openzfsplugin:latest --push .

docker-local:
	docker build -f cmd/openzfsplugin/Dockerfile --platform linux/arm64 -t localhost:5001/openzfsplugin:latest .
	docker push localhost:5001/openzfsplugin:latest 

