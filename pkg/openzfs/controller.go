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
	"strings"
	"time"

	"golang.org/x/net/context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/informers"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/cache"
	"k8s.io/klog/v2"
	signals "sigs.k8s.io/controller-runtime/pkg/manager/signals"

	"github.com/container-storage-interface/spec/lib/go/csi"
)

type Controller struct {
	csi.UnimplementedControllerServer
	driver       *OpenZFS
	capabilities []*csi.ControllerServiceCapability
	nodeInformer cache.SharedIndexInformer
	client       *kubernetes.Clientset
}

func NewController(openzfs *OpenZFS) csi.ControllerServer {
	ctrl := &Controller{
		driver:       openzfs,
		capabilities: newControllerCapabilities(),
		nodeInformer: nil,
	}
	if err := ctrl.init(); err != nil {
		klog.Fatalf("init controller failed: %v", err)
	}
	return ctrl
}

func (ctrl *Controller) init() error {
	config, err := rest.InClusterConfig()
	if err != nil {
		klog.Errorf("Error while creating the in cluster config %v", err.Error())
		return err
	}

	if ctrl.client, err = kubernetes.NewForConfig(config); err != nil {
		klog.Errorf("Error while creating the clientset %v", err.Error())
		return err
	}

	informerFactory := informers.NewSharedInformerFactory(ctrl.client, time.Second*10)

	stopCh := signals.SetupSignalHandler()

	ctrl.nodeInformer = informerFactory.Core().V1().Nodes().Informer()

	go ctrl.nodeInformer.Run(stopCh.Done())

	klog.Info("waiting for node informer caches to be synced")
	cache.WaitForCacheSync(stopCh.Done(), ctrl.nodeInformer.HasSynced)
	klog.Info("synced node informer caches")
	return nil
}

func newControllerCapabilities() []*csi.ControllerServiceCapability {
	capTypes := []csi.ControllerServiceCapability_RPC_Type{
		csi.ControllerServiceCapability_RPC_CREATE_DELETE_VOLUME,
	}

	capabilities := make([]*csi.ControllerServiceCapability, len(capTypes))
	for i, cap := range capTypes {
		capabilities[i] = &csi.ControllerServiceCapability{
			Type: &csi.ControllerServiceCapability_Rpc{
				Rpc: &csi.ControllerServiceCapability_RPC{Type: cap},
			},
		}
	}
	return capabilities
}

// This implements csi.ControllerServer
func (ctrl *Controller) ControllerGetCapabilities(ctx context.Context, req *csi.ControllerGetCapabilitiesRequest) (*csi.ControllerGetCapabilitiesResponse, error) {
	return &csi.ControllerGetCapabilitiesResponse{
		Capabilities: ctrl.capabilities,
	}, nil
}

// This implements csi.ControllerServer
func (ctrl *Controller) ValidateVolumeCapabilities(ctx context.Context, req *csi.ValidateVolumeCapabilitiesRequest) (*csi.ValidateVolumeCapabilitiesResponse, error) {
	volumeID := strings.ToLower(req.GetVolumeId())
	if len(volumeID) == 0 {
		return nil, status.Error(codes.InvalidArgument, "Volume ID not provided")
	}

	volCaps := req.GetVolumeCapabilities()
	if len(volCaps) == 0 {
		return nil, status.Error(codes.InvalidArgument, "Volume capabilities not provided")
	}

	return &csi.ValidateVolumeCapabilitiesResponse{
		Confirmed: &csi.ValidateVolumeCapabilitiesResponse_Confirmed{},
	}, nil
}

// TODO
// create mechanism of creating nodes
// 1. get all nodes in the cluster
// 2. validate on which node create volume
// 3. check if volume does not exist, if exists handle err
// 4. execute volume creation

// This implements csi.ControllerServer
func (ctrl *Controller) CreateVolume(ctx context.Context, req *csi.CreateVolumeRequest) (*csi.CreateVolumeResponse, error) {
	nodes, err := ctrl.client.CoreV1().Nodes().List(context.TODO(), v1.ListOptions{})
	if err != nil {
		return nil, status.Errorf(codes.Aborted, "Could not list nodes: %s", err.Error())
	}
	klog.Infof("Hello from my nodes: %+v", nodes)
	return nil, status.Errorf(codes.Unimplemented, "unsupported %+v", req)
}

// This implements csi.ControllerServer
func (ctrl *Controller) DeleteVolume(ctx context.Context, req *csi.DeleteVolumeRequest) (*csi.DeleteVolumeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "unsupported %+v", req)
}

// This implements csi.ControllerServer
func (ctrl *Controller) ControllerPublishVolume(ctx context.Context, req *csi.ControllerPublishVolumeRequest) (*csi.ControllerPublishVolumeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "unsupported %+v", req)
}

// This implements csi.ControllerServer
func (ctrl *Controller) ControllerUnpublishVolume(ctx context.Context, req *csi.ControllerUnpublishVolumeRequest) (*csi.ControllerUnpublishVolumeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "unsupported %+v", req)
}

// This implements csi.ControllerServer
func (ctrl *Controller) GetCapacity(ctx context.Context, req *csi.GetCapacityRequest) (*csi.GetCapacityResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "unsupported %+v", req)
}

// This implements csi.ControllerServer
func (ctrl *Controller) ListVolumes(ctx context.Context, req *csi.ListVolumesRequest) (*csi.ListVolumesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "unsupported %+v", req)
}

// This implements csi.ControllerServer
func (ctrl *Controller) CreateSnapshot(ctx context.Context, req *csi.CreateSnapshotRequest) (*csi.CreateSnapshotResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "unsupported %+v", req)
}

// This implements csi.ControllerServer
func (ctrl *Controller) DeleteSnapshot(ctx context.Context, req *csi.DeleteSnapshotRequest) (*csi.DeleteSnapshotResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "unsupported %+v", req)
}

// This implements csi.ControllerServer
func (ctrl *Controller) ListSnapshots(ctx context.Context, req *csi.ListSnapshotsRequest) (*csi.ListSnapshotsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "unsupported %+v", req)
}

// This implements csi.ControllerServer
func (ctrl *Controller) ControllerExpandVolume(ctx context.Context, req *csi.ControllerExpandVolumeRequest) (*csi.ControllerExpandVolumeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "unsupported %+v", req)
}
