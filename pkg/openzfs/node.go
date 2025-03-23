package openzfs

import (
	"github.com/container-storage-interface/spec/lib/go/csi"
	"golang.org/x/net/context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Node struct {
	csi.UnimplementedNodeServer
	driver *OpenZFS
}

func NewNode(openzfs *OpenZFS) csi.NodeServer {
	return &Node{
		driver: openzfs,
	}
}

// This implements csi.NodeServer
func (node *Node) NodePublishVolume(ctx context.Context, req *csi.NodePublishVolumeRequest) (resp *csi.NodePublishVolumeResponse, finalErr error) {
	return nil, status.Errorf(codes.Aborted, "unsupported %+v", req)
}

// This implements csi.NodeServer
func (node *Node) NodeUnpublishVolume(ctx context.Context, req *csi.NodeUnpublishVolumeRequest) (resp *csi.NodeUnpublishVolumeResponse, finalErr error) {
	return nil, status.Errorf(codes.Aborted, "unsupported %+v", req)
}

// This implements csi.NodeServer
func (node *Node) NodeGetInfo(ctx context.Context, req *csi.NodeGetInfoRequest) (resp *csi.NodeGetInfoResponse, finalErr error) {
	return nil, status.Errorf(codes.Aborted, "unsupported %+v", req)
}

// This implements csi.NodeServer
func (node *Node) NodeGetVolumeStats(ctx context.Context, req *csi.NodeGetVolumeStatsRequest) (resp *csi.NodeGetVolumeStatsResponse, finalErr error) {
	return nil, status.Errorf(codes.Aborted, "unsupported %+v", req)
}

// This implements csi.NodeServer
func (node *Node) NodeGetCapabilities(ctx context.Context, req *csi.NodeGetCapabilitiesRequest) (resp *csi.NodeGetCapabilitiesResponse, finalErr error) {
	return nil, status.Errorf(codes.Aborted, "unsupported %+v", req)
}

// This implements csi.NodeServer
func (node *Node) NodeStageVolume(ctx context.Context, req *csi.NodeStageVolumeRequest) (resp *csi.NodeStageVolumeResponse, finalErr error) {
	return nil, status.Errorf(codes.Aborted, "unsupported %+v", req)
}

// This implements csi.NodeServer
func (node *Node) NodeUnstageVolume(ctx context.Context, req *csi.NodeUnstageVolumeRequest) (resp *csi.NodeUnstageVolumeResponse, finalErr error) {
	return nil, status.Errorf(codes.Aborted, "unsupported %+v", req)
}

// This implements csi.NodeServer
func (node *Node) NodeExpandVolume(ctx context.Context, req *csi.NodeExpandVolumeRequest) (resp *csi.NodeExpandVolumeResponse, finalErr error) {
	return nil, status.Errorf(codes.Aborted, "unsupported %+v", req)
}
