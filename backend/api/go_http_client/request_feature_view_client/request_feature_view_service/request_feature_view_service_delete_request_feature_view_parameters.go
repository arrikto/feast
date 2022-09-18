// Code generated by go-swagger; DO NOT EDIT.

package request_feature_view_service

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
)

// NewRequestFeatureViewServiceDeleteRequestFeatureViewParams creates a new RequestFeatureViewServiceDeleteRequestFeatureViewParams object
// with the default values initialized.
func NewRequestFeatureViewServiceDeleteRequestFeatureViewParams() *RequestFeatureViewServiceDeleteRequestFeatureViewParams {
	var ()
	return &RequestFeatureViewServiceDeleteRequestFeatureViewParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewRequestFeatureViewServiceDeleteRequestFeatureViewParamsWithTimeout creates a new RequestFeatureViewServiceDeleteRequestFeatureViewParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewRequestFeatureViewServiceDeleteRequestFeatureViewParamsWithTimeout(timeout time.Duration) *RequestFeatureViewServiceDeleteRequestFeatureViewParams {
	var ()
	return &RequestFeatureViewServiceDeleteRequestFeatureViewParams{

		timeout: timeout,
	}
}

// NewRequestFeatureViewServiceDeleteRequestFeatureViewParamsWithContext creates a new RequestFeatureViewServiceDeleteRequestFeatureViewParams object
// with the default values initialized, and the ability to set a context for a request
func NewRequestFeatureViewServiceDeleteRequestFeatureViewParamsWithContext(ctx context.Context) *RequestFeatureViewServiceDeleteRequestFeatureViewParams {
	var ()
	return &RequestFeatureViewServiceDeleteRequestFeatureViewParams{

		Context: ctx,
	}
}

// NewRequestFeatureViewServiceDeleteRequestFeatureViewParamsWithHTTPClient creates a new RequestFeatureViewServiceDeleteRequestFeatureViewParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewRequestFeatureViewServiceDeleteRequestFeatureViewParamsWithHTTPClient(client *http.Client) *RequestFeatureViewServiceDeleteRequestFeatureViewParams {
	var ()
	return &RequestFeatureViewServiceDeleteRequestFeatureViewParams{
		HTTPClient: client,
	}
}

/*RequestFeatureViewServiceDeleteRequestFeatureViewParams contains all the parameters to send to the API endpoint
for the request feature view service delete request feature view operation typically these are written to a http.Request
*/
type RequestFeatureViewServiceDeleteRequestFeatureViewParams struct {

	/*Name*/
	Name *string
	/*Project*/
	Project *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the request feature view service delete request feature view params
func (o *RequestFeatureViewServiceDeleteRequestFeatureViewParams) WithTimeout(timeout time.Duration) *RequestFeatureViewServiceDeleteRequestFeatureViewParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the request feature view service delete request feature view params
func (o *RequestFeatureViewServiceDeleteRequestFeatureViewParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the request feature view service delete request feature view params
func (o *RequestFeatureViewServiceDeleteRequestFeatureViewParams) WithContext(ctx context.Context) *RequestFeatureViewServiceDeleteRequestFeatureViewParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the request feature view service delete request feature view params
func (o *RequestFeatureViewServiceDeleteRequestFeatureViewParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the request feature view service delete request feature view params
func (o *RequestFeatureViewServiceDeleteRequestFeatureViewParams) WithHTTPClient(client *http.Client) *RequestFeatureViewServiceDeleteRequestFeatureViewParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the request feature view service delete request feature view params
func (o *RequestFeatureViewServiceDeleteRequestFeatureViewParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithName adds the name to the request feature view service delete request feature view params
func (o *RequestFeatureViewServiceDeleteRequestFeatureViewParams) WithName(name *string) *RequestFeatureViewServiceDeleteRequestFeatureViewParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the request feature view service delete request feature view params
func (o *RequestFeatureViewServiceDeleteRequestFeatureViewParams) SetName(name *string) {
	o.Name = name
}

// WithProject adds the project to the request feature view service delete request feature view params
func (o *RequestFeatureViewServiceDeleteRequestFeatureViewParams) WithProject(project *string) *RequestFeatureViewServiceDeleteRequestFeatureViewParams {
	o.SetProject(project)
	return o
}

// SetProject adds the project to the request feature view service delete request feature view params
func (o *RequestFeatureViewServiceDeleteRequestFeatureViewParams) SetProject(project *string) {
	o.Project = project
}

// WriteToRequest writes these params to a swagger request
func (o *RequestFeatureViewServiceDeleteRequestFeatureViewParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Name != nil {

		// query param name
		var qrName string
		if o.Name != nil {
			qrName = *o.Name
		}
		qName := qrName
		if qName != "" {
			if err := r.SetQueryParam("name", qName); err != nil {
				return err
			}
		}

	}

	if o.Project != nil {

		// query param project
		var qrProject string
		if o.Project != nil {
			qrProject = *o.Project
		}
		qProject := qrProject
		if qProject != "" {
			if err := r.SetQueryParam("project", qProject); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}