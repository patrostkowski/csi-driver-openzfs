FROM registry.k8s.io/build-image/debian-base:bookworm-v1.0.4

RUN apt update && apt upgrade -y

LABEL maintainers="patrostkowski"
LABEL description="OpenZFS CSI Driver"
ARG ARCH=amd64
ARG binary=./_output/${ARCH}/openzfsplugin
COPY ${binary} /openzfsplugin

ENTRYPOINT ["/openzfsplugin"]