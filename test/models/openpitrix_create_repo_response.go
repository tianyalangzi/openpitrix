// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// OpenpitrixCreateRepoResponse openpitrix create repo response
// swagger:model openpitrixCreateRepoResponse
type OpenpitrixCreateRepoResponse struct {

	// repo id
	RepoID string `json:"repo_id,omitempty"`
}

// Validate validates this openpitrix create repo response
func (m *OpenpitrixCreateRepoResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *OpenpitrixCreateRepoResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *OpenpitrixCreateRepoResponse) UnmarshalBinary(b []byte) error {
	var res OpenpitrixCreateRepoResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
