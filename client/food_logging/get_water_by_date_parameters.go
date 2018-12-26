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

// NewGetWaterByDateParams creates a new GetWaterByDateParams object
// with the default values initialized.
func NewGetWaterByDateParams() *GetWaterByDateParams {
	var ()
	return &GetWaterByDateParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetWaterByDateParamsWithTimeout creates a new GetWaterByDateParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetWaterByDateParamsWithTimeout(timeout time.Duration) *GetWaterByDateParams {
	var ()
	return &GetWaterByDateParams{

		timeout: timeout,
	}
}

// NewGetWaterByDateParamsWithContext creates a new GetWaterByDateParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetWaterByDateParamsWithContext(ctx context.Context) *GetWaterByDateParams {
	var ()
	return &GetWaterByDateParams{

		Context: ctx,
	}
}

// NewGetWaterByDateParamsWithHTTPClient creates a new GetWaterByDateParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetWaterByDateParamsWithHTTPClient(client *http.Client) *GetWaterByDateParams {
	var ()
	return &GetWaterByDateParams{
		HTTPClient: client,
	}
}

/*GetWaterByDateParams contains all the parameters to send to the API endpoint
for the get water by date operation typically these are written to a http.Request
*/
type GetWaterByDateParams struct {

	/*Date
	  The date of records to be returned. In the format yyyy-MM-dd.

	*/
	Date strfmt.Date

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get water by date params
func (o *GetWaterByDateParams) WithTimeout(timeout time.Duration) *GetWaterByDateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get water by date params
func (o *GetWaterByDateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get water by date params
func (o *GetWaterByDateParams) WithContext(ctx context.Context) *GetWaterByDateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get water by date params
func (o *GetWaterByDateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get water by date params
func (o *GetWaterByDateParams) WithHTTPClient(client *http.Client) *GetWaterByDateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get water by date params
func (o *GetWaterByDateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDate adds the date to the get water by date params
func (o *GetWaterByDateParams) WithDate(date strfmt.Date) *GetWaterByDateParams {
	o.SetDate(date)
	return o
}

// SetDate adds the date to the get water by date params
func (o *GetWaterByDateParams) SetDate(date strfmt.Date) {
	o.Date = date
}

// WriteToRequest writes these params to a swagger request
func (o *GetWaterByDateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param date
	if err := r.SetPathParam("date", o.Date.String()); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
