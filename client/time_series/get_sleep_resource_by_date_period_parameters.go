// Code generated by go-swagger; DO NOT EDIT.

package time_series

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

// NewGetSleepResourceByDatePeriodParams creates a new GetSleepResourceByDatePeriodParams object
// with the default values initialized.
func NewGetSleepResourceByDatePeriodParams() *GetSleepResourceByDatePeriodParams {
	var (
		resourcePathDefault = string("startTime")
	)
	return &GetSleepResourceByDatePeriodParams{
		ResourcePath: resourcePathDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewGetSleepResourceByDatePeriodParamsWithTimeout creates a new GetSleepResourceByDatePeriodParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetSleepResourceByDatePeriodParamsWithTimeout(timeout time.Duration) *GetSleepResourceByDatePeriodParams {
	var (
		resourcePathDefault = string("startTime")
	)
	return &GetSleepResourceByDatePeriodParams{
		ResourcePath: resourcePathDefault,

		timeout: timeout,
	}
}

// NewGetSleepResourceByDatePeriodParamsWithContext creates a new GetSleepResourceByDatePeriodParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetSleepResourceByDatePeriodParamsWithContext(ctx context.Context) *GetSleepResourceByDatePeriodParams {
	var (
		resourcePathDefault = string("startTime")
	)
	return &GetSleepResourceByDatePeriodParams{
		ResourcePath: resourcePathDefault,

		Context: ctx,
	}
}

// NewGetSleepResourceByDatePeriodParamsWithHTTPClient creates a new GetSleepResourceByDatePeriodParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetSleepResourceByDatePeriodParamsWithHTTPClient(client *http.Client) *GetSleepResourceByDatePeriodParams {
	var (
		resourcePathDefault = string("startTime")
	)
	return &GetSleepResourceByDatePeriodParams{
		ResourcePath: resourcePathDefault,
		HTTPClient:   client,
	}
}

/*GetSleepResourceByDatePeriodParams contains all the parameters to send to the API endpoint
for the get sleep resource by date period operation typically these are written to a http.Request
*/
type GetSleepResourceByDatePeriodParams struct {

	/*Date
	  The end date of the period specified in the format yyyy-MM-dd or today.

	*/
	Date strfmt.Date
	/*Period
	  The range for which data will be returned. Options are 1d, 7d, 30d, 1w, 1m, 3m, 6m, 1y, or max.

	*/
	Period string
	/*ResourcePath
	  The resource path; see the Resource Path Options in the full documentation.

	*/
	ResourcePath string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get sleep resource by date period params
func (o *GetSleepResourceByDatePeriodParams) WithTimeout(timeout time.Duration) *GetSleepResourceByDatePeriodParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get sleep resource by date period params
func (o *GetSleepResourceByDatePeriodParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get sleep resource by date period params
func (o *GetSleepResourceByDatePeriodParams) WithContext(ctx context.Context) *GetSleepResourceByDatePeriodParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get sleep resource by date period params
func (o *GetSleepResourceByDatePeriodParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get sleep resource by date period params
func (o *GetSleepResourceByDatePeriodParams) WithHTTPClient(client *http.Client) *GetSleepResourceByDatePeriodParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get sleep resource by date period params
func (o *GetSleepResourceByDatePeriodParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDate adds the date to the get sleep resource by date period params
func (o *GetSleepResourceByDatePeriodParams) WithDate(date strfmt.Date) *GetSleepResourceByDatePeriodParams {
	o.SetDate(date)
	return o
}

// SetDate adds the date to the get sleep resource by date period params
func (o *GetSleepResourceByDatePeriodParams) SetDate(date strfmt.Date) {
	o.Date = date
}

// WithPeriod adds the period to the get sleep resource by date period params
func (o *GetSleepResourceByDatePeriodParams) WithPeriod(period string) *GetSleepResourceByDatePeriodParams {
	o.SetPeriod(period)
	return o
}

// SetPeriod adds the period to the get sleep resource by date period params
func (o *GetSleepResourceByDatePeriodParams) SetPeriod(period string) {
	o.Period = period
}

// WithResourcePath adds the resourcePath to the get sleep resource by date period params
func (o *GetSleepResourceByDatePeriodParams) WithResourcePath(resourcePath string) *GetSleepResourceByDatePeriodParams {
	o.SetResourcePath(resourcePath)
	return o
}

// SetResourcePath adds the resourcePath to the get sleep resource by date period params
func (o *GetSleepResourceByDatePeriodParams) SetResourcePath(resourcePath string) {
	o.ResourcePath = resourcePath
}

// WriteToRequest writes these params to a swagger request
func (o *GetSleepResourceByDatePeriodParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param date
	if err := r.SetPathParam("date", o.Date.String()); err != nil {
		return err
	}

	// path param period
	if err := r.SetPathParam("period", o.Period); err != nil {
		return err
	}

	// path param resource-path
	if err := r.SetPathParam("resource-path", o.ResourcePath); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
