// Code generated by go-swagger; DO NOT EDIT.

package script

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// Delete1Reader is a Reader for the Delete1 structure.
type Delete1Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *Delete1Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewDelete1NoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewDelete1NotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDelete1NoContent creates a Delete1NoContent with default headers values
func NewDelete1NoContent() *Delete1NoContent {
	return &Delete1NoContent{}
}

/* Delete1NoContent describes a response with status code 204, with default header values.

Script was deleted
*/
type Delete1NoContent struct {
}

func (o *Delete1NoContent) Error() string {
	return fmt.Sprintf("[DELETE /v1/script/{name}][%d] delete1NoContent ", 204)
}

func (o *Delete1NoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDelete1NotFound creates a Delete1NotFound with default headers values
func NewDelete1NotFound() *Delete1NotFound {
	return &Delete1NotFound{}
}

/* Delete1NotFound describes a response with status code 404, with default header values.

No script with the specified name
*/
type Delete1NotFound struct {
}

func (o *Delete1NotFound) Error() string {
	return fmt.Sprintf("[DELETE /v1/script/{name}][%d] delete1NotFound ", 404)
}

func (o *Delete1NotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
