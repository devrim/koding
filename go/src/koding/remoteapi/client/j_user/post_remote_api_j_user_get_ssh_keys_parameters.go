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
)

// NewPostRemoteAPIJUserGetSSHKeysParams creates a new PostRemoteAPIJUserGetSSHKeysParams object
// with the default values initialized.
func NewPostRemoteAPIJUserGetSSHKeysParams() *PostRemoteAPIJUserGetSSHKeysParams {

	return &PostRemoteAPIJUserGetSSHKeysParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostRemoteAPIJUserGetSSHKeysParamsWithTimeout creates a new PostRemoteAPIJUserGetSSHKeysParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostRemoteAPIJUserGetSSHKeysParamsWithTimeout(timeout time.Duration) *PostRemoteAPIJUserGetSSHKeysParams {

	return &PostRemoteAPIJUserGetSSHKeysParams{

		timeout: timeout,
	}
}

// NewPostRemoteAPIJUserGetSSHKeysParamsWithContext creates a new PostRemoteAPIJUserGetSSHKeysParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostRemoteAPIJUserGetSSHKeysParamsWithContext(ctx context.Context) *PostRemoteAPIJUserGetSSHKeysParams {

	return &PostRemoteAPIJUserGetSSHKeysParams{

		Context: ctx,
	}
}

/*PostRemoteAPIJUserGetSSHKeysParams contains all the parameters to send to the API endpoint
for the post remote API j user get SSH keys operation typically these are written to a http.Request
*/
type PostRemoteAPIJUserGetSSHKeysParams struct {
	timeout time.Duration
	Context context.Context
}

// WithTimeout adds the timeout to the post remote API j user get SSH keys params
func (o *PostRemoteAPIJUserGetSSHKeysParams) WithTimeout(timeout time.Duration) *PostRemoteAPIJUserGetSSHKeysParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post remote API j user get SSH keys params
func (o *PostRemoteAPIJUserGetSSHKeysParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post remote API j user get SSH keys params
func (o *PostRemoteAPIJUserGetSSHKeysParams) WithContext(ctx context.Context) *PostRemoteAPIJUserGetSSHKeysParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post remote API j user get SSH keys params
func (o *PostRemoteAPIJUserGetSSHKeysParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WriteToRequest writes these params to a swagger request
func (o *PostRemoteAPIJUserGetSSHKeysParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	r.SetTimeout(o.timeout)
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
