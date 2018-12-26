// Code generated by go-swagger; DO NOT EDIT.

package body_and_weight

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// GetWeightByDateReader is a Reader for the GetWeightByDate structure.
type GetWeightByDateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetWeightByDateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetWeightByDateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewGetWeightByDateBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewGetWeightByDateUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetWeightByDateOK creates a GetWeightByDateOK with default headers values
func NewGetWeightByDateOK() *GetWeightByDateOK {
	return &GetWeightByDateOK{}
}

/*GetWeightByDateOK handles this case with default header values.

A successful request.
*/
type GetWeightByDateOK struct {
}

func (o *GetWeightByDateOK) Error() string {
	return fmt.Sprintf("[GET /1/user/-/body/log/weight/date/{date}.json][%d] getWeightByDateOK ", 200)
}

func (o *GetWeightByDateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetWeightByDateBadRequest creates a GetWeightByDateBadRequest with default headers values
func NewGetWeightByDateBadRequest() *GetWeightByDateBadRequest {
	return &GetWeightByDateBadRequest{}
}

/*GetWeightByDateBadRequest handles this case with default header values.

The request had bad syntax or was inherently impossible to be satified.
*/
type GetWeightByDateBadRequest struct {
}

func (o *GetWeightByDateBadRequest) Error() string {
	return fmt.Sprintf("[GET /1/user/-/body/log/weight/date/{date}.json][%d] getWeightByDateBadRequest ", 400)
}

func (o *GetWeightByDateBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetWeightByDateUnauthorized creates a GetWeightByDateUnauthorized with default headers values
func NewGetWeightByDateUnauthorized() *GetWeightByDateUnauthorized {
	return &GetWeightByDateUnauthorized{}
}

/*GetWeightByDateUnauthorized handles this case with default header values.

The request requires user authentication.
*/
type GetWeightByDateUnauthorized struct {
}

func (o *GetWeightByDateUnauthorized) Error() string {
	return fmt.Sprintf("[GET /1/user/-/body/log/weight/date/{date}.json][%d] getWeightByDateUnauthorized ", 401)
}

func (o *GetWeightByDateUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}