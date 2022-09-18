// Code generated by go-swagger; DO NOT EDIT.

package feature_view_model

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/validate"
)

// ValueTypeEnum value type enum
// swagger:model ValueTypeEnum
type ValueTypeEnum string

const (

	// ValueTypeEnumINVALID captures enum value "INVALID"
	ValueTypeEnumINVALID ValueTypeEnum = "INVALID"

	// ValueTypeEnumBYTES captures enum value "BYTES"
	ValueTypeEnumBYTES ValueTypeEnum = "BYTES"

	// ValueTypeEnumSTRING captures enum value "STRING"
	ValueTypeEnumSTRING ValueTypeEnum = "STRING"

	// ValueTypeEnumINT32 captures enum value "INT32"
	ValueTypeEnumINT32 ValueTypeEnum = "INT32"

	// ValueTypeEnumINT64 captures enum value "INT64"
	ValueTypeEnumINT64 ValueTypeEnum = "INT64"

	// ValueTypeEnumDOUBLE captures enum value "DOUBLE"
	ValueTypeEnumDOUBLE ValueTypeEnum = "DOUBLE"

	// ValueTypeEnumFLOAT captures enum value "FLOAT"
	ValueTypeEnumFLOAT ValueTypeEnum = "FLOAT"

	// ValueTypeEnumBOOL captures enum value "BOOL"
	ValueTypeEnumBOOL ValueTypeEnum = "BOOL"

	// ValueTypeEnumUNIXTIMESTAMP captures enum value "UNIX_TIMESTAMP"
	ValueTypeEnumUNIXTIMESTAMP ValueTypeEnum = "UNIX_TIMESTAMP"

	// ValueTypeEnumBYTESLIST captures enum value "BYTES_LIST"
	ValueTypeEnumBYTESLIST ValueTypeEnum = "BYTES_LIST"

	// ValueTypeEnumSTRINGLIST captures enum value "STRING_LIST"
	ValueTypeEnumSTRINGLIST ValueTypeEnum = "STRING_LIST"

	// ValueTypeEnumINT32LIST captures enum value "INT32_LIST"
	ValueTypeEnumINT32LIST ValueTypeEnum = "INT32_LIST"

	// ValueTypeEnumINT64LIST captures enum value "INT64_LIST"
	ValueTypeEnumINT64LIST ValueTypeEnum = "INT64_LIST"

	// ValueTypeEnumDOUBLELIST captures enum value "DOUBLE_LIST"
	ValueTypeEnumDOUBLELIST ValueTypeEnum = "DOUBLE_LIST"

	// ValueTypeEnumFLOATLIST captures enum value "FLOAT_LIST"
	ValueTypeEnumFLOATLIST ValueTypeEnum = "FLOAT_LIST"

	// ValueTypeEnumBOOLLIST captures enum value "BOOL_LIST"
	ValueTypeEnumBOOLLIST ValueTypeEnum = "BOOL_LIST"

	// ValueTypeEnumUNIXTIMESTAMPLIST captures enum value "UNIX_TIMESTAMP_LIST"
	ValueTypeEnumUNIXTIMESTAMPLIST ValueTypeEnum = "UNIX_TIMESTAMP_LIST"

	// ValueTypeEnumNULL captures enum value "NULL"
	ValueTypeEnumNULL ValueTypeEnum = "NULL"
)

// for schema
var valueTypeEnumEnum []interface{}

func init() {
	var res []ValueTypeEnum
	if err := json.Unmarshal([]byte(`["INVALID","BYTES","STRING","INT32","INT64","DOUBLE","FLOAT","BOOL","UNIX_TIMESTAMP","BYTES_LIST","STRING_LIST","INT32_LIST","INT64_LIST","DOUBLE_LIST","FLOAT_LIST","BOOL_LIST","UNIX_TIMESTAMP_LIST","NULL"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		valueTypeEnumEnum = append(valueTypeEnumEnum, v)
	}
}

func (m ValueTypeEnum) validateValueTypeEnumEnum(path, location string, value ValueTypeEnum) error {
	if err := validate.Enum(path, location, value, valueTypeEnumEnum); err != nil {
		return err
	}
	return nil
}

// Validate validates this value type enum
func (m ValueTypeEnum) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateValueTypeEnumEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}