// Code generated by go-swagger; DO NOT EDIT.

package infra_object_model

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// APIInfraObject api infra object
// swagger:model apiInfraObject
type APIInfraObject struct {

	// The infrastructure object.
	// Protobuf object transformed to a JSON string.
	InfraObject string `json:"infra_object,omitempty"`

	// Represents the Python class for the infrastructure object.
	InfraObjectClassType string `json:"infra_object_class_type,omitempty"`
}

// Validate validates this api infra object
func (m *APIInfraObject) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *APIInfraObject) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *APIInfraObject) UnmarshalBinary(b []byte) error {
	var res APIInfraObject
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
