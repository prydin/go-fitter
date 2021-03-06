// Code generated by go-swagger; DO NOT EDIT.

package body_and_weight

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// GetWeightByDateRangeReader is a Reader for the GetWeightByDateRange structure.
type GetWeightByDateRangeReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetWeightByDateRangeReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetWeightByDateRangeOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewGetWeightByDateRangeBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewGetWeightByDateRangeUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetWeightByDateRangeOK creates a GetWeightByDateRangeOK with default headers values
func NewGetWeightByDateRangeOK() *GetWeightByDateRangeOK {
	return &GetWeightByDateRangeOK{}
}

/*GetWeightByDateRangeOK handles this case with default header values.

A successful request.
*/
type GetWeightByDateRangeOK struct {
}

func (o *GetWeightByDateRangeOK) Error() string {
	return fmt.Sprintf("[GET /1/user/-/body/log/weight/date/{base-date}/{end-date}.json][%d] getWeightByDateRangeOK ", 200)
}

func (o *GetWeightByDateRangeOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetWeightByDateRangeBadRequest creates a GetWeightByDateRangeBadRequest with default headers values
func NewGetWeightByDateRangeBadRequest() *GetWeightByDateRangeBadRequest {
	return &GetWeightByDateRangeBadRequest{}
}

/*GetWeightByDateRangeBadRequest handles this case with default header values.

The request had bad syntax or was inherently impossible to be satified.
*/
type GetWeightByDateRangeBadRequest struct {
}

func (o *GetWeightByDateRangeBadRequest) Error() string {
	return fmt.Sprintf("[GET /1/user/-/body/log/weight/date/{base-date}/{end-date}.json][%d] getWeightByDateRangeBadRequest ", 400)
}

func (o *GetWeightByDateRangeBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetWeightByDateRangeUnauthorized creates a GetWeightByDateRangeUnauthorized with default headers values
func NewGetWeightByDateRangeUnauthorized() *GetWeightByDateRangeUnauthorized {
	return &GetWeightByDateRangeUnauthorized{}
}

/*GetWeightByDateRangeUnauthorized handles this case with default header values.

The request requires user authentication.
*/
type GetWeightByDateRangeUnauthorized struct {
}

func (o *GetWeightByDateRangeUnauthorized) Error() string {
	return fmt.Sprintf("[GET /1/user/-/body/log/weight/date/{base-date}/{end-date}.json][%d] getWeightByDateRangeUnauthorized ", 401)
}

func (o *GetWeightByDateRangeUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
