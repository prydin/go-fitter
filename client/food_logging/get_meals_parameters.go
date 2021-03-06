// Code generated by go-swagger; DO NOT EDIT.

package food_logging

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

// NewGetMealsParams creates a new GetMealsParams object
// with the default values initialized.
func NewGetMealsParams() *GetMealsParams {

	return &GetMealsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetMealsParamsWithTimeout creates a new GetMealsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetMealsParamsWithTimeout(timeout time.Duration) *GetMealsParams {

	return &GetMealsParams{

		timeout: timeout,
	}
}

// NewGetMealsParamsWithContext creates a new GetMealsParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetMealsParamsWithContext(ctx context.Context) *GetMealsParams {

	return &GetMealsParams{

		Context: ctx,
	}
}

// NewGetMealsParamsWithHTTPClient creates a new GetMealsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetMealsParamsWithHTTPClient(client *http.Client) *GetMealsParams {

	return &GetMealsParams{
		HTTPClient: client,
	}
}

/*GetMealsParams contains all the parameters to send to the API endpoint
for the get meals operation typically these are written to a http.Request
*/
type GetMealsParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get meals params
func (o *GetMealsParams) WithTimeout(timeout time.Duration) *GetMealsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get meals params
func (o *GetMealsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get meals params
func (o *GetMealsParams) WithContext(ctx context.Context) *GetMealsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get meals params
func (o *GetMealsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get meals params
func (o *GetMealsParams) WithHTTPClient(client *http.Client) *GetMealsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get meals params
func (o *GetMealsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *GetMealsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
