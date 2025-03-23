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
	identity   csi.IdentityServer
	controller csi.ControllerServer
	node       csi.NodeServer
	config     Config
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

	driver := &OpenZFS{config: *cfg}

	switch cfg.Plugin {
	case "controller":
		driver.controller = NewController(driver)
	case "node":
		driver.node = NewNode(driver)
	default:
		return nil, errors.New("invalid or missing plugin name")
	}

	driver.identity = NewIdentity(driver)

	klog.Infof("Driver: %v", cfg.DriverName)
	klog.Infof("Version: %s", cfg.VendorVersion)

	return driver, nil
}

func (openzfs *OpenZFS) Run() error {
	s := NewNonBlockingGRPCServer()
	s.Start(openzfs.config.Endpoint, openzfs.identity, openzfs.controller, openzfs.node, nil, nil)
	s.Wait()
	return nil
}
