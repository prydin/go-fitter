// Code generated by go-swagger; DO NOT EDIT.

package body_and_weight

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// GetBodyGoalsReader is a Reader for the GetBodyGoals structure.
type GetBodyGoalsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetBodyGoalsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetBodyGoalsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewGetBodyGoalsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewGetBodyGoalsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetBodyGoalsOK creates a GetBodyGoalsOK with default headers values
func NewGetBodyGoalsOK() *GetBodyGoalsOK {
	return &GetBodyGoalsOK{}
}

/*GetBodyGoalsOK handles this case with default header values.

A successful request.
*/
type GetBodyGoalsOK struct {
}

func (o *GetBodyGoalsOK) Error() string {
	return fmt.Sprintf("[GET /1/user/-/body/log/{goal-type}/goal.json][%d] getBodyGoalsOK ", 200)
}

func (o *GetBodyGoalsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetBodyGoalsBadRequest creates a GetBodyGoalsBadRequest with default headers values
func NewGetBodyGoalsBadRequest() *GetBodyGoalsBadRequest {
	return &GetBodyGoalsBadRequest{}
}

/*GetBodyGoalsBadRequest handles this case with default header values.

The request had bad syntax or was inherently impossible to be satified.
*/
type GetBodyGoalsBadRequest struct {
}

func (o *GetBodyGoalsBadRequest) Error() string {
	return fmt.Sprintf("[GET /1/user/-/body/log/{goal-type}/goal.json][%d] getBodyGoalsBadRequest ", 400)
}

func (o *GetBodyGoalsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetBodyGoalsUnauthorized creates a GetBodyGoalsUnauthorized with default headers values
func NewGetBodyGoalsUnauthorized() *GetBodyGoalsUnauthorized {
	return &GetBodyGoalsUnauthorized{}
}

/*GetBodyGoalsUnauthorized handles this case with default header values.

The request requires user authentication.
*/
type GetBodyGoalsUnauthorized struct {
}

func (o *GetBodyGoalsUnauthorized) Error() string {
	return fmt.Sprintf("[GET /1/user/-/body/log/{goal-type}/goal.json][%d] getBodyGoalsUnauthorized ", 401)
}

func (o *GetBodyGoalsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
