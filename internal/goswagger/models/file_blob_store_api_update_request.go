// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// FileBlobStoreAPIUpdateRequest file blob store Api update request
//
// swagger:model FileBlobStoreApiUpdateRequest
type FileBlobStoreAPIUpdateRequest struct {

	// The path to the blobstore contents. This can be an absolute path to anywhere on the system Nexus Repository Manager has access to or it can be a path relative to the sonatype-work directory.
	Path string `json:"path,omitempty"`

	// Settings to control the soft quota
	SoftQuota *BlobStoreAPISoftQuota `json:"softQuota,omitempty"`
}

// Validate validates this file blob store Api update request
func (m *FileBlobStoreAPIUpdateRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSoftQuota(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *FileBlobStoreAPIUpdateRequest) validateSoftQuota(formats strfmt.Registry) error {
	if swag.IsZero(m.SoftQuota) { // not required
		return nil
	}

	if m.SoftQuota != nil {
		if err := m.SoftQuota.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("softQuota")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("softQuota")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this file blob store Api update request based on the context it is used
func (m *FileBlobStoreAPIUpdateRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateSoftQuota(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *FileBlobStoreAPIUpdateRequest) contextValidateSoftQuota(ctx context.Context, formats strfmt.Registry) error {

	if m.SoftQuota != nil {
		if err := m.SoftQuota.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("softQuota")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("softQuota")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *FileBlobStoreAPIUpdateRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *FileBlobStoreAPIUpdateRequest) UnmarshalBinary(b []byte) error {
	var res FileBlobStoreAPIUpdateRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
