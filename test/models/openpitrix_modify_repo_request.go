// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// OpenpitrixModifyRepoRequest openpitrix modify repo request
// swagger:model openpitrixModifyRepoRequest
type OpenpitrixModifyRepoRequest struct {

	// category id
	CategoryID string `json:"category_id,omitempty"`

	// credential
	Credential string `json:"credential,omitempty"`

	// description
	Description string `json:"description,omitempty"`

	// labels
	Labels []*OpenpitrixRepoLabel `json:"labels"`

	// name
	Name string `json:"name,omitempty"`

	// providers
	Providers []string `json:"providers"`

	// repo id
	RepoID string `json:"repo_id,omitempty"`

	// selectors
	Selectors []*OpenpitrixRepoSelector `json:"selectors"`

	// type
	Type string `json:"type,omitempty"`

	// url
	URL string `json:"url,omitempty"`

	// visibility
	Visibility string `json:"visibility,omitempty"`
}

// Validate validates this openpitrix modify repo request
func (m *OpenpitrixModifyRepoRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLabels(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSelectors(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *OpenpitrixModifyRepoRequest) validateLabels(formats strfmt.Registry) error {

	if swag.IsZero(m.Labels) { // not required
		return nil
	}

	for i := 0; i < len(m.Labels); i++ {
		if swag.IsZero(m.Labels[i]) { // not required
			continue
		}

		if m.Labels[i] != nil {
			if err := m.Labels[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("labels" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *OpenpitrixModifyRepoRequest) validateSelectors(formats strfmt.Registry) error {

	if swag.IsZero(m.Selectors) { // not required
		return nil
	}

	for i := 0; i < len(m.Selectors); i++ {
		if swag.IsZero(m.Selectors[i]) { // not required
			continue
		}

		if m.Selectors[i] != nil {
			if err := m.Selectors[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("selectors" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *OpenpitrixModifyRepoRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *OpenpitrixModifyRepoRequest) UnmarshalBinary(b []byte) error {
	var res OpenpitrixModifyRepoRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
