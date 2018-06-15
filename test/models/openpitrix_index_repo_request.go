// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// OpenpitrixIndexRepoRequest openpitrix index repo request
// swagger:model openpitrixIndexRepoRequest
type OpenpitrixIndexRepoRequest struct {

	// repo id
	RepoID string `json:"repo_id,omitempty"`
}

// Validate validates this openpitrix index repo request
func (m *OpenpitrixIndexRepoRequest) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *OpenpitrixIndexRepoRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *OpenpitrixIndexRepoRequest) UnmarshalBinary(b []byte) error {
	var res OpenpitrixIndexRepoRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
