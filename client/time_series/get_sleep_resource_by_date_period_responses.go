// Code generated by go-swagger; DO NOT EDIT.

package time_series

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// GetSleepResourceByDatePeriodReader is a Reader for the GetSleepResourceByDatePeriod structure.
type GetSleepResourceByDatePeriodReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetSleepResourceByDatePeriodReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetSleepResourceByDatePeriodOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 401:
		result := NewGetSleepResourceByDatePeriodUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewGetSleepResourceByDatePeriodForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetSleepResourceByDatePeriodOK creates a GetSleepResourceByDatePeriodOK with default headers values
func NewGetSleepResourceByDatePeriodOK() *GetSleepResourceByDatePeriodOK {
	return &GetSleepResourceByDatePeriodOK{}
}

/*GetSleepResourceByDatePeriodOK handles this case with default header values.

Successful request.
*/
type GetSleepResourceByDatePeriodOK struct {
}

func (o *GetSleepResourceByDatePeriodOK) Error() string {
	return fmt.Sprintf("[GET /1/user/-/sleep/{resource-path}/date/{date}/{period}.json][%d] getSleepResourceByDatePeriodOK ", 200)
}

func (o *GetSleepResourceByDatePeriodOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetSleepResourceByDatePeriodUnauthorized creates a GetSleepResourceByDatePeriodUnauthorized with default headers values
func NewGetSleepResourceByDatePeriodUnauthorized() *GetSleepResourceByDatePeriodUnauthorized {
	return &GetSleepResourceByDatePeriodUnauthorized{}
}

/*GetSleepResourceByDatePeriodUnauthorized handles this case with default header values.

The request requires user authentication.
*/
type GetSleepResourceByDatePeriodUnauthorized struct {
}

func (o *GetSleepResourceByDatePeriodUnauthorized) Error() string {
	return fmt.Sprintf("[GET /1/user/-/sleep/{resource-path}/date/{date}/{period}.json][%d] getSleepResourceByDatePeriodUnauthorized ", 401)
}

func (o *GetSleepResourceByDatePeriodUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetSleepResourceByDatePeriodForbidden creates a GetSleepResourceByDatePeriodForbidden with default headers values
func NewGetSleepResourceByDatePeriodForbidden() *GetSleepResourceByDatePeriodForbidden {
	return &GetSleepResourceByDatePeriodForbidden{}
}

/*GetSleepResourceByDatePeriodForbidden handles this case with default header values.

The server understood the request, but is refusing to fulfill it.
*/
type GetSleepResourceByDatePeriodForbidden struct {
}

func (o *GetSleepResourceByDatePeriodForbidden) Error() string {
	return fmt.Sprintf("[GET /1/user/-/sleep/{resource-path}/date/{date}/{period}.json][%d] getSleepResourceByDatePeriodForbidden ", 403)
}

func (o *GetSleepResourceByDatePeriodForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
