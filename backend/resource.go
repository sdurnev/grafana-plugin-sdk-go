package backend

import (
	"context"

	bproto "github.com/grafana/grafana-plugin-sdk-go/genproto/go/grafana_plugin"
)

type ResourceRequest struct {
	Headers map[string]string
	Method  string
	Path    string
	Body    []byte
}

func resourceRequestFromProtobuf(req *bproto.ResourceRequest) *ResourceRequest {
	return &ResourceRequest{
		Headers: req.Headers,
		Method:  req.Method,
		Path:    req.Path,
		Body:    req.Body,
	}
}

type ResourceResponse struct {
	Headers map[string]string
	Code    int32
	Body    []byte
}

func (rr *ResourceResponse) toProtobuf() *bproto.ResourceResponse {
	return &bproto.ResourceResponse{
		Headers: rr.Headers,
		Code:    rr.Code,
		Body:    rr.Body,
	}
}

// ResourceHandler handles backend plugin checks.
type ResourceHandler interface {
	Resource(ctx context.Context, pc PluginConfig, req *ResourceRequest) (*ResourceResponse, error)
}

func (p *backendPluginWrapper) Resource(ctx context.Context, req *bproto.ResourceRequest) (*bproto.ResourceResponse, error) {
	pc := pluginConfigFromProto(req.Config)
	resourceReq := resourceRequestFromProtobuf(req)
	res, err := p.resourceHandler.Resource(ctx, pc, resourceReq)
	if err != nil {
		return nil, err
	}
	return res.toProtobuf(), nil

}
