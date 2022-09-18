// Code generated by go-swagger; DO NOT EDIT.

package saved_dataset_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	saved_dataset_model "github.com/feast-dev/feast/backend/api/go_http_client/saved_dataset_model"
)

// SavedDatasetServiceUpdateSavedDatasetReader is a Reader for the SavedDatasetServiceUpdateSavedDataset structure.
type SavedDatasetServiceUpdateSavedDatasetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SavedDatasetServiceUpdateSavedDatasetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewSavedDatasetServiceUpdateSavedDatasetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewSavedDatasetServiceUpdateSavedDatasetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewSavedDatasetServiceUpdateSavedDatasetOK creates a SavedDatasetServiceUpdateSavedDatasetOK with default headers values
func NewSavedDatasetServiceUpdateSavedDatasetOK() *SavedDatasetServiceUpdateSavedDatasetOK {
	return &SavedDatasetServiceUpdateSavedDatasetOK{}
}

/*SavedDatasetServiceUpdateSavedDatasetOK handles this case with default header values.

A successful response.
*/
type SavedDatasetServiceUpdateSavedDatasetOK struct {
	Payload *saved_dataset_model.APISavedDataset
}

func (o *SavedDatasetServiceUpdateSavedDatasetOK) Error() string {
	return fmt.Sprintf("[POST /UpdateSavedDataset][%d] savedDatasetServiceUpdateSavedDatasetOK  %+v", 200, o.Payload)
}

func (o *SavedDatasetServiceUpdateSavedDatasetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(saved_dataset_model.APISavedDataset)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSavedDatasetServiceUpdateSavedDatasetDefault creates a SavedDatasetServiceUpdateSavedDatasetDefault with default headers values
func NewSavedDatasetServiceUpdateSavedDatasetDefault(code int) *SavedDatasetServiceUpdateSavedDatasetDefault {
	return &SavedDatasetServiceUpdateSavedDatasetDefault{
		_statusCode: code,
	}
}

/*SavedDatasetServiceUpdateSavedDatasetDefault handles this case with default header values.

An unexpected error response.
*/
type SavedDatasetServiceUpdateSavedDatasetDefault struct {
	_statusCode int

	Payload *saved_dataset_model.GatewayruntimeError
}

// Code gets the status code for the saved dataset service update saved dataset default response
func (o *SavedDatasetServiceUpdateSavedDatasetDefault) Code() int {
	return o._statusCode
}

func (o *SavedDatasetServiceUpdateSavedDatasetDefault) Error() string {
	return fmt.Sprintf("[POST /UpdateSavedDataset][%d] SavedDatasetService_UpdateSavedDataset default  %+v", o._statusCode, o.Payload)
}

func (o *SavedDatasetServiceUpdateSavedDatasetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(saved_dataset_model.GatewayruntimeError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}