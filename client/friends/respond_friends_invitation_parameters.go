// Code generated by go-swagger; DO NOT EDIT.

package friends

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

// NewRespondFriendsInvitationParams creates a new RespondFriendsInvitationParams object
// with the default values initialized.
func NewRespondFriendsInvitationParams() *RespondFriendsInvitationParams {
	var ()
	return &RespondFriendsInvitationParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewRespondFriendsInvitationParamsWithTimeout creates a new RespondFriendsInvitationParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewRespondFriendsInvitationParamsWithTimeout(timeout time.Duration) *RespondFriendsInvitationParams {
	var ()
	return &RespondFriendsInvitationParams{

		timeout: timeout,
	}
}

// NewRespondFriendsInvitationParamsWithContext creates a new RespondFriendsInvitationParams object
// with the default values initialized, and the ability to set a context for a request
func NewRespondFriendsInvitationParamsWithContext(ctx context.Context) *RespondFriendsInvitationParams {
	var ()
	return &RespondFriendsInvitationParams{

		Context: ctx,
	}
}

// NewRespondFriendsInvitationParamsWithHTTPClient creates a new RespondFriendsInvitationParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewRespondFriendsInvitationParamsWithHTTPClient(client *http.Client) *RespondFriendsInvitationParams {
	var ()
	return &RespondFriendsInvitationParams{
		HTTPClient: client,
	}
}

/*RespondFriendsInvitationParams contains all the parameters to send to the API endpoint
for the respond friends invitation operation typically these are written to a http.Request
*/
type RespondFriendsInvitationParams struct {

	/*Accept
	  Accept or reject invitation; true or false.

	*/
	Accept string
	/*FromUserID
	  The encoded ID of a user from which to accept or reject invitation.

	*/
	FromUserID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the respond friends invitation params
func (o *RespondFriendsInvitationParams) WithTimeout(timeout time.Duration) *RespondFriendsInvitationParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the respond friends invitation params
func (o *RespondFriendsInvitationParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the respond friends invitation params
func (o *RespondFriendsInvitationParams) WithContext(ctx context.Context) *RespondFriendsInvitationParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the respond friends invitation params
func (o *RespondFriendsInvitationParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the respond friends invitation params
func (o *RespondFriendsInvitationParams) WithHTTPClient(client *http.Client) *RespondFriendsInvitationParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the respond friends invitation params
func (o *RespondFriendsInvitationParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAccept adds the accept to the respond friends invitation params
func (o *RespondFriendsInvitationParams) WithAccept(accept string) *RespondFriendsInvitationParams {
	o.SetAccept(accept)
	return o
}

// SetAccept adds the accept to the respond friends invitation params
func (o *RespondFriendsInvitationParams) SetAccept(accept string) {
	o.Accept = accept
}

// WithFromUserID adds the fromUserID to the respond friends invitation params
func (o *RespondFriendsInvitationParams) WithFromUserID(fromUserID string) *RespondFriendsInvitationParams {
	o.SetFromUserID(fromUserID)
	return o
}

// SetFromUserID adds the fromUserId to the respond friends invitation params
func (o *RespondFriendsInvitationParams) SetFromUserID(fromUserID string) {
	o.FromUserID = fromUserID
}

// WriteToRequest writes these params to a swagger request
func (o *RespondFriendsInvitationParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// query param accept
	qrAccept := o.Accept
	qAccept := qrAccept
	if qAccept != "" {
		if err := r.SetQueryParam("accept", qAccept); err != nil {
			return err
		}
	}

	// path param from-user-id
	if err := r.SetPathParam("from-user-id", o.FromUserID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}