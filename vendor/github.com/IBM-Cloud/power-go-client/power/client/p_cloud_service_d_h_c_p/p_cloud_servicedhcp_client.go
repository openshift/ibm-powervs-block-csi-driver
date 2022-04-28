// Code generated by go-swagger; DO NOT EDIT.

package p_cloud_service_d_h_c_p

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new p cloud service d h c p API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for p cloud service d h c p API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	PcloudDhcpDelete(params *PcloudDhcpDeleteParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PcloudDhcpDeleteAccepted, error)

	PcloudDhcpGet(params *PcloudDhcpGetParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PcloudDhcpGetOK, error)

	PcloudDhcpGetall(params *PcloudDhcpGetallParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PcloudDhcpGetallOK, error)

	PcloudDhcpPost(params *PcloudDhcpPostParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PcloudDhcpPostAccepted, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  PcloudDhcpDelete deletes d h c p server open shift internal use only
*/
func (a *Client) PcloudDhcpDelete(params *PcloudDhcpDeleteParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PcloudDhcpDeleteAccepted, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPcloudDhcpDeleteParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "pcloud.dhcp.delete",
		Method:             "DELETE",
		PathPattern:        "/pcloud/v1/cloud-instances/{cloud_instance_id}/services/dhcp/{dhcp_id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &PcloudDhcpDeleteReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*PcloudDhcpDeleteAccepted)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for pcloud.dhcp.delete: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  PcloudDhcpGet gets d h c p server information open shift internal use only
*/
func (a *Client) PcloudDhcpGet(params *PcloudDhcpGetParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PcloudDhcpGetOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPcloudDhcpGetParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "pcloud.dhcp.get",
		Method:             "GET",
		PathPattern:        "/pcloud/v1/cloud-instances/{cloud_instance_id}/services/dhcp/{dhcp_id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &PcloudDhcpGetReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*PcloudDhcpGetOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for pcloud.dhcp.get: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  PcloudDhcpGetall gets all d h c p servers information open shift internal use only
*/
func (a *Client) PcloudDhcpGetall(params *PcloudDhcpGetallParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PcloudDhcpGetallOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPcloudDhcpGetallParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "pcloud.dhcp.getall",
		Method:             "GET",
		PathPattern:        "/pcloud/v1/cloud-instances/{cloud_instance_id}/services/dhcp",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &PcloudDhcpGetallReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*PcloudDhcpGetallOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for pcloud.dhcp.getall: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  PcloudDhcpPost creates a d h c p server open shift internal use only
*/
func (a *Client) PcloudDhcpPost(params *PcloudDhcpPostParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PcloudDhcpPostAccepted, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPcloudDhcpPostParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "pcloud.dhcp.post",
		Method:             "POST",
		PathPattern:        "/pcloud/v1/cloud-instances/{cloud_instance_id}/services/dhcp",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &PcloudDhcpPostReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*PcloudDhcpPostAccepted)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for pcloud.dhcp.post: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
