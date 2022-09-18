// Code generated by go-swagger; DO NOT EDIT.

package entity_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	entity_model "github.com/feast-dev/feast/backend/api/go_http_client/entity_model"
)

// EntityServiceCreateEntityReader is a Reader for the EntityServiceCreateEntity structure.
type EntityServiceCreateEntityReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *EntityServiceCreateEntityReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewEntityServiceCreateEntityOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewEntityServiceCreateEntityDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewEntityServiceCreateEntityOK creates a EntityServiceCreateEntityOK with default headers values
func NewEntityServiceCreateEntityOK() *EntityServiceCreateEntityOK {
	return &EntityServiceCreateEntityOK{}
}

/*EntityServiceCreateEntityOK handles this case with default header values.

A successful response.
*/
type EntityServiceCreateEntityOK struct {
	Payload *entity_model.APIEntity
}

func (o *EntityServiceCreateEntityOK) Error() string {
	return fmt.Sprintf("[POST /CreateEntity][%d] entityServiceCreateEntityOK  %+v", 200, o.Payload)
}

func (o *EntityServiceCreateEntityOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(entity_model.APIEntity)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEntityServiceCreateEntityDefault creates a EntityServiceCreateEntityDefault with default headers values
func NewEntityServiceCreateEntityDefault(code int) *EntityServiceCreateEntityDefault {
	return &EntityServiceCreateEntityDefault{
		_statusCode: code,
	}
}

/*EntityServiceCreateEntityDefault handles this case with default header values.

An unexpected error response.
*/
type EntityServiceCreateEntityDefault struct {
	_statusCode int

	Payload *entity_model.GatewayruntimeError
}

// Code gets the status code for the entity service create entity default response
func (o *EntityServiceCreateEntityDefault) Code() int {
	return o._statusCode
}

func (o *EntityServiceCreateEntityDefault) Error() string {
	return fmt.Sprintf("[POST /CreateEntity][%d] EntityService_CreateEntity default  %+v", o._statusCode, o.Payload)
}

func (o *EntityServiceCreateEntityDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(entity_model.GatewayruntimeError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}