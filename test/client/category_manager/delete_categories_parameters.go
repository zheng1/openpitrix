// Code generated by go-swagger; DO NOT EDIT.

package category_manager

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

	models "openpitrix.io/openpitrix/test/models"
)

// NewDeleteCategoriesParams creates a new DeleteCategoriesParams object
// with the default values initialized.
func NewDeleteCategoriesParams() *DeleteCategoriesParams {
	var ()
	return &DeleteCategoriesParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteCategoriesParamsWithTimeout creates a new DeleteCategoriesParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteCategoriesParamsWithTimeout(timeout time.Duration) *DeleteCategoriesParams {
	var ()
	return &DeleteCategoriesParams{

		timeout: timeout,
	}
}

// NewDeleteCategoriesParamsWithContext creates a new DeleteCategoriesParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteCategoriesParamsWithContext(ctx context.Context) *DeleteCategoriesParams {
	var ()
	return &DeleteCategoriesParams{

		Context: ctx,
	}
}

// NewDeleteCategoriesParamsWithHTTPClient creates a new DeleteCategoriesParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteCategoriesParamsWithHTTPClient(client *http.Client) *DeleteCategoriesParams {
	var ()
	return &DeleteCategoriesParams{
		HTTPClient: client,
	}
}

/*DeleteCategoriesParams contains all the parameters to send to the API endpoint
for the delete categories operation typically these are written to a http.Request
*/
type DeleteCategoriesParams struct {

	/*Body*/
	Body *models.OpenpitrixDeleteCategoriesRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the delete categories params
func (o *DeleteCategoriesParams) WithTimeout(timeout time.Duration) *DeleteCategoriesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete categories params
func (o *DeleteCategoriesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete categories params
func (o *DeleteCategoriesParams) WithContext(ctx context.Context) *DeleteCategoriesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete categories params
func (o *DeleteCategoriesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete categories params
func (o *DeleteCategoriesParams) WithHTTPClient(client *http.Client) *DeleteCategoriesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete categories params
func (o *DeleteCategoriesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the delete categories params
func (o *DeleteCategoriesParams) WithBody(body *models.OpenpitrixDeleteCategoriesRequest) *DeleteCategoriesParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the delete categories params
func (o *DeleteCategoriesParams) SetBody(body *models.OpenpitrixDeleteCategoriesRequest) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteCategoriesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
