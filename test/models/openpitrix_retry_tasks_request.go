// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// OpenpitrixRetryTasksRequest openpitrix retry tasks request
// swagger:model openpitrixRetryTasksRequest
type OpenpitrixRetryTasksRequest struct {

	// task id
	TaskID []string `json:"task_id"`
}

// Validate validates this openpitrix retry tasks request
func (m *OpenpitrixRetryTasksRequest) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *OpenpitrixRetryTasksRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *OpenpitrixRetryTasksRequest) UnmarshalBinary(b []byte) error {
	var res OpenpitrixRetryTasksRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
