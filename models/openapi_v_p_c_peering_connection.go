// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// OpenapiVPCPeeringConnection openapi v p c peering connection
//
// swagger:model openapiVPCPeeringConnection
type OpenapiVPCPeeringConnection struct {

	// The host of VPC peering connection.
	// Example: private-tidb.f69f3808.acea1f2a.us-east-1.shared.aws.tidbcloud.com
	Host string `json:"host,omitempty"`

	// The TiDB port for connection. The port must be in the range of 1024-65535 except 10080.
	//
	// **Limitations**:
	// - For a Developer Tier cluster, only port `4000` is available.
	// Example: 4000
	// Maximum: 65535
	// Minimum: 1024
	Port int32 `json:"port,omitempty"`
}

// Validate validates this openapi v p c peering connection
func (m *OpenapiVPCPeeringConnection) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validatePort(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *OpenapiVPCPeeringConnection) validatePort(formats strfmt.Registry) error {
	if swag.IsZero(m.Port) { // not required
		return nil
	}

	if err := validate.MinimumInt("port", "body", int64(m.Port), 1024, false); err != nil {
		return err
	}

	if err := validate.MaximumInt("port", "body", int64(m.Port), 65535, false); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this openapi v p c peering connection based on context it is used
func (m *OpenapiVPCPeeringConnection) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *OpenapiVPCPeeringConnection) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *OpenapiVPCPeeringConnection) UnmarshalBinary(b []byte) error {
	var res OpenapiVPCPeeringConnection
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
