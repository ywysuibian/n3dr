// Code generated by go-swagger; DO NOT EDIT.

package repository_management

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

// NewGetRepository14Params creates a new GetRepository14Params object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetRepository14Params() *GetRepository14Params {
	return &GetRepository14Params{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetRepository14ParamsWithTimeout creates a new GetRepository14Params object
// with the ability to set a timeout on a request.
func NewGetRepository14ParamsWithTimeout(timeout time.Duration) *GetRepository14Params {
	return &GetRepository14Params{
		timeout: timeout,
	}
}

// NewGetRepository14ParamsWithContext creates a new GetRepository14Params object
// with the ability to set a context for a request.
func NewGetRepository14ParamsWithContext(ctx context.Context) *GetRepository14Params {
	return &GetRepository14Params{
		Context: ctx,
	}
}

// NewGetRepository14ParamsWithHTTPClient creates a new GetRepository14Params object
// with the ability to set a custom HTTPClient for a request.
func NewGetRepository14ParamsWithHTTPClient(client *http.Client) *GetRepository14Params {
	return &GetRepository14Params{
		HTTPClient: client,
	}
}

/* GetRepository14Params contains all the parameters to send to the API endpoint
   for the get repository 14 operation.

   Typically these are written to a http.Request.
*/
type GetRepository14Params struct {

	// RepositoryName.
	RepositoryName string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get repository 14 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetRepository14Params) WithDefaults() *GetRepository14Params {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get repository 14 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetRepository14Params) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get repository 14 params
func (o *GetRepository14Params) WithTimeout(timeout time.Duration) *GetRepository14Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get repository 14 params
func (o *GetRepository14Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get repository 14 params
func (o *GetRepository14Params) WithContext(ctx context.Context) *GetRepository14Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get repository 14 params
func (o *GetRepository14Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get repository 14 params
func (o *GetRepository14Params) WithHTTPClient(client *http.Client) *GetRepository14Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get repository 14 params
func (o *GetRepository14Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithRepositoryName adds the repositoryName to the get repository 14 params
func (o *GetRepository14Params) WithRepositoryName(repositoryName string) *GetRepository14Params {
	o.SetRepositoryName(repositoryName)
	return o
}

// SetRepositoryName adds the repositoryName to the get repository 14 params
func (o *GetRepository14Params) SetRepositoryName(repositoryName string) {
	o.RepositoryName = repositoryName
}

// WriteToRequest writes these params to a swagger request
func (o *GetRepository14Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param repositoryName
	if err := r.SetPathParam("repositoryName", o.RepositoryName); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
