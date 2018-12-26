// Code generated by go-swagger; DO NOT EDIT.

package body_and_weight

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// UpdateBodyFatGoalReader is a Reader for the UpdateBodyFatGoal structure.
type UpdateBodyFatGoalReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateBodyFatGoalReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewUpdateBodyFatGoalOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 201:
		result := NewUpdateBodyFatGoalCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewUpdateBodyFatGoalBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewUpdateBodyFatGoalUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewUpdateBodyFatGoalOK creates a UpdateBodyFatGoalOK with default headers values
func NewUpdateBodyFatGoalOK() *UpdateBodyFatGoalOK {
	return &UpdateBodyFatGoalOK{}
}

/*UpdateBodyFatGoalOK handles this case with default header values.

A successful request.
*/
type UpdateBodyFatGoalOK struct {
}

func (o *UpdateBodyFatGoalOK) Error() string {
	return fmt.Sprintf("[POST /1/user/-/body/log/fat/goal.json][%d] updateBodyFatGoalOK ", 200)
}

func (o *UpdateBodyFatGoalOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateBodyFatGoalCreated creates a UpdateBodyFatGoalCreated with default headers values
func NewUpdateBodyFatGoalCreated() *UpdateBodyFatGoalCreated {
	return &UpdateBodyFatGoalCreated{}
}

/*UpdateBodyFatGoalCreated handles this case with default header values.

The request has been fulfilled and resulted in a new resouce being created.
*/
type UpdateBodyFatGoalCreated struct {
}

func (o *UpdateBodyFatGoalCreated) Error() string {
	return fmt.Sprintf("[POST /1/user/-/body/log/fat/goal.json][%d] updateBodyFatGoalCreated ", 201)
}

func (o *UpdateBodyFatGoalCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateBodyFatGoalBadRequest creates a UpdateBodyFatGoalBadRequest with default headers values
func NewUpdateBodyFatGoalBadRequest() *UpdateBodyFatGoalBadRequest {
	return &UpdateBodyFatGoalBadRequest{}
}

/*UpdateBodyFatGoalBadRequest handles this case with default header values.

The request had bad syntax or was inherently impossible to be satified.
*/
type UpdateBodyFatGoalBadRequest struct {
}

func (o *UpdateBodyFatGoalBadRequest) Error() string {
	return fmt.Sprintf("[POST /1/user/-/body/log/fat/goal.json][%d] updateBodyFatGoalBadRequest ", 400)
}

func (o *UpdateBodyFatGoalBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateBodyFatGoalUnauthorized creates a UpdateBodyFatGoalUnauthorized with default headers values
func NewUpdateBodyFatGoalUnauthorized() *UpdateBodyFatGoalUnauthorized {
	return &UpdateBodyFatGoalUnauthorized{}
}

/*UpdateBodyFatGoalUnauthorized handles this case with default header values.

The request requires user authentication.
*/
type UpdateBodyFatGoalUnauthorized struct {
}

func (o *UpdateBodyFatGoalUnauthorized) Error() string {
	return fmt.Sprintf("[POST /1/user/-/body/log/fat/goal.json][%d] updateBodyFatGoalUnauthorized ", 401)
}

func (o *UpdateBodyFatGoalUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
