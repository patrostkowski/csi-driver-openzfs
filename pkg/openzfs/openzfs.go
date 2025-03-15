package openzfs

import (
	"errors"
	"time"

	"k8s.io/klog/v2"
)

type Config struct {
	Endpoint   string
	DriverName string
	NodeID     string
}

type OpenZFS struct {
	config Config
}

func NewOpenZFSDriver(cfg *Config) (*OpenZFS, error) {
	if cfg.DriverName == "" {
		return nil, errors.New("no driver name provided")
	}

	if cfg.NodeID == "" {
		return nil, errors.New("no node id provided")
	}

	if cfg.Endpoint == "" {
		return nil, errors.New("no driver endpoint provided")
	}
	return &OpenZFS{config: *cfg}, nil
}

func (openzfs *OpenZFS) Run() error {
	klog.Infof("Hello from %+v", openzfs)

	for {
		time.Sleep(time.Hour)
	}

	return nil // Unreachable
}
