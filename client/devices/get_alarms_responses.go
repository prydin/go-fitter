// Code generated by go-swagger; DO NOT EDIT.

package devices

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// GetAlarmsReader is a Reader for the GetAlarms structure.
type GetAlarmsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetAlarmsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetAlarmsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewGetAlarmsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewGetAlarmsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetAlarmsOK creates a GetAlarmsOK with default headers values
func NewGetAlarmsOK() *GetAlarmsOK {
	return &GetAlarmsOK{}
}

/*GetAlarmsOK handles this case with default header values.

A successful request.
*/
type GetAlarmsOK struct {
}

func (o *GetAlarmsOK) Error() string {
	return fmt.Sprintf("[GET /1/user/-/devices/tracker/{tracker-id}/alarms.json][%d] getAlarmsOK ", 200)
}

func (o *GetAlarmsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetAlarmsBadRequest creates a GetAlarmsBadRequest with default headers values
func NewGetAlarmsBadRequest() *GetAlarmsBadRequest {
	return &GetAlarmsBadRequest{}
}

/*GetAlarmsBadRequest handles this case with default header values.

The request had bad syntax or was inherently impossible to be satified.
*/
type GetAlarmsBadRequest struct {
}

func (o *GetAlarmsBadRequest) Error() string {
	return fmt.Sprintf("[GET /1/user/-/devices/tracker/{tracker-id}/alarms.json][%d] getAlarmsBadRequest ", 400)
}

func (o *GetAlarmsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetAlarmsUnauthorized creates a GetAlarmsUnauthorized with default headers values
func NewGetAlarmsUnauthorized() *GetAlarmsUnauthorized {
	return &GetAlarmsUnauthorized{}
}

/*GetAlarmsUnauthorized handles this case with default header values.

The request requires user authentication.
*/
type GetAlarmsUnauthorized struct {
}

func (o *GetAlarmsUnauthorized) Error() string {
	return fmt.Sprintf("[GET /1/user/-/devices/tracker/{tracker-id}/alarms.json][%d] getAlarmsUnauthorized ", 401)
}

func (o *GetAlarmsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
