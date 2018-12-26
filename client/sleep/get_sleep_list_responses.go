// Code generated by go-swagger; DO NOT EDIT.

package sleep

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// GetSleepListReader is a Reader for the GetSleepList structure.
type GetSleepListReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetSleepListReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetSleepListOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 401:
		result := NewGetSleepListUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewGetSleepListForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetSleepListOK creates a GetSleepListOK with default headers values
func NewGetSleepListOK() *GetSleepListOK {
	return &GetSleepListOK{}
}

/*GetSleepListOK handles this case with default header values.

Successful request.
*/
type GetSleepListOK struct {
}

func (o *GetSleepListOK) Error() string {
	return fmt.Sprintf("[GET /1.2/user/-/sleep/list.json][%d] getSleepListOK ", 200)
}

func (o *GetSleepListOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetSleepListUnauthorized creates a GetSleepListUnauthorized with default headers values
func NewGetSleepListUnauthorized() *GetSleepListUnauthorized {
	return &GetSleepListUnauthorized{}
}

/*GetSleepListUnauthorized handles this case with default header values.

The request requires user authentication.
*/
type GetSleepListUnauthorized struct {
}

func (o *GetSleepListUnauthorized) Error() string {
	return fmt.Sprintf("[GET /1.2/user/-/sleep/list.json][%d] getSleepListUnauthorized ", 401)
}

func (o *GetSleepListUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetSleepListForbidden creates a GetSleepListForbidden with default headers values
func NewGetSleepListForbidden() *GetSleepListForbidden {
	return &GetSleepListForbidden{}
}

/*GetSleepListForbidden handles this case with default header values.

The server understood the request, but is refusing to fulfill it.
*/
type GetSleepListForbidden struct {
}

func (o *GetSleepListForbidden) Error() string {
	return fmt.Sprintf("[GET /1.2/user/-/sleep/list.json][%d] getSleepListForbidden ", 403)
}

func (o *GetSleepListForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
