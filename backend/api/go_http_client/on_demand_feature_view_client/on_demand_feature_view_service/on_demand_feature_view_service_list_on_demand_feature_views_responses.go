// Code generated by go-swagger; DO NOT EDIT.

package on_demand_feature_view_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	on_demand_feature_view_model "github.com/feast-dev/feast/backend/api/go_http_client/on_demand_feature_view_model"
)

// OnDemandFeatureViewServiceListOnDemandFeatureViewsReader is a Reader for the OnDemandFeatureViewServiceListOnDemandFeatureViews structure.
type OnDemandFeatureViewServiceListOnDemandFeatureViewsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *OnDemandFeatureViewServiceListOnDemandFeatureViewsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewOnDemandFeatureViewServiceListOnDemandFeatureViewsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewOnDemandFeatureViewServiceListOnDemandFeatureViewsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewOnDemandFeatureViewServiceListOnDemandFeatureViewsOK creates a OnDemandFeatureViewServiceListOnDemandFeatureViewsOK with default headers values
func NewOnDemandFeatureViewServiceListOnDemandFeatureViewsOK() *OnDemandFeatureViewServiceListOnDemandFeatureViewsOK {
	return &OnDemandFeatureViewServiceListOnDemandFeatureViewsOK{}
}

/*OnDemandFeatureViewServiceListOnDemandFeatureViewsOK handles this case with default header values.

A successful response.
*/
type OnDemandFeatureViewServiceListOnDemandFeatureViewsOK struct {
	Payload *on_demand_feature_view_model.APIListOnDemandFeatureViewsResponse
}

func (o *OnDemandFeatureViewServiceListOnDemandFeatureViewsOK) Error() string {
	return fmt.Sprintf("[GET /ListOnDemandFeatureViews][%d] onDemandFeatureViewServiceListOnDemandFeatureViewsOK  %+v", 200, o.Payload)
}

func (o *OnDemandFeatureViewServiceListOnDemandFeatureViewsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(on_demand_feature_view_model.APIListOnDemandFeatureViewsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewOnDemandFeatureViewServiceListOnDemandFeatureViewsDefault creates a OnDemandFeatureViewServiceListOnDemandFeatureViewsDefault with default headers values
func NewOnDemandFeatureViewServiceListOnDemandFeatureViewsDefault(code int) *OnDemandFeatureViewServiceListOnDemandFeatureViewsDefault {
	return &OnDemandFeatureViewServiceListOnDemandFeatureViewsDefault{
		_statusCode: code,
	}
}

/*OnDemandFeatureViewServiceListOnDemandFeatureViewsDefault handles this case with default header values.

An unexpected error response.
*/
type OnDemandFeatureViewServiceListOnDemandFeatureViewsDefault struct {
	_statusCode int

	Payload *on_demand_feature_view_model.GatewayruntimeError
}

// Code gets the status code for the on demand feature view service list on demand feature views default response
func (o *OnDemandFeatureViewServiceListOnDemandFeatureViewsDefault) Code() int {
	return o._statusCode
}

func (o *OnDemandFeatureViewServiceListOnDemandFeatureViewsDefault) Error() string {
	return fmt.Sprintf("[GET /ListOnDemandFeatureViews][%d] OnDemandFeatureViewService_ListOnDemandFeatureViews default  %+v", o._statusCode, o.Payload)
}

func (o *OnDemandFeatureViewServiceListOnDemandFeatureViewsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(on_demand_feature_view_model.GatewayruntimeError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}