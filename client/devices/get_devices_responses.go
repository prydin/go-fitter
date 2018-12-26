// Code generated by go-swagger; DO NOT EDIT.

package devices

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// GetDevicesReader is a Reader for the GetDevices structure.
type GetDevicesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetDevicesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetDevicesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewGetDevicesBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewGetDevicesUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetDevicesOK creates a GetDevicesOK with default headers values
func NewGetDevicesOK() *GetDevicesOK {
	return &GetDevicesOK{}
}

/*GetDevicesOK handles this case with default header values.

A successful request.
*/
type GetDevicesOK struct {
}

func (o *GetDevicesOK) Error() string {
	return fmt.Sprintf("[GET /1/user/-/devices.json][%d] getDevicesOK ", 200)
}

func (o *GetDevicesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetDevicesBadRequest creates a GetDevicesBadRequest with default headers values
func NewGetDevicesBadRequest() *GetDevicesBadRequest {
	return &GetDevicesBadRequest{}
}

/*GetDevicesBadRequest handles this case with default header values.

The request had bad syntax or was inherently impossible to be satified.
*/
type GetDevicesBadRequest struct {
}

func (o *GetDevicesBadRequest) Error() string {
	return fmt.Sprintf("[GET /1/user/-/devices.json][%d] getDevicesBadRequest ", 400)
}

func (o *GetDevicesBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetDevicesUnauthorized creates a GetDevicesUnauthorized with default headers values
func NewGetDevicesUnauthorized() *GetDevicesUnauthorized {
	return &GetDevicesUnauthorized{}
}

/*GetDevicesUnauthorized handles this case with default header values.

The request requires user authentication.
*/
type GetDevicesUnauthorized struct {
}

func (o *GetDevicesUnauthorized) Error() string {
	return fmt.Sprintf("[GET /1/user/-/devices.json][%d] getDevicesUnauthorized ", 401)
}

func (o *GetDevicesUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}