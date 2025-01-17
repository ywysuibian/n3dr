// Code generated by go-swagger; DO NOT EDIT.

package routing_rules

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// DeleteRoutingRuleReader is a Reader for the DeleteRoutingRule structure.
type DeleteRoutingRuleReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteRoutingRuleReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewDeleteRoutingRuleNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewDeleteRoutingRuleForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewDeleteRoutingRuleNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDeleteRoutingRuleNoContent creates a DeleteRoutingRuleNoContent with default headers values
func NewDeleteRoutingRuleNoContent() *DeleteRoutingRuleNoContent {
	return &DeleteRoutingRuleNoContent{}
}

/* DeleteRoutingRuleNoContent describes a response with status code 204, with default header values.

Routing rule was successfully deleted
*/
type DeleteRoutingRuleNoContent struct {
}

func (o *DeleteRoutingRuleNoContent) Error() string {
	return fmt.Sprintf("[DELETE /v1/routing-rules/{name}][%d] deleteRoutingRuleNoContent ", 204)
}

func (o *DeleteRoutingRuleNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteRoutingRuleForbidden creates a DeleteRoutingRuleForbidden with default headers values
func NewDeleteRoutingRuleForbidden() *DeleteRoutingRuleForbidden {
	return &DeleteRoutingRuleForbidden{}
}

/* DeleteRoutingRuleForbidden describes a response with status code 403, with default header values.

Insufficient permissions to delete routing rules
*/
type DeleteRoutingRuleForbidden struct {
}

func (o *DeleteRoutingRuleForbidden) Error() string {
	return fmt.Sprintf("[DELETE /v1/routing-rules/{name}][%d] deleteRoutingRuleForbidden ", 403)
}

func (o *DeleteRoutingRuleForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteRoutingRuleNotFound creates a DeleteRoutingRuleNotFound with default headers values
func NewDeleteRoutingRuleNotFound() *DeleteRoutingRuleNotFound {
	return &DeleteRoutingRuleNotFound{}
}

/* DeleteRoutingRuleNotFound describes a response with status code 404, with default header values.

Routing rule not found
*/
type DeleteRoutingRuleNotFound struct {
}

func (o *DeleteRoutingRuleNotFound) Error() string {
	return fmt.Sprintf("[DELETE /v1/routing-rules/{name}][%d] deleteRoutingRuleNotFound ", 404)
}

func (o *DeleteRoutingRuleNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
