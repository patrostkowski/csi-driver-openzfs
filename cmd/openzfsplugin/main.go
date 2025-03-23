package main

import (
	"flag"
	"os"

	"k8s.io/klog/v2"

	"github.com/patrostkowski/csi-driver-openzfs/pkg/openzfs"
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
