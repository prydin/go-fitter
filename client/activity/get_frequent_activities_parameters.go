// Code generated by go-swagger; DO NOT EDIT.

package activity

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

// NewGetFrequentActivitiesParams creates a new GetFrequentActivitiesParams object
// with the default values initialized.
func NewGetFrequentActivitiesParams() *GetFrequentActivitiesParams {

	return &GetFrequentActivitiesParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetFrequentActivitiesParamsWithTimeout creates a new GetFrequentActivitiesParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetFrequentActivitiesParamsWithTimeout(timeout time.Duration) *GetFrequentActivitiesParams {

	return &GetFrequentActivitiesParams{

		timeout: timeout,
	}
}

// NewGetFrequentActivitiesParamsWithContext creates a new GetFrequentActivitiesParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetFrequentActivitiesParamsWithContext(ctx context.Context) *GetFrequentActivitiesParams {

	return &GetFrequentActivitiesParams{

		Context: ctx,
	}
}

// NewGetFrequentActivitiesParamsWithHTTPClient creates a new GetFrequentActivitiesParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetFrequentActivitiesParamsWithHTTPClient(client *http.Client) *GetFrequentActivitiesParams {

	return &GetFrequentActivitiesParams{
		HTTPClient: client,
	}
}

/*GetFrequentActivitiesParams contains all the parameters to send to the API endpoint
for the get frequent activities operation typically these are written to a http.Request
*/
type GetFrequentActivitiesParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get frequent activities params
func (o *GetFrequentActivitiesParams) WithTimeout(timeout time.Duration) *GetFrequentActivitiesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get frequent activities params
func (o *GetFrequentActivitiesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get frequent activities params
func (o *GetFrequentActivitiesParams) WithContext(ctx context.Context) *GetFrequentActivitiesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get frequent activities params
func (o *GetFrequentActivitiesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get frequent activities params
func (o *GetFrequentActivitiesParams) WithHTTPClient(client *http.Client) *GetFrequentActivitiesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get frequent activities params
func (o *GetFrequentActivitiesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *GetFrequentActivitiesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}