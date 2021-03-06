// Code generated by go-swagger; DO NOT EDIT.

package food_logging

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// AddUpdateFoodsGoalReader is a Reader for the AddUpdateFoodsGoal structure.
type AddUpdateFoodsGoalReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AddUpdateFoodsGoalReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 201:
		result := NewAddUpdateFoodsGoalCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewAddUpdateFoodsGoalBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewAddUpdateFoodsGoalUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewAddUpdateFoodsGoalForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewAddUpdateFoodsGoalCreated creates a AddUpdateFoodsGoalCreated with default headers values
func NewAddUpdateFoodsGoalCreated() *AddUpdateFoodsGoalCreated {
	return &AddUpdateFoodsGoalCreated{}
}

/*AddUpdateFoodsGoalCreated handles this case with default header values.

The request has been fulfilled and resulted in a new resource being created.
*/
type AddUpdateFoodsGoalCreated struct {
}

func (o *AddUpdateFoodsGoalCreated) Error() string {
	return fmt.Sprintf("[POST /1/user/-/foods/log/goal.json][%d] addUpdateFoodsGoalCreated ", 201)
}

func (o *AddUpdateFoodsGoalCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewAddUpdateFoodsGoalBadRequest creates a AddUpdateFoodsGoalBadRequest with default headers values
func NewAddUpdateFoodsGoalBadRequest() *AddUpdateFoodsGoalBadRequest {
	return &AddUpdateFoodsGoalBadRequest{}
}

/*AddUpdateFoodsGoalBadRequest handles this case with default header values.

The request had bad syntax or was inherently impossible to be satified.
*/
type AddUpdateFoodsGoalBadRequest struct {
}

func (o *AddUpdateFoodsGoalBadRequest) Error() string {
	return fmt.Sprintf("[POST /1/user/-/foods/log/goal.json][%d] addUpdateFoodsGoalBadRequest ", 400)
}

func (o *AddUpdateFoodsGoalBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewAddUpdateFoodsGoalUnauthorized creates a AddUpdateFoodsGoalUnauthorized with default headers values
func NewAddUpdateFoodsGoalUnauthorized() *AddUpdateFoodsGoalUnauthorized {
	return &AddUpdateFoodsGoalUnauthorized{}
}

/*AddUpdateFoodsGoalUnauthorized handles this case with default header values.

The request requires user authentication.
*/
type AddUpdateFoodsGoalUnauthorized struct {
}

func (o *AddUpdateFoodsGoalUnauthorized) Error() string {
	return fmt.Sprintf("[POST /1/user/-/foods/log/goal.json][%d] addUpdateFoodsGoalUnauthorized ", 401)
}

func (o *AddUpdateFoodsGoalUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewAddUpdateFoodsGoalForbidden creates a AddUpdateFoodsGoalForbidden with default headers values
func NewAddUpdateFoodsGoalForbidden() *AddUpdateFoodsGoalForbidden {
	return &AddUpdateFoodsGoalForbidden{}
}

/*AddUpdateFoodsGoalForbidden handles this case with default header values.

The server understood the request, but is refusing to fulfill it.
*/
type AddUpdateFoodsGoalForbidden struct {
}

func (o *AddUpdateFoodsGoalForbidden) Error() string {
	return fmt.Sprintf("[POST /1/user/-/foods/log/goal.json][%d] addUpdateFoodsGoalForbidden ", 403)
}

func (o *AddUpdateFoodsGoalForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
