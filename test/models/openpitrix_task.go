// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// OpenpitrixTask openpitrix task
// swagger:model openpitrixTask
type OpenpitrixTask struct {

	// create time
	// Format: date-time
	CreateTime strfmt.DateTime `json:"create_time,omitempty"`

	// directive
	Directive string `json:"directive,omitempty"`

	// error code
	ErrorCode *ProtobufUint32Value `json:"error_code,omitempty"`

	// executor
	Executor string `json:"executor,omitempty"`

	// failure allowed
	FailureAllowed bool `json:"failure_allowed,omitempty"`

	// job id
	JobID string `json:"job_id,omitempty"`

	// node id
	NodeID string `json:"node_id,omitempty"`

	// owner
	Owner string `json:"owner,omitempty"`

	// status
	Status string `json:"status,omitempty"`

	// status time
	// Format: date-time
	StatusTime strfmt.DateTime `json:"status_time,omitempty"`

	// target
	Target string `json:"target,omitempty"`

	// task action
	TaskAction string `json:"task_action,omitempty"`

	// task id
	TaskID string `json:"task_id,omitempty"`
}

// Validate validates this openpitrix task
func (m *OpenpitrixTask) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCreateTime(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateErrorCode(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStatusTime(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *OpenpitrixTask) validateCreateTime(formats strfmt.Registry) error {

	if swag.IsZero(m.CreateTime) { // not required
		return nil
	}

	if err := validate.FormatOf("create_time", "body", "date-time", m.CreateTime.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *OpenpitrixTask) validateErrorCode(formats strfmt.Registry) error {

	if swag.IsZero(m.ErrorCode) { // not required
		return nil
	}

	if m.ErrorCode != nil {
		if err := m.ErrorCode.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("error_code")
			}
			return err
		}
	}

	return nil
}

func (m *OpenpitrixTask) validateStatusTime(formats strfmt.Registry) error {

	if swag.IsZero(m.StatusTime) { // not required
		return nil
	}

	if err := validate.FormatOf("status_time", "body", "date-time", m.StatusTime.String(), formats); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *OpenpitrixTask) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *OpenpitrixTask) UnmarshalBinary(b []byte) error {
	var res OpenpitrixTask
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
