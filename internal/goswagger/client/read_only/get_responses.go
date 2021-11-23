// Code generated by go-swagger; DO NOT EDIT.

package read_only

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/030/n3dr/internal/goswagger/models"
)

// GetReader is a Reader for the Get structure.
type GetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetOK creates a GetOK with default headers values
func NewGetOK() *GetOK {
	return &GetOK{}
}

/* GetOK describes a response with status code 200, with default header values.

successful operation
*/
type GetOK struct {
	Payload *models.ReadOnlyState
}

func (o *GetOK) Error() string {
	return fmt.Sprintf("[GET /v1/read-only][%d] getOK  %+v", 200, o.Payload)
}
func (o *GetOK) GetPayload() *models.ReadOnlyState {
	return o.Payload
}

func (o *GetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ReadOnlyState)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
