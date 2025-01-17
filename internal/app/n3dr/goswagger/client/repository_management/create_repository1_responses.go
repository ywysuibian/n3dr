// Code generated by go-swagger; DO NOT EDIT.

package repository_management

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// CreateRepository1Reader is a Reader for the CreateRepository1 structure.
type CreateRepository1Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateRepository1Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewCreateRepository1Created()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewCreateRepository1Unauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewCreateRepository1Forbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCreateRepository1Created creates a CreateRepository1Created with default headers values
func NewCreateRepository1Created() *CreateRepository1Created {
	return &CreateRepository1Created{}
}

/* CreateRepository1Created describes a response with status code 201, with default header values.

Repository created
*/
type CreateRepository1Created struct {
}

func (o *CreateRepository1Created) Error() string {
	return fmt.Sprintf("[POST /v1/repositories/maven/hosted][%d] createRepository1Created ", 201)
}

func (o *CreateRepository1Created) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCreateRepository1Unauthorized creates a CreateRepository1Unauthorized with default headers values
func NewCreateRepository1Unauthorized() *CreateRepository1Unauthorized {
	return &CreateRepository1Unauthorized{}
}

/* CreateRepository1Unauthorized describes a response with status code 401, with default header values.

Authentication required
*/
type CreateRepository1Unauthorized struct {
}

func (o *CreateRepository1Unauthorized) Error() string {
	return fmt.Sprintf("[POST /v1/repositories/maven/hosted][%d] createRepository1Unauthorized ", 401)
}

func (o *CreateRepository1Unauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCreateRepository1Forbidden creates a CreateRepository1Forbidden with default headers values
func NewCreateRepository1Forbidden() *CreateRepository1Forbidden {
	return &CreateRepository1Forbidden{}
}

/* CreateRepository1Forbidden describes a response with status code 403, with default header values.

Insufficient permissions
*/
type CreateRepository1Forbidden struct {
}

func (o *CreateRepository1Forbidden) Error() string {
	return fmt.Sprintf("[POST /v1/repositories/maven/hosted][%d] createRepository1Forbidden ", 403)
}

func (o *CreateRepository1Forbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
