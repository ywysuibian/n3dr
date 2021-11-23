// Code generated by go-swagger; DO NOT EDIT.

package repository_management

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// RebuildIndexReader is a Reader for the RebuildIndex structure.
type RebuildIndexReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *RebuildIndexReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewRebuildIndexNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewRebuildIndexBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewRebuildIndexUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewRebuildIndexForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewRebuildIndexNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewRebuildIndexNoContent creates a RebuildIndexNoContent with default headers values
func NewRebuildIndexNoContent() *RebuildIndexNoContent {
	return &RebuildIndexNoContent{}
}

/* RebuildIndexNoContent describes a response with status code 204, with default header values.

Repository search index rebuild has been scheduled
*/
type RebuildIndexNoContent struct {
}

func (o *RebuildIndexNoContent) Error() string {
	return fmt.Sprintf("[POST /v1/repositories/{repositoryName}/rebuild-index][%d] rebuildIndexNoContent ", 204)
}

func (o *RebuildIndexNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewRebuildIndexBadRequest creates a RebuildIndexBadRequest with default headers values
func NewRebuildIndexBadRequest() *RebuildIndexBadRequest {
	return &RebuildIndexBadRequest{}
}

/* RebuildIndexBadRequest describes a response with status code 400, with default header values.

Repository is not of hosted or proxy type
*/
type RebuildIndexBadRequest struct {
}

func (o *RebuildIndexBadRequest) Error() string {
	return fmt.Sprintf("[POST /v1/repositories/{repositoryName}/rebuild-index][%d] rebuildIndexBadRequest ", 400)
}

func (o *RebuildIndexBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewRebuildIndexUnauthorized creates a RebuildIndexUnauthorized with default headers values
func NewRebuildIndexUnauthorized() *RebuildIndexUnauthorized {
	return &RebuildIndexUnauthorized{}
}

/* RebuildIndexUnauthorized describes a response with status code 401, with default header values.

Authentication required
*/
type RebuildIndexUnauthorized struct {
}

func (o *RebuildIndexUnauthorized) Error() string {
	return fmt.Sprintf("[POST /v1/repositories/{repositoryName}/rebuild-index][%d] rebuildIndexUnauthorized ", 401)
}

func (o *RebuildIndexUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewRebuildIndexForbidden creates a RebuildIndexForbidden with default headers values
func NewRebuildIndexForbidden() *RebuildIndexForbidden {
	return &RebuildIndexForbidden{}
}

/* RebuildIndexForbidden describes a response with status code 403, with default header values.

Insufficient permissions
*/
type RebuildIndexForbidden struct {
}

func (o *RebuildIndexForbidden) Error() string {
	return fmt.Sprintf("[POST /v1/repositories/{repositoryName}/rebuild-index][%d] rebuildIndexForbidden ", 403)
}

func (o *RebuildIndexForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewRebuildIndexNotFound creates a RebuildIndexNotFound with default headers values
func NewRebuildIndexNotFound() *RebuildIndexNotFound {
	return &RebuildIndexNotFound{}
}

/* RebuildIndexNotFound describes a response with status code 404, with default header values.

Repository not found
*/
type RebuildIndexNotFound struct {
}

func (o *RebuildIndexNotFound) Error() string {
	return fmt.Sprintf("[POST /v1/repositories/{repositoryName}/rebuild-index][%d] rebuildIndexNotFound ", 404)
}

func (o *RebuildIndexNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
