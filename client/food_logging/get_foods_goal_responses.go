// Code generated by go-swagger; DO NOT EDIT.

package food_logging

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// GetFoodsGoalReader is a Reader for the GetFoodsGoal structure.
type GetFoodsGoalReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetFoodsGoalReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetFoodsGoalOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewGetFoodsGoalBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewGetFoodsGoalUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewGetFoodsGoalForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetFoodsGoalOK creates a GetFoodsGoalOK with default headers values
func NewGetFoodsGoalOK() *GetFoodsGoalOK {
	return &GetFoodsGoalOK{}
}

/*GetFoodsGoalOK handles this case with default header values.

A successful request.
*/
type GetFoodsGoalOK struct {
}

func (o *GetFoodsGoalOK) Error() string {
	return fmt.Sprintf("[GET /1/user/-/foods/log/goal.json][%d] getFoodsGoalOK ", 200)
}

func (o *GetFoodsGoalOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetFoodsGoalBadRequest creates a GetFoodsGoalBadRequest with default headers values
func NewGetFoodsGoalBadRequest() *GetFoodsGoalBadRequest {
	return &GetFoodsGoalBadRequest{}
}

/*GetFoodsGoalBadRequest handles this case with default header values.

The request had bad syntax or was inherently impossible to be satified.
*/
type GetFoodsGoalBadRequest struct {
}

func (o *GetFoodsGoalBadRequest) Error() string {
	return fmt.Sprintf("[GET /1/user/-/foods/log/goal.json][%d] getFoodsGoalBadRequest ", 400)
}

func (o *GetFoodsGoalBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetFoodsGoalUnauthorized creates a GetFoodsGoalUnauthorized with default headers values
func NewGetFoodsGoalUnauthorized() *GetFoodsGoalUnauthorized {
	return &GetFoodsGoalUnauthorized{}
}

/*GetFoodsGoalUnauthorized handles this case with default header values.

The request requires user authentication.
*/
type GetFoodsGoalUnauthorized struct {
}

func (o *GetFoodsGoalUnauthorized) Error() string {
	return fmt.Sprintf("[GET /1/user/-/foods/log/goal.json][%d] getFoodsGoalUnauthorized ", 401)
}

func (o *GetFoodsGoalUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetFoodsGoalForbidden creates a GetFoodsGoalForbidden with default headers values
func NewGetFoodsGoalForbidden() *GetFoodsGoalForbidden {
	return &GetFoodsGoalForbidden{}
}

/*GetFoodsGoalForbidden handles this case with default header values.

The server understood the request, but is refusing to fulfill it.
*/
type GetFoodsGoalForbidden struct {
}

func (o *GetFoodsGoalForbidden) Error() string {
	return fmt.Sprintf("[GET /1/user/-/foods/log/goal.json][%d] getFoodsGoalForbidden ", 403)
}

func (o *GetFoodsGoalForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}