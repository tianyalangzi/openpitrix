// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// OpenpitrixRuntime openpitrix runtime
// swagger:model openpitrixRuntime
type OpenpitrixRuntime struct {

	// create time
	CreateTime strfmt.DateTime `json:"create_time,omitempty"`

	// description
	Description string `json:"description,omitempty"`

	// labels
	Labels OpenpitrixRuntimeLabels `json:"labels"`

	// name
	Name string `json:"name,omitempty"`

	// owner
	Owner string `json:"owner,omitempty"`

	// provider
	Provider string `json:"provider,omitempty"`

	// runtime credential
	RuntimeCredential string `json:"runtime_credential,omitempty"`

	// runtime id
	RuntimeID string `json:"runtime_id,omitempty"`

	// runtime url
	RuntimeURL string `json:"runtime_url,omitempty"`

	// status
	Status string `json:"status,omitempty"`

	// status time
	StatusTime strfmt.DateTime `json:"status_time,omitempty"`

	// zone
	Zone string `json:"zone,omitempty"`
}

// Validate validates this openpitrix runtime
func (m *OpenpitrixRuntime) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *OpenpitrixRuntime) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *OpenpitrixRuntime) UnmarshalBinary(b []byte) error {
	var res OpenpitrixRuntime
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}