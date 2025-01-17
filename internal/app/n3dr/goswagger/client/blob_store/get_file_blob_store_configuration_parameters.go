// Code generated by go-swagger; DO NOT EDIT.

package blob_store

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

// NewGetFileBlobStoreConfigurationParams creates a new GetFileBlobStoreConfigurationParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetFileBlobStoreConfigurationParams() *GetFileBlobStoreConfigurationParams {
	return &GetFileBlobStoreConfigurationParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetFileBlobStoreConfigurationParamsWithTimeout creates a new GetFileBlobStoreConfigurationParams object
// with the ability to set a timeout on a request.
func NewGetFileBlobStoreConfigurationParamsWithTimeout(timeout time.Duration) *GetFileBlobStoreConfigurationParams {
	return &GetFileBlobStoreConfigurationParams{
		timeout: timeout,
	}
}

// NewGetFileBlobStoreConfigurationParamsWithContext creates a new GetFileBlobStoreConfigurationParams object
// with the ability to set a context for a request.
func NewGetFileBlobStoreConfigurationParamsWithContext(ctx context.Context) *GetFileBlobStoreConfigurationParams {
	return &GetFileBlobStoreConfigurationParams{
		Context: ctx,
	}
}

// NewGetFileBlobStoreConfigurationParamsWithHTTPClient creates a new GetFileBlobStoreConfigurationParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetFileBlobStoreConfigurationParamsWithHTTPClient(client *http.Client) *GetFileBlobStoreConfigurationParams {
	return &GetFileBlobStoreConfigurationParams{
		HTTPClient: client,
	}
}

/* GetFileBlobStoreConfigurationParams contains all the parameters to send to the API endpoint
   for the get file blob store configuration operation.

   Typically these are written to a http.Request.
*/
type GetFileBlobStoreConfigurationParams struct {

	/* Name.

	   The name of the file blob store to read
	*/
	Name string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get file blob store configuration params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetFileBlobStoreConfigurationParams) WithDefaults() *GetFileBlobStoreConfigurationParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get file blob store configuration params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetFileBlobStoreConfigurationParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get file blob store configuration params
func (o *GetFileBlobStoreConfigurationParams) WithTimeout(timeout time.Duration) *GetFileBlobStoreConfigurationParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get file blob store configuration params
func (o *GetFileBlobStoreConfigurationParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get file blob store configuration params
func (o *GetFileBlobStoreConfigurationParams) WithContext(ctx context.Context) *GetFileBlobStoreConfigurationParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get file blob store configuration params
func (o *GetFileBlobStoreConfigurationParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get file blob store configuration params
func (o *GetFileBlobStoreConfigurationParams) WithHTTPClient(client *http.Client) *GetFileBlobStoreConfigurationParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get file blob store configuration params
func (o *GetFileBlobStoreConfigurationParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithName adds the name to the get file blob store configuration params
func (o *GetFileBlobStoreConfigurationParams) WithName(name string) *GetFileBlobStoreConfigurationParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the get file blob store configuration params
func (o *GetFileBlobStoreConfigurationParams) SetName(name string) {
	o.Name = name
}

// WriteToRequest writes these params to a swagger request
func (o *GetFileBlobStoreConfigurationParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param name
	if err := r.SetPathParam("name", o.Name); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
