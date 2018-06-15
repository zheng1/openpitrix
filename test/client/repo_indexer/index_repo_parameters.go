// Code generated by go-swagger; DO NOT EDIT.

package repo_indexer

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

// NewIndexRepoParams creates a new IndexRepoParams object
// with the default values initialized.
func NewIndexRepoParams() *IndexRepoParams {
	var ()
	return &IndexRepoParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewIndexRepoParamsWithTimeout creates a new IndexRepoParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewIndexRepoParamsWithTimeout(timeout time.Duration) *IndexRepoParams {
	var ()
	return &IndexRepoParams{

		timeout: timeout,
	}
}

// NewIndexRepoParamsWithContext creates a new IndexRepoParams object
// with the default values initialized, and the ability to set a context for a request
func NewIndexRepoParamsWithContext(ctx context.Context) *IndexRepoParams {
	var ()
	return &IndexRepoParams{

		Context: ctx,
	}
}

// NewIndexRepoParamsWithHTTPClient creates a new IndexRepoParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewIndexRepoParamsWithHTTPClient(client *http.Client) *IndexRepoParams {
	var ()
	return &IndexRepoParams{
		HTTPClient: client,
	}
}

/*IndexRepoParams contains all the parameters to send to the API endpoint
for the index repo operation typically these are written to a http.Request
*/
type IndexRepoParams struct {

	/*Body*/
	Body *models.OpenpitrixIndexRepoRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the index repo params
func (o *IndexRepoParams) WithTimeout(timeout time.Duration) *IndexRepoParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the index repo params
func (o *IndexRepoParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the index repo params
func (o *IndexRepoParams) WithContext(ctx context.Context) *IndexRepoParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the index repo params
func (o *IndexRepoParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the index repo params
func (o *IndexRepoParams) WithHTTPClient(client *http.Client) *IndexRepoParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the index repo params
func (o *IndexRepoParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the index repo params
func (o *IndexRepoParams) WithBody(body *models.OpenpitrixIndexRepoRequest) *IndexRepoParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the index repo params
func (o *IndexRepoParams) SetBody(body *models.OpenpitrixIndexRepoRequest) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *IndexRepoParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
