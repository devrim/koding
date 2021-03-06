package j_provisioner

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

// NewPostRemoteAPIJProvisionerSetAccessIDParams creates a new PostRemoteAPIJProvisionerSetAccessIDParams object
// with the default values initialized.
func NewPostRemoteAPIJProvisionerSetAccessIDParams() *PostRemoteAPIJProvisionerSetAccessIDParams {
	var ()
	return &PostRemoteAPIJProvisionerSetAccessIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostRemoteAPIJProvisionerSetAccessIDParamsWithTimeout creates a new PostRemoteAPIJProvisionerSetAccessIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostRemoteAPIJProvisionerSetAccessIDParamsWithTimeout(timeout time.Duration) *PostRemoteAPIJProvisionerSetAccessIDParams {
	var ()
	return &PostRemoteAPIJProvisionerSetAccessIDParams{

		timeout: timeout,
	}
}

// NewPostRemoteAPIJProvisionerSetAccessIDParamsWithContext creates a new PostRemoteAPIJProvisionerSetAccessIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostRemoteAPIJProvisionerSetAccessIDParamsWithContext(ctx context.Context) *PostRemoteAPIJProvisionerSetAccessIDParams {
	var ()
	return &PostRemoteAPIJProvisionerSetAccessIDParams{

		Context: ctx,
	}
}

/*PostRemoteAPIJProvisionerSetAccessIDParams contains all the parameters to send to the API endpoint
for the post remote API j provisioner set access ID operation typically these are written to a http.Request
*/
type PostRemoteAPIJProvisionerSetAccessIDParams struct {

	/*ID
	  Mongo ID of target instance

	*/
	ID string

	timeout time.Duration
	Context context.Context
}

// WithTimeout adds the timeout to the post remote API j provisioner set access ID params
func (o *PostRemoteAPIJProvisionerSetAccessIDParams) WithTimeout(timeout time.Duration) *PostRemoteAPIJProvisionerSetAccessIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post remote API j provisioner set access ID params
func (o *PostRemoteAPIJProvisionerSetAccessIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post remote API j provisioner set access ID params
func (o *PostRemoteAPIJProvisionerSetAccessIDParams) WithContext(ctx context.Context) *PostRemoteAPIJProvisionerSetAccessIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post remote API j provisioner set access ID params
func (o *PostRemoteAPIJProvisionerSetAccessIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithID adds the id to the post remote API j provisioner set access ID params
func (o *PostRemoteAPIJProvisionerSetAccessIDParams) WithID(id string) *PostRemoteAPIJProvisionerSetAccessIDParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the post remote API j provisioner set access ID params
func (o *PostRemoteAPIJProvisionerSetAccessIDParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *PostRemoteAPIJProvisionerSetAccessIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
