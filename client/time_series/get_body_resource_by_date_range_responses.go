// Code generated by go-swagger; DO NOT EDIT.

package time_series

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// GetBodyResourceByDateRangeReader is a Reader for the GetBodyResourceByDateRange structure.
type GetBodyResourceByDateRangeReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetBodyResourceByDateRangeReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetBodyResourceByDateRangeOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewGetBodyResourceByDateRangeBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewGetBodyResourceByDateRangeUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetBodyResourceByDateRangeOK creates a GetBodyResourceByDateRangeOK with default headers values
func NewGetBodyResourceByDateRangeOK() *GetBodyResourceByDateRangeOK {
	return &GetBodyResourceByDateRangeOK{}
}

/*GetBodyResourceByDateRangeOK handles this case with default header values.

A successful request.
*/
type GetBodyResourceByDateRangeOK struct {
}

func (o *GetBodyResourceByDateRangeOK) Error() string {
	return fmt.Sprintf("[GET /1/user/-/body/{resource-path}/date/{base-date}/{end-date}.json][%d] getBodyResourceByDateRangeOK ", 200)
}

func (o *GetBodyResourceByDateRangeOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetBodyResourceByDateRangeBadRequest creates a GetBodyResourceByDateRangeBadRequest with default headers values
func NewGetBodyResourceByDateRangeBadRequest() *GetBodyResourceByDateRangeBadRequest {
	return &GetBodyResourceByDateRangeBadRequest{}
}

/*GetBodyResourceByDateRangeBadRequest handles this case with default header values.

The request had bad syntax or was inherently impossible to be satified.
*/
type GetBodyResourceByDateRangeBadRequest struct {
}

func (o *GetBodyResourceByDateRangeBadRequest) Error() string {
	return fmt.Sprintf("[GET /1/user/-/body/{resource-path}/date/{base-date}/{end-date}.json][%d] getBodyResourceByDateRangeBadRequest ", 400)
}

func (o *GetBodyResourceByDateRangeBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetBodyResourceByDateRangeUnauthorized creates a GetBodyResourceByDateRangeUnauthorized with default headers values
func NewGetBodyResourceByDateRangeUnauthorized() *GetBodyResourceByDateRangeUnauthorized {
	return &GetBodyResourceByDateRangeUnauthorized{}
}

/*GetBodyResourceByDateRangeUnauthorized handles this case with default header values.

The request requires user authentication.
*/
type GetBodyResourceByDateRangeUnauthorized struct {
}

func (o *GetBodyResourceByDateRangeUnauthorized) Error() string {
	return fmt.Sprintf("[GET /1/user/-/body/{resource-path}/date/{base-date}/{end-date}.json][%d] getBodyResourceByDateRangeUnauthorized ", 401)
}

func (o *GetBodyResourceByDateRangeUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
