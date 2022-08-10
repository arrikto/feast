// Code generated by go-swagger; DO NOT EDIT.

package feature_view_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	feature_view_model "github.com/feast-dev/feast/backend/api/go_http_client/feature_view_model"
)

// FeatureViewServiceGetFeatureViewReader is a Reader for the FeatureViewServiceGetFeatureView structure.
type FeatureViewServiceGetFeatureViewReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *FeatureViewServiceGetFeatureViewReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewFeatureViewServiceGetFeatureViewOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewFeatureViewServiceGetFeatureViewDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewFeatureViewServiceGetFeatureViewOK creates a FeatureViewServiceGetFeatureViewOK with default headers values
func NewFeatureViewServiceGetFeatureViewOK() *FeatureViewServiceGetFeatureViewOK {
	return &FeatureViewServiceGetFeatureViewOK{}
}

/*FeatureViewServiceGetFeatureViewOK handles this case with default header values.

A successful response.
*/
type FeatureViewServiceGetFeatureViewOK struct {
	Payload *feature_view_model.APIFeatureView
}

func (o *FeatureViewServiceGetFeatureViewOK) Error() string {
	return fmt.Sprintf("[GET /GetFeatureView][%d] featureViewServiceGetFeatureViewOK  %+v", 200, o.Payload)
}

func (o *FeatureViewServiceGetFeatureViewOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(feature_view_model.APIFeatureView)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewFeatureViewServiceGetFeatureViewDefault creates a FeatureViewServiceGetFeatureViewDefault with default headers values
func NewFeatureViewServiceGetFeatureViewDefault(code int) *FeatureViewServiceGetFeatureViewDefault {
	return &FeatureViewServiceGetFeatureViewDefault{
		_statusCode: code,
	}
}

/*FeatureViewServiceGetFeatureViewDefault handles this case with default header values.

An unexpected error response.
*/
type FeatureViewServiceGetFeatureViewDefault struct {
	_statusCode int

	Payload *feature_view_model.GatewayruntimeError
}

// Code gets the status code for the feature view service get feature view default response
func (o *FeatureViewServiceGetFeatureViewDefault) Code() int {
	return o._statusCode
}

func (o *FeatureViewServiceGetFeatureViewDefault) Error() string {
	return fmt.Sprintf("[GET /GetFeatureView][%d] FeatureViewService_GetFeatureView default  %+v", o._statusCode, o.Payload)
}

func (o *FeatureViewServiceGetFeatureViewDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(feature_view_model.GatewayruntimeError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
