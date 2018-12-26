// Code generated by go-swagger; DO NOT EDIT.

package auth

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// OauthTokenReader is a Reader for the OauthToken structure.
type OauthTokenReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *OauthTokenReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewOauthTokenOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewOauthTokenBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewOauthTokenUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 409:
		result := NewOauthTokenConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewOauthTokenOK creates a OauthTokenOK with default headers values
func NewOauthTokenOK() *OauthTokenOK {
	return &OauthTokenOK{}
}

/*OauthTokenOK handles this case with default header values.

A successful request.
*/
type OauthTokenOK struct {
}

func (o *OauthTokenOK) Error() string {
	return fmt.Sprintf("[POST /oauth2/token][%d] oauthTokenOK ", 200)
}

func (o *OauthTokenOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewOauthTokenBadRequest creates a OauthTokenBadRequest with default headers values
func NewOauthTokenBadRequest() *OauthTokenBadRequest {
	return &OauthTokenBadRequest{}
}

/*OauthTokenBadRequest handles this case with default header values.

The request had bad syntax or was inherently impossible to be satisfied.
*/
type OauthTokenBadRequest struct {
}

func (o *OauthTokenBadRequest) Error() string {
	return fmt.Sprintf("[POST /oauth2/token][%d] oauthTokenBadRequest ", 400)
}

func (o *OauthTokenBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewOauthTokenUnauthorized creates a OauthTokenUnauthorized with default headers values
func NewOauthTokenUnauthorized() *OauthTokenUnauthorized {
	return &OauthTokenUnauthorized{}
}

/*OauthTokenUnauthorized handles this case with default header values.

Authentication was unsuccessful due to invalid client credentials.
*/
type OauthTokenUnauthorized struct {
}

func (o *OauthTokenUnauthorized) Error() string {
	return fmt.Sprintf("[POST /oauth2/token][%d] oauthTokenUnauthorized ", 401)
}

func (o *OauthTokenUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewOauthTokenConflict creates a OauthTokenConflict with default headers values
func NewOauthTokenConflict() *OauthTokenConflict {
	return &OauthTokenConflict{}
}

/*OauthTokenConflict handles this case with default header values.

Request conflict due to multiple clients attempting to refresh the same access token.
*/
type OauthTokenConflict struct {
}

func (o *OauthTokenConflict) Error() string {
	return fmt.Sprintf("[POST /oauth2/token][%d] oauthTokenConflict ", 409)
}

func (o *OauthTokenConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
