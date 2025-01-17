// Code generated by go-swagger; DO NOT EDIT.

package assets

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

// NewDeleteAssetParams creates a new DeleteAssetParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeleteAssetParams() *DeleteAssetParams {
	return &DeleteAssetParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteAssetParamsWithTimeout creates a new DeleteAssetParams object
// with the ability to set a timeout on a request.
func NewDeleteAssetParamsWithTimeout(timeout time.Duration) *DeleteAssetParams {
	return &DeleteAssetParams{
		timeout: timeout,
	}
}

// NewDeleteAssetParamsWithContext creates a new DeleteAssetParams object
// with the ability to set a context for a request.
func NewDeleteAssetParamsWithContext(ctx context.Context) *DeleteAssetParams {
	return &DeleteAssetParams{
		Context: ctx,
	}
}

// NewDeleteAssetParamsWithHTTPClient creates a new DeleteAssetParams object
// with the ability to set a custom HTTPClient for a request.
func NewDeleteAssetParamsWithHTTPClient(client *http.Client) *DeleteAssetParams {
	return &DeleteAssetParams{
		HTTPClient: client,
	}
}

/* DeleteAssetParams contains all the parameters to send to the API endpoint
   for the delete asset operation.

   Typically these are written to a http.Request.
*/
type DeleteAssetParams struct {

	/* ID.

	   Id of the asset to delete
	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the delete asset params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteAssetParams) WithDefaults() *DeleteAssetParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the delete asset params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteAssetParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the delete asset params
func (o *DeleteAssetParams) WithTimeout(timeout time.Duration) *DeleteAssetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete asset params
func (o *DeleteAssetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete asset params
func (o *DeleteAssetParams) WithContext(ctx context.Context) *DeleteAssetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete asset params
func (o *DeleteAssetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete asset params
func (o *DeleteAssetParams) WithHTTPClient(client *http.Client) *DeleteAssetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete asset params
func (o *DeleteAssetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the delete asset params
func (o *DeleteAssetParams) WithID(id string) *DeleteAssetParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the delete asset params
func (o *DeleteAssetParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteAssetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
