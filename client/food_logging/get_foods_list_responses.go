// Code generated by go-swagger; DO NOT EDIT.

package food_logging

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// GetFoodsListReader is a Reader for the GetFoodsList structure.
type GetFoodsListReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetFoodsListReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetFoodsListOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewGetFoodsListBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewGetFoodsListUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewGetFoodsListForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetFoodsListOK creates a GetFoodsListOK with default headers values
func NewGetFoodsListOK() *GetFoodsListOK {
	return &GetFoodsListOK{}
}

/*GetFoodsListOK handles this case with default header values.

A successful request.
*/
type GetFoodsListOK struct {
}

func (o *GetFoodsListOK) Error() string {
	return fmt.Sprintf("[GET /1/foods/search.json][%d] getFoodsListOK ", 200)
}

func (o *GetFoodsListOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetFoodsListBadRequest creates a GetFoodsListBadRequest with default headers values
func NewGetFoodsListBadRequest() *GetFoodsListBadRequest {
	return &GetFoodsListBadRequest{}
}

/*GetFoodsListBadRequest handles this case with default header values.

The request had bad syntax or was inherently impossible to be satified.
*/
type GetFoodsListBadRequest struct {
}

func (o *GetFoodsListBadRequest) Error() string {
	return fmt.Sprintf("[GET /1/foods/search.json][%d] getFoodsListBadRequest ", 400)
}

func (o *GetFoodsListBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetFoodsListUnauthorized creates a GetFoodsListUnauthorized with default headers values
func NewGetFoodsListUnauthorized() *GetFoodsListUnauthorized {
	return &GetFoodsListUnauthorized{}
}

/*GetFoodsListUnauthorized handles this case with default header values.

The request requires user authentication.
*/
type GetFoodsListUnauthorized struct {
}

func (o *GetFoodsListUnauthorized) Error() string {
	return fmt.Sprintf("[GET /1/foods/search.json][%d] getFoodsListUnauthorized ", 401)
}

func (o *GetFoodsListUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetFoodsListForbidden creates a GetFoodsListForbidden with default headers values
func NewGetFoodsListForbidden() *GetFoodsListForbidden {
	return &GetFoodsListForbidden{}
}

/*GetFoodsListForbidden handles this case with default header values.

The server understood the request, but is refusing to fulfill it.
*/
type GetFoodsListForbidden struct {
}

func (o *GetFoodsListForbidden) Error() string {
	return fmt.Sprintf("[GET /1/foods/search.json][%d] getFoodsListForbidden ", 403)
}

func (o *GetFoodsListForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
