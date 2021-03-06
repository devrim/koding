package j_account

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"
)

// NewPostRemoteAPIJAccountPushNotificationIDParams creates a new PostRemoteAPIJAccountPushNotificationIDParams object
// with the default values initialized.
func NewPostRemoteAPIJAccountPushNotificationIDParams() *PostRemoteAPIJAccountPushNotificationIDParams {
	var ()
	return &PostRemoteAPIJAccountPushNotificationIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostRemoteAPIJAccountPushNotificationIDParamsWithTimeout creates a new PostRemoteAPIJAccountPushNotificationIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostRemoteAPIJAccountPushNotificationIDParamsWithTimeout(timeout time.Duration) *PostRemoteAPIJAccountPushNotificationIDParams {
	var ()
	return &PostRemoteAPIJAccountPushNotificationIDParams{

		timeout: timeout,
	}
}

// NewPostRemoteAPIJAccountPushNotificationIDParamsWithContext creates a new PostRemoteAPIJAccountPushNotificationIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostRemoteAPIJAccountPushNotificationIDParamsWithContext(ctx context.Context) *PostRemoteAPIJAccountPushNotificationIDParams {
	var ()
	return &PostRemoteAPIJAccountPushNotificationIDParams{

		Context: ctx,
	}
}

/*PostRemoteAPIJAccountPushNotificationIDParams contains all the parameters to send to the API endpoint
for the post remote API j account push notification ID operation typically these are written to a http.Request
*/
type PostRemoteAPIJAccountPushNotificationIDParams struct {

	/*ID
	  Mongo ID of target instance

	*/
	ID string

	timeout time.Duration
	Context context.Context
}

// WithTimeout adds the timeout to the post remote API j account push notification ID params
func (o *PostRemoteAPIJAccountPushNotificationIDParams) WithTimeout(timeout time.Duration) *PostRemoteAPIJAccountPushNotificationIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post remote API j account push notification ID params
func (o *PostRemoteAPIJAccountPushNotificationIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post remote API j account push notification ID params
func (o *PostRemoteAPIJAccountPushNotificationIDParams) WithContext(ctx context.Context) *PostRemoteAPIJAccountPushNotificationIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post remote API j account push notification ID params
func (o *PostRemoteAPIJAccountPushNotificationIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithID adds the id to the post remote API j account push notification ID params
func (o *PostRemoteAPIJAccountPushNotificationIDParams) WithID(id string) *PostRemoteAPIJAccountPushNotificationIDParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the post remote API j account push notification ID params
func (o *PostRemoteAPIJAccountPushNotificationIDParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *PostRemoteAPIJAccountPushNotificationIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	r.SetTimeout(o.timeout)
	var res []error

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
