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

// SavedDatasetServiceCreateSavedDatasetReader is a Reader for the SavedDatasetServiceCreateSavedDataset structure.
type SavedDatasetServiceCreateSavedDatasetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SavedDatasetServiceCreateSavedDatasetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewSavedDatasetServiceCreateSavedDatasetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewSavedDatasetServiceCreateSavedDatasetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewSavedDatasetServiceCreateSavedDatasetOK creates a SavedDatasetServiceCreateSavedDatasetOK with default headers values
func NewSavedDatasetServiceCreateSavedDatasetOK() *SavedDatasetServiceCreateSavedDatasetOK {
	return &SavedDatasetServiceCreateSavedDatasetOK{}
}

/*SavedDatasetServiceCreateSavedDatasetOK handles this case with default header values.

A successful response.
*/
type SavedDatasetServiceCreateSavedDatasetOK struct {
	Payload *saved_dataset_model.APISavedDataset
}

func (o *SavedDatasetServiceCreateSavedDatasetOK) Error() string {
	return fmt.Sprintf("[POST /CreateSavedDataset][%d] savedDatasetServiceCreateSavedDatasetOK  %+v", 200, o.Payload)
}

func (o *SavedDatasetServiceCreateSavedDatasetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(saved_dataset_model.APISavedDataset)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSavedDatasetServiceCreateSavedDatasetDefault creates a SavedDatasetServiceCreateSavedDatasetDefault with default headers values
func NewSavedDatasetServiceCreateSavedDatasetDefault(code int) *SavedDatasetServiceCreateSavedDatasetDefault {
	return &SavedDatasetServiceCreateSavedDatasetDefault{
		_statusCode: code,
	}
}

/*SavedDatasetServiceCreateSavedDatasetDefault handles this case with default header values.

An unexpected error response.
*/
type SavedDatasetServiceCreateSavedDatasetDefault struct {
	_statusCode int

	Payload *saved_dataset_model.GatewayruntimeError
}

// Code gets the status code for the saved dataset service create saved dataset default response
func (o *SavedDatasetServiceCreateSavedDatasetDefault) Code() int {
	return o._statusCode
}

func (o *SavedDatasetServiceCreateSavedDatasetDefault) Error() string {
	return fmt.Sprintf("[POST /CreateSavedDataset][%d] SavedDatasetService_CreateSavedDataset default  %+v", o._statusCode, o.Payload)
}

func (o *SavedDatasetServiceCreateSavedDatasetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(saved_dataset_model.GatewayruntimeError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}