package main

import (
	"flag"
	"os"

	"k8s.io/klog/v2"

	"github.com/patrostkowski/csi-driver-openzfs/pkg/openzfs"
)

func main() {
	var cfg openzfs.Config

	flag.StringVar(&cfg.Endpoint, "endpoint", "unix://tmp/csi.sock", "CSI endpoint")
	flag.StringVar(&cfg.DriverName, "drivername", "hostpath.csi.k8s.io", "name of the driver")
	flag.StringVar(&cfg.NodeID, "nodeid", "", "node id")

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
