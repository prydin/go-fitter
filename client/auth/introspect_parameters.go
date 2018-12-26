// Code generated by go-swagger; DO NOT EDIT.

package auth

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"
)

// NewIntrospectParams creates a new IntrospectParams object
// with the default values initialized.
func NewIntrospectParams() *IntrospectParams {
	var ()
	return &IntrospectParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewIntrospectParamsWithTimeout creates a new IntrospectParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewIntrospectParamsWithTimeout(timeout time.Duration) *IntrospectParams {
	var ()
	return &IntrospectParams{

		timeout: timeout,
	}
}

// NewIntrospectParamsWithContext creates a new IntrospectParams object
// with the default values initialized, and the ability to set a context for a request
func NewIntrospectParamsWithContext(ctx context.Context) *IntrospectParams {
	var ()
	return &IntrospectParams{

		Context: ctx,
	}
}

// NewIntrospectParamsWithHTTPClient creates a new IntrospectParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewIntrospectParamsWithHTTPClient(client *http.Client) *IntrospectParams {
	var ()
	return &IntrospectParams{
		HTTPClient: client,
	}
}

/*IntrospectParams contains all the parameters to send to the API endpoint
for the introspect operation typically these are written to a http.Request
*/
type IntrospectParams struct {

	/*Token
	  OAuth 2.0 token to retrieve the state of

	*/
	Token string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the introspect params
func (o *IntrospectParams) WithTimeout(timeout time.Duration) *IntrospectParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the introspect params
func (o *IntrospectParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the introspect params
func (o *IntrospectParams) WithContext(ctx context.Context) *IntrospectParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the introspect params
func (o *IntrospectParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the introspect params
func (o *IntrospectParams) WithHTTPClient(client *http.Client) *IntrospectParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the introspect params
func (o *IntrospectParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithToken adds the token to the introspect params
func (o *IntrospectParams) WithToken(token string) *IntrospectParams {
	o.SetToken(token)
	return o
}

// SetToken adds the token to the introspect params
func (o *IntrospectParams) SetToken(token string) {
	o.Token = token
}

// WriteToRequest writes these params to a swagger request
func (o *IntrospectParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// form param token
	frToken := o.Token
	fToken := frToken
	if fToken != "" {
		if err := r.SetFormParam("token", fToken); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
