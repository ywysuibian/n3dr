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

// NewEnableRepositoryHealthCheckParams creates a new EnableRepositoryHealthCheckParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewEnableRepositoryHealthCheckParams() *EnableRepositoryHealthCheckParams {
	return &EnableRepositoryHealthCheckParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewEnableRepositoryHealthCheckParamsWithTimeout creates a new EnableRepositoryHealthCheckParams object
// with the ability to set a timeout on a request.
func NewEnableRepositoryHealthCheckParamsWithTimeout(timeout time.Duration) *EnableRepositoryHealthCheckParams {
	return &EnableRepositoryHealthCheckParams{
		timeout: timeout,
	}
}

// NewEnableRepositoryHealthCheckParamsWithContext creates a new EnableRepositoryHealthCheckParams object
// with the ability to set a context for a request.
func NewEnableRepositoryHealthCheckParamsWithContext(ctx context.Context) *EnableRepositoryHealthCheckParams {
	return &EnableRepositoryHealthCheckParams{
		Context: ctx,
	}
}

// NewEnableRepositoryHealthCheckParamsWithHTTPClient creates a new EnableRepositoryHealthCheckParams object
// with the ability to set a custom HTTPClient for a request.
func NewEnableRepositoryHealthCheckParamsWithHTTPClient(client *http.Client) *EnableRepositoryHealthCheckParams {
	return &EnableRepositoryHealthCheckParams{
		HTTPClient: client,
	}
}

/* EnableRepositoryHealthCheckParams contains all the parameters to send to the API endpoint
   for the enable repository health check operation.

   Typically these are written to a http.Request.
*/
type EnableRepositoryHealthCheckParams struct {

	/* RepositoryName.

	   Name of the repository to enable Repository Health Check for
	*/
	RepositoryName string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the enable repository health check params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *EnableRepositoryHealthCheckParams) WithDefaults() *EnableRepositoryHealthCheckParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the enable repository health check params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *EnableRepositoryHealthCheckParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the enable repository health check params
func (o *EnableRepositoryHealthCheckParams) WithTimeout(timeout time.Duration) *EnableRepositoryHealthCheckParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the enable repository health check params
func (o *EnableRepositoryHealthCheckParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the enable repository health check params
func (o *EnableRepositoryHealthCheckParams) WithContext(ctx context.Context) *EnableRepositoryHealthCheckParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the enable repository health check params
func (o *EnableRepositoryHealthCheckParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the enable repository health check params
func (o *EnableRepositoryHealthCheckParams) WithHTTPClient(client *http.Client) *EnableRepositoryHealthCheckParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the enable repository health check params
func (o *EnableRepositoryHealthCheckParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithRepositoryName adds the repositoryName to the enable repository health check params
func (o *EnableRepositoryHealthCheckParams) WithRepositoryName(repositoryName string) *EnableRepositoryHealthCheckParams {
	o.SetRepositoryName(repositoryName)
	return o
}

// SetRepositoryName adds the repositoryName to the enable repository health check params
func (o *EnableRepositoryHealthCheckParams) SetRepositoryName(repositoryName string) {
	o.RepositoryName = repositoryName
}

// WriteToRequest writes these params to a swagger request
func (o *EnableRepositoryHealthCheckParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
