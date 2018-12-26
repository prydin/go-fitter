// Code generated by go-swagger; DO NOT EDIT.

package food_logging

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// AddFavoriteFoodReader is a Reader for the AddFavoriteFood structure.
type AddFavoriteFoodReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AddFavoriteFoodReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewAddFavoriteFoodOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewAddFavoriteFoodBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewAddFavoriteFoodUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewAddFavoriteFoodForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewAddFavoriteFoodOK creates a AddFavoriteFoodOK with default headers values
func NewAddFavoriteFoodOK() *AddFavoriteFoodOK {
	return &AddFavoriteFoodOK{}
}

/*AddFavoriteFoodOK handles this case with default header values.

Successful request.
*/
type AddFavoriteFoodOK struct {
}

func (o *AddFavoriteFoodOK) Error() string {
	return fmt.Sprintf("[POST /1/user/-/foods/log/favorite/{food-id}.json][%d] addFavoriteFoodOK ", 200)
}

func (o *AddFavoriteFoodOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewAddFavoriteFoodBadRequest creates a AddFavoriteFoodBadRequest with default headers values
func NewAddFavoriteFoodBadRequest() *AddFavoriteFoodBadRequest {
	return &AddFavoriteFoodBadRequest{}
}

/*AddFavoriteFoodBadRequest handles this case with default header values.

The request had bad syntax or was inherently impossible to be satified.
*/
type AddFavoriteFoodBadRequest struct {
}

func (o *AddFavoriteFoodBadRequest) Error() string {
	return fmt.Sprintf("[POST /1/user/-/foods/log/favorite/{food-id}.json][%d] addFavoriteFoodBadRequest ", 400)
}

func (o *AddFavoriteFoodBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewAddFavoriteFoodUnauthorized creates a AddFavoriteFoodUnauthorized with default headers values
func NewAddFavoriteFoodUnauthorized() *AddFavoriteFoodUnauthorized {
	return &AddFavoriteFoodUnauthorized{}
}

/*AddFavoriteFoodUnauthorized handles this case with default header values.

The request requires user authentication.
*/
type AddFavoriteFoodUnauthorized struct {
}

func (o *AddFavoriteFoodUnauthorized) Error() string {
	return fmt.Sprintf("[POST /1/user/-/foods/log/favorite/{food-id}.json][%d] addFavoriteFoodUnauthorized ", 401)
}

func (o *AddFavoriteFoodUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewAddFavoriteFoodForbidden creates a AddFavoriteFoodForbidden with default headers values
func NewAddFavoriteFoodForbidden() *AddFavoriteFoodForbidden {
	return &AddFavoriteFoodForbidden{}
}

/*AddFavoriteFoodForbidden handles this case with default header values.

The server understood the request, but is refusing to fulfill it.
*/
type AddFavoriteFoodForbidden struct {
}

func (o *AddFavoriteFoodForbidden) Error() string {
	return fmt.Sprintf("[POST /1/user/-/foods/log/favorite/{food-id}.json][%d] addFavoriteFoodForbidden ", 403)
}

func (o *AddFavoriteFoodForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
