// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// OpenapiImportSourceFormat ImportSourceFormat
//
// ImportSourceFormat represents the settings for reading and parsing source data in the correct format.
//
// swagger:model openapiImportSourceFormat
type OpenapiImportSourceFormat struct {

	// csv config
	CsvConfig *OpenapiImportSourceFormatCsvConfig `json:"csv_config,omitempty"`

	// The format type of an import source.
	// Example: CSV
	// Required: true
	// Enum: [CSV PARQUET SQL AURORA_SNAPSHOT]
	Type *string `json:"type"`
}

// Validate validates this openapi import source format
func (m *OpenapiImportSourceFormat) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCsvConfig(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *OpenapiImportSourceFormat) validateCsvConfig(formats strfmt.Registry) error {
	if swag.IsZero(m.CsvConfig) { // not required
		return nil
	}

	if m.CsvConfig != nil {
		if err := m.CsvConfig.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("csv_config")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("csv_config")
			}
			return err
		}
	}

	return nil
}

var openapiImportSourceFormatTypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["CSV","PARQUET","SQL","AURORA_SNAPSHOT"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		openapiImportSourceFormatTypeTypePropEnum = append(openapiImportSourceFormatTypeTypePropEnum, v)
	}
}

const (

	// OpenapiImportSourceFormatTypeCSV captures enum value "CSV"
	OpenapiImportSourceFormatTypeCSV string = "CSV"

	// OpenapiImportSourceFormatTypePARQUET captures enum value "PARQUET"
	OpenapiImportSourceFormatTypePARQUET string = "PARQUET"

	// OpenapiImportSourceFormatTypeSQL captures enum value "SQL"
	OpenapiImportSourceFormatTypeSQL string = "SQL"

	// OpenapiImportSourceFormatTypeAURORASNAPSHOT captures enum value "AURORA_SNAPSHOT"
	OpenapiImportSourceFormatTypeAURORASNAPSHOT string = "AURORA_SNAPSHOT"
)

// prop value enum
func (m *OpenapiImportSourceFormat) validateTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, openapiImportSourceFormatTypeTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *OpenapiImportSourceFormat) validateType(formats strfmt.Registry) error {

	if err := validate.Required("type", "body", m.Type); err != nil {
		return err
	}

	// value enum
	if err := m.validateTypeEnum("type", "body", *m.Type); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this openapi import source format based on the context it is used
func (m *OpenapiImportSourceFormat) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateCsvConfig(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *OpenapiImportSourceFormat) contextValidateCsvConfig(ctx context.Context, formats strfmt.Registry) error {

	if m.CsvConfig != nil {

		if swag.IsZero(m.CsvConfig) { // not required
			return nil
		}

		if err := m.CsvConfig.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("csv_config")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("csv_config")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *OpenapiImportSourceFormat) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *OpenapiImportSourceFormat) UnmarshalBinary(b []byte) error {
	var res OpenapiImportSourceFormat
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// OpenapiImportSourceFormatCsvConfig ImportSourceCSVConfig
//
// The CSV format settings to parse the source CSV files. This field is only needed if the source format is CSV.
//
// swagger:model OpenapiImportSourceFormatCsvConfig
type OpenapiImportSourceFormatCsvConfig struct {

	// Whether a backslash (`\`) symbol followed by a character should be combined as a whole and treated as an escape sequence in a CSV field. For example, if this parameter is set to `true`, `\n` will be treated as a 'new-line' character. If it is set to `false`, `\n` will be treated as two separate characters: backslash and `n`.
	//
	// Currently, these are several supported escape sequences: `\0`, `\b`, `\n`, `\r`, `\t`, and `\Z`. If the parameter is set to `true`, but the backslash escape sequence is not recognized, the backslash character is ignored.
	BackslashEscape *bool `json:"backslash_escape,omitempty"`

	// The delimiter character used to separate fields in the CSV data.
	Delimiter *string `json:"delimiter,omitempty"`

	// Whether the CSV data has a header row, which is not part of the data. If it is set to `true`, the import task will use the column names in the header row to match the column names in the target table.
	HasHeaderRow *bool `json:"has_header_row,omitempty"`

	// The character used to quote the fields in the CSV data.
	Quote *string `json:"quote,omitempty"`
}

// Validate validates this openapi import source format csv config
func (m *OpenapiImportSourceFormatCsvConfig) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this openapi import source format csv config based on context it is used
func (m *OpenapiImportSourceFormatCsvConfig) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *OpenapiImportSourceFormatCsvConfig) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *OpenapiImportSourceFormatCsvConfig) UnmarshalBinary(b []byte) error {
	var res OpenapiImportSourceFormatCsvConfig
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
