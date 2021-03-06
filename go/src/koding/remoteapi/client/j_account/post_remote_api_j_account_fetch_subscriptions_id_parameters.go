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

// NewPostRemoteAPIJAccountFetchSubscriptionsIDParams creates a new PostRemoteAPIJAccountFetchSubscriptionsIDParams object
// with the default values initialized.
func NewPostRemoteAPIJAccountFetchSubscriptionsIDParams() *PostRemoteAPIJAccountFetchSubscriptionsIDParams {
	var ()
	return &PostRemoteAPIJAccountFetchSubscriptionsIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostRemoteAPIJAccountFetchSubscriptionsIDParamsWithTimeout creates a new PostRemoteAPIJAccountFetchSubscriptionsIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostRemoteAPIJAccountFetchSubscriptionsIDParamsWithTimeout(timeout time.Duration) *PostRemoteAPIJAccountFetchSubscriptionsIDParams {
	var ()
	return &PostRemoteAPIJAccountFetchSubscriptionsIDParams{

		timeout: timeout,
	}
}

// NewPostRemoteAPIJAccountFetchSubscriptionsIDParamsWithContext creates a new PostRemoteAPIJAccountFetchSubscriptionsIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostRemoteAPIJAccountFetchSubscriptionsIDParamsWithContext(ctx context.Context) *PostRemoteAPIJAccountFetchSubscriptionsIDParams {
	var ()
	return &PostRemoteAPIJAccountFetchSubscriptionsIDParams{

		Context: ctx,
	}
}

/*PostRemoteAPIJAccountFetchSubscriptionsIDParams contains all the parameters to send to the API endpoint
for the post remote API j account fetch subscriptions ID operation typically these are written to a http.Request
*/
type PostRemoteAPIJAccountFetchSubscriptionsIDParams struct {

	/*ID
	  Mongo ID of target instance

	*/
	ID string

	timeout time.Duration
	Context context.Context
}

// WithTimeout adds the timeout to the post remote API j account fetch subscriptions ID params
func (o *PostRemoteAPIJAccountFetchSubscriptionsIDParams) WithTimeout(timeout time.Duration) *PostRemoteAPIJAccountFetchSubscriptionsIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post remote API j account fetch subscriptions ID params
func (o *PostRemoteAPIJAccountFetchSubscriptionsIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post remote API j account fetch subscriptions ID params
func (o *PostRemoteAPIJAccountFetchSubscriptionsIDParams) WithContext(ctx context.Context) *PostRemoteAPIJAccountFetchSubscriptionsIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post remote API j account fetch subscriptions ID params
func (o *PostRemoteAPIJAccountFetchSubscriptionsIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithID adds the id to the post remote API j account fetch subscriptions ID params
func (o *PostRemoteAPIJAccountFetchSubscriptionsIDParams) WithID(id string) *PostRemoteAPIJAccountFetchSubscriptionsIDParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the post remote API j account fetch subscriptions ID params
func (o *PostRemoteAPIJAccountFetchSubscriptionsIDParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *PostRemoteAPIJAccountFetchSubscriptionsIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
