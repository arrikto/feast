// Code generated by go-swagger; DO NOT EDIT.

package infra_object_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new infra object service API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for infra object service API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
InfraObjectServiceListInfraObjects infra object service list infra objects API
*/
func (a *Client) InfraObjectServiceListInfraObjects(params *InfraObjectServiceListInfraObjectsParams, authInfo runtime.ClientAuthInfoWriter) (*InfraObjectServiceListInfraObjectsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewInfraObjectServiceListInfraObjectsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "InfraObjectService_ListInfraObjects",
		Method:             "GET",
		PathPattern:        "/ListInfraObjects",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &InfraObjectServiceListInfraObjectsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*InfraObjectServiceListInfraObjectsOK), nil

}

/*
InfraObjectServiceUpdateInfraObjects infra object service update infra objects API
*/
func (a *Client) InfraObjectServiceUpdateInfraObjects(params *InfraObjectServiceUpdateInfraObjectsParams, authInfo runtime.ClientAuthInfoWriter) (*InfraObjectServiceUpdateInfraObjectsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewInfraObjectServiceUpdateInfraObjectsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "InfraObjectService_UpdateInfraObjects",
		Method:             "POST",
		PathPattern:        "/UpdateInfraObjects",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &InfraObjectServiceUpdateInfraObjectsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*InfraObjectServiceUpdateInfraObjectsOK), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}