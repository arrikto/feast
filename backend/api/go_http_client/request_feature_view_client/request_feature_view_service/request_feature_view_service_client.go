// Code generated by go-swagger; DO NOT EDIT.

package request_feature_view_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new request feature view service API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for request feature view service API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
RequestFeatureViewServiceCreateRequestFeatureView request feature view service create request feature view API
*/
func (a *Client) RequestFeatureViewServiceCreateRequestFeatureView(params *RequestFeatureViewServiceCreateRequestFeatureViewParams, authInfo runtime.ClientAuthInfoWriter) (*RequestFeatureViewServiceCreateRequestFeatureViewOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewRequestFeatureViewServiceCreateRequestFeatureViewParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "RequestFeatureViewService_CreateRequestFeatureView",
		Method:             "POST",
		PathPattern:        "/CreateRequestFeatureView",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &RequestFeatureViewServiceCreateRequestFeatureViewReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*RequestFeatureViewServiceCreateRequestFeatureViewOK), nil

}

/*
RequestFeatureViewServiceDeleteRequestFeatureView request feature view service delete request feature view API
*/
func (a *Client) RequestFeatureViewServiceDeleteRequestFeatureView(params *RequestFeatureViewServiceDeleteRequestFeatureViewParams, authInfo runtime.ClientAuthInfoWriter) (*RequestFeatureViewServiceDeleteRequestFeatureViewOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewRequestFeatureViewServiceDeleteRequestFeatureViewParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "RequestFeatureViewService_DeleteRequestFeatureView",
		Method:             "DELETE",
		PathPattern:        "/DeleteRequestFeatureView",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &RequestFeatureViewServiceDeleteRequestFeatureViewReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*RequestFeatureViewServiceDeleteRequestFeatureViewOK), nil

}

/*
RequestFeatureViewServiceGetRequestFeatureView request feature view service get request feature view API
*/
func (a *Client) RequestFeatureViewServiceGetRequestFeatureView(params *RequestFeatureViewServiceGetRequestFeatureViewParams, authInfo runtime.ClientAuthInfoWriter) (*RequestFeatureViewServiceGetRequestFeatureViewOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewRequestFeatureViewServiceGetRequestFeatureViewParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "RequestFeatureViewService_GetRequestFeatureView",
		Method:             "GET",
		PathPattern:        "/GetRequestFeatureView",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &RequestFeatureViewServiceGetRequestFeatureViewReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*RequestFeatureViewServiceGetRequestFeatureViewOK), nil

}

/*
RequestFeatureViewServiceListRequestFeatureViews request feature view service list request feature views API
*/
func (a *Client) RequestFeatureViewServiceListRequestFeatureViews(params *RequestFeatureViewServiceListRequestFeatureViewsParams, authInfo runtime.ClientAuthInfoWriter) (*RequestFeatureViewServiceListRequestFeatureViewsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewRequestFeatureViewServiceListRequestFeatureViewsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "RequestFeatureViewService_ListRequestFeatureViews",
		Method:             "GET",
		PathPattern:        "/ListRequestFeatureViews",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &RequestFeatureViewServiceListRequestFeatureViewsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*RequestFeatureViewServiceListRequestFeatureViewsOK), nil

}

/*
RequestFeatureViewServiceUpdateRequestFeatureView request feature view service update request feature view API
*/
func (a *Client) RequestFeatureViewServiceUpdateRequestFeatureView(params *RequestFeatureViewServiceUpdateRequestFeatureViewParams, authInfo runtime.ClientAuthInfoWriter) (*RequestFeatureViewServiceUpdateRequestFeatureViewOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewRequestFeatureViewServiceUpdateRequestFeatureViewParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "RequestFeatureViewService_UpdateRequestFeatureView",
		Method:             "POST",
		PathPattern:        "/UpdateRequestFeatureView",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &RequestFeatureViewServiceUpdateRequestFeatureViewReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*RequestFeatureViewServiceUpdateRequestFeatureViewOK), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
