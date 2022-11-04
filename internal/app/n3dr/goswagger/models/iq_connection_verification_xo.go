// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// IqConnectionVerificationXo iq connection verification xo
//
// swagger:model IqConnectionVerificationXo
type IqConnectionVerificationXo struct {

	// reason
	Reason string `json:"reason,omitempty"`

	// success
	Success bool `json:"success,omitempty"`
}

// Validate validates this iq connection verification xo
func (m *IqConnectionVerificationXo) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this iq connection verification xo based on context it is used
func (m *IqConnectionVerificationXo) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *IqConnectionVerificationXo) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *IqConnectionVerificationXo) UnmarshalBinary(b []byte) error {
	var res IqConnectionVerificationXo
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}