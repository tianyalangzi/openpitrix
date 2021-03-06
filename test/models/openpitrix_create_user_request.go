// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// OpenpitrixCreateUserRequest openpitrix create user request
// swagger:model openpitrixCreateUserRequest
type OpenpitrixCreateUserRequest struct {

	// description
	Description string `json:"description,omitempty"`

	// email
	Email string `json:"email,omitempty"`

	// password
	Password string `json:"password,omitempty"`

	// phone number
	PhoneNumber string `json:"phone_number,omitempty"`

	// role id
	RoleID string `json:"role_id,omitempty"`
}

// Validate validates this openpitrix create user request
func (m *OpenpitrixCreateUserRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *OpenpitrixCreateUserRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *OpenpitrixCreateUserRequest) UnmarshalBinary(b []byte) error {
	var res OpenpitrixCreateUserRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
