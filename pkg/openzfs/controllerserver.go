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

func (openzfs *OpenZFS) CreateVolume(ctx context.Context, req *csi.CreateVolumeRequest) (resp *csi.CreateVolumeResponse, finalErr error) {
	return nil, status.Errorf(codes.Aborted, "unsupported %+v", req)
}

func (openzfs *OpenZFS) DeleteVolume(ctx context.Context, req *csi.DeleteVolumeRequest) (*csi.DeleteVolumeResponse, error) {
	return nil, status.Errorf(codes.Aborted, "unsupported %+v", req)
}

func (openzfs *OpenZFS) ControllerGetCapabilities(ctx context.Context, req *csi.ControllerGetCapabilitiesRequest) (*csi.ControllerGetCapabilitiesResponse, error) {
	return &csi.ControllerGetCapabilitiesResponse{
		Capabilities: openzfs.getControllerServiceCapabilities(),
	}, nil
}

func (openzfs *OpenZFS) ValidateVolumeCapabilities(ctx context.Context, req *csi.ValidateVolumeCapabilitiesRequest) (*csi.ValidateVolumeCapabilitiesResponse, error) {
	return nil, status.Errorf(codes.Aborted, "unsupported %+v", req)
}

func (openzfs *OpenZFS) ControllerPublishVolume(ctx context.Context, req *csi.ControllerPublishVolumeRequest) (*csi.ControllerPublishVolumeResponse, error) {
	return nil, status.Errorf(codes.Aborted, "unsupported %+v", req)
}

func (openzfs *OpenZFS) ControllerUnpublishVolume(ctx context.Context, req *csi.ControllerUnpublishVolumeRequest) (*csi.ControllerUnpublishVolumeResponse, error) {
	return nil, status.Errorf(codes.Aborted, "unsupported %+v", req)
}

func (openzfs *OpenZFS) GetCapacity(ctx context.Context, req *csi.GetCapacityRequest) (*csi.GetCapacityResponse, error) {
	return nil, status.Errorf(codes.Aborted, "unsupported %+v", req)
}

func (openzfs *OpenZFS) ListVolumes(ctx context.Context, req *csi.ListVolumesRequest) (*csi.ListVolumesResponse, error) {
	return nil, status.Errorf(codes.Aborted, "unsupported %+v", req)
}

func (openzfs *OpenZFS) ControllerGetVolume(ctx context.Context, req *csi.ControllerGetVolumeRequest) (*csi.ControllerGetVolumeResponse, error) {
	return nil, status.Errorf(codes.Aborted, "unsupported %+v", req)
}

func (openzfs *OpenZFS) ControllerModifyVolume(ctx context.Context, req *csi.ControllerModifyVolumeRequest) (*csi.ControllerModifyVolumeResponse, error) {
	return nil, status.Errorf(codes.Aborted, "unsupported %+v", req)
}

func (openzfs *OpenZFS) CreateSnapshot(ctx context.Context, req *csi.CreateSnapshotRequest) (*csi.CreateSnapshotResponse, error) {
	return nil, status.Errorf(codes.Aborted, "unsupported %+v", req)
}

func (openzfs *OpenZFS) DeleteSnapshot(ctx context.Context, req *csi.DeleteSnapshotRequest) (*csi.DeleteSnapshotResponse, error) {
	return nil, status.Errorf(codes.Aborted, "unsupported %+v", req)
}

func (openzfs *OpenZFS) ListSnapshots(ctx context.Context, req *csi.ListSnapshotsRequest) (*csi.ListSnapshotsResponse, error) {
	return nil, status.Errorf(codes.Aborted, "unsupported %+v", req)
}

func (openzfs *OpenZFS) ControllerExpandVolume(ctx context.Context, req *csi.ControllerExpandVolumeRequest) (*csi.ControllerExpandVolumeResponse, error) {
	return nil, status.Errorf(codes.Aborted, "unsupported %+v", req)
}

func (openzfs *OpenZFS) getControllerServiceCapabilities() []*csi.ControllerServiceCapability {
	var csc []*csi.ControllerServiceCapability
	return csc
}
