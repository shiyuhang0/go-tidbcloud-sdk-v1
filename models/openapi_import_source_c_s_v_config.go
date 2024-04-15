// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// OpenapiImportSourceCSVConfig ImportSourceCSVConfig
//
// ImportSourceCSVConfig represents the settings for parsing the CSV source data of an import task.
//
// swagger:model openapiImportSourceCSVConfig
type OpenapiImportSourceCSVConfig struct {

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

// Validate validates this openapi import source c s v config
func (m *OpenapiImportSourceCSVConfig) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this openapi import source c s v config based on context it is used
func (m *OpenapiImportSourceCSVConfig) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *OpenapiImportSourceCSVConfig) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *OpenapiImportSourceCSVConfig) UnmarshalBinary(b []byte) error {
	var res OpenapiImportSourceCSVConfig
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
