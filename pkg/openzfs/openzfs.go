package openzfs

import (
	"errors"

	"github.com/container-storage-interface/spec/lib/go/csi"
	"k8s.io/klog/v2"
)

type Config struct {
	Endpoint      string
	DriverName    string
	NodeID        string
	VendorVersion string
	Plugin        string
}

type OpenZFS struct {
	csi.UnimplementedIdentityServer
	csi.UnimplementedControllerServer
	csi.UnimplementedNodeServer
	csi.UnimplementedGroupControllerServer
	csi.UnimplementedSnapshotMetadataServer
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

	switch cfg.Plugin {
	case "controller":
		break
	case "node":
		break
	default:
		return nil, errors.New("no plugin name provided")
	}

	klog.Infof("Driver: %v ", cfg.DriverName)
	klog.Infof("Version: %s", cfg.VendorVersion)

	return &OpenZFS{config: *cfg}, nil
}

func (openzfs *OpenZFS) Run() error {
	s := NewNonBlockingGRPCServer()

	s.Start(openzfs.config.Endpoint, openzfs, openzfs, openzfs, openzfs, nil)
	s.Wait()

	return nil
}
