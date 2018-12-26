// Code generated by go-swagger; DO NOT EDIT.

package food_logging

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// GetFoodsByDateReader is a Reader for the GetFoodsByDate structure.
type GetFoodsByDateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetFoodsByDateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetFoodsByDateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewGetFoodsByDateBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewGetFoodsByDateUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewGetFoodsByDateForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetFoodsByDateOK creates a GetFoodsByDateOK with default headers values
func NewGetFoodsByDateOK() *GetFoodsByDateOK {
	return &GetFoodsByDateOK{}
}

/*GetFoodsByDateOK handles this case with default header values.

A successful request.
*/
type GetFoodsByDateOK struct {
}

func (o *GetFoodsByDateOK) Error() string {
	return fmt.Sprintf("[GET /1/user/-/foods/log/date/{date}.json][%d] getFoodsByDateOK ", 200)
}

func (o *GetFoodsByDateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetFoodsByDateBadRequest creates a GetFoodsByDateBadRequest with default headers values
func NewGetFoodsByDateBadRequest() *GetFoodsByDateBadRequest {
	return &GetFoodsByDateBadRequest{}
}

/*GetFoodsByDateBadRequest handles this case with default header values.

The request had bad syntax or was inherently impossible to be satified.
*/
type GetFoodsByDateBadRequest struct {
}

func (o *GetFoodsByDateBadRequest) Error() string {
	return fmt.Sprintf("[GET /1/user/-/foods/log/date/{date}.json][%d] getFoodsByDateBadRequest ", 400)
}

func (o *GetFoodsByDateBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetFoodsByDateUnauthorized creates a GetFoodsByDateUnauthorized with default headers values
func NewGetFoodsByDateUnauthorized() *GetFoodsByDateUnauthorized {
	return &GetFoodsByDateUnauthorized{}
}

/*GetFoodsByDateUnauthorized handles this case with default header values.

The request requires user authentication.
*/
type GetFoodsByDateUnauthorized struct {
}

func (o *GetFoodsByDateUnauthorized) Error() string {
	return fmt.Sprintf("[GET /1/user/-/foods/log/date/{date}.json][%d] getFoodsByDateUnauthorized ", 401)
}

func (o *GetFoodsByDateUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetFoodsByDateForbidden creates a GetFoodsByDateForbidden with default headers values
func NewGetFoodsByDateForbidden() *GetFoodsByDateForbidden {
	return &GetFoodsByDateForbidden{}
}

/*GetFoodsByDateForbidden handles this case with default header values.

The server understood the request, but is refusing to fulfill it.
*/
type GetFoodsByDateForbidden struct {
}

func (o *GetFoodsByDateForbidden) Error() string {
	return fmt.Sprintf("[GET /1/user/-/foods/log/date/{date}.json][%d] getFoodsByDateForbidden ", 403)
}

func (o *GetFoodsByDateForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}