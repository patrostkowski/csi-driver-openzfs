CMDS=openzfsplugin
all: build

build:
	go build -o ${CMDS} ./cmd/openzfsplugin/main.go