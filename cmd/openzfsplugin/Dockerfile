# Build stage (cross-compilation for ARM64)
FROM golang:1.23 AS builder

WORKDIR /app

# Copy Go modules and dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the source code
COPY . .

RUN go build -o openzfsplugin ./cmd/openzfsplugin/main.go

# Runtime stage (running on ARM64)
FROM registry.k8s.io/build-image/debian-base:bookworm-v1.0.4

# Update base image
RUN apt update && apt upgrade -y

LABEL maintainers="patrostkowski"
LABEL description="OpenZFS CSI Driver"

# Copy the built binary from the builder stage
COPY --from=builder /app/openzfsplugin /openzfsplugin

# Set executable permission
RUN chmod +x /openzfsplugin

# Set entrypoint
ENTRYPOINT ["/openzfsplugin"]
