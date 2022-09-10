// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/validate"
)

// OpenapiNodeStatus openapi node status
//
// swagger:model openapiNodeStatus
type OpenapiNodeStatus string

func NewOpenapiNodeStatus(value OpenapiNodeStatus) *OpenapiNodeStatus {
	return &value
}

// Pointer returns a pointer to a freshly-allocated OpenapiNodeStatus.
func (m OpenapiNodeStatus) Pointer() *OpenapiNodeStatus {
	return &m
}

const (

	// OpenapiNodeStatusNODESTATUSAVAILABLE captures enum value "NODE_STATUS_AVAILABLE"
	OpenapiNodeStatusNODESTATUSAVAILABLE OpenapiNodeStatus = "NODE_STATUS_AVAILABLE"

	// OpenapiNodeStatusNODESTATUSUNAVAILABLE captures enum value "NODE_STATUS_UNAVAILABLE"
	OpenapiNodeStatusNODESTATUSUNAVAILABLE OpenapiNodeStatus = "NODE_STATUS_UNAVAILABLE"

	// OpenapiNodeStatusNODESTATUSCREATING captures enum value "NODE_STATUS_CREATING"
	OpenapiNodeStatusNODESTATUSCREATING OpenapiNodeStatus = "NODE_STATUS_CREATING"

	// OpenapiNodeStatusNODESTATUSDELETING captures enum value "NODE_STATUS_DELETING"
	OpenapiNodeStatusNODESTATUSDELETING OpenapiNodeStatus = "NODE_STATUS_DELETING"
)

// for schema
var openapiNodeStatusEnum []interface{}

func init() {
	var res []OpenapiNodeStatus
	if err := json.Unmarshal([]byte(`["NODE_STATUS_AVAILABLE","NODE_STATUS_UNAVAILABLE","NODE_STATUS_CREATING","NODE_STATUS_DELETING"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		openapiNodeStatusEnum = append(openapiNodeStatusEnum, v)
	}
}

func (m OpenapiNodeStatus) validateOpenapiNodeStatusEnum(path, location string, value OpenapiNodeStatus) error {
	if err := validate.EnumCase(path, location, value, openapiNodeStatusEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this openapi node status
func (m OpenapiNodeStatus) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateOpenapiNodeStatusEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this openapi node status based on context it is used
func (m OpenapiNodeStatus) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}
