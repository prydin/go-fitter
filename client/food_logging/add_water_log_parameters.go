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
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"
)

// NewAddWaterLogParams creates a new AddWaterLogParams object
// with the default values initialized.
func NewAddWaterLogParams() *AddWaterLogParams {
	var ()
	return &AddWaterLogParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewAddWaterLogParamsWithTimeout creates a new AddWaterLogParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewAddWaterLogParamsWithTimeout(timeout time.Duration) *AddWaterLogParams {
	var ()
	return &AddWaterLogParams{

		timeout: timeout,
	}
}

// NewAddWaterLogParamsWithContext creates a new AddWaterLogParams object
// with the default values initialized, and the ability to set a context for a request
func NewAddWaterLogParamsWithContext(ctx context.Context) *AddWaterLogParams {
	var ()
	return &AddWaterLogParams{

		Context: ctx,
	}
}

// NewAddWaterLogParamsWithHTTPClient creates a new AddWaterLogParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewAddWaterLogParamsWithHTTPClient(client *http.Client) *AddWaterLogParams {
	var ()
	return &AddWaterLogParams{
		HTTPClient: client,
	}
}

/*AddWaterLogParams contains all the parameters to send to the API endpoint
for the add water log operation typically these are written to a http.Request
*/
type AddWaterLogParams struct {

	/*Amount
	  The amount consumption in the format X.XX and in the specified waterUnit or in the unit system that corresponds to the Accept-Language header provided.

	*/
	Amount int64
	/*Date
	  The date of records to be returned in the format yyyy-MM-dd.

	*/
	Date strfmt.Date
	/*Unit
	  Water measurement unit; `ml`, `fl oz`, or `cup`.

	*/
	Unit *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the add water log params
func (o *AddWaterLogParams) WithTimeout(timeout time.Duration) *AddWaterLogParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the add water log params
func (o *AddWaterLogParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the add water log params
func (o *AddWaterLogParams) WithContext(ctx context.Context) *AddWaterLogParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the add water log params
func (o *AddWaterLogParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the add water log params
func (o *AddWaterLogParams) WithHTTPClient(client *http.Client) *AddWaterLogParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the add water log params
func (o *AddWaterLogParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAmount adds the amount to the add water log params
func (o *AddWaterLogParams) WithAmount(amount int64) *AddWaterLogParams {
	o.SetAmount(amount)
	return o
}

// SetAmount adds the amount to the add water log params
func (o *AddWaterLogParams) SetAmount(amount int64) {
	o.Amount = amount
}

// WithDate adds the date to the add water log params
func (o *AddWaterLogParams) WithDate(date strfmt.Date) *AddWaterLogParams {
	o.SetDate(date)
	return o
}

// SetDate adds the date to the add water log params
func (o *AddWaterLogParams) SetDate(date strfmt.Date) {
	o.Date = date
}

// WithUnit adds the unit to the add water log params
func (o *AddWaterLogParams) WithUnit(unit *string) *AddWaterLogParams {
	o.SetUnit(unit)
	return o
}

// SetUnit adds the unit to the add water log params
func (o *AddWaterLogParams) SetUnit(unit *string) {
	o.Unit = unit
}

// WriteToRequest writes these params to a swagger request
func (o *AddWaterLogParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// query param amount
	qrAmount := o.Amount
	qAmount := swag.FormatInt64(qrAmount)
	if qAmount != "" {
		if err := r.SetQueryParam("amount", qAmount); err != nil {
			return err
		}
	}

	// query param date
	qrDate := o.Date
	qDate := qrDate.String()
	if qDate != "" {
		if err := r.SetQueryParam("date", qDate); err != nil {
			return err
		}
	}

	if o.Unit != nil {

		// query param unit
		var qrUnit string
		if o.Unit != nil {
			qrUnit = *o.Unit
		}
		qUnit := qrUnit
		if qUnit != "" {
			if err := r.SetQueryParam("unit", qUnit); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
