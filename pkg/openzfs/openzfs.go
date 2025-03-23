/*
Copyright 2025 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

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
