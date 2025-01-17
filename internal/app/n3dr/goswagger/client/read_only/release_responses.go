// Code generated by go-swagger; DO NOT EDIT.

package read_only

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// ReleaseReader is a Reader for the Release structure.
type ReleaseReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ReleaseReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewReleaseNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewReleaseForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewReleaseNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewReleaseNoContent creates a ReleaseNoContent with default headers values
func NewReleaseNoContent() *ReleaseNoContent {
	return &ReleaseNoContent{}
}

/* ReleaseNoContent describes a response with status code 204, with default header values.

System is no longer read-only
*/
type ReleaseNoContent struct {
}

func (o *ReleaseNoContent) Error() string {
	return fmt.Sprintf("[POST /v1/read-only/release][%d] releaseNoContent ", 204)
}

func (o *ReleaseNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewReleaseForbidden creates a ReleaseForbidden with default headers values
func NewReleaseForbidden() *ReleaseForbidden {
	return &ReleaseForbidden{}
}

/* ReleaseForbidden describes a response with status code 403, with default header values.

Authentication required
*/
type ReleaseForbidden struct {
}

func (o *ReleaseForbidden) Error() string {
	return fmt.Sprintf("[POST /v1/read-only/release][%d] releaseForbidden ", 403)
}

func (o *ReleaseForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewReleaseNotFound creates a ReleaseNotFound with default headers values
func NewReleaseNotFound() *ReleaseNotFound {
	return &ReleaseNotFound{}
}

/* ReleaseNotFound describes a response with status code 404, with default header values.

No change to read-only state
*/
type ReleaseNotFound struct {
}

func (o *ReleaseNotFound) Error() string {
	return fmt.Sprintf("[POST /v1/read-only/release][%d] releaseNotFound ", 404)
}

func (o *ReleaseNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
