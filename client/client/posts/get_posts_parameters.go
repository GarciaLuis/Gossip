// Code generated by go-swagger; DO NOT EDIT.

package posts

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

// NewGetPostsParams creates a new GetPostsParams object
// with the default values initialized.
func NewGetPostsParams() *GetPostsParams {

	return &GetPostsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetPostsParamsWithTimeout creates a new GetPostsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetPostsParamsWithTimeout(timeout time.Duration) *GetPostsParams {

	return &GetPostsParams{

		timeout: timeout,
	}
}

// NewGetPostsParamsWithContext creates a new GetPostsParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetPostsParamsWithContext(ctx context.Context) *GetPostsParams {

	return &GetPostsParams{

		Context: ctx,
	}
}

// NewGetPostsParamsWithHTTPClient creates a new GetPostsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetPostsParamsWithHTTPClient(client *http.Client) *GetPostsParams {

	return &GetPostsParams{
		HTTPClient: client,
	}
}

/*GetPostsParams contains all the parameters to send to the API endpoint
for the get posts operation typically these are written to a http.Request
*/
type GetPostsParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get posts params
func (o *GetPostsParams) WithTimeout(timeout time.Duration) *GetPostsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get posts params
func (o *GetPostsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get posts params
func (o *GetPostsParams) WithContext(ctx context.Context) *GetPostsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get posts params
func (o *GetPostsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get posts params
func (o *GetPostsParams) WithHTTPClient(client *http.Client) *GetPostsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get posts params
func (o *GetPostsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *GetPostsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}