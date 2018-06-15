// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// OpenpitrixDeleteClusterNodesRequest openpitrix delete cluster nodes request
// swagger:model openpitrixDeleteClusterNodesRequest
type OpenpitrixDeleteClusterNodesRequest struct {

	// advanced param
	AdvancedParam []string `json:"advanced_param"`

	// cluster id
	ClusterID string `json:"cluster_id,omitempty"`

	// node id
	NodeID []string `json:"node_id"`

	// role
	Role string `json:"role,omitempty"`
}

// Validate validates this openpitrix delete cluster nodes request
func (m *OpenpitrixDeleteClusterNodesRequest) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *OpenpitrixDeleteClusterNodesRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *OpenpitrixDeleteClusterNodesRequest) UnmarshalBinary(b []byte) error {
	var res OpenpitrixDeleteClusterNodesRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
