// Code generated by go-swagger; DO NOT EDIT.

package manage_i_q_server_configuration

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// EnableIqReader is a Reader for the EnableIq structure.
type EnableIqReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *EnableIqReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewEnableIqNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewEnableIqBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewEnableIqNoContent creates a EnableIqNoContent with default headers values
func NewEnableIqNoContent() *EnableIqNoContent {
	return &EnableIqNoContent{}
}

/* EnableIqNoContent describes a response with status code 204, with default header values.

IQ server has been enabled
*/
type EnableIqNoContent struct {
}

func (o *EnableIqNoContent) Error() string {
	return fmt.Sprintf("[POST /v1/iq/enable][%d] enableIqNoContent ", 204)
}

func (o *EnableIqNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewEnableIqBadRequest creates a EnableIqBadRequest with default headers values
func NewEnableIqBadRequest() *EnableIqBadRequest {
	return &EnableIqBadRequest{}
}

/* EnableIqBadRequest describes a response with status code 400, with default header values.

IQ server connection not configured
*/
type EnableIqBadRequest struct {
}

func (o *EnableIqBadRequest) Error() string {
	return fmt.Sprintf("[POST /v1/iq/enable][%d] enableIqBadRequest ", 400)
}

func (o *EnableIqBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
