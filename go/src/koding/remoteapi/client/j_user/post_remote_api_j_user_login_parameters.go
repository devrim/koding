package j_user

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"

	"koding/remoteapi/models"
)

// NewPostRemoteAPIJUserLoginParams creates a new PostRemoteAPIJUserLoginParams object
// with the default values initialized.
func NewPostRemoteAPIJUserLoginParams() *PostRemoteAPIJUserLoginParams {
	var ()
	return &PostRemoteAPIJUserLoginParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostRemoteAPIJUserLoginParamsWithTimeout creates a new PostRemoteAPIJUserLoginParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostRemoteAPIJUserLoginParamsWithTimeout(timeout time.Duration) *PostRemoteAPIJUserLoginParams {
	var ()
	return &PostRemoteAPIJUserLoginParams{

		timeout: timeout,
	}
}

// NewPostRemoteAPIJUserLoginParamsWithContext creates a new PostRemoteAPIJUserLoginParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostRemoteAPIJUserLoginParamsWithContext(ctx context.Context) *PostRemoteAPIJUserLoginParams {
	var ()
	return &PostRemoteAPIJUserLoginParams{

		Context: ctx,
	}
}

/*PostRemoteAPIJUserLoginParams contains all the parameters to send to the API endpoint
for the post remote API j user login operation typically these are written to a http.Request
*/
type PostRemoteAPIJUserLoginParams struct {

	/*Body
	  body of the request

	*/
	Body *models.DefaultSelector

	timeout time.Duration
	Context context.Context
}

// WithTimeout adds the timeout to the post remote API j user login params
func (o *PostRemoteAPIJUserLoginParams) WithTimeout(timeout time.Duration) *PostRemoteAPIJUserLoginParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post remote API j user login params
func (o *PostRemoteAPIJUserLoginParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post remote API j user login params
func (o *PostRemoteAPIJUserLoginParams) WithContext(ctx context.Context) *PostRemoteAPIJUserLoginParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post remote API j user login params
func (o *PostRemoteAPIJUserLoginParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithBody adds the body to the post remote API j user login params
func (o *PostRemoteAPIJUserLoginParams) WithBody(body *models.DefaultSelector) *PostRemoteAPIJUserLoginParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the post remote API j user login params
func (o *PostRemoteAPIJUserLoginParams) SetBody(body *models.DefaultSelector) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *PostRemoteAPIJUserLoginParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	r.SetTimeout(o.timeout)
	var res []error

	if o.Body == nil {
		o.Body = new(models.DefaultSelector)
	}

	if err := r.SetBodyParam(o.Body); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
