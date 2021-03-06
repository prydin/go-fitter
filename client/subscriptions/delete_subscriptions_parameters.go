// Code generated by go-swagger; DO NOT EDIT.

package subscriptions

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

// NewDeleteSubscriptionsParams creates a new DeleteSubscriptionsParams object
// with the default values initialized.
func NewDeleteSubscriptionsParams() *DeleteSubscriptionsParams {
	var ()
	return &DeleteSubscriptionsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteSubscriptionsParamsWithTimeout creates a new DeleteSubscriptionsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteSubscriptionsParamsWithTimeout(timeout time.Duration) *DeleteSubscriptionsParams {
	var ()
	return &DeleteSubscriptionsParams{

		timeout: timeout,
	}
}

// NewDeleteSubscriptionsParamsWithContext creates a new DeleteSubscriptionsParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteSubscriptionsParamsWithContext(ctx context.Context) *DeleteSubscriptionsParams {
	var ()
	return &DeleteSubscriptionsParams{

		Context: ctx,
	}
}

// NewDeleteSubscriptionsParamsWithHTTPClient creates a new DeleteSubscriptionsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteSubscriptionsParamsWithHTTPClient(client *http.Client) *DeleteSubscriptionsParams {
	var ()
	return &DeleteSubscriptionsParams{
		HTTPClient: client,
	}
}

/*DeleteSubscriptionsParams contains all the parameters to send to the API endpoint
for the delete subscriptions operation typically these are written to a http.Request
*/
type DeleteSubscriptionsParams struct {

	/*CollectionPath
	  This is the resource of the collection to receive notifications from (foods, activities, sleep, or body). If not present, subscription will be created for all collections. If you have both all and specific collection subscriptions, you will get duplicate notifications on that collections' updates. Each subscriber can have only one subscription for a specific user's collection.

	*/
	CollectionPath string
	/*SubscriptionID
	  This is the resource of the collection to receive notifications from (foods, activities, sleep, or body). If not present, subscription will be created for all collections. If you have both all and specific collection subscriptions, you will get duplicate notifications on that collections' updates. Each subscriber can have only one subscription for a specific user's collection.

	*/
	SubscriptionID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the delete subscriptions params
func (o *DeleteSubscriptionsParams) WithTimeout(timeout time.Duration) *DeleteSubscriptionsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete subscriptions params
func (o *DeleteSubscriptionsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete subscriptions params
func (o *DeleteSubscriptionsParams) WithContext(ctx context.Context) *DeleteSubscriptionsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete subscriptions params
func (o *DeleteSubscriptionsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete subscriptions params
func (o *DeleteSubscriptionsParams) WithHTTPClient(client *http.Client) *DeleteSubscriptionsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete subscriptions params
func (o *DeleteSubscriptionsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCollectionPath adds the collectionPath to the delete subscriptions params
func (o *DeleteSubscriptionsParams) WithCollectionPath(collectionPath string) *DeleteSubscriptionsParams {
	o.SetCollectionPath(collectionPath)
	return o
}

// SetCollectionPath adds the collectionPath to the delete subscriptions params
func (o *DeleteSubscriptionsParams) SetCollectionPath(collectionPath string) {
	o.CollectionPath = collectionPath
}

// WithSubscriptionID adds the subscriptionID to the delete subscriptions params
func (o *DeleteSubscriptionsParams) WithSubscriptionID(subscriptionID string) *DeleteSubscriptionsParams {
	o.SetSubscriptionID(subscriptionID)
	return o
}

// SetSubscriptionID adds the subscriptionId to the delete subscriptions params
func (o *DeleteSubscriptionsParams) SetSubscriptionID(subscriptionID string) {
	o.SubscriptionID = subscriptionID
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteSubscriptionsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param collection-path
	if err := r.SetPathParam("collection-path", o.CollectionPath); err != nil {
		return err
	}

	// path param subscription-id
	if err := r.SetPathParam("subscription-id", o.SubscriptionID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
