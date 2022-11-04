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

// TaskXO task x o
//
// swagger:model TaskXO
type TaskXO struct {

	// current state
	CurrentState string `json:"currentState,omitempty"`

	// id
	ID string `json:"id,omitempty"`

	// last run
	// Format: date-time
	LastRun strfmt.DateTime `json:"lastRun,omitempty"`

	// last run result
	LastRunResult string `json:"lastRunResult,omitempty"`

	// message
	Message string `json:"message,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// next run
	// Format: date-time
	NextRun strfmt.DateTime `json:"nextRun,omitempty"`

	// type
	Type string `json:"type,omitempty"`
}

// Validate validates this task x o
func (m *TaskXO) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLastRun(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNextRun(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TaskXO) validateLastRun(formats strfmt.Registry) error {
	if swag.IsZero(m.LastRun) { // not required
		return nil
	}

	if err := validate.FormatOf("lastRun", "body", "date-time", m.LastRun.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *TaskXO) validateNextRun(formats strfmt.Registry) error {
	if swag.IsZero(m.NextRun) { // not required
		return nil
	}

	if err := validate.FormatOf("nextRun", "body", "date-time", m.NextRun.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this task x o based on context it is used
func (m *TaskXO) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *TaskXO) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TaskXO) UnmarshalBinary(b []byte) error {
	var res TaskXO
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}