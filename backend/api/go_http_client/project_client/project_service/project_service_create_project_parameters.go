// Code generated by go-swagger; DO NOT EDIT.

package project_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"

	project_model "github.com/feast-dev/feast/backend/api/go_http_client/project_model"
)

// NewProjectServiceCreateProjectParams creates a new ProjectServiceCreateProjectParams object
// with the default values initialized.
func NewProjectServiceCreateProjectParams() *ProjectServiceCreateProjectParams {
	var ()
	return &ProjectServiceCreateProjectParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewProjectServiceCreateProjectParamsWithTimeout creates a new ProjectServiceCreateProjectParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewProjectServiceCreateProjectParamsWithTimeout(timeout time.Duration) *ProjectServiceCreateProjectParams {
	var ()
	return &ProjectServiceCreateProjectParams{

		timeout: timeout,
	}
}

// NewProjectServiceCreateProjectParamsWithContext creates a new ProjectServiceCreateProjectParams object
// with the default values initialized, and the ability to set a context for a request
func NewProjectServiceCreateProjectParamsWithContext(ctx context.Context) *ProjectServiceCreateProjectParams {
	var ()
	return &ProjectServiceCreateProjectParams{

		Context: ctx,
	}
}

// NewProjectServiceCreateProjectParamsWithHTTPClient creates a new ProjectServiceCreateProjectParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewProjectServiceCreateProjectParamsWithHTTPClient(client *http.Client) *ProjectServiceCreateProjectParams {
	var ()
	return &ProjectServiceCreateProjectParams{
		HTTPClient: client,
	}
}

/*ProjectServiceCreateProjectParams contains all the parameters to send to the API endpoint
for the project service create project operation typically these are written to a http.Request
*/
type ProjectServiceCreateProjectParams struct {

	/*Body*/
	Body *project_model.APIProject

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the project service create project params
func (o *ProjectServiceCreateProjectParams) WithTimeout(timeout time.Duration) *ProjectServiceCreateProjectParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the project service create project params
func (o *ProjectServiceCreateProjectParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the project service create project params
func (o *ProjectServiceCreateProjectParams) WithContext(ctx context.Context) *ProjectServiceCreateProjectParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the project service create project params
func (o *ProjectServiceCreateProjectParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the project service create project params
func (o *ProjectServiceCreateProjectParams) WithHTTPClient(client *http.Client) *ProjectServiceCreateProjectParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the project service create project params
func (o *ProjectServiceCreateProjectParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the project service create project params
func (o *ProjectServiceCreateProjectParams) WithBody(body *project_model.APIProject) *ProjectServiceCreateProjectParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the project service create project params
func (o *ProjectServiceCreateProjectParams) SetBody(body *project_model.APIProject) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *ProjectServiceCreateProjectParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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