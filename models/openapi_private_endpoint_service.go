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

// OpenapiPrivateEndpointService openapi private endpoint service
//
// swagger:model openapiPrivateEndpointService
type OpenapiPrivateEndpointService struct {

	// Availability zones for the private endpoint service. This field is only applicable when the `cloud_provider` is `"AWS"`.
	// Example: ["usw2-az2","usw2-az1"]
	AzIds []string `json:"az_ids"`

	// The cloud provider on which the private endpoint service is hosted.
	// - `"AWS"`: the Amazon Web Services cloud provider
	// - `"GCP"`: the Google Cloud cloud provider
	// Example: AWS
	// Enum: [AWS GCP]
	CloudProvider string `json:"cloud_provider,omitempty"`

	// The DNS name of the private endpoint service.
	// Example: privatelink-tidb.01234567890.clusters.tidb-cloud.com
	DNSName string `json:"dns_name,omitempty"`

	// The name of the private endpoint service, which is used for connection.
	// Example: com.amazonaws.vpce.us-west-2.vpce-svc-01234567890123456
	Name string `json:"name,omitempty"`

	// The port of the private endpoint service.
	// Example: 4000
	Port int32 `json:"port,omitempty"`

	// The status of the private endpoint service.
	// Example: ACTIVE
	// Enum: [CREATING ACTIVE DELETING]
	Status string `json:"status,omitempty"`
}

// Validate validates this openapi private endpoint service
func (m *OpenapiPrivateEndpointService) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCloudProvider(formats); err != nil {
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

var openapiPrivateEndpointServiceTypeCloudProviderPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["AWS","GCP"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		openapiPrivateEndpointServiceTypeCloudProviderPropEnum = append(openapiPrivateEndpointServiceTypeCloudProviderPropEnum, v)
	}
}

const (

	// OpenapiPrivateEndpointServiceCloudProviderAWS captures enum value "AWS"
	OpenapiPrivateEndpointServiceCloudProviderAWS string = "AWS"

	// OpenapiPrivateEndpointServiceCloudProviderGCP captures enum value "GCP"
	OpenapiPrivateEndpointServiceCloudProviderGCP string = "GCP"
)

// prop value enum
func (m *OpenapiPrivateEndpointService) validateCloudProviderEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, openapiPrivateEndpointServiceTypeCloudProviderPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *OpenapiPrivateEndpointService) validateCloudProvider(formats strfmt.Registry) error {
	if swag.IsZero(m.CloudProvider) { // not required
		return nil
	}

	// value enum
	if err := m.validateCloudProviderEnum("cloud_provider", "body", m.CloudProvider); err != nil {
		return err
	}

	return nil
}

var openapiPrivateEndpointServiceTypeStatusPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["CREATING","ACTIVE","DELETING"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		openapiPrivateEndpointServiceTypeStatusPropEnum = append(openapiPrivateEndpointServiceTypeStatusPropEnum, v)
	}
}

const (

	// OpenapiPrivateEndpointServiceStatusCREATING captures enum value "CREATING"
	OpenapiPrivateEndpointServiceStatusCREATING string = "CREATING"

	// OpenapiPrivateEndpointServiceStatusACTIVE captures enum value "ACTIVE"
	OpenapiPrivateEndpointServiceStatusACTIVE string = "ACTIVE"

	// OpenapiPrivateEndpointServiceStatusDELETING captures enum value "DELETING"
	OpenapiPrivateEndpointServiceStatusDELETING string = "DELETING"
)

// prop value enum
func (m *OpenapiPrivateEndpointService) validateStatusEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, openapiPrivateEndpointServiceTypeStatusPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *OpenapiPrivateEndpointService) validateStatus(formats strfmt.Registry) error {
	if swag.IsZero(m.Status) { // not required
		return nil
	}

	// value enum
	if err := m.validateStatusEnum("status", "body", m.Status); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this openapi private endpoint service based on context it is used
func (m *OpenapiPrivateEndpointService) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *OpenapiPrivateEndpointService) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *OpenapiPrivateEndpointService) UnmarshalBinary(b []byte) error {
	var res OpenapiPrivateEndpointService
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
