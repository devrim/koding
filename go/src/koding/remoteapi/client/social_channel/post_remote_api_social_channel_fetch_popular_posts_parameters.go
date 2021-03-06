package social_channel

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

// NewPostRemoteAPISocialChannelFetchPopularPostsParams creates a new PostRemoteAPISocialChannelFetchPopularPostsParams object
// with the default values initialized.
func NewPostRemoteAPISocialChannelFetchPopularPostsParams() *PostRemoteAPISocialChannelFetchPopularPostsParams {
	var ()
	return &PostRemoteAPISocialChannelFetchPopularPostsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostRemoteAPISocialChannelFetchPopularPostsParamsWithTimeout creates a new PostRemoteAPISocialChannelFetchPopularPostsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostRemoteAPISocialChannelFetchPopularPostsParamsWithTimeout(timeout time.Duration) *PostRemoteAPISocialChannelFetchPopularPostsParams {
	var ()
	return &PostRemoteAPISocialChannelFetchPopularPostsParams{

		timeout: timeout,
	}
}

// NewPostRemoteAPISocialChannelFetchPopularPostsParamsWithContext creates a new PostRemoteAPISocialChannelFetchPopularPostsParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostRemoteAPISocialChannelFetchPopularPostsParamsWithContext(ctx context.Context) *PostRemoteAPISocialChannelFetchPopularPostsParams {
	var ()
	return &PostRemoteAPISocialChannelFetchPopularPostsParams{

		Context: ctx,
	}
}

/*PostRemoteAPISocialChannelFetchPopularPostsParams contains all the parameters to send to the API endpoint
for the post remote API social channel fetch popular posts operation typically these are written to a http.Request
*/
type PostRemoteAPISocialChannelFetchPopularPostsParams struct {

	/*Body
	  body of the request

	*/
	Body *models.DefaultSelector

	timeout time.Duration
	Context context.Context
}

// WithTimeout adds the timeout to the post remote API social channel fetch popular posts params
func (o *PostRemoteAPISocialChannelFetchPopularPostsParams) WithTimeout(timeout time.Duration) *PostRemoteAPISocialChannelFetchPopularPostsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post remote API social channel fetch popular posts params
func (o *PostRemoteAPISocialChannelFetchPopularPostsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post remote API social channel fetch popular posts params
func (o *PostRemoteAPISocialChannelFetchPopularPostsParams) WithContext(ctx context.Context) *PostRemoteAPISocialChannelFetchPopularPostsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post remote API social channel fetch popular posts params
func (o *PostRemoteAPISocialChannelFetchPopularPostsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithBody adds the body to the post remote API social channel fetch popular posts params
func (o *PostRemoteAPISocialChannelFetchPopularPostsParams) WithBody(body *models.DefaultSelector) *PostRemoteAPISocialChannelFetchPopularPostsParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the post remote API social channel fetch popular posts params
func (o *PostRemoteAPISocialChannelFetchPopularPostsParams) SetBody(body *models.DefaultSelector) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *PostRemoteAPISocialChannelFetchPopularPostsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
