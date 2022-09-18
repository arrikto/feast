// Code generated by go-swagger; DO NOT EDIT.

package feature_service_model

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// APIListFeatureServicesResponse api list feature services response
// swagger:model apiListFeatureServicesResponse
type APIListFeatureServicesResponse struct {

	// feature services
	FeatureServices []*APIFeatureService `json:"feature_services"`
}

// Validate validates this api list feature services response
func (m *APIListFeatureServicesResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateFeatureServices(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *APIListFeatureServicesResponse) validateFeatureServices(formats strfmt.Registry) error {

	if swag.IsZero(m.FeatureServices) { // not required
		return nil
	}

	for i := 0; i < len(m.FeatureServices); i++ {
		if swag.IsZero(m.FeatureServices[i]) { // not required
			continue
		}

		if m.FeatureServices[i] != nil {
			if err := m.FeatureServices[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("feature_services" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *APIListFeatureServicesResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *APIListFeatureServicesResponse) UnmarshalBinary(b []byte) error {
	var res APIListFeatureServicesResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}