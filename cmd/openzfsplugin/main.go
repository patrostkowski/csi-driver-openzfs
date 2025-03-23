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

package main

import (
	"flag"
	"os"

	"github.com/patrostkowski/csi-driver-openzfs/pkg/openzfs"
	"k8s.io/klog/v2"
)

var (
	version = "dev"
)

func main() {
	cfg := openzfs.Config{
		VendorVersion: version,
	}

	flag.StringVar(&cfg.Endpoint, "endpoint", "", "CSI endpoint")
	flag.StringVar(&cfg.DriverName, "drivername", "openzfs.csi.k8s.io", "name of the driver")
	flag.StringVar(&cfg.NodeID, "nodeid", "", "node id")
	flag.StringVar(&cfg.Plugin, "plugin", "", "plugin type: controller or node")

	klog.InitFlags(nil)
	flag.Parse()

	driver, err := openzfs.NewOpenZFSDriver(&cfg)
	if err != nil {
		klog.Errorf("Failed to initialize driver: %s", err.Error())
		os.Exit(1)
	}

	if err := driver.Run(); err != nil {
		klog.Errorf("Failed to run driver: %s", err.Error())
		os.Exit(1)
	}
}
