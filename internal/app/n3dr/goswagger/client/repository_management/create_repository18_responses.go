// Code generated by go-swagger; DO NOT EDIT.

package repository_management

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// CreateRepository18Reader is a Reader for the CreateRepository18 structure.
type CreateRepository18Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateRepository18Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewCreateRepository18Created()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewCreateRepository18Unauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewCreateRepository18Forbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCreateRepository18Created creates a CreateRepository18Created with default headers values
func NewCreateRepository18Created() *CreateRepository18Created {
	return &CreateRepository18Created{}
}

/* CreateRepository18Created describes a response with status code 201, with default header values.

Repository created
*/
type CreateRepository18Created struct {
}

func (o *CreateRepository18Created) Error() string {
	return fmt.Sprintf("[POST /v1/repositories/docker/hosted][%d] createRepository18Created ", 201)
}

func (o *CreateRepository18Created) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCreateRepository18Unauthorized creates a CreateRepository18Unauthorized with default headers values
func NewCreateRepository18Unauthorized() *CreateRepository18Unauthorized {
	return &CreateRepository18Unauthorized{}
}

/* CreateRepository18Unauthorized describes a response with status code 401, with default header values.

Authentication required
*/
type CreateRepository18Unauthorized struct {
}

func (o *CreateRepository18Unauthorized) Error() string {
	return fmt.Sprintf("[POST /v1/repositories/docker/hosted][%d] createRepository18Unauthorized ", 401)
}

func (o *CreateRepository18Unauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCreateRepository18Forbidden creates a CreateRepository18Forbidden with default headers values
func NewCreateRepository18Forbidden() *CreateRepository18Forbidden {
	return &CreateRepository18Forbidden{}
}

/* CreateRepository18Forbidden describes a response with status code 403, with default header values.

Repository not found
*/
type CreateRepository18Forbidden struct {
}

func (o *CreateRepository18Forbidden) Error() string {
	return fmt.Sprintf("[POST /v1/repositories/docker/hosted][%d] createRepository18Forbidden ", 403)
}

func (o *CreateRepository18Forbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
