// Code generated by go-swagger; DO NOT EDIT.

package food_logging

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// GetFavoriteFoodsReader is a Reader for the GetFavoriteFoods structure.
type GetFavoriteFoodsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetFavoriteFoodsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetFavoriteFoodsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewGetFavoriteFoodsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewGetFavoriteFoodsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewGetFavoriteFoodsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetFavoriteFoodsOK creates a GetFavoriteFoodsOK with default headers values
func NewGetFavoriteFoodsOK() *GetFavoriteFoodsOK {
	return &GetFavoriteFoodsOK{}
}

/*GetFavoriteFoodsOK handles this case with default header values.

Successful request.
*/
type GetFavoriteFoodsOK struct {
}

func (o *GetFavoriteFoodsOK) Error() string {
	return fmt.Sprintf("[GET /1/user/-/foods/log/favorite.json][%d] getFavoriteFoodsOK ", 200)
}

func (o *GetFavoriteFoodsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetFavoriteFoodsBadRequest creates a GetFavoriteFoodsBadRequest with default headers values
func NewGetFavoriteFoodsBadRequest() *GetFavoriteFoodsBadRequest {
	return &GetFavoriteFoodsBadRequest{}
}

/*GetFavoriteFoodsBadRequest handles this case with default header values.

The request had bad syntax or was inherently impossible to be satified.
*/
type GetFavoriteFoodsBadRequest struct {
}

func (o *GetFavoriteFoodsBadRequest) Error() string {
	return fmt.Sprintf("[GET /1/user/-/foods/log/favorite.json][%d] getFavoriteFoodsBadRequest ", 400)
}

func (o *GetFavoriteFoodsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetFavoriteFoodsUnauthorized creates a GetFavoriteFoodsUnauthorized with default headers values
func NewGetFavoriteFoodsUnauthorized() *GetFavoriteFoodsUnauthorized {
	return &GetFavoriteFoodsUnauthorized{}
}

/*GetFavoriteFoodsUnauthorized handles this case with default header values.

The request requires user authentication.
*/
type GetFavoriteFoodsUnauthorized struct {
}

func (o *GetFavoriteFoodsUnauthorized) Error() string {
	return fmt.Sprintf("[GET /1/user/-/foods/log/favorite.json][%d] getFavoriteFoodsUnauthorized ", 401)
}

func (o *GetFavoriteFoodsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetFavoriteFoodsForbidden creates a GetFavoriteFoodsForbidden with default headers values
func NewGetFavoriteFoodsForbidden() *GetFavoriteFoodsForbidden {
	return &GetFavoriteFoodsForbidden{}
}

/*GetFavoriteFoodsForbidden handles this case with default header values.

The server understood the request, but is refusing to fulfill it.
*/
type GetFavoriteFoodsForbidden struct {
}

func (o *GetFavoriteFoodsForbidden) Error() string {
	return fmt.Sprintf("[GET /1/user/-/foods/log/favorite.json][%d] getFavoriteFoodsForbidden ", 403)
}

func (o *GetFavoriteFoodsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
