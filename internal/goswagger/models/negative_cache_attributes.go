// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// NegativeCacheAttributes negative cache attributes
//
// swagger:model NegativeCacheAttributes
type NegativeCacheAttributes struct {

	// Whether to cache responses for content not present in the proxied repository
	// Example: true
	// Required: true
	Enabled *bool `json:"enabled"`

	// How long to cache the fact that a file was not found in the repository (in minutes)
	// Example: 1440
	// Required: true
	TimeToLive *int32 `json:"timeToLive"`
}

// Validate validates this negative cache attributes
func (m *NegativeCacheAttributes) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateEnabled(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTimeToLive(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *NegativeCacheAttributes) validateEnabled(formats strfmt.Registry) error {

	if err := validate.Required("enabled", "body", m.Enabled); err != nil {
		return err
	}

	return nil
}

func (m *NegativeCacheAttributes) validateTimeToLive(formats strfmt.Registry) error {

	if err := validate.Required("timeToLive", "body", m.TimeToLive); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this negative cache attributes based on context it is used
func (m *NegativeCacheAttributes) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *NegativeCacheAttributes) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NegativeCacheAttributes) UnmarshalBinary(b []byte) error {
	var res NegativeCacheAttributes
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
