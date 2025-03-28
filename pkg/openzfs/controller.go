/*
Copyright 2017 The Kubernetes Authors.

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
	"golang.org/x/net/context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/container-storage-interface/spec/lib/go/csi"
)

type Controller struct {
	csi.UnimplementedControllerServer
	driver *OpenZFS
}

func NewController(openzfs *OpenZFS) csi.ControllerServer {
	return &Controller{
		driver: openzfs,
	}
}

// This implements csi.ControllerServer
func (ctrl *Controller) ControllerGetCapabilities(ctx context.Context, req *csi.ControllerGetCapabilitiesRequest) (*csi.ControllerGetCapabilitiesResponse, error) {
	var csc []*csi.ControllerServiceCapability
	return &csi.ControllerGetCapabilitiesResponse{
		Capabilities: csc,
	}, nil
}

// This implements csi.ControllerServer
func (ctrl *Controller) CreateVolume(ctx context.Context, req *csi.CreateVolumeRequest) (*csi.CreateVolumeResponse, error) {
	return nil, status.Errorf(codes.Aborted, "unsupported %+v", req)
}

// This implements csi.ControllerServer
func (ctrl *Controller) DeleteVolume(ctx context.Context, req *csi.DeleteVolumeRequest) (*csi.DeleteVolumeResponse, error) {
	return nil, status.Errorf(codes.Aborted, "unsupported %+v", req)
}

// This implements csi.ControllerServer
func (ctrl *Controller) ValidateVolumeCapabilities(ctx context.Context, req *csi.ValidateVolumeCapabilitiesRequest) (*csi.ValidateVolumeCapabilitiesResponse, error) {
	return nil, status.Errorf(codes.Aborted, "unsupported %+v", req)
}

// This implements csi.ControllerServer
func (ctrl *Controller) ControllerPublishVolume(ctx context.Context, req *csi.ControllerPublishVolumeRequest) (*csi.ControllerPublishVolumeResponse, error) {
	return nil, status.Errorf(codes.Aborted, "unsupported %+v", req)
}

// This implements csi.ControllerServer
func (ctrl *Controller) ControllerUnpublishVolume(ctx context.Context, req *csi.ControllerUnpublishVolumeRequest) (*csi.ControllerUnpublishVolumeResponse, error) {
	return nil, status.Errorf(codes.Aborted, "unsupported %+v", req)
}

// This implements csi.ControllerServer
func (ctrl *Controller) GetCapacity(ctx context.Context, req *csi.GetCapacityRequest) (*csi.GetCapacityResponse, error) {
	return nil, status.Errorf(codes.Aborted, "unsupported %+v", req)
}

// This implements csi.ControllerServer
func (ctrl *Controller) ListVolumes(ctx context.Context, req *csi.ListVolumesRequest) (*csi.ListVolumesResponse, error) {
	return nil, status.Errorf(codes.Aborted, "unsupported %+v", req)
}

// This implements csi.ControllerServer
func (ctrl *Controller) CreateSnapshot(ctx context.Context, req *csi.CreateSnapshotRequest) (*csi.CreateSnapshotResponse, error) {
	return nil, status.Errorf(codes.Aborted, "unsupported %+v", req)
}

// This implements csi.ControllerServer
func (ctrl *Controller) DeleteSnapshot(ctx context.Context, req *csi.DeleteSnapshotRequest) (*csi.DeleteSnapshotResponse, error) {
	return nil, status.Errorf(codes.Aborted, "unsupported %+v", req)
}

// This implements csi.ControllerServer
func (ctrl *Controller) ListSnapshots(ctx context.Context, req *csi.ListSnapshotsRequest) (*csi.ListSnapshotsResponse, error) {
	return nil, status.Errorf(codes.Aborted, "unsupported %+v", req)
}

// This implements csi.ControllerServer
func (ctrl *Controller) ControllerExpandVolume(ctx context.Context, req *csi.ControllerExpandVolumeRequest) (*csi.ControllerExpandVolumeResponse, error) {
	return nil, status.Errorf(codes.Aborted, "unsupported %+v", req)
}
