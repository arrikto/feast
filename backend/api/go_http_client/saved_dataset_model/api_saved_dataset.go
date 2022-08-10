// Code generated by go-swagger; DO NOT EDIT.

package saved_dataset_model

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// APISavedDataset api saved dataset
// swagger:model apiSavedDataset
type APISavedDataset struct {

	// Creation time of the saved dataset.
	// Format: date-time
	CreatedTimestamp strfmt.DateTime `json:"created_timestamp,omitempty"`

	// Optional and only populated if generated from a feature service fetch.
	FeatureServiceName string `json:"feature_service_name,omitempty"`

	// List of feature references with format "<view name>:<feature name>".
	Features []string `json:"features"`

	// Whether full feature names are used in stored data.
	FullFeatureNames bool `json:"full_feature_names,omitempty"`

	// Entity columns + request columns from all feature views used during retrieval.
	JoinKeys []string `json:"join_keys"`

	// Last update time of the saved dataset.
	// Format: date-time
	LastUpdatedTimestamp strfmt.DateTime `json:"last_updated_timestamp,omitempty"`

	// Max timestamp in the dataset (needed for retrieval).
	// Format: date-time
	MaxEventTimestamp strfmt.DateTime `json:"max_event_timestamp,omitempty"`

	// Min timestamp in the dataset (needed for retrieval).
	// Format: date-time
	MinEventTimestamp strfmt.DateTime `json:"min_event_timestamp,omitempty"`

	// Name of the dataset. Must be unique since it's possible to overwrite dataset by name.
	Name string `json:"name,omitempty"`

	// Name of Feast project that this dataset belongs to.
	Project string `json:"project,omitempty"`

	// Storage location of the saved dataset.
	// Protobuf object transformed to a JSON string.
	Storage string `json:"storage,omitempty"`

	// User defined metadata.
	Tags map[string]string `json:"tags,omitempty"`
}

// Validate validates this api saved dataset
func (m *APISavedDataset) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCreatedTimestamp(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLastUpdatedTimestamp(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMaxEventTimestamp(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMinEventTimestamp(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *APISavedDataset) validateCreatedTimestamp(formats strfmt.Registry) error {

	if swag.IsZero(m.CreatedTimestamp) { // not required
		return nil
	}

	if err := validate.FormatOf("created_timestamp", "body", "date-time", m.CreatedTimestamp.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *APISavedDataset) validateLastUpdatedTimestamp(formats strfmt.Registry) error {

	if swag.IsZero(m.LastUpdatedTimestamp) { // not required
		return nil
	}

	if err := validate.FormatOf("last_updated_timestamp", "body", "date-time", m.LastUpdatedTimestamp.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *APISavedDataset) validateMaxEventTimestamp(formats strfmt.Registry) error {

	if swag.IsZero(m.MaxEventTimestamp) { // not required
		return nil
	}

	if err := validate.FormatOf("max_event_timestamp", "body", "date-time", m.MaxEventTimestamp.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *APISavedDataset) validateMinEventTimestamp(formats strfmt.Registry) error {

	if swag.IsZero(m.MinEventTimestamp) { // not required
		return nil
	}

	if err := validate.FormatOf("min_event_timestamp", "body", "date-time", m.MinEventTimestamp.String(), formats); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *APISavedDataset) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *APISavedDataset) UnmarshalBinary(b []byte) error {
	var res APISavedDataset
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
