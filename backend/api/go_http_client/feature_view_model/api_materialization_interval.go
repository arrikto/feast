// Code generated by go-swagger; DO NOT EDIT.

package feature_view_model

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// APIMaterializationInterval api materialization interval
// swagger:model apiMaterializationInterval
type APIMaterializationInterval struct {

	// end time
	// Format: date-time
	EndTime strfmt.DateTime `json:"end_time,omitempty"`

	// start time
	// Format: date-time
	StartTime strfmt.DateTime `json:"start_time,omitempty"`
}

// Validate validates this api materialization interval
func (m *APIMaterializationInterval) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateEndTime(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStartTime(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *APIMaterializationInterval) validateEndTime(formats strfmt.Registry) error {

	if swag.IsZero(m.EndTime) { // not required
		return nil
	}

	if err := validate.FormatOf("end_time", "body", "date-time", m.EndTime.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *APIMaterializationInterval) validateStartTime(formats strfmt.Registry) error {

	if swag.IsZero(m.StartTime) { // not required
		return nil
	}

	if err := validate.FormatOf("start_time", "body", "date-time", m.StartTime.String(), formats); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *APIMaterializationInterval) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *APIMaterializationInterval) UnmarshalBinary(b []byte) error {
	var res APIMaterializationInterval
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}