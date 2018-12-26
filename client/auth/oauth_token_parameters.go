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

// NewOauthTokenParams creates a new OauthTokenParams object
// with the default values initialized.
func NewOauthTokenParams() *OauthTokenParams {
	var ()
	return &OauthTokenParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewOauthTokenParamsWithTimeout creates a new OauthTokenParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewOauthTokenParamsWithTimeout(timeout time.Duration) *OauthTokenParams {
	var ()
	return &OauthTokenParams{

		timeout: timeout,
	}
}

// NewOauthTokenParamsWithContext creates a new OauthTokenParams object
// with the default values initialized, and the ability to set a context for a request
func NewOauthTokenParamsWithContext(ctx context.Context) *OauthTokenParams {
	var ()
	return &OauthTokenParams{

		Context: ctx,
	}
}

// NewOauthTokenParamsWithHTTPClient creates a new OauthTokenParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewOauthTokenParamsWithHTTPClient(client *http.Client) *OauthTokenParams {
	var ()
	return &OauthTokenParams{
		HTTPClient: client,
	}
}

/*OauthTokenParams contains all the parameters to send to the API endpoint
for the oauth token operation typically these are written to a http.Request
*/
type OauthTokenParams struct {

	/*Authorization
	  The Authorization header must be set to 'Basic' followed by a space, then the Base64 encoded string of your application's client id and secret concatenated with a colon. For example, 'Basic Y2xpZW50X2lkOmNsaWVudCBzZWNyZXQ='. The Base64 encoded string, 'Y2xpZW50X2lkOmNsaWVudCBzZWNyZXQ=', is decoded as 'client_id:client secret'.

	*/
	Authorization *string
	/*ClientID
	  This is your Fitbit API application id from your settings on dev.fitbit.com.

	*/
	ClientID string
	/*Code
	  Authorization code received in the redirect as URI parameter. Required if using the Authorization Code flow.

	*/
	Code *string
	/*ExpiresIn
	  Specify the desired access token lifetime. Defaults to 28800 for 8 hours. The other valid value is 3600 for 1 hour.

	*/
	ExpiresIn *string
	/*GrantType
	  Authorization grant type. Valid values are 'authorization_code' and 'refresh_token'.

	*/
	GrantType string
	/*RedirectURI
	  Uri to which the access token will be sent if the request is successful. Required if specified in the redirect to the authorization page. Must be exact match.

	*/
	RedirectURI *string
	/*RefreshToken
	  Refresh token issued by Fitbit. Required if 'grant_type' is 'refresh_token'.

	*/
	RefreshToken *string
	/*State
	  Required if specified in the redirect uri of the authorization page. Must be an exact match.

	*/
	State *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the oauth token params
func (o *OauthTokenParams) WithTimeout(timeout time.Duration) *OauthTokenParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the oauth token params
func (o *OauthTokenParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the oauth token params
func (o *OauthTokenParams) WithContext(ctx context.Context) *OauthTokenParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the oauth token params
func (o *OauthTokenParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the oauth token params
func (o *OauthTokenParams) WithHTTPClient(client *http.Client) *OauthTokenParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the oauth token params
func (o *OauthTokenParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAuthorization adds the authorization to the oauth token params
func (o *OauthTokenParams) WithAuthorization(authorization *string) *OauthTokenParams {
	o.SetAuthorization(authorization)
	return o
}

// SetAuthorization adds the authorization to the oauth token params
func (o *OauthTokenParams) SetAuthorization(authorization *string) {
	o.Authorization = authorization
}

// WithClientID adds the clientID to the oauth token params
func (o *OauthTokenParams) WithClientID(clientID string) *OauthTokenParams {
	o.SetClientID(clientID)
	return o
}

// SetClientID adds the clientId to the oauth token params
func (o *OauthTokenParams) SetClientID(clientID string) {
	o.ClientID = clientID
}

// WithCode adds the code to the oauth token params
func (o *OauthTokenParams) WithCode(code *string) *OauthTokenParams {
	o.SetCode(code)
	return o
}

// SetCode adds the code to the oauth token params
func (o *OauthTokenParams) SetCode(code *string) {
	o.Code = code
}

// WithExpiresIn adds the expiresIn to the oauth token params
func (o *OauthTokenParams) WithExpiresIn(expiresIn *string) *OauthTokenParams {
	o.SetExpiresIn(expiresIn)
	return o
}

// SetExpiresIn adds the expiresIn to the oauth token params
func (o *OauthTokenParams) SetExpiresIn(expiresIn *string) {
	o.ExpiresIn = expiresIn
}

// WithGrantType adds the grantType to the oauth token params
func (o *OauthTokenParams) WithGrantType(grantType string) *OauthTokenParams {
	o.SetGrantType(grantType)
	return o
}

// SetGrantType adds the grantType to the oauth token params
func (o *OauthTokenParams) SetGrantType(grantType string) {
	o.GrantType = grantType
}

// WithRedirectURI adds the redirectURI to the oauth token params
func (o *OauthTokenParams) WithRedirectURI(redirectURI *string) *OauthTokenParams {
	o.SetRedirectURI(redirectURI)
	return o
}

// SetRedirectURI adds the redirectUri to the oauth token params
func (o *OauthTokenParams) SetRedirectURI(redirectURI *string) {
	o.RedirectURI = redirectURI
}

// WithRefreshToken adds the refreshToken to the oauth token params
func (o *OauthTokenParams) WithRefreshToken(refreshToken *string) *OauthTokenParams {
	o.SetRefreshToken(refreshToken)
	return o
}

// SetRefreshToken adds the refreshToken to the oauth token params
func (o *OauthTokenParams) SetRefreshToken(refreshToken *string) {
	o.RefreshToken = refreshToken
}

// WithState adds the state to the oauth token params
func (o *OauthTokenParams) WithState(state *string) *OauthTokenParams {
	o.SetState(state)
	return o
}

// SetState adds the state to the oauth token params
func (o *OauthTokenParams) SetState(state *string) {
	o.State = state
}

// WriteToRequest writes these params to a swagger request
func (o *OauthTokenParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Authorization != nil {

		// header param Authorization
		if err := r.SetHeaderParam("Authorization", *o.Authorization); err != nil {
			return err
		}

	}

	// query param client_id
	qrClientID := o.ClientID
	qClientID := qrClientID
	if qClientID != "" {
		if err := r.SetQueryParam("client_id", qClientID); err != nil {
			return err
		}
	}

	if o.Code != nil {

		// query param code
		var qrCode string
		if o.Code != nil {
			qrCode = *o.Code
		}
		qCode := qrCode
		if qCode != "" {
			if err := r.SetQueryParam("code", qCode); err != nil {
				return err
			}
		}

	}

	if o.ExpiresIn != nil {

		// query param expires_in
		var qrExpiresIn string
		if o.ExpiresIn != nil {
			qrExpiresIn = *o.ExpiresIn
		}
		qExpiresIn := qrExpiresIn
		if qExpiresIn != "" {
			if err := r.SetQueryParam("expires_in", qExpiresIn); err != nil {
				return err
			}
		}

	}

	// query param grant_type
	qrGrantType := o.GrantType
	qGrantType := qrGrantType
	if qGrantType != "" {
		if err := r.SetQueryParam("grant_type", qGrantType); err != nil {
			return err
		}
	}

	if o.RedirectURI != nil {

		// query param redirect_uri
		var qrRedirectURI string
		if o.RedirectURI != nil {
			qrRedirectURI = *o.RedirectURI
		}
		qRedirectURI := qrRedirectURI
		if qRedirectURI != "" {
			if err := r.SetQueryParam("redirect_uri", qRedirectURI); err != nil {
				return err
			}
		}

	}

	if o.RefreshToken != nil {

		// query param refresh_token
		var qrRefreshToken string
		if o.RefreshToken != nil {
			qrRefreshToken = *o.RefreshToken
		}
		qRefreshToken := qrRefreshToken
		if qRefreshToken != "" {
			if err := r.SetQueryParam("refresh_token", qRefreshToken); err != nil {
				return err
			}
		}

	}

	if o.State != nil {

		// query param state
		var qrState string
		if o.State != nil {
			qrState = *o.State
		}
		qState := qrState
		if qState != "" {
			if err := r.SetQueryParam("state", qState); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
