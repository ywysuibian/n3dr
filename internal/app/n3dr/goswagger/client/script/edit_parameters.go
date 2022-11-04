// Code generated by go-swagger; DO NOT EDIT.

package script

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

	"github.com/030/n3dr/internal/app/n3dr/goswagger/models"
)

// NewEditParams creates a new EditParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewEditParams() *EditParams {
	return &EditParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewEditParamsWithTimeout creates a new EditParams object
// with the ability to set a timeout on a request.
func NewEditParamsWithTimeout(timeout time.Duration) *EditParams {
	return &EditParams{
		timeout: timeout,
	}
}

// NewEditParamsWithContext creates a new EditParams object
// with the ability to set a context for a request.
func NewEditParamsWithContext(ctx context.Context) *EditParams {
	return &EditParams{
		Context: ctx,
	}
}

// NewEditParamsWithHTTPClient creates a new EditParams object
// with the ability to set a custom HTTPClient for a request.
func NewEditParamsWithHTTPClient(client *http.Client) *EditParams {
	return &EditParams{
		HTTPClient: client,
	}
}

/* EditParams contains all the parameters to send to the API endpoint
   for the edit operation.

   Typically these are written to a http.Request.
*/
type EditParams struct {

	// Body.
	Body *models.ScriptXO

	// Name.
	Name string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the edit params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *EditParams) WithDefaults() *EditParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the edit params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *EditParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the edit params
func (o *EditParams) WithTimeout(timeout time.Duration) *EditParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the edit params
func (o *EditParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the edit params
func (o *EditParams) WithContext(ctx context.Context) *EditParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the edit params
func (o *EditParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the edit params
func (o *EditParams) WithHTTPClient(client *http.Client) *EditParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the edit params
func (o *EditParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the edit params
func (o *EditParams) WithBody(body *models.ScriptXO) *EditParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the edit params
func (o *EditParams) SetBody(body *models.ScriptXO) {
	o.Body = body
}

// WithName adds the name to the edit params
func (o *EditParams) WithName(name string) *EditParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the edit params
func (o *EditParams) SetName(name string) {
	o.Name = name
}

// WriteToRequest writes these params to a swagger request
func (o *EditParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param name
	if err := r.SetPathParam("name", o.Name); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}