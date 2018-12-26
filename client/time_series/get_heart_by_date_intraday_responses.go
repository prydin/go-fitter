// Code generated by go-swagger; DO NOT EDIT.

package time_series

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// GetHeartByDateIntradayReader is a Reader for the GetHeartByDateIntraday structure.
type GetHeartByDateIntradayReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetHeartByDateIntradayReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetHeartByDateIntradayOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 401:
		result := NewGetHeartByDateIntradayUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewGetHeartByDateIntradayForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetHeartByDateIntradayOK creates a GetHeartByDateIntradayOK with default headers values
func NewGetHeartByDateIntradayOK() *GetHeartByDateIntradayOK {
	return &GetHeartByDateIntradayOK{}
}

/*GetHeartByDateIntradayOK handles this case with default header values.

Successful request.
*/
type GetHeartByDateIntradayOK struct {
}

func (o *GetHeartByDateIntradayOK) Error() string {
	return fmt.Sprintf("[GET /1/user/-/activities/heart/date/{date}/1d/{detail-level}.json][%d] getHeartByDateIntradayOK ", 200)
}

func (o *GetHeartByDateIntradayOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetHeartByDateIntradayUnauthorized creates a GetHeartByDateIntradayUnauthorized with default headers values
func NewGetHeartByDateIntradayUnauthorized() *GetHeartByDateIntradayUnauthorized {
	return &GetHeartByDateIntradayUnauthorized{}
}

/*GetHeartByDateIntradayUnauthorized handles this case with default header values.

The request requires user authentication.
*/
type GetHeartByDateIntradayUnauthorized struct {
}

func (o *GetHeartByDateIntradayUnauthorized) Error() string {
	return fmt.Sprintf("[GET /1/user/-/activities/heart/date/{date}/1d/{detail-level}.json][%d] getHeartByDateIntradayUnauthorized ", 401)
}

func (o *GetHeartByDateIntradayUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetHeartByDateIntradayForbidden creates a GetHeartByDateIntradayForbidden with default headers values
func NewGetHeartByDateIntradayForbidden() *GetHeartByDateIntradayForbidden {
	return &GetHeartByDateIntradayForbidden{}
}

/*GetHeartByDateIntradayForbidden handles this case with default header values.

The server understood the request, but is refusing to fulfill it.
*/
type GetHeartByDateIntradayForbidden struct {
}

func (o *GetHeartByDateIntradayForbidden) Error() string {
	return fmt.Sprintf("[GET /1/user/-/activities/heart/date/{date}/1d/{detail-level}.json][%d] getHeartByDateIntradayForbidden ", 403)
}

func (o *GetHeartByDateIntradayForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
