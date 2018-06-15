// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// OpenpitrixResizeClusterRequest openpitrix resize cluster request
// swagger:model openpitrixResizeClusterRequest
type OpenpitrixResizeClusterRequest struct {

	// advanced param
	AdvancedParam []string `json:"advanced_param"`

	// cluster id
	ClusterID string `json:"cluster_id,omitempty"`

	// cpu
	CPU *ProtobufUint32Value `json:"cpu,omitempty"`

	// memory
	Memory *ProtobufUint32Value `json:"memory,omitempty"`

	// role
	Role string `json:"role,omitempty"`
}

// Validate validates this openpitrix resize cluster request
func (m *OpenpitrixResizeClusterRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCPU(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMemory(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *OpenpitrixResizeClusterRequest) validateCPU(formats strfmt.Registry) error {

	if swag.IsZero(m.CPU) { // not required
		return nil
	}

	if m.CPU != nil {
		if err := m.CPU.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("cpu")
			}
			return err
		}
	}

	return nil
}

func (m *OpenpitrixResizeClusterRequest) validateMemory(formats strfmt.Registry) error {

	if swag.IsZero(m.Memory) { // not required
		return nil
	}

	if m.Memory != nil {
		if err := m.Memory.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("memory")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *OpenpitrixResizeClusterRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *OpenpitrixResizeClusterRequest) UnmarshalBinary(b []byte) error {
	var res OpenpitrixResizeClusterRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
