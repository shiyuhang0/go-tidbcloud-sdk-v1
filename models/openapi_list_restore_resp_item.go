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

// OpenapiListRestoreRespItem ListRestoreRespItem
//
// The items of all restore tasks.
//
// swagger:model openapiListRestoreRespItem
type OpenapiListRestoreRespItem struct {

	// The ID of the backup.
	// Example: 1
	BackupID string `json:"backup_id,omitempty"`

	// cluster
	Cluster *OpenapiListRestoreRespItemCluster `json:"cluster,omitempty"`

	// The cluster ID of the backup.
	// Example: 1
	ClusterID string `json:"cluster_id,omitempty"`

	// The creation time of the backup in UTC.
	//
	// The time format follows the [ISO8601](http://en.wikipedia.org/wiki/ISO_8601) standard, which is `YYYY-MM-DD` (year-month-day) + T +`HH:MM:SS` (hour-minutes-seconds) + Z. For example, `2020-01-01T00:00:00Z`.
	// Example: 2020-01-01T00:00:00Z
	// Format: date-time
	CreateTimestamp strfmt.DateTime `json:"create_timestamp,omitempty"`

	// The error message of restore if failed.
	// Example: Please contact support.
	ErrorMessage string `json:"error_message,omitempty"`

	// The ID of the restore task.
	// Example: 1
	ID string `json:"id,omitempty"`

	// The status of the restore task.
	// Example: PENDING
	// Enum: [PENDING RUNNING FAILED SUCCESS]
	Status string `json:"status,omitempty"`
}

// Validate validates this openapi list restore resp item
func (m *OpenapiListRestoreRespItem) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCluster(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCreateTimestamp(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *OpenapiListRestoreRespItem) validateCluster(formats strfmt.Registry) error {
	if swag.IsZero(m.Cluster) { // not required
		return nil
	}

	if m.Cluster != nil {
		if err := m.Cluster.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("cluster")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("cluster")
			}
			return err
		}
	}

	return nil
}

func (m *OpenapiListRestoreRespItem) validateCreateTimestamp(formats strfmt.Registry) error {
	if swag.IsZero(m.CreateTimestamp) { // not required
		return nil
	}

	if err := validate.FormatOf("create_timestamp", "body", "date-time", m.CreateTimestamp.String(), formats); err != nil {
		return err
	}

	return nil
}

var openapiListRestoreRespItemTypeStatusPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["PENDING","RUNNING","FAILED","SUCCESS"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		openapiListRestoreRespItemTypeStatusPropEnum = append(openapiListRestoreRespItemTypeStatusPropEnum, v)
	}
}

const (

	// OpenapiListRestoreRespItemStatusPENDING captures enum value "PENDING"
	OpenapiListRestoreRespItemStatusPENDING string = "PENDING"

	// OpenapiListRestoreRespItemStatusRUNNING captures enum value "RUNNING"
	OpenapiListRestoreRespItemStatusRUNNING string = "RUNNING"

	// OpenapiListRestoreRespItemStatusFAILED captures enum value "FAILED"
	OpenapiListRestoreRespItemStatusFAILED string = "FAILED"

	// OpenapiListRestoreRespItemStatusSUCCESS captures enum value "SUCCESS"
	OpenapiListRestoreRespItemStatusSUCCESS string = "SUCCESS"
)

// prop value enum
func (m *OpenapiListRestoreRespItem) validateStatusEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, openapiListRestoreRespItemTypeStatusPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *OpenapiListRestoreRespItem) validateStatus(formats strfmt.Registry) error {
	if swag.IsZero(m.Status) { // not required
		return nil
	}

	// value enum
	if err := m.validateStatusEnum("status", "body", m.Status); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this openapi list restore resp item based on the context it is used
func (m *OpenapiListRestoreRespItem) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateCluster(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *OpenapiListRestoreRespItem) contextValidateCluster(ctx context.Context, formats strfmt.Registry) error {

	if m.Cluster != nil {
		if err := m.Cluster.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("cluster")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("cluster")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *OpenapiListRestoreRespItem) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *OpenapiListRestoreRespItem) UnmarshalBinary(b []byte) error {
	var res OpenapiListRestoreRespItem
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// OpenapiListRestoreRespItemCluster ClusterInfoOfRestore
//
// The information of the restored cluster. The restored cluster is the new cluster your backup data is restored to.
//
// swagger:model OpenapiListRestoreRespItemCluster
type OpenapiListRestoreRespItemCluster struct {

	// The ID of the restored cluster. The restored cluster is the new cluster your backup data is restored to.
	// Example: 1
	ID string `json:"id,omitempty"`

	// The name of the restored cluster. The restored cluster is the new cluster your backup data is restored to.
	// Example: cluster-1
	Name string `json:"name,omitempty"`

	// The status of the restored cluster. Possible values are `"AVAILABLE"`, `"CREATING"`, `"MODIFYING"`, `"PAUSED"`, `"RESUMING"`, `"UNAVAILABLE"`, `"IMPORTING"`, and `"CLEARED"`.
	// Example: AVAILABLE
	Status string `json:"status,omitempty"`
}

// Validate validates this openapi list restore resp item cluster
func (m *OpenapiListRestoreRespItemCluster) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this openapi list restore resp item cluster based on context it is used
func (m *OpenapiListRestoreRespItemCluster) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *OpenapiListRestoreRespItemCluster) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *OpenapiListRestoreRespItemCluster) UnmarshalBinary(b []byte) error {
	var res OpenapiListRestoreRespItemCluster
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
