// Code generated by go-swagger; DO NOT EDIT.

package time_series

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// GetActivitiesResourceByDatePeriodReader is a Reader for the GetActivitiesResourceByDatePeriod structure.
type GetActivitiesResourceByDatePeriodReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetActivitiesResourceByDatePeriodReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetActivitiesResourceByDatePeriodOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewGetActivitiesResourceByDatePeriodBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewGetActivitiesResourceByDatePeriodUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 409:
		result := NewGetActivitiesResourceByDatePeriodConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetActivitiesResourceByDatePeriodOK creates a GetActivitiesResourceByDatePeriodOK with default headers values
func NewGetActivitiesResourceByDatePeriodOK() *GetActivitiesResourceByDatePeriodOK {
	return &GetActivitiesResourceByDatePeriodOK{}
}

/*GetActivitiesResourceByDatePeriodOK handles this case with default header values.

A successful request.
*/
type GetActivitiesResourceByDatePeriodOK struct {
}

func (o *GetActivitiesResourceByDatePeriodOK) Error() string {
	return fmt.Sprintf("[GET /1/user/-/activities/{resource-path}/date/{date}/{period}.json][%d] getActivitiesResourceByDatePeriodOK ", 200)
}

func (o *GetActivitiesResourceByDatePeriodOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetActivitiesResourceByDatePeriodBadRequest creates a GetActivitiesResourceByDatePeriodBadRequest with default headers values
func NewGetActivitiesResourceByDatePeriodBadRequest() *GetActivitiesResourceByDatePeriodBadRequest {
	return &GetActivitiesResourceByDatePeriodBadRequest{}
}

/*GetActivitiesResourceByDatePeriodBadRequest handles this case with default header values.

The request had bad syntax or was inherently impossible to be satified.
*/
type GetActivitiesResourceByDatePeriodBadRequest struct {
}

func (o *GetActivitiesResourceByDatePeriodBadRequest) Error() string {
	return fmt.Sprintf("[GET /1/user/-/activities/{resource-path}/date/{date}/{period}.json][%d] getActivitiesResourceByDatePeriodBadRequest ", 400)
}

func (o *GetActivitiesResourceByDatePeriodBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetActivitiesResourceByDatePeriodUnauthorized creates a GetActivitiesResourceByDatePeriodUnauthorized with default headers values
func NewGetActivitiesResourceByDatePeriodUnauthorized() *GetActivitiesResourceByDatePeriodUnauthorized {
	return &GetActivitiesResourceByDatePeriodUnauthorized{}
}

/*GetActivitiesResourceByDatePeriodUnauthorized handles this case with default header values.

The request requires user authentication.
*/
type GetActivitiesResourceByDatePeriodUnauthorized struct {
}

func (o *GetActivitiesResourceByDatePeriodUnauthorized) Error() string {
	return fmt.Sprintf("[GET /1/user/-/activities/{resource-path}/date/{date}/{period}.json][%d] getActivitiesResourceByDatePeriodUnauthorized ", 401)
}

func (o *GetActivitiesResourceByDatePeriodUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetActivitiesResourceByDatePeriodConflict creates a GetActivitiesResourceByDatePeriodConflict with default headers values
func NewGetActivitiesResourceByDatePeriodConflict() *GetActivitiesResourceByDatePeriodConflict {
	return &GetActivitiesResourceByDatePeriodConflict{}
}

/*GetActivitiesResourceByDatePeriodConflict handles this case with default header values.

Returned if the given user is already subscribed to this stream using a different subscription ID, OR if the given subscription ID is already used to identify a subscription to a different stream.
*/
type GetActivitiesResourceByDatePeriodConflict struct {
}

func (o *GetActivitiesResourceByDatePeriodConflict) Error() string {
	return fmt.Sprintf("[GET /1/user/-/activities/{resource-path}/date/{date}/{period}.json][%d] getActivitiesResourceByDatePeriodConflict ", 409)
}

func (o *GetActivitiesResourceByDatePeriodConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
