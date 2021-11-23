// Code generated by go-swagger; DO NOT EDIT.

package read_only

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

// NewFreezeParams creates a new FreezeParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewFreezeParams() *FreezeParams {
	return &FreezeParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewFreezeParamsWithTimeout creates a new FreezeParams object
// with the ability to set a timeout on a request.
func NewFreezeParamsWithTimeout(timeout time.Duration) *FreezeParams {
	return &FreezeParams{
		timeout: timeout,
	}
}

// NewFreezeParamsWithContext creates a new FreezeParams object
// with the ability to set a context for a request.
func NewFreezeParamsWithContext(ctx context.Context) *FreezeParams {
	return &FreezeParams{
		Context: ctx,
	}
}

// NewFreezeParamsWithHTTPClient creates a new FreezeParams object
// with the ability to set a custom HTTPClient for a request.
func NewFreezeParamsWithHTTPClient(client *http.Client) *FreezeParams {
	return &FreezeParams{
		HTTPClient: client,
	}
}

/* FreezeParams contains all the parameters to send to the API endpoint
   for the freeze operation.

   Typically these are written to a http.Request.
*/
type FreezeParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the freeze params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *FreezeParams) WithDefaults() *FreezeParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the freeze params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *FreezeParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the freeze params
func (o *FreezeParams) WithTimeout(timeout time.Duration) *FreezeParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the freeze params
func (o *FreezeParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the freeze params
func (o *FreezeParams) WithContext(ctx context.Context) *FreezeParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the freeze params
func (o *FreezeParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the freeze params
func (o *FreezeParams) WithHTTPClient(client *http.Client) *FreezeParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the freeze params
func (o *FreezeParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *FreezeParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
