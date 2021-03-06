// Code generated by go-swagger; DO NOT EDIT.

package runtime_manager

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"

	"openpitrix.io/openpitrix/test/models"
)

// NewDeleteRuntimesParams creates a new DeleteRuntimesParams object
// with the default values initialized.
func NewDeleteRuntimesParams() *DeleteRuntimesParams {
	var ()
	return &DeleteRuntimesParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteRuntimesParamsWithTimeout creates a new DeleteRuntimesParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteRuntimesParamsWithTimeout(timeout time.Duration) *DeleteRuntimesParams {
	var ()
	return &DeleteRuntimesParams{

		timeout: timeout,
	}
}

// NewDeleteRuntimesParamsWithContext creates a new DeleteRuntimesParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteRuntimesParamsWithContext(ctx context.Context) *DeleteRuntimesParams {
	var ()
	return &DeleteRuntimesParams{

		Context: ctx,
	}
}

// NewDeleteRuntimesParamsWithHTTPClient creates a new DeleteRuntimesParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteRuntimesParamsWithHTTPClient(client *http.Client) *DeleteRuntimesParams {
	var ()
	return &DeleteRuntimesParams{
		HTTPClient: client,
	}
}

/*DeleteRuntimesParams contains all the parameters to send to the API endpoint
for the delete runtimes operation typically these are written to a http.Request
*/
type DeleteRuntimesParams struct {

	/*Body*/
	Body *models.OpenpitrixDeleteRuntimesRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the delete runtimes params
func (o *DeleteRuntimesParams) WithTimeout(timeout time.Duration) *DeleteRuntimesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete runtimes params
func (o *DeleteRuntimesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete runtimes params
func (o *DeleteRuntimesParams) WithContext(ctx context.Context) *DeleteRuntimesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete runtimes params
func (o *DeleteRuntimesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete runtimes params
func (o *DeleteRuntimesParams) WithHTTPClient(client *http.Client) *DeleteRuntimesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete runtimes params
func (o *DeleteRuntimesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the delete runtimes params
func (o *DeleteRuntimesParams) WithBody(body *models.OpenpitrixDeleteRuntimesRequest) *DeleteRuntimesParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the delete runtimes params
func (o *DeleteRuntimesParams) SetBody(body *models.OpenpitrixDeleteRuntimesRequest) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteRuntimesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
