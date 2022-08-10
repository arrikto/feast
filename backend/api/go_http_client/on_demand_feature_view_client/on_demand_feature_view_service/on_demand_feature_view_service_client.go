// Code generated by go-swagger; DO NOT EDIT.

package on_demand_feature_view_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new on demand feature view service API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for on demand feature view service API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
OnDemandFeatureViewServiceCreateOnDemandFeatureView on demand feature view service create on demand feature view API
*/
func (a *Client) OnDemandFeatureViewServiceCreateOnDemandFeatureView(params *OnDemandFeatureViewServiceCreateOnDemandFeatureViewParams, authInfo runtime.ClientAuthInfoWriter) (*OnDemandFeatureViewServiceCreateOnDemandFeatureViewOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewOnDemandFeatureViewServiceCreateOnDemandFeatureViewParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "OnDemandFeatureViewService_CreateOnDemandFeatureView",
		Method:             "POST",
		PathPattern:        "/CreateOnDemandFeatureView",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &OnDemandFeatureViewServiceCreateOnDemandFeatureViewReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*OnDemandFeatureViewServiceCreateOnDemandFeatureViewOK), nil

}

/*
OnDemandFeatureViewServiceDeleteOnDemandFeatureView on demand feature view service delete on demand feature view API
*/
func (a *Client) OnDemandFeatureViewServiceDeleteOnDemandFeatureView(params *OnDemandFeatureViewServiceDeleteOnDemandFeatureViewParams, authInfo runtime.ClientAuthInfoWriter) (*OnDemandFeatureViewServiceDeleteOnDemandFeatureViewOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewOnDemandFeatureViewServiceDeleteOnDemandFeatureViewParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "OnDemandFeatureViewService_DeleteOnDemandFeatureView",
		Method:             "DELETE",
		PathPattern:        "/DeleteOnDemandFeatureView",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &OnDemandFeatureViewServiceDeleteOnDemandFeatureViewReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*OnDemandFeatureViewServiceDeleteOnDemandFeatureViewOK), nil

}

/*
OnDemandFeatureViewServiceGetOnDemandFeatureView on demand feature view service get on demand feature view API
*/
func (a *Client) OnDemandFeatureViewServiceGetOnDemandFeatureView(params *OnDemandFeatureViewServiceGetOnDemandFeatureViewParams, authInfo runtime.ClientAuthInfoWriter) (*OnDemandFeatureViewServiceGetOnDemandFeatureViewOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewOnDemandFeatureViewServiceGetOnDemandFeatureViewParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "OnDemandFeatureViewService_GetOnDemandFeatureView",
		Method:             "GET",
		PathPattern:        "/GetOnDemandFeatureView",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &OnDemandFeatureViewServiceGetOnDemandFeatureViewReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*OnDemandFeatureViewServiceGetOnDemandFeatureViewOK), nil

}

/*
OnDemandFeatureViewServiceListOnDemandFeatureViews on demand feature view service list on demand feature views API
*/
func (a *Client) OnDemandFeatureViewServiceListOnDemandFeatureViews(params *OnDemandFeatureViewServiceListOnDemandFeatureViewsParams, authInfo runtime.ClientAuthInfoWriter) (*OnDemandFeatureViewServiceListOnDemandFeatureViewsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewOnDemandFeatureViewServiceListOnDemandFeatureViewsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "OnDemandFeatureViewService_ListOnDemandFeatureViews",
		Method:             "GET",
		PathPattern:        "/ListOnDemandFeatureViews",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &OnDemandFeatureViewServiceListOnDemandFeatureViewsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*OnDemandFeatureViewServiceListOnDemandFeatureViewsOK), nil

}

/*
OnDemandFeatureViewServiceUpdateOnDemandFeatureView on demand feature view service update on demand feature view API
*/
func (a *Client) OnDemandFeatureViewServiceUpdateOnDemandFeatureView(params *OnDemandFeatureViewServiceUpdateOnDemandFeatureViewParams, authInfo runtime.ClientAuthInfoWriter) (*OnDemandFeatureViewServiceUpdateOnDemandFeatureViewOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewOnDemandFeatureViewServiceUpdateOnDemandFeatureViewParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "OnDemandFeatureViewService_UpdateOnDemandFeatureView",
		Method:             "POST",
		PathPattern:        "/UpdateOnDemandFeatureView",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &OnDemandFeatureViewServiceUpdateOnDemandFeatureViewReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*OnDemandFeatureViewServiceUpdateOnDemandFeatureViewOK), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
